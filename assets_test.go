package ciderui

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func get(t *testing.T, target, encoding string) *http.Response {
	t.Helper()
	request := httptest.NewRequest(http.MethodGet, target, nil)
	request.Header.Set("Accept-Encoding", encoding)
	recorder := httptest.NewRecorder()
	Assets().ServeHTTP(recorder, request)
	return recorder.Result()
}

func body(t *testing.T, response *http.Response) []byte {
	t.Helper()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestAssetsServesEachSetAtItsContentsURL(t *testing.T) {
	for _, c := range []struct {
		url  string
		rule string
	}{{SiteStylesheet(), ".cui-btn"}, {ConsoleStylesheet(), ".cui-console-btn"}} {
		if !regexp.MustCompile(`^/ciderui/(site|console)\.[0-9a-f]{12}\.css$`).MatchString(c.url) {
			t.Errorf("url %s", c.url)
		}
		plain := get(t, c.url, "")
		if plain.StatusCode != 200 || plain.Header.Get("Content-Type") != "text/css; charset=utf-8" || !strings.Contains(plain.Header.Get("Cache-Control"), "immutable") || plain.Header.Get("Vary") != "Accept-Encoding" {
			t.Errorf("%s: %d %v", c.url, plain.StatusCode, plain.Header)
		}
		text := body(t, plain)
		// @invariant: the URL's stem is the body's hash: a browser holding the URL
		// holds this body, none older
		sum := sha256.Sum256(text)
		if !strings.HasSuffix(c.url, "."+hex.EncodeToString(sum[:6])+".css") {
			t.Errorf("%s is not the URL of the body it served (sha256 %s)", c.url, hex.EncodeToString(sum[:6]))
		}
		if !bytes.Contains(text, []byte(c.rule)) {
			t.Errorf("%s has no %s", c.url, c.rule)
		}
		zipped := get(t, c.url, "br, gzip")
		if zipped.Header.Get("Content-Encoding") != "gzip" {
			t.Fatalf("%s: not gzip", c.url)
		}
		reader, err := gzip.NewReader(bytes.NewReader(body(t, zipped)))
		if err != nil {
			t.Fatal(err)
		}
		unzipped, err := io.ReadAll(reader)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(unzipped, text) {
			t.Errorf("%s: gzip differs", c.url)
		}
	}
	if SiteStylesheet() == ConsoleStylesheet() {
		t.Error("one URL for both sets")
	}
}

// a stylesheet's relative font URLs resolve beside it
func TestAssetsServesEveryFontAStylesheetNames(t *testing.T) {
	for _, url := range []string{SiteStylesheet(), ConsoleStylesheet()} {
		fonts := regexp.MustCompile(`url\("?(fonts/[^")]+)"?\)`).FindAllSubmatch(body(t, get(t, url, "")), -1)
		if len(fonts) == 0 {
			t.Fatalf("%s names no font", url)
		}
		for _, font := range fonts {
			response := get(t, AssetsPath+string(font[1]), "")
			// @trap: a font's name carries no hash: cached a day, never immutable,
			// or an upgrade leaves every page its old faces
			if response.StatusCode != 200 || response.Header.Get("Content-Type") != "font/woff2" || strings.Contains(response.Header.Get("Cache-Control"), "immutable") || len(body(t, response)) == 0 {
				t.Errorf("%s: %d", font[1], response.StatusCode)
			}
		}
	}
}

func TestAssetsFindsNothingElse(t *testing.T) {
	for _, target := range []string{"/ciderui/", "/ciderui/site.css", "/ciderui/fonts/none.woff2", "/ciderui/fonts/../index.css", "/ciderui/index.css"} {
		if got := get(t, target, "").StatusCode; got != 404 {
			t.Errorf("%s: %d", target, got)
		}
	}
}

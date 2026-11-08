package ciderui_test

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/at-least/ciderui"
)

func get(t *testing.T, target, encoding string) *http.Response {
	t.Helper()
	request := httptest.NewRequest(http.MethodGet, target, nil)
	request.Header.Set("Accept-Encoding", encoding)
	recorder := httptest.NewRecorder()
	ciderui.Assets().ServeHTTP(recorder, request)
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

type set struct{ name, url, rule string }

func sets() []set {
	return []set{
		{"site", ciderui.SiteStylesheet(), ".cui-btn"},
		{"console", ciderui.ConsoleStylesheet(), ".cui-console-btn"},
	}
}

// @invariant: the URL's stem is the body's hash: a browser holding the URL
// holds this body, none older
func TestAStylesheetsURLIsItsContentsHash(t *testing.T) {
	for _, s := range sets() {
		if !regexp.MustCompile(`^/ciderui/` + s.name + `\.[0-9a-f]{12}\.css$`).MatchString(s.url) {
			t.Errorf("the %s stylesheet's URL is %s, not /ciderui/%s.<12 hex>.css", s.name, s.url, s.name)
		}
		response := get(t, s.url, "")
		text := body(t, response)
		if response.StatusCode != http.StatusOK || !bytes.Contains(text, []byte(s.rule)) {
			t.Fatalf("%s served %d with no %s", s.url, response.StatusCode, s.rule)
		}
		sum := sha256.Sum256(text)
		if !strings.HasSuffix(s.url, "."+hex.EncodeToString(sum[:6])+".css") {
			t.Errorf("%s is not the URL of the body it served (sha256 %s)", s.url, hex.EncodeToString(sum[:6]))
		}
	}
}

func TestAStylesheetIsCachedForeverAsCSS(t *testing.T) {
	for _, s := range sets() {
		header := get(t, s.url, "").Header
		for name, want := range map[string]string{
			"Content-Type":  "text/css; charset=utf-8",
			"Cache-Control": "public, max-age=31536000, immutable",
			"Vary":          "Accept-Encoding",
		} {
			if got := header.Get(name); got != want {
				t.Errorf("%s: %s is %q, want %q", s.url, name, got, want)
			}
		}
	}
}

func TestAStylesheetIsGzippedForABrowserThatAccepts(t *testing.T) {
	for _, s := range sets() {
		plain := get(t, s.url, "")
		if got := plain.Header.Get("Content-Encoding"); got != "" {
			t.Errorf("%s with no Accept-Encoding came as %s", s.url, got)
		}
		text := body(t, plain)
		zipped := get(t, s.url, "br, gzip")
		if got := zipped.Header.Get("Content-Encoding"); got != "gzip" {
			t.Fatalf("%s for a browser that accepts gzip came as %q", s.url, got)
		}
		reader, err := gzip.NewReader(bytes.NewReader(body(t, zipped)))
		if err != nil {
			t.Fatal(err)
		}
		unzipped, err := io.ReadAll(reader)
		if err != nil {
			t.Fatalf("%s: the gzip body does not unzip: %v", s.url, err)
		}
		if !bytes.Equal(unzipped, text) {
			t.Errorf("%s: the gzip body unzips to another stylesheet", s.url)
		}
	}
}

// a stylesheet's relative font URLs resolve beside it
func TestAssetsServesEveryFontAStylesheetNames(t *testing.T) {
	for _, s := range sets() {
		fonts := regexp.MustCompile(`url\("?(fonts/[^")]+)"?\)`).FindAllSubmatch(body(t, get(t, s.url, "")), -1)
		if len(fonts) == 0 {
			t.Fatalf("%s names no font", s.url)
		}
		for _, font := range fonts {
			response := get(t, ciderui.AssetsPath+string(font[1]), "")
			if response.StatusCode != http.StatusOK || response.Header.Get("Content-Type") != "font/woff2" || len(body(t, response)) == 0 {
				t.Errorf("%s: served %d as %q", font[1], response.StatusCode, response.Header.Get("Content-Type"))
			}
			// @trap: a font's name carries no hash: cached a day, never immutable,
			// or an upgrade leaves every page its old faces
			if got := response.Header.Get("Cache-Control"); got != "public, max-age=86400" {
				t.Errorf("%s: Cache-Control is %q, want a day", font[1], got)
			}
		}
	}
}

func TestAssetsFindsNothingElse(t *testing.T) {
	for _, target := range []string{"/ciderui/", "/ciderui/site.css", "/ciderui/fonts/none.woff2", "/ciderui/fonts/../index.css", "/ciderui/index.css"} {
		if got := get(t, target, "").StatusCode; got != http.StatusNotFound {
			t.Errorf("%s served %d, want 404", target, got)
		}
	}
}

// @invariant: target/ciderui is the command's own example: a dir not made yet
func TestWriteCSSWritesUtilitiesAndWhatItImports(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "target", "ciderui")
	if err := ciderui.WriteCSS(dir); err != nil {
		t.Fatalf("WriteCSS into a dir not made yet: %v", err)
	}
	utilities, err := os.ReadFile("utilities.css")
	if err != nil {
		t.Fatal(err)
	}
	files := []string{"utilities.css"}
	for _, match := range regexp.MustCompile(`@import "\./([^"]+)"`).FindAllStringSubmatch(string(utilities), -1) {
		files = append(files, match[1])
	}
	if len(files) < 3 {
		t.Fatalf("utilities.css imports %v", files[1:])
	}
	for _, file := range files {
		written, err := os.ReadFile(filepath.Join(dir, file))
		if err != nil {
			t.Fatalf("WriteCSS wrote no %s: %v", file, err)
		}
		kit, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(written, kit) {
			t.Errorf("the %s WriteCSS wrote differs from the kit's", file)
		}
	}
}

// an error names the directory it could not make: MkdirAll's own error names
// only the component that failed
func TestWriteCSSErrorsNameTheDir(t *testing.T) {
	blocked := filepath.Join(t.TempDir(), "not-a-dir")
	if err := os.WriteFile(blocked, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	err := ciderui.WriteCSS(filepath.Join(blocked, "under"))
	if err == nil || !strings.Contains(err.Error(), filepath.Join(blocked, "under")) {
		t.Errorf("WriteCSS under a file gave %v, want an error naming %s", err, filepath.Join(blocked, "under"))
	}
}

func TestWriteCSSFailsOnAFileItCannotWrite(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "utilities.css"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := ciderui.WriteCSS(dir); err == nil {
		t.Error("WriteCSS over a directory named utilities.css reported no error")
	}
}

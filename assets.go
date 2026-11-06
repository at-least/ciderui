package ciderui

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"net/http"
	"path"
	"strings"
	"sync"
)

// @external: written by scripts/build-css.mjs; npm test fails a stale one
//
//go:embed dist
var dist embed.FS

// AssetsPath is where a site mounts [Assets]:
// mux.Handle(ciderui.AssetsPath, ciderui.Assets()). The stylesheets' url("fonts/…")
// are relative to them, so a stylesheet and its fonts share this one path.
const AssetsPath = "/ciderui/"

type stylesheet struct {
	url          string
	plain, gzipd []byte
}

func gzipped(plain []byte) []byte {
	var out bytes.Buffer
	zipper := gzip.NewWriter(&out)
	if _, err := zipper.Write(plain); err != nil {
		panic(err)
	}
	if err := zipper.Close(); err != nil {
		panic(err)
	}
	return out.Bytes()
}

func built(set string) stylesheet {
	plain, err := dist.ReadFile("dist/" + set + ".css")
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(plain)
	return stylesheet{url: AssetsPath + set + "." + hex.EncodeToString(sum[:6]) + ".css", plain: plain, gzipd: gzipped(plain)}
}

var (
	site    = sync.OnceValue(func() stylesheet { return built("site") })
	console = sync.OnceValue(func() stylesheet { return built("console") })
)

// SiteStylesheet is the URL of the site set's stylesheet, changing with its
// content, so a browser never holds an old one after a site upgrades the module.
func SiteStylesheet() string { return site().url }

// ConsoleStylesheet is [SiteStylesheet] for the console set.
func ConsoleStylesheet() string { return console().url }

func (s stylesheet) serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	w.Header().Set("Vary", "Accept-Encoding")
	body := s.plain
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		body = s.gzipd
	}
	// @trap: a client gone mid-response has nothing to be told
	_, _ = w.Write(body)
}

// Assets serves the sets' stylesheets and the fonts under [AssetsPath].
func Assets() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SiteStylesheet():
			site().serve(w, r)
			return
		case ConsoleStylesheet():
			console().serve(w, r)
			return
		}
		name, isFont := strings.CutPrefix(r.URL.Path, AssetsPath+"fonts/")
		if !isFont || path.Base(name) != name {
			http.NotFound(w, r)
			return
		}
		// @trap: not immutable: the name carries no hash
		w.Header().Set("Cache-Control", "public, max-age=86400")
		http.ServeFileFS(w, r, kit, "fonts/"+name)
	})
}

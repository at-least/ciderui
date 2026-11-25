// Command scheme: BROWSER=safari|ios|chrome|chromium|firefox go run ./reference/scheme
// Does light-dark() follow the color-scheme a section sets, and back again inside it?
// Serves page.html on 127.0.0.1 (the browsers: internal/probe); the result is
// results/<browser>.json.
package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	browser := probe.Browser()
	html, err := os.ReadFile(filepath.Join(probe.Dir("scheme"), "page.html"))
	if err != nil {
		return err
	}
	src, err := os.ReadFile(filepath.Join(probe.Dir("scheme"), "scheme.js"))
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf-8")
		w.Write(html)
	})}
	go func() {
		if err := server.Serve(l); !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintln(os.Stderr, "serve:", err)
		}
	}()
	defer server.Close()
	page, err := probe.Open(browser, fmt.Sprintf("http://127.0.0.1:%d/", l.Addr().(*net.TCPAddr).Port), nil)
	if err != nil {
		return err
	}
	v, err := page.Run("(async () => {\n" + string(src) + "\nreturn scheme();\n})()")
	label := page.Label()
	if err := closeAfter(page, err); err != nil {
		return err
	}
	result, ok := v.(ojson.Object)
	if !ok {
		return fmt.Errorf("scheme() gave %T", v)
	}
	if err := probe.Record("scheme", browser+".json", ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}}.With(result...)); err != nil {
		return err
	}
	light := []string{"rgb(29, 29, 31)", "rgb(255, 255, 255)"}
	dark := []string{"rgb(245, 245, 247)", "rgb(0, 0, 0)"}
	want := []struct {
		id   string
		pair []string
	}{{"page", light}, {"dark", dark}, {"in-dark", dark}, {"back", light}}
	fmt.Printf("%s\n%s\n", label, result.Get("ua"))
	got, _ := result.Get("got").(ojson.Object)
	pass := true
	for _, w := range want {
		var have []string
		if a, ok := got.Get(w.id).(ojson.Array); ok {
			for _, e := range a {
				s, _ := e.(ojson.String)
				have = append(have, string(s))
			}
		}
		ok := slices.Equal(have, w.pair)
		pass = pass && ok
		line := fmt.Sprintf("ok   %s: %s", w.id, strings.Join(have, " on "))
		if !ok {
			line = fmt.Sprintf("FAIL %s: %s (want %s)", w.id, strings.Join(have, " on "), strings.Join(w.pair, " on "))
		}
		fmt.Println(line)
	}
	if !pass {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("PASS")
	return nil
}

// @invariant: the page closes whether or not its run failed, and the run's error comes
// first
func closeAfter(page probe.Page, err error) error {
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	return err
}

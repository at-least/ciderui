// Command stack: BROWSER=safari|ios|chrome|chromium|firefox EXPECT=sf|inter go run ./reference/stack
// Serves the repository on 127.0.0.1 and opens page.html (the browsers:
// internal/probe). EXPECT=sf: the stack draws the system font and no Inter
// face was fetched; EXPECT=inter: the stack draws Inter, and its Latin face
// was fetched. The result is results/<browser>.json.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/at-least/ciderui/internal/build"
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
	expect := os.Getenv("EXPECT")
	if expect != "sf" && expect != "inter" {
		return fmt.Errorf("EXPECT must be sf or inter, not %s", expect)
	}
	browser := probe.Browser()

	types := map[string]string{".html": "text/html; charset=utf-8", ".css": "text/css", ".woff2": "font/woff2"}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, err := url.PathUnescape(r.URL.EscapedPath())
		if err != nil || strings.Contains(path, "..") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := os.ReadFile(filepath.Join(build.Root, filepath.FromSlash(path)))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		t, ok := types[filepath.Ext(path)]
		if !ok {
			t = "application/octet-stream"
		}
		w.Header().Set("content-type", t)
		w.Write(body)
	})}
	go server.Serve(l)
	defer server.Close()
	pageURL := fmt.Sprintf("http://127.0.0.1:%d/reference/stack/page.html", l.Addr().(*net.TCPAddr).Port)

	src, err := os.ReadFile(filepath.Join(probe.Dir("stack"), "stack.js"))
	if err != nil {
		return err
	}
	page, err := probe.Open(browser, pageURL, nil)
	if err != nil {
		return err
	}
	v, err := page.Run("(async () => {\n" + string(src) + "\nreturn stack();\n})()")
	label := page.Label()
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return err
	}
	result, ok := v.(ojson.Object)
	if !ok {
		return fmt.Errorf("stack() gave %T", v)
	}
	if err := probe.Record("stack", browser+".json", ojson.Object{{Key: "browser", Value: ojson.String(label)}, {Key: "date", Value: probe.Date()}, {Key: "expect", Value: ojson.String(expect)}}.With(result...)); err != nil {
		return err
	}

	before, _ := result.Get("before").(ojson.Object)
	fmt.Printf("%s\n%s\n", label, plain(result.Get("ua")))
	fmt.Println("fetched by first paint:", compact(before.Get("fetched")))
	fmt.Println("Inter's faces then:", compact(before.Get("inter")))
	stackWidth, systemWidth, interWidth := number(result.Get("stackWidth")), number(result.Get("systemWidth")), number(result.Get("interWidth"))
	fmt.Printf("width: stack %spx, system-ui %spx, Inter %spx\n", toFixed(stackWidth, 3), toFixed(systemWidth, 3), toFixed(interWidth, 3))
	same := func(a, b float64) bool { return math.Abs(a-b) < 0.01 }
	fetched := stringsOf(before.Get("fetched"))
	inter := stringsOf(before.Get("inter"))
	interFetched := slices.ContainsFunc(fetched, func(n string) bool { return strings.HasPrefix(n, "inter-") })
	type check struct {
		name string
		ok   bool
	}
	var checks []check
	if expect == "sf" {
		checks = []check{
			{"the stack draws the system font", same(stackWidth, systemWidth)},
			{"the system font is not Inter", !same(systemWidth, interWidth)},
			{"no Inter face fetched", !interFetched},
			{"every Inter face unloaded", len(inter) > 0 && !slices.ContainsFunc(inter, func(s string) bool { return s != "unloaded" })},
		}
	} else {
		checks = []check{
			{"the stack draws Inter", same(stackWidth, interWidth)},
			{"Inter is not the system font here", !same(systemWidth, interWidth)},
			{"Inter's Latin face fetched", slices.Contains(fetched, "inter-roman-latin.woff2")},
		}
	}
	pass := true
	for _, c := range checks {
		pass = pass && c.ok
		mark := "ok  "
		if !c.ok {
			mark = "FAIL"
		}
		fmt.Printf("%s %s\n", mark, c.name)
	}
	if !pass {
		fmt.Printf("FAIL (EXPECT=%s)\n", expect)
		os.Exit(1)
	}
	fmt.Printf("PASS (EXPECT=%s)\n", expect)
	return nil
}

// @invariant: a value as JavaScript prints it in a template: a string bare
func plain(v ojson.Value) string {
	if s, ok := v.(ojson.String); ok {
		return string(s)
	}
	return compact(v)
}

// @invariant: JSON.stringify(v), with no indent, its keys in order
func compact(v ojson.Value) string {
	var b bytes.Buffer
	if err := json.Compact(&b, ojson.Marshal(v)); err != nil {
		panic(err)
	}
	return b.String()
}

func number(v ojson.Value) float64 {
	var f float64
	if err := json.Unmarshal(ojson.Marshal(v), &f); err != nil {
		panic(fmt.Sprintf("%s is no number", ojson.Marshal(v)))
	}
	return f
}

func stringsOf(v ojson.Value) []string {
	var s []string
	if a, ok := v.(ojson.Array); ok {
		for _, e := range a {
			str, _ := e.(ojson.String)
			s = append(s, string(str))
		}
	}
	return s
}

// @trap: JavaScript's toFixed takes the larger of two equally near: 512.8125 is
// 512.813 there, where %.3f rounds to even (512.812); the widths are 64ths of a
// pixel, so such ties occur
func toFixed(x float64, digits int) string {
	if x < 0 {
		return "-" + toFixed(-x, digits)
	}
	scale := new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits)), nil))
	r := new(big.Rat).Mul(new(big.Rat).SetFloat64(x), scale)
	r.Add(r, big.NewRat(1, 2))
	n := new(big.Int).Quo(r.Num(), r.Denom())
	return new(big.Rat).SetFrac(n, scale.Num()).FloatString(digits)
}

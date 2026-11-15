// Command tracking: how far system-ui's widths sit from apple.com's own cut (SF Pro
// Text, SF Pro Display) at apple.com's sizes, and the letter-spacing that lands the
// kit's font stack on them. One step a run:
//
//	BROWSER=safari|ios|chrome|chromium|firefox MODE=measure|verify|inter go run ./reference/tracking run
//	    the page is https://www.apple.com/ (the browsers: internal/probe); the
//	    result is results/<browser>-<mode>.json, then analyzed
//	BROWSER=… MODE=… go run ./reference/tracking analyze
//	    run's analysis of the results/<browser>-<mode>.json written before
//	go run ./reference/tracking derive
//	    table/table.go, the table (table.CUITracking), from the Safari and Chrome
//	    measurements
//	go run ./reference/tracking tables
//	    MEASURED.md's two tables, from results/ and the table
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go/format"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
	"github.com/at-least/ciderui/internal/reference"
	"github.com/at-least/ciderui/reference/tracking/table"
)

// @external: apple.com's scale (measured, ../apple.com/: the store's t- ladder and the
// home page's 56px, 2026-10-19): the size, the cut it sets at that size, and
// its letter-spacing in em. 32px and 48px are the store's and the product
// page's (+0.004em, -0.003em), not the home page's (+0.002em, -0.002em).
var scale = []struct {
	size          int
	cut           string
	letterSpacing float64
}{
	{12, "Text", -0.01}, {14, "Text", -0.016}, {17, "Text", -0.022},
	{19, "Display", 0.012}, {21, "Display", 0.011}, {24, "Display", 0.009}, {28, "Display", 0.007},
	{32, "Display", 0.004}, {40, "Display", 0}, {48, "Display", -0.003}, {56, "Display", -0.005},
	{64, "Display", -0.009}, {80, "Display", -0.015}, {96, "Display", -0.015},
}

func scaleAt(size float64) (cut string, letterSpacing float64) {
	for _, s := range scale {
		if float64(s.size) == size {
			return s.cut, s.letterSpacing
		}
	}
	panic(fmt.Sprintf("tracking: no size %v in the scale", size))
}

var weights = []int{400, 600}

const sample = "Sphinx of black quartz judge my vow 0123456789 AVATAR Wave"

// @invariant: decided before the run: system-ui matches apple.com's cut within 0.003em per
// character (apple.com's smallest nonzero letter-spacing is 0.002em)
const tolerance = 0.003

// @invariant: verify: decided before the run, the kit's own font stack (tokens.css's
// --cui-font, whichever of its names the browser has: -apple-system in
// Safari, BlinkMacSystemFont in Chrome, which has no -apple-system) at
// ciderui's letter-spacing within 0.0015em per character of apple.com's text
// (the table averages two weights and two browsers; at 56px 0.0015em is 0.08px
// a character); each name read beside it, system-ui too, which derived it
const verifyTolerance = 0.0015

var stack = func() string {
	css, err := os.ReadFile(filepath.Join(build.Root, "tokens.css"))
	if err != nil {
		panic(err)
	}
	return regexp.MustCompile(`--cui-font:\s*([^;]+);`).FindStringSubmatch(string(css))[1]
}()

var families = []string{"system-ui", "-apple-system", "BlinkMacSystemFont", stack}

type row struct {
	Size    float64            `json:"size"`
	Weight  float64            `json:"weight"`
	Text    float64            `json:"text"`
	Display float64            `json:"display"`
	Sys     float64            `json:"sys"`
	Fams    map[string]float64 `json:"fams"`
	Web     float64            `json:"web"`
	Inter   float64            `json:"inter"`
	Inter0  float64            `json:"inter0"`
	Web0    float64            `json:"web0"`
}

type result struct {
	Browser string `json:"browser"`
	UA      string `json:"ua"`
	Rows    []row  `json:"rows"`
}

func (r result) find(size, weight float64) row {
	for _, x := range r.Rows {
		if x.Size == size && x.Weight == weight {
			return x
		}
	}
	panic(fmt.Sprintf("tracking: %s has no row at %vpx and %v", r.Browser, size, weight))
}

var n = float64(utf8.RuneCountInString(sample))

func em(a, b, size float64) float64 { return (a - b) / (n * size) }

// @trap: JavaScript's Number.prototype.toFixed: of the two nearest, the larger in
// magnitude where x lies halfway, which strconv's 'f' (to the even one) is not
func toFixed(x float64, digits int) string {
	if math.IsNaN(x) {
		return "NaN"
	}
	if x < 0 {
		return "-" + toFixed(-x, digits)
	}
	exact := new(big.Float).SetFloat64(x).Text('f', 1100)
	whole, frac, _ := strings.Cut(exact, ".")
	kept, rest := frac[:digits], frac[digits:]
	n, ok := new(big.Int).SetString(whole+kept, 10)
	if !ok {
		panic(exact)
	}
	if rest[0] >= '5' {
		n.Add(n, big.NewInt(1))
	}
	s := n.String()
	if len(s) <= digits {
		s = strings.Repeat("0", digits-len(s)+1) + s
	}
	if digits == 0 {
		return s
	}
	return s[:len(s)-digits] + "." + s[len(s)-digits:]
}

func signed(x float64) string {
	if x >= 0 {
		return "+" + toFixed(x, 4)
	}
	return toFixed(x, 4)
}

func number(x float64) string { return string(ojson.Float(x)) }

func padStart(s string, width int) string {
	return strings.Repeat(" ", max(0, width-utf8.RuneCountInString(s))) + s
}

func padEnd(s string, width int) string {
	return s + strings.Repeat(" ", max(0, width-utf8.RuneCountInString(s)))
}

func analyze(label string, r result) bool {
	fmt.Printf("\n%s\n%s\n", label, r.UA)
	fmt.Println("size weight cut      apple-ls  Δ(sys−cut)em  Δ(sys−Text)  Δ(sys−Display)  ok")
	pass := true
	for _, x := range r.Rows {
		cut, ls := scaleAt(x.Size)
		against := x.Display
		if cut == "Text" {
			against = x.Text
		}
		d := em(x.Sys, against, x.Size)
		ok := math.Abs(d) < tolerance
		pass = pass && ok
		yes := "NO"
		if ok {
			yes = "yes"
		}
		fmt.Printf("%s %s %s %s  %s  %s  %s  %s\n", padStart(number(x.Size), 4), padStart(number(x.Weight), 6), padEnd(cut, 8), padStart(signed(ls), 8),
			padStart(signed(d), 12), padStart(signed(em(x.Sys, x.Text, x.Size)), 11), padStart(signed(em(x.Sys, x.Display, x.Size)), 14), yes)
	}
	if pass {
		fmt.Printf("PASS: system-ui within %sem of apple.com's cut at every size\n", number(tolerance))
	} else {
		fmt.Printf("FAIL: system-ui off apple.com's cut by %sem or more somewhere\n", number(tolerance))
	}
	return pass
}

func analyzeVerify(label string, r result) bool {
	fmt.Printf("\n%s\n%s\n", label, r.UA)
	fmt.Println("Δ(name at ciderui ls − apple.com cut at apple.com ls), em per character")
	head := "size weight"
	for _, f := range families {
		name := f
		if f == stack {
			name = "the stack"
		}
		head += padStart(name, 20)
	}
	fmt.Println(head)
	worst := map[string]float64{}
	for _, x := range r.Rows {
		cells := ""
		for _, f := range families {
			// @invariant: a name the run did not read (a record from before tokens.css's
			// stack changed) is NaN, as JavaScript's undefined made it: its column NaN, the
			// run a FAIL
			fam, ok := x.Fams[f]
			if !ok {
				fam = math.NaN()
			}
			d := em(fam, x.Web, x.Size)
			worst[f] = math.Max(worst[f], math.Abs(d))
			cells += padStart(signed(d), 20)
		}
		fmt.Printf("%s %s%s\n", padStart(number(x.Size), 4), padStart(number(x.Weight), 6), cells)
	}
	for _, f := range families {
		name := f
		if f == stack {
			name = "the stack (" + stack + ")"
		}
		within := "OUTSIDE"
		if worst[f] < verifyTolerance {
			within = "within"
		}
		fmt.Printf("worst |Δ| %s: %sem %s %s\n", name, toFixed(worst[f], 4), within, number(verifyTolerance))
	}
	pass := worst[stack] < verifyTolerance
	if pass {
		fmt.Println("PASS: the kit's font stack at ciderui's letter-spacing lands on apple.com's widths")
	} else {
		fmt.Println("FAIL: the kit's font stack at ciderui's letter-spacing misses apple.com's widths")
	}
	return pass
}

// @invariant: inter: no pass mark for the distance, which is what is being measured; the
// run fails only if Inter did not load (inter.js throws) or drew nothing.
func analyzeInter(label string, r result) bool {
	fmt.Printf("\n%s\n%s\n", label, r.UA)
	fmt.Println("size weight  Δ(Inter at ciderui ls − apple.com)  Inter alone: ls to land on apple.com")
	drew := true
	worst := math.Inf(-1)
	for _, x := range r.Rows {
		drew = drew && x.Inter0 > 0
		d := em(x.Inter, x.Web, x.Size)
		_, ls := scaleAt(x.Size)
		own := ls - em(x.Inter0, x.Web0, x.Size)
		worst = math.Max(worst, math.Abs(d))
		fmt.Printf("%s %s  %s  %s\n", padStart(number(x.Size), 4), padStart(number(x.Weight), 6), padStart(signed(d), 33), padStart(signed(own), 36))
	}
	fmt.Printf("worst |Δ| Inter at ciderui's letter-spacing: %sem a character\n", toFixed(worst, 4))
	return drew
}

var analyses = map[string]func(string, result) bool{"measure": analyze, "verify": analyzeVerify, "inter": analyzeInter}

func main() {
	steps := map[string]func() (bool, error){"run": run, "analyze": reanalyze, "derive": derive, "tables": tables}
	if len(os.Args) < 2 || steps[os.Args[1]] == nil {
		fmt.Fprintln(os.Stderr, "usage: go run ./reference/tracking run|analyze|derive|tables")
		os.Exit(2)
	}
	pass, err := steps[os.Args[1]]()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !pass {
		os.Exit(1)
	}
}

func mode() string {
	if m := os.Getenv("MODE"); m != "" {
		return m
	}
	return "measure"
}

func dir(name string) string { return filepath.Join(probe.Dir("tracking"), name) }

func load(name string) (result, error) {
	var r result
	raw, err := os.ReadFile(dir("results/" + name + ".json"))
	if err != nil {
		return r, err
	}
	if err := json.Unmarshal(raw, &r); err != nil {
		return r, fmt.Errorf("results/%s.json: %w", name, err)
	}
	return r, nil
}

// @invariant: the table as the page reads it: {"<size>": em, …}, in its order
func tableJSON(rows []reference.Tracking) ojson.Object {
	o := ojson.Object{}
	for _, r := range rows {
		o = append(o, ojson.Member{Key: strconv.Itoa(r.Size), Value: ojson.Float(r.Em)})
	}
	return o
}

func numbers[T int | float64](list []T) ojson.Array {
	a := ojson.Array{}
	for _, v := range list {
		a = append(a, ojson.Float(float64(v)))
	}
	return a
}

func strs(list []string) ojson.Array {
	a := ojson.Array{}
	for _, s := range list {
		a = append(a, ojson.String(s))
	}
	return a
}

// @invariant: the scale as its object literal writes it: the size, then the cut and its
// letter-spacing
func scaleJSON() ojson.Object {
	o := ojson.Object{}
	for _, s := range scale {
		o = append(o, ojson.Member{Key: strconv.Itoa(s.size), Value: ojson.Array{ojson.String(s.cut), ojson.Float(s.letterSpacing)}})
	}
	return o
}

func run() (bool, error) {
	browser := probe.Browser()
	m := mode()
	analysis := analyses[m]
	src, err := os.ReadFile(dir(m + ".js"))
	if err != nil || analysis == nil {
		return false, fmt.Errorf("MODE=%s: no such mode (measure, verify, inter): %v", m, err)
	}
	var args ojson.Array
	switch m {
	case "measure":
		sizes := []int{}
		for _, s := range scale {
			sizes = append(sizes, s.size)
		}
		args = ojson.Array{numbers(sizes), numbers(weights), ojson.String(sample)}
	case "verify":
		args = ojson.Array{scaleJSON(), tableJSON(table.CUITracking), numbers(weights), ojson.String(sample), strs(families)}
	case "inter":
		// @invariant: INTER names the font file, by default the kit's Latin subset
		font := os.Getenv("INTER")
		if font == "" {
			font = filepath.Join(build.Root, "fonts", "inter-roman-latin.woff2")
		}
		bytes, err := os.ReadFile(font)
		if err != nil {
			return false, err
		}
		args = ojson.Array{scaleJSON(), tableJSON(table.CUITracking), numbers(weights), ojson.String(sample), ojson.String(base64.StdEncoding.EncodeToString(bytes))}
	}
	expr := fmt.Sprintf("(async () => {\n%s\nreturn %s(...%s);\n})()", src, m, ojson.Marshal(args))
	page, err := probe.Open(browser, "https://www.apple.com/", nil)
	if err != nil {
		return false, err
	}
	v, err := page.Run(expr)
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return false, err
	}
	got, ok := v.(ojson.Object)
	if !ok {
		return false, fmt.Errorf("%s() gave %T", m, v)
	}
	record := ojson.Object{{Key: "browser", Value: ojson.String(page.Label())}, {Key: "date", Value: probe.Date()}}.With(got...)
	if err := probe.Record("tracking", browser+"-"+m+".json", record); err != nil {
		return false, err
	}
	var r result
	if err := json.Unmarshal(ojson.Marshal(got), &r); err != nil {
		return false, err
	}
	return analysis(page.Label(), r), nil
}

func reanalyze() (bool, error) {
	m := mode()
	analysis := analyses[m]
	if analysis == nil {
		return false, fmt.Errorf("MODE=%s: no such mode (measure, verify, inter)", m)
	}
	r, err := load(probe.Browser() + "-" + m)
	if err != nil {
		return false, err
	}
	return analysis(r.Browser, r), nil
}

// @invariant: derive: from the Safari and Chrome runs, how far apart the two browsers are, and
// the letter-spacing system-ui needs per size to land on apple.com's widths
// (apple.com's value minus the measured Δ, averaged over both browsers and both
// weights).
func derive() (bool, error) {
	var runs []result
	for _, b := range []string{"safari", "chrome"} {
		r, err := load(b + "-measure")
		if err != nil {
			return false, err
		}
		runs = append(runs, r)
	}
	delta := func(x row) float64 {
		cut, _ := scaleAt(x.Size)
		if cut == "Text" {
			return (x.Sys - x.Text) / (n * x.Size)
		}
		return (x.Sys - x.Display) / (n * x.Size)
	}
	maxGap := 0.0
	var rows []reference.Tracking
	for _, s := range scale {
		var ds []float64
		for _, w := range weights {
			sa, ch := delta(runs[0].find(float64(s.size), float64(w))), delta(runs[1].find(float64(s.size), float64(w)))
			maxGap = math.Max(maxGap, math.Abs(sa-ch))
			ds = append(ds, sa, ch)
		}
		sum := ds[0]
		for _, d := range ds[1:] {
			sum += d
		}
		mean := sum / float64(len(ds))
		v, err := strconv.ParseFloat(toFixed(s.letterSpacing-mean, 4), 64)
		if err != nil {
			return false, err
		}
		rows = append(rows, reference.Tracking{Size: s.size, Em: v})
	}
	fmt.Println(runs[0].Browser + " | " + runs[1].Browser)
	fmt.Println("max |Δ safari − Δ chrome| over all sizes and weights:", toFixed(maxGap, 5), "em")
	fmt.Println("letter-spacing for system-ui (em):", compactJSON(tableJSON(rows)))
	src, err := tableFile(rows)
	if err != nil {
		return false, err
	}
	return true, os.WriteFile(dir("table/table.go"), src, 0o644)
}

// @invariant: table/table.go: rows as table.CUITracking, in their order, gofmt'd
func tableFile(rows []reference.Tracking) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("// Code generated by go run ./reference/tracking derive; DO NOT EDIT.\n\n")
	b.WriteString("// Package table is ciderui's letter-spacing per size, from reference/tracking's\n// Safari and Chrome measurements.\npackage table\n\nimport \"github.com/at-least/ciderui/internal/reference\"\n\n")
	b.WriteString("// CUITracking is the letter-spacing at each of apple.com's sizes, ascending.\nvar CUITracking = []reference.Tracking{\n")
	for _, r := range rows {
		fmt.Fprintf(&b, "{Size: %d, Em: %s},\n", r.Size, strconv.FormatFloat(r.Em, 'g', -1, 64))
	}
	b.WriteString("}\n")
	return format.Source(b.Bytes())
}

// @invariant: JSON.stringify(table) with no indent: a flat object of numbers
func compactJSON(o ojson.Object) string {
	parts := make([]string, len(o))
	for i, m := range o {
		parts[i] = string(ojson.Marshal(ojson.String(m.Key))) + ":" + string(ojson.Marshal(m.Value))
	}
	return "{" + strings.Join(parts, ",") + "}"
}

func tables() (bool, error) {
	cui := map[int]float64{}
	for _, r := range table.CUITracking {
		cui[r.Size] = r.Em
	}
	f := func(x float64) string {
		if x >= 0 {
			return "+" + toFixed(math.Abs(x), 4)
		}
		return "−" + toFixed(math.Abs(x), 4)
	}
	runs := map[string]result{}
	for _, b := range []string{"safari", "chrome", "ios"} {
		var err error
		if runs[b], err = load(b + "-measure"); err != nil {
			return false, err
		}
	}
	delta := func(r result, size, w float64) float64 {
		x := r.find(size, w)
		cut, _ := scaleAt(size)
		if cut == "Text" {
			return em(x.Sys, x.Text, size)
		}
		return em(x.Sys, x.Display, size)
	}
	fmt.Println("| size | apple.com cut | apple.com letter-spacing | Δ at 400 (Safari / Chrome / iOS) | Δ at 600 (Safari / Chrome / iOS) | ciderui letter-spacing |")
	fmt.Println("| ---: | --- | ---: | ---: | ---: | ---: |")
	for _, s := range scale {
		var cells []string
		for _, w := range weights {
			var bs []string
			for _, b := range []string{"safari", "chrome", "ios"} {
				bs = append(bs, f(delta(runs[b], float64(s.size), float64(w))))
			}
			cells = append(cells, strings.Join(bs, " / "))
		}
		ls, ok := cui[s.size]
		if !ok {
			return false, fmt.Errorf("table.CUITracking has no %d", s.size)
		}
		fmt.Printf("| %dpx | %s | %s | %s | %s | %s |\n", s.size, s.cut, f(s.letterSpacing), cells[0], cells[1], f(ls))
	}
	fmt.Println()
	inter, err := load("safari-inter")
	if err != nil {
		return false, err
	}
	fmt.Println("| size | Δ at 400 | Δ at 600 | Inter alone at 400 | Inter alone at 600 |")
	fmt.Println("| ---: | ---: | ---: | ---: | ---: |")
	for _, s := range scale {
		size := float64(s.size)
		var ds, alone []string
		for _, w := range weights {
			x := inter.find(size, float64(w))
			ds = append(ds, f(em(x.Inter, x.Web, size)))
			alone = append(alone, f(s.letterSpacing-em(x.Inter0, x.Web0, size)))
		}
		fmt.Printf("| %dpx | %s | %s |\n", s.size, strings.Join(ds, " | "), strings.Join(alone, " | "))
	}
	return true, nil
}

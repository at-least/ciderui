// Command anchor: go run ./reference/anchor: where apple.com lands a footnote's note
// under its sticky local nav, and what puts it there. Three pages in Playwright's
// Chromium (internal/probe), at 1440 by 900 and at 390 by 900: iPhone 18 Pro's,
// whose local nav is a newer one (nav#localnav, 72px), and MacBook Air's and AirPods
// Pro's, whose local nav is the older one (#ac-localnav, 52px, 48px to 833px). Each
// scrolled 2000px down, its first footnote mark with a box (a link to #footnote-…)
// clicked, and read 2s on: the note's top against the nav's foot, the page at its
// end or not, the root's scroll-padding-top, the note's scroll-margin-top before the
// click and after (its inline style too: what a script sets), and the rules in the
// page's stylesheets that set it on the note (their selector, value and sheet; a
// sheet the page cannot read, another origin's, counted). The result is
// results/apple.json; it records what each page does, and fails only where a page
// has no mark, note or nav to read.
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/probe"
)

var pages = []string{"iphone-18-pro", "macbook-air", "airpods-pro"}

// @external: the rules that set the note's scroll-margin-top, through grouping rules
// (@media, @supports, @layer)
const rules = `(note) => {
  const found = [];
  let unreadable = 0;
  const walk = (rules, sheet) => {
    for (const rule of rules) {
      if (rule.selectorText && rule.style.scrollMarginTop && note.matches(rule.selectorText)) found.push({ selector: rule.selectorText, value: rule.style.scrollMarginTop, sheet });
      if (rule.cssRules) walk(rule.cssRules, sheet);
    }
  };
  for (const s of document.styleSheets) {
    let rules;
    try { rules = s.cssRules; } catch { unreadable++; continue; }
    walk(rules, s.href ?? '(inline)');
  }
  return { found, unreadable };
}`

func main() {
	ok, err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !ok {
		fmt.Println("FAIL a page had no footnote mark, note or local nav to read (results/apple.json)")
		os.Exit(1)
	}
}

func quote(s string) string { return string(ojson.Marshal(ojson.String(s))) }

func num(v ojson.Value) float64 {
	n, _ := v.(ojson.Number)
	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		panic(fmt.Sprintf("anchor: %v is no number", v))
	}
	return f
}

// @invariant: a JavaScript value as a template literal writes it
func text(v ojson.Value) string {
	switch v := v.(type) {
	case ojson.String:
		return string(v)
	case ojson.Number:
		return string(ojson.Float(num(v)))
	case ojson.Bool:
		return strconv.FormatBool(bool(v))
	case nil:
		return "undefined"
	case ojson.Null:
		return "null"
	}
	panic(fmt.Sprintf("anchor: no text for %T", v))
}

func run() (bool, error) {
	got := ojson.Object{}
	set := func(k string, v ojson.Value) { got = got.With(ojson.Member{Key: k, Value: v}) }
	ok := true
	page, err := probe.Open("chromium", "https://www.apple.com/", nil)
	if err != nil {
		return false, err
	}
	err = func() error {
		set("browser", ojson.String(page.Label()))
		for _, name := range pages {
			for _, width := range []int{1440, 390} {
				key := fmt.Sprintf("%s at %d", name, width)
				if err := page.Size(width, 900); err != nil {
					return err
				}
				if err := page.Go("https://www.apple.com/" + name + "/"); err != nil {
					return err
				}
				if _, err := page.Run(probe.Wait(3000)); err != nil {
					return err
				}
				hrefV, err := page.Run(`Promise.resolve((() => {
        const a = [...document.querySelectorAll('a[href^="#footnote"]')].find((a) => a.getBoundingClientRect().width > 0);
        if (!a) return null;
        a.setAttribute('data-probe-mark', '');
        return a.getAttribute('href');
      })())`)
				if err != nil {
					return err
				}
				href, isString := hrefV.(ojson.String)
				if !isString || href == "" {
					set(key, ojson.Object{{Key: "mark", Value: ojson.Null{}}})
					ok = false
					continue
				}
				id := string(href)[1:]
				before, err := page.Run(fmt.Sprintf("Promise.resolve(getComputedStyle(document.getElementById(%s)).scrollMarginTop)", quote(id)))
				if err != nil {
					return err
				}
				if _, err := page.Run("Promise.resolve((scrollTo(0, 2000), true))"); err != nil {
					return err
				}
				if _, err := page.Run(probe.Wait(800)); err != nil {
					return err
				}
				if err := page.Click("[data-probe-mark]"); err != nil {
					return err
				}
				if _, err := page.Run(probe.Wait(2000)); err != nil {
					return err
				}
				rv, err := page.Run(fmt.Sprintf(`Promise.resolve((() => {
        const note = document.getElementById(%s);
        const nav = document.querySelector('#ac-localnav, nav#localnav');
        if (!note || !nav) return { note: Boolean(note), nav: Boolean(nav) };
        return {
          href: %s,
          hash: location.hash,
          nav: { id: nav.id, position: getComputedStyle(nav).position, height: nav.getBoundingClientRect().height, bottom: +nav.getBoundingClientRect().bottom.toFixed(2) },
          note: +note.getBoundingClientRect().top.toFixed(2),
          end: scrollY >= document.documentElement.scrollHeight - innerHeight - 1,
          rootScrollPaddingTop: getComputedStyle(document.documentElement).scrollPaddingTop,
          noteScrollMarginTop: getComputedStyle(note).scrollMarginTop,
          noteInlineScrollMarginTop: note.style.scrollMarginTop,
          rules: (%s)(note),
        };
      })())`, quote(id), quote(string(href)), rules))
				if err != nil {
					return err
				}
				r, _ := rv.(ojson.Object)
				if _, isNumber := r.Get("note").(ojson.Number); !isNumber {
					set(key, r)
					ok = false
					continue
				}
				r = r.With(ojson.Member{Key: "noteScrollMarginTopBefore", Value: before})
				nav, _ := r.Get("nav").(ojson.Object)
				note, bottom := num(r.Get("note")), num(nav.Get("bottom"))
				var lands string
				switch {
				case math.Abs(note-bottom) <= 1:
					lands = "at the nav's foot"
				case note < bottom:
					lands = "under the nav"
				case r.Get("end") == ojson.Bool(true):
					lands = "below the nav, the page at its end"
				default:
					lands = "below the nav"
				}
				r = r.With(ojson.Member{Key: "lands", Value: ojson.String(lands)})
				set(key, r)
				found, _ := r.Get("rules").(ojson.Object)
				var sets []string
				list, _ := found.Get("found").(ojson.Array)
				for _, f := range list {
					f, _ := f.(ojson.Object)
					sets = append(sets, fmt.Sprintf("%s { %s }", text(f.Get("selector")), text(f.Get("value"))))
				}
				setBy := strings.Join(sets, ", ")
				if setBy == "" {
					setBy = "no readable rule"
				}
				inline := text(r.Get("noteInlineScrollMarginTop"))
				if inline == "" {
					inline = "none"
				}
				fmt.Printf("%s at %dpx (%s, %s, %spx): the note %s, its top %spx, the nav's foot %spx; scroll-margin-top %s before the click, %s after (inline: %s; set by %s; %s sheets unreadable); the root's scroll-padding-top %s\n",
					name, width, text(nav.Get("id")), text(nav.Get("position")), text(nav.Get("height")), lands, text(r.Get("note")), text(nav.Get("bottom")),
					text(before), text(r.Get("noteScrollMarginTop")), inline, setBy, text(found.Get("unreadable")), text(r.Get("rootScrollPaddingTop")))
			}
		}
		return nil
	}()
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return false, err
	}
	record := ojson.Object{{Key: "date", Value: probe.Date()}}.With(got...)
	if err := probe.Record("anchor", "apple.json", record); err != nil {
		return false, err
	}
	return ok, nil
}

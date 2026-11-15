package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/internal/pagespec"
	"github.com/at-least/ciderui/internal/probe"
	"github.com/at-least/ciderui/internal/reference"
	"github.com/at-least/ciderui/reference/apple.com/pages"
	"github.com/mxschmitt/playwright-go"
)

// @invariant: in the frame: tag the probe's element and the one its states are forced on;
// its outerHTML's head, or null
const place = `async (probe) => {
  const y = probe.scroll ?? 0;
  if (window.scrollY !== y) {
    window.scrollTo(0, y);
    await new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900))));
  }
  if (probe.blur) document.activeElement?.blur();
  for (const [selector, name, value] of probe.set ?? []) {
    const target = document.querySelector(selector);
    if (!target) return null;
    if (name === 'value') target.value = value;
    else if (name === 'classList') target.classList.add(value);
    else if (value === null) target.removeAttribute(name);
    else target.setAttribute(name, value);
  }
  // @external: what a blur or a set moves (the page's script checking a field, a
  // label's transition) settled before it is read
  if (probe.blur || probe.set) await new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, 900))));
  const roots = (root) => [root, ...[...root.querySelectorAll('*')].filter(e => e.shadowRoot).flatMap(e => roots(e.shadowRoot))];
  for (const root of roots(document)) {
    root.querySelectorAll('[data-cui-probe]').forEach(e => e.removeAttribute('data-cui-probe'));
    root.querySelectorAll('[data-cui-state]').forEach(e => e.removeAttribute('data-cui-state'));
  }
  document.querySelectorAll('[data-cui-injected]').forEach(e => e.remove());
  let scope = document;
  let el;
  if (probe.inject) {
    const host = probe.into ? document.querySelector(probe.into) : document.body;
    if (!host) return null;
    const t = document.createElement('template');
    t.innerHTML = probe.inject.trim();
    const node = t.content.firstElementChild;
    node.setAttribute('data-cui-injected', '');
    host.append(node);
    scope = node;
    el = probe.select ? (node.matches(probe.select) ? node : node.querySelector(probe.select)) : node;
  } else {
    el = probe.select.split(' >>> ').reduce((root, part) => root && (root === document ? root : root.shadowRoot)?.querySelector(part), document);
  }
  if (!el) return null;
  const on = probe.stateOn ? scope.querySelector(probe.stateOn) ?? document.querySelector(probe.stateOn) : el;
  if (!on) return null;
  el.setAttribute('data-cui-probe', '');
  on.setAttribute('data-cui-state', '');
  // @external: without the probe's tags, and without query strings (the store's carry
  // session ids)
  return el.outerHTML.replace(/ data-cui-[a-z]+=""/g, '').replace(/(href|src|srcset)="([^"?]*)\?[^"]*"/g, '$1="$2?…"').slice(0, 240);
}`

var (
	pageText   = regexp.MustCompile(`>[^<]*`)
	pageWords  = regexp.MustCompile(` (aria-label|alt|title|placeholder)="[^"]+("|$)`)
	wordsMoved = ` ${1}="…${2}`
)

// @invariant: the head as a result records it, without the page's words: its text, and
// the values of the attributes that hold words (one the head cuts too), so a
// result says where a value was read, not what the page says
func strip(head string) string {
	return pageWords.ReplaceAllString(pageText.ReplaceAllString(head, ">"), wordsMoved)
}

// @invariant: in the frame: from now on, every 4ms, describe every animation not seen
// yet whose element is inside `within`
const watchAnimations = `(within) => {
  const seen = new Set();
  const t0 = performance.now();
  window.cuiAnimations = [];
  const name = (e) => ` + "`${e.tagName.toLowerCase()}${e.id ? '#' + e.id : ''}${[...e.classList].slice(0, 3).map(c => '.' + c).join('')}`" + `;
  const look = () => {
    for (const a of document.getAnimations()) {
      const target = a.effect?.target;
      if (seen.has(a) || !target?.closest?.(within)) continue;
      seen.add(a);
      const timing = a.effect.getTiming();
      const frames = a.effect.getKeyframes();
      const what = a.transitionProperty ?? a.animationName;
      const prop = Object.keys(frames[0] ?? {}).find(k => !['offset', 'computedOffset', 'easing', 'composite'].includes(k));
      window.cuiAnimations.push({
        kind: a.constructor.name, what, on: name(target), duration: timing.duration, delay: Math.round(timing.delay * 100) / 100, direction: timing.direction,
        // @external: a transition's timing is its effect's, an animation's its keyframes'
        easing: a.transitionProperty ? timing.easing : frames[0]?.easing,
        from: prop ? frames[0][prop] : null, to: prop ? frames.at(-1)[prop] : null,
        seen: Math.round(performance.now() - t0),
      });
    }
  };
  window.cuiWatch = setInterval(look, 4);
}`

const read = `(probe) => {
  // @external: in the document or a shadow root in it
  const walk = (root, s) => root.querySelector(s) ?? [...root.querySelectorAll('*')].filter(e => e.shadowRoot).map(e => walk(e.shadowRoot, s)).find(Boolean) ?? null;
  const el = walk(document, '[data-cui-probe]');
  const cs = getComputedStyle(el, probe.pseudo ?? null);
  const out = { styles: {} };
  for (const p of probe.props ?? []) out.styles[p] = cs.getPropertyValue(p);
  for (const v of probe.vars ?? []) out.styles[v] = cs.getPropertyValue(v).trim();
  for (const p of probe.image ?? []) {
    const m = cs.getPropertyValue(p).match(/^url\("data:image\/svg\+xml;base64,([A-Za-z0-9+/=]+)"\)$/);
    if (!m) throw new Error(` + "`${probe.name}: ${p} is not a base64 SVG`" + `);
    const svg = new DOMParser().parseFromString(atob(m[1]), 'image/svg+xml').documentElement;
    const paint = (e, a) => e.closest(` + "`[${a}]`" + `)?.getAttribute(a) ?? null;
    out.styles[p] = { viewBox: svg.getAttribute('viewBox'), shapes: [...svg.querySelectorAll('circle, ellipse, rect, path, polygon, polyline, line')].map(e => ({ shape: e.tagName, fill: paint(e, 'fill'), stroke: paint(e, 'stroke') })) };
  }
  if (probe.keyframes) {
    out.keyframes = {};
    for (const name of cs.animationName.split(',').map(n => n.trim()).filter(n => n !== 'none')) {
      let rule = null;
      for (const sheet of document.styleSheets) {
        // @external: a stylesheet from another origin cannot be read: the rule, if it is
        // there, stays not found, and the run fails
        let rules;
        try { rules = sheet.cssRules; } catch { continue; }
        rule = [...rules].find(r => r.type === CSSRule.KEYFRAMES_RULE && r.name === name) ?? rule;
      }
      out.keyframes[name] = rule && [...rule.cssRules].map(k => ({ at: k.keyText, style: k.style.cssText }));
    }
  }
  if (probe.sections) {
    const [section, attribute, row] = probe.sections;
    out.sections = [...el.querySelectorAll(section)].map(s => [s.getAttribute(attribute), s.querySelectorAll(row).length]);
  }
  if (probe.matches) out.matches = Object.fromEntries(probe.matches.map(s => [s, el.matches(s)]));
  if (probe.box) {
    const r = el.getBoundingClientRect();
    out.box = { width: r.width, height: r.height, ...(probe.position ? { top: r.top, left: r.left } : {}) };
  }
  return out;
}`

// @invariant: two frames, then ms
const settleFor = `(ms) => new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => setTimeout(r, ms))))`

const still = `*, *::before, *::after { transition: none !important; }`

// @external: a document's style does not reach into a shadow root: one there too
const stillInShadow = `() => {
  const walk = (root, s) => root.querySelector(s) ?? [...root.querySelectorAll('*')].filter(e => e.shadowRoot).map(e => walk(e.shadowRoot, s)).find(Boolean) ?? null;
  const root = walk(document, '[data-cui-state]').getRootNode();
  if (root === document) return;
  const still = document.createElement('style');
  still.textContent = '*, *::before, *::after { transition: none !important; }';
  still.setAttribute('data-cui-still', '');
  root.append(still);
}`

// @invariant: the state's last change must land while transitions are still off, or
// it starts one that a later probe reads in flight: flush, then remove
const unstill = `() => {
  const walk = (root, s) => root.querySelector(s) ?? [...root.querySelectorAll('*')].filter(e => e.shadowRoot).map(e => walk(e.shadowRoot, s)).find(Boolean) ?? null;
  void document.body.offsetHeight;
  for (let still; (still = walk(document, '[data-cui-still]'));) still.remove();
}`

// @invariant: an in-page function's value crosses as the page's own JSON.stringify
// text, so an object keeps its keys' order; its argument crosses as JSON text
// (playwright-go's argument serializer panics on a map other than map[string]any)
func call(f playwright.Frame, fn string, arg any) (ojson.Value, error) {
	b, err := json.Marshal(arg)
	if err != nil {
		return nil, err
	}
	v, err := f.Evaluate(`async (arg) => { const v = await (`+fn+`)(JSON.parse(arg)); return v === undefined ? undefined : JSON.stringify(v); }`, string(b))
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("the page gave %T, no JSON text", v)
	}
	return ojson.Parse([]byte(s))
}

type session struct {
	pw  playwright.Page
	cdp playwright.CDPSession
}

func (s session) send(method string, params map[string]any) (map[string]any, error) {
	v, err := s.cdp.Send(method, params)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}
	m, ok := v.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s gave %T", method, v)
	}
	return m, nil
}

// @external: a DevTools node's attributes are its names and values in one list
func hasAttribute(n map[string]any, name string) bool {
	attrs, _ := n["attributes"].([]any)
	return slices.Contains(attrs, any(name))
}

func children(n map[string]any, keys ...string) []map[string]any {
	var out []map[string]any
	for _, k := range keys {
		switch v := n[k].(type) {
		case []any:
			for _, c := range v {
				if m, ok := c.(map[string]any); ok {
					out = append(out, m)
				}
			}
		case map[string]any:
			out = append(out, v)
		}
	}
	return out
}

func find(n map[string]any, attribute string, keys ...string) map[string]any {
	if hasAttribute(n, attribute) {
		return n
	}
	for _, c := range children(n, keys...) {
		if f := find(c, attribute, keys...); f != nil {
			return f
		}
	}
	return nil
}

// @external: the DevTools node of the element tagged data-cui-state in this frame: a
// frame's document is reached through its iframe in the page's whole tree
func (s session) stateNode(frame playwright.Frame) (any, error) {
	if frame == s.pw.MainFrame() {
		doc, err := s.send("DOM.getDocument", map[string]any{"depth": 0})
		if err != nil {
			return nil, err
		}
		root := doc["root"].(map[string]any)
		q, err := s.send("DOM.querySelector", map[string]any{"nodeId": root["nodeId"], "selector": "[data-cui-state]"})
		if err != nil {
			return nil, err
		}
		if id, _ := q["nodeId"].(float64); id != 0 {
			return q["nodeId"], nil
		}
		// @external: in a shadow root, which DOM.querySelector does not enter
		whole, err := s.send("DOM.getDocument", map[string]any{"depth": -1, "pierce": true})
		if err != nil {
			return nil, err
		}
		if n := find(whole["root"].(map[string]any), "data-cui-state", "children", "shadowRoots"); n != nil {
			return n["nodeId"], nil
		}
		return nil, nil
	}
	iframe, err := frame.FrameElement()
	if err != nil {
		return nil, err
	}
	if _, err := iframe.Evaluate(`e => e.setAttribute('data-cui-frame', '')`); err != nil {
		return nil, err
	}
	whole, err := s.send("DOM.getDocument", map[string]any{"depth": -1, "pierce": true})
	if err != nil {
		return nil, err
	}
	if _, err := iframe.Evaluate(`e => e.removeAttribute('data-cui-frame')`); err != nil {
		return nil, err
	}
	var doc map[string]any
	if n := find(whole["root"].(map[string]any), "data-cui-frame", "children", "contentDocument", "shadowRoots"); n != nil {
		doc, _ = n["contentDocument"].(map[string]any)
	}
	if doc == nil {
		return nil, fmt.Errorf("the page's session does not reach the document of %s", frame.URL())
	}
	q, err := s.send("DOM.querySelector", map[string]any{"nodeId": doc["nodeId"], "selector": "[data-cui-state]"})
	if err != nil {
		return nil, err
	}
	return q["nodeId"], nil
}

func (s session) frameOf(p pagespec.Probe) (playwright.Frame, error) {
	if p.Frame == "" {
		return s.pw.MainFrame(), nil
	}
	for _, f := range s.pw.Frames() {
		if strings.Contains(f.URL(), p.Frame) {
			return f, nil
		}
	}
	return nil, fmt.Errorf("no frame whose URL holds %s", p.Frame)
}

func measure(name string) error {
	build, ok := pages.Specs[name]
	if !ok {
		return fmt.Errorf("PAGE=%q names no spec in pages/", name)
	}
	spec := build()
	widths := spec.ReadWidths()
	browser := spec.Browser
	if browser == "" {
		browser = "chromium"
	}
	page, err := probe.Open(browser, spec.URL, nil)
	if err != nil {
		return err
	}
	result, nulls, err := measureOpen(page, spec, widths)
	if cerr := page.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return err
	}
	if err := reference.WriteRecord(filepath.Join(probe.Dir("apple.com"), "results"), name, result); err != nil {
		return err
	}
	var ws []string
	for _, w := range widths {
		ws = append(ws, strconv.Itoa(w))
	}
	fmt.Printf("%s: %d probes at %spx, %d null; %s; %s\n", name, len(spec.Probes), strings.Join(ws, ", "), len(nulls), result.Get("browser"), result.Get("finalUrl"))
	if len(nulls) > 0 {
		os.Exit(1)
	}
	return nil
}

func measureOpen(page probe.Page, spec pagespec.Spec, widths []int) (ojson.Object, []string, error) {
	pw := page.Playwright()
	if pw == nil {
		return nil, nil, fmt.Errorf("%s: measure needs Playwright's page", page.Label())
	}
	main := pw.MainFrame()
	reload := func() error {
		_, err := pw.Reload(playwright.PageReloadOptions{WaitUntil: playwright.WaitUntilStateLoad})
		return err
	}
	// @external: `colorScheme`: the page read in that scheme (prefers-color-scheme), from its load
	if spec.ColorScheme != "" {
		schemes := map[string]*playwright.ColorScheme{"dark": playwright.ColorSchemeDark, "light": playwright.ColorSchemeLight}
		scheme, ok := schemes[spec.ColorScheme]
		if !ok {
			return nil, nil, fmt.Errorf("colorScheme %q: dark or light", spec.ColorScheme)
		}
		if err := pw.EmulateMedia(playwright.PageEmulateMediaOptions{ColorScheme: scheme}); err != nil {
			return nil, nil, err
		}
		if err := reload(); err != nil {
			return nil, nil, err
		}
	}
	cdp, err := pw.Context().NewCDPSession(pw)
	if err != nil {
		return nil, nil, err
	}
	s := session{pw: pw, cdp: cdp}
	// @external: `touch`: the page read on a touch screen, from its load
	if spec.Touch {
		if _, err := s.send("Emulation.setTouchEmulationEnabled", map[string]any{"enabled": true, "maxTouchPoints": 5}); err != nil {
			return nil, nil, err
		}
		if err := reload(); err != nil {
			return nil, nil, err
		}
		hover, err := call(main, `() => matchMedia('(hover: hover)').matches`, nil)
		if err != nil {
			return nil, nil, err
		}
		if hover == ojson.Bool(true) {
			return nil, nil, errors.New("(hover: hover) matches on the emulated touch screen")
		}
	}
	if _, err := s.send("DOM.enable", map[string]any{}); err != nil {
		return nil, nil, err
	}
	if _, err := s.send("CSS.enable", map[string]any{}); err != nil {
		return nil, nil, err
	}

	final, err := url.Parse(pw.URL())
	if err != nil {
		return nil, nil, err
	}
	path := final.EscapedPath()
	if path == "" {
		path = "/"
	}
	date := probe.Date()
	clicked := map[string]bool{}
	// @invariant: what a JavaScript object of width keys holds: integer keys in
	// ascending order, whatever order they were set in
	at := map[int]ojson.Object{}
	var nulls []string
	var used []playwright.Frame
	settle := func() error {
		_, err := call(main, settleFor, 900)
		return err
	}
	pointer := func(p pagespec.Probe, width int) error {
		to := [2]float64{0, 899}
		if p.Pointer != nil {
			to = p.Pointer(width)
		}
		return pw.Mouse().Move(to[0], to[1])
	}
	if spec.Wait > 0 {
		pw.WaitForTimeout(float64(spec.Wait))
	}
	settleMs := spec.Settle
	if settleMs == 0 {
		settleMs = 800
	}
	for _, width := range widths {
		if err := pw.SetViewportSize(width, 900); err != nil {
			return nil, nil, err
		}
		if _, err := call(main, settleFor, settleMs); err != nil {
			return nil, nil, err
		}
		got := ojson.Object{}
		for _, p := range spec.Probes {
			if p.Widths != nil && !slices.Contains(p.Widths, width) {
				continue
			}
			frame, err := s.frameOf(p)
			if err != nil {
				return nil, nil, err
			}
			// @external: a click once, at the first width the probe is read: what it opens
			// stays so as the window narrows
			clicking := (p.Click != "" || p.ClickAt != nil || len(p.Type) > 0 || p.Tap != "" || p.Hover != "" || p.HoverAt != nil) && !clicked[p.Name]
			watch := func() error {
				if clicking && p.Animations != "" {
					_, err := call(frame, watchAnimations, p.Animations)
					return err
				}
				return nil
			}
			switch {
			case p.Hover != "" || p.HoverAt != nil:
				// @external: at every width: what the pointer opens a window's resize may close
				if spec.Touch {
					return nil, nil, fmt.Errorf("%s: a hover on a touch spec", p.Name)
				}
				clicked[p.Name] = true
				if err := watch(); err != nil {
					return nil, nil, err
				}
				// @invariant: a probe's scroll before the pointer moves, so what it hovers is in
				// the window (the read's own scroll then finds it there)
				if p.Scroll != 0 {
					if _, err := call(frame, `(y) => window.scrollTo(0, y)`, p.Scroll); err != nil {
						return nil, nil, err
					}
					if err := settle(); err != nil {
						return nil, nil, err
					}
				}
				var to [2]float64
				if p.HoverAt != nil {
					to = p.HoverAt(width)
				} else {
					r, err := frame.Locator(p.Hover).First().BoundingBox()
					if err != nil {
						return nil, nil, err
					}
					if r == nil || r.Width == 0 || r.Height == 0 {
						return nil, nil, fmt.Errorf("%s: %s has no box to hover", p.Name, p.Hover)
					}
					to = [2]float64{r.X + r.Width/2, r.Y + r.Height/2}
				}
				if err := pw.Mouse().Move(to[0], to[1], playwright.MouseMoveOptions{Steps: playwright.Int(12)}); err != nil {
					return nil, nil, err
				}
				hovered := p.Hovered
				if hovered == 0 {
					hovered = 1500
				}
				pw.WaitForTimeout(float64(hovered))
				if p.Pointer != nil {
					to := p.Pointer(width)
					if err := pw.Mouse().Move(to[0], to[1], playwright.MouseMoveOptions{Steps: playwright.Int(6)}); err != nil {
						return nil, nil, err
					}
					if err := settle(); err != nil {
						return nil, nil, err
					}
				}
			case clicking && p.Tap != "":
				if !spec.Touch {
					return nil, nil, fmt.Errorf("%s: a tap on a spec that is not touch", p.Name)
				}
				clicked[p.Name] = true
				if err := watch(); err != nil {
					return nil, nil, err
				}
				target := frame.Locator(p.Tap).First()
				if _, err := target.Evaluate(`e => e.scrollIntoView({ block: 'center' })`, nil); err != nil {
					return nil, nil, err
				}
				if err := settle(); err != nil {
					return nil, nil, err
				}
				r, err := target.BoundingBox()
				if err != nil {
					return nil, nil, err
				}
				if r == nil || r.Width == 0 || r.Height == 0 {
					return nil, nil, fmt.Errorf("%s: %s has no box to tap", p.Name, p.Tap)
				}
				if _, err := s.send("Input.dispatchTouchEvent", map[string]any{"type": "touchStart", "touchPoints": []any{map[string]any{"x": r.X + r.Width/2, "y": r.Y + r.Height/2}}}); err != nil {
					return nil, nil, err
				}
				if _, err := s.send("Input.dispatchTouchEvent", map[string]any{"type": "touchEnd", "touchPoints": []any{}}); err != nil {
					return nil, nil, err
				}
				if err := settle(); err != nil {
					return nil, nil, err
				}
			case clicking && len(p.Type) > 0:
				clicked[p.Name] = true
				if err := watch(); err != nil {
					return nil, nil, err
				}
				for _, key := range p.Type {
					if err := pw.Keyboard().Press(key); err != nil {
						return nil, nil, err
					}
					pw.WaitForTimeout(150)
				}
				typed := p.Typed
				if typed == 0 {
					typed = 2500
				}
				pw.WaitForTimeout(float64(typed))
			case clicking:
				clicked[p.Name] = true
				if err := watch(); err != nil {
					return nil, nil, err
				}
				// @invariant: a point: the window clicked there (what a layer over the
				// element would take a click on it for)
				if p.ClickAt != nil {
					to := p.ClickAt(width)
					if err := pw.Mouse().Click(to[0], to[1]); err != nil {
						return nil, nil, err
					}
				} else if err := frame.Click(p.Click); err != nil {
					return nil, nil, err
				}
				// @external: the pointer off what it clicked, which it would leave hovered, to the
				// window's bottom corner; or, once what the click opened has opened, to
				// `pointer`, a point inside it, which the pointer leaving would close
				// (moved at once, it lands beyond what is still growing)
				if p.Pointer != nil {
					if err := settle(); err != nil {
						return nil, nil, err
					}
				}
				if err := pointer(p, width); err != nil {
					return nil, nil, err
				}
				if p.Clicked != 0 {
					pw.WaitForTimeout(float64(p.Clicked))
				} else if err := settle(); err != nil {
					return nil, nil, err
				}
			}
			if !slices.Contains(used, frame) {
				used = append(used, frame)
			}
			// @invariant: what the page is handed of the probe: its data, not its points
			html, err := call(frame, place, p)
			if err != nil {
				return nil, nil, err
			}
			if html == (ojson.Null{}) {
				got = got.With(ojson.Member{Key: p.Name, Value: ojson.Null{}})
				nulls = append(nulls, fmt.Sprintf("%d %s", width, p.Name))
				fmt.Printf("NULL %dpx %s: %s %s %s %s\n", width, p.Name, p.Select, p.Inject, p.Into, p.StateOn)
				continue
			}
			head, ok := html.(ojson.String)
			if !ok {
				return nil, nil, fmt.Errorf("%s: place gave %T", p.Name, html)
			}
			values, err := call(frame, read, p)
			if err != nil {
				return nil, nil, err
			}
			valuesObject, ok := values.(ojson.Object)
			if !ok {
				return nil, nil, fmt.Errorf("%s: read gave %T", p.Name, values)
			}
			entry := ojson.Object{{Key: "html", Value: ojson.String(strip(string(head)))}}.With(valuesObject...)
			if kf, ok := entry.Get("keyframes").(ojson.Object); ok {
				for _, m := range kf {
					if _, null := m.Value.(ojson.Null); null {
						nulls = append(nulls, fmt.Sprintf("%d %s @keyframes %s", width, p.Name, m.Key))
					}
				}
			}
			if clicking && p.Animations != "" {
				animations, err := call(frame, `() => { clearInterval(window.cuiWatch); return window.cuiAnimations; }`, nil)
				if err != nil {
					return nil, nil, err
				}
				entry = entry.With(ojson.Member{Key: "animations", Value: animations})
			}
			if len(p.States) > 0 {
				nodeID, err := s.stateNode(frame)
				if err != nil {
					return nil, nil, err
				}
				if id, _ := nodeID.(float64); id == 0 {
					return nil, nil, fmt.Errorf("%s: the DevTools session does not see the tagged element", p.Name)
				}
				tag, err := frame.AddStyleTag(playwright.FrameAddStyleTagOptions{Content: playwright.String(still)})
				if err != nil {
					return nil, nil, err
				}
				if _, err := tag.Evaluate(`e => e.setAttribute('data-cui-still', '')`); err != nil {
					return nil, nil, err
				}
				if _, err := call(frame, stillInShadow, nil); err != nil {
					return nil, nil, err
				}
				states := ojson.Object{}
				for _, state := range p.States {
					// @invariant: a state may be several at once: 'focus+focus-visible'
					var forced []any
					for _, c := range strings.Split(state, "+") {
						forced = append(forced, c)
					}
					if _, err := s.send("CSS.forcePseudoState", map[string]any{"nodeId": nodeID, "forcedPseudoClasses": forced}); err != nil {
						return nil, nil, err
					}
					v, err := call(frame, read, p)
					if err != nil {
						return nil, nil, err
					}
					o, ok := v.(ojson.Object)
					if !ok {
						return nil, nil, fmt.Errorf("%s: read gave %T", p.Name, v)
					}
					states = states.With(ojson.Member{Key: state, Value: o.Get("styles")})
					if _, err := s.send("CSS.forcePseudoState", map[string]any{"nodeId": nodeID, "forcedPseudoClasses": []any{}}); err != nil {
						return nil, nil, err
					}
				}
				entry = entry.With(ojson.Member{Key: "states", Value: states})
				if _, err := call(frame, unstill, nil); err != nil {
					return nil, nil, err
				}
			}
			got = got.With(ojson.Member{Key: p.Name, Value: entry})
		}
		at[width] = got
		for _, f := range used {
			if _, err := call(f, `() => document.querySelectorAll('[data-cui-injected]').forEach(e => e.remove())`, nil); err != nil {
				return nil, nil, err
			}
		}
	}
	keys := make([]int, 0, len(at))
	for w := range at {
		keys = append(keys, w)
	}
	sort.Ints(keys)
	byWidth := ojson.Object{}
	for _, w := range keys {
		byWidth = append(byWidth, ojson.Member{Key: strconv.Itoa(w), Value: at[w]})
	}
	return ojson.Object{
		{Key: "url", Value: ojson.String(spec.URL)},
		{Key: "finalUrl", Value: ojson.String(final.Scheme + "://" + final.Host + path)},
		{Key: "browser", Value: ojson.String(page.Label())},
		{Key: "date", Value: date},
		{Key: "widths", Value: byWidth},
	}, nulls, nil
}

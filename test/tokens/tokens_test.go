package tokens

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/at-least/ciderui/internal/reference"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
	asc "github.com/at-least/ciderui/reference/asc/results"
	"github.com/at-least/ciderui/reference/tracking/table"
)

const W = 1200

// @invariant: one side of a contract: an apple.com reference
// (page, probe, width, property, state) or Fitted's reason
type source struct {
	page        reference.Page
	probe       string
	width       int
	prop, state string
	fitted      string
}

func appleAt(page reference.Page, probe string, width int, prop string) source {
	return source{page: page, probe: probe, width: width, prop: prop, state: "", fitted: ""}
}

func appleIn(page reference.Page, probe string, width int, prop, state string) source {
	return source{page: page, probe: probe, width: width, prop: prop, state: state, fitted: ""}
}

func fitted(why string) source {
	return source{page: reference.Page{}, probe: "", width: 0, prop: "", state: "", fitted: why}
}

// @invariant: each color's light and dark side: [page, probe, width, property, state?]
// where apple.com shows it, or 'fitted: why', held by the contrast checks
var contract = []struct {
	name        string
	light, dark source
}{
	{"--cui-c-bg", appleAt(apple.Home, "page.body", W, "background-color"), appleAt(apple.Product, "page.body", W, "background-color")},
	{"--cui-c-bg-alt", appleAt(apple.Home, "hero.tile", W, "background-color"), appleAt(apple.Product, "section.dark", W, "background-color")},
	{"--cui-c-text", appleAt(apple.Home, "page.body", W, "color"), appleAt(apple.Home, "tile.dark-headline", W, "color")},
	{"--cui-c-text-2", appleAt(apple.Product, "stat.caption", W, "color"), fitted("fitted: apple.com's dark #6e6e73 is below 4.5:1")},
	{"--cui-c-text-3", appleAt(apple.Buy, "option", W, "border-top-color"), appleAt(apple.Product, "localnav.explore", W, "box-shadow")},
	{"--cui-c-border", appleAt(apple.Buy, "option", W, "border-top-color"), appleAt(apple.Product, "localnav.explore", W, "box-shadow")},
	{"--cui-c-accent", appleAt(apple.Home, "button", W, "background-color"), appleAt(apple.Home, "tile.dark-button", W, "background-color")},
	{"--cui-c-accent-hover", appleIn(apple.Home, "button", W, "background-color", "hover"), appleIn(apple.Home, "tile.dark-button", W, "background-color", "hover")},
	{"--cui-c-accent-active", appleIn(apple.Home, "button", W, "background-color", "active"), appleIn(apple.Home, "tile.dark-button", W, "background-color", "active")},
	{"--cui-c-on-accent", appleAt(apple.Home, "button", W, "color"), appleAt(apple.Home, "tile.dark-button", W, "color")},
	{"--cui-c-link", appleAt(apple.Home, "button.secondary", W, "color"), appleAt(apple.Home, "tile.dark-button-secondary", W, "color")},
	// @external: the neutral in a dark block: a gallery item's (li.media-gallery-item.theme-dark)
	{"--cui-c-neutral", appleAt(apple.Store, "button.neutral", W, "background-color"), appleAt(apple.Home, "gallery.button-neutral", W, "background-color")},
	{"--cui-c-neutral-hover", appleIn(apple.Store, "button.neutral", W, "background-color", "hover"), appleIn(apple.Home, "gallery.button-neutral", W, "background-color", "hover")},
	{"--cui-c-neutral-active", appleIn(apple.Store, "button.neutral", W, "background-color", "active"), appleIn(apple.Home, "gallery.button-neutral", W, "background-color", "active")},
	{"--cui-c-on-neutral", appleAt(apple.Store, "button.neutral", W, "color"), appleAt(apple.Home, "gallery.button-neutral", W, "color")},
	// @external: the global nav's: the home page's, light, and the Newsroom's, a dark nav;
	// its menu opened on a phone
	{"--cui-c-bar", appleAt(apple.Home, "nav", W, "background-color"), appleAt(apple.Newsroom, "nav", W, "background-color")},
	{"--cui-c-on-bar", appleAt(apple.Home, "nav.link", W, "color"), appleAt(apple.Newsroom, "nav.link", W, "color")},
	{"--cui-c-on-bar-hover", appleIn(apple.Home, "nav.link", W, "color", "hover"), appleIn(apple.Newsroom, "nav.link", W, "color", "hover")},
	{"--cui-c-menu", appleAt(apple.Home, "menu.content", 390, "background-color"), appleAt(apple.Newsroom, "menu.content", 390, "background-color")},
	{"--cui-c-on-menu", appleAt(apple.Home, "menu.link", 390, "color"), appleAt(apple.Newsroom, "menu.link", 390, "color")},
	// @external: apple.com's own field (the retail page's form-textbox), and in a dark
	// color scheme
	{"--cui-c-field", appleAt(apple.Retail, "field.input", W, "background-color"), appleAt(apple.RetailDark, "field.input", W, "background-color")},
	// @external: a field's states: the gift card's page's (an empty required field the
	// focus left, one disabled, the currency field's), and copies on the retail
	// page in a dark color scheme; the currency's ground read on the gift
	// card's page alone (a field is light on any page)
	{"--cui-c-field-error", appleAt(apple.Giftcard, "error", W, "background-color"), appleAt(apple.RetailDark, "field.error", W, "background-color")},
	{"--cui-c-field-disabled", appleAt(apple.Giftcard, "field.disabled", 1600, "background-color"), appleAt(apple.RetailDark, "field.disabled", W, "background-color")},
	{"--cui-c-field-prefix", appleAt(apple.Giftcard, "prefix.currency", 1600, "background-color"), appleAt(apple.Giftcard, "prefix.currency", 1600, "background-color")},
	{"--cui-c-border-disabled", appleAt(apple.Giftcard, "field.disabled", 1600, "border-top-color"), appleAt(apple.RetailDark, "field.disabled", W, "border-top-color")},
	{"--cui-c-error", appleAt(apple.Giftcard, "error.message", W, "color"), fitted("fitted: apple.com's #e30000 is below 4.5:1 on black")},
	// @external: the bag's accordion's rule, and a rule between copy blocks on iPhone 18
	// Pro's dark page
	{"--cui-c-divider", appleAt(apple.Bag, "accordion.item", W, "border-top-color"), appleAt(apple.Product, "divider", W, "border-top-color")},
	// @external: the footer: the home page's, and the dark product page's, light as well
	{"--cui-c-on-footer", appleAt(apple.Home, "footer", W, "color"), appleAt(apple.Product, "footer", W, "color")},
	{"--cui-c-on-footer-link", appleAt(apple.Home, "footer.link", W, "color"), appleAt(apple.Product, "footer.link", W, "color")},
	{"--cui-c-on-footer-title", appleAt(apple.Home, "footer.title", W, "color"), appleAt(apple.Product, "footer.title", W, "color")},
	{"--cui-c-footer-separator", appleAt(apple.Home, "footer.section-rule", 390, "background-color"), appleAt(apple.Home, "footer.section-rule", 390, "background-color")},
	// @external: the store's overlay (iPhone 18 Pro's buy page, its Show more opened): one
	// value each, read at a wide window and a phone's
	{"--cui-c-scrim", appleAt(apple.Overlay, "overlay", W, "background-color"), appleAt(apple.Overlay, "overlay", 390, "background-color")},
	// @external: tv.apple.com's video player, paused (decision 28): one value each, a
	// video dark on any page
	{"--cui-c-video", appleAt(apple.TvPlayer, "dialog", W, "background-color"), appleAt(apple.TvPlayer, "dialog", 1600, "background-color")},
	{"--cui-c-on-video", appleAt(apple.TvPlayer, "play", W, "color"), appleAt(apple.TvPlayer, "play", 1600, "color")},
	{"--cui-c-on-video-control", appleAt(apple.TvPlayer, "subtitles", W, "color"), appleAt(apple.TvPlayer, "fullscreen", W, "color")},
	{"--cui-c-on-video-time", appleAt(apple.TvPlayer, "elapsed", W, "color"), appleAt(apple.TvPlayer, "remaining", W, "color")},
	{"--cui-c-video-track", appleAt(apple.TvPlayer, "volume.range", W, "--progress-track-color"), appleAt(apple.TvPlayer, "volume.range", 1600, "--progress-track-color")},
}

// @invariant: the console's colors: light the toolkit's, dark derived (decision 18:
// not measured). A side its job's contrast finds short is fitted (decision 21):
// FITTED's value, a step of the same ramp (or the same step at another alpha),
// held by the contrast checks.
var sources = map[string]string{
	"--cui-c-console-brand-fill-hover":        "component-Button-variant-colorRole-brand-effect-hover-variant-filled-backgroundColor",
	"--cui-c-console-brand-fill-press":        "component-Button-variant-colorRole-brand-effect-press-variant-filled-backgroundColor",
	"--cui-c-console-brand-label-hover":       "component-Button-variant-colorRole-brand-effect-hover-variant-gray-color",
	"--cui-c-console-brand-label-press":       "component-Button-variant-colorRole-brand-effect-press-variant-plain-color",
	"--cui-c-console-on-accent-press":         "component-Button-variant-colorRole-brand-effect-press-variant-filled-color",
	"--cui-c-console-gray-fill-hover":         "component-Button-variant-colorRole-neutral-effect-hover-variant-gray-backgroundColor",
	"--cui-c-console-gray-fill-press":         "component-Button-variant-colorRole-neutral-effect-press-variant-gray-backgroundColor",
	"--cui-c-console-neutral-label-hover":     "component-Button-variant-colorRole-neutral-effect-hover-variant-gray-color",
	"--cui-c-console-neutral-fill-hover":      "component-Button-variant-colorRole-neutral-effect-hover-variant-filled-backgroundColor",
	"--cui-c-console-neutral-on-accent-press": "component-Button-variant-colorRole-neutral-effect-press-variant-filled-color",
	"--cui-c-console-error-fill-hover":        "component-Button-variant-colorRole-error-effect-hover-variant-filled-backgroundColor",
	"--cui-c-console-error-fill-press":        "component-Button-variant-colorRole-error-effect-press-variant-filled-backgroundColor",
	"--cui-c-console-error-label-hover":       "component-Button-variant-colorRole-error-effect-hover-variant-gray-color",
	"--cui-c-console-error-label-press":       "component-Button-variant-colorRole-error-effect-press-variant-gray-color",
	"--cui-c-console-card":                    "color-v2-bgPlaceholder",
	"--cui-c-console-field-border":            "component-TextInput-variant-colorRole-neutral-focus-off-variant-stroked-border-color",
	"--cui-c-console-brand-label-on-tint":     "component-Link-variant-colorRole-brand-effect-default-variant-inline-color",
	"--cui-c-console-dialog-stroke":           "component-Modal-variant-variant-modal-stroke",
}

type fit struct{ name, light, dark, why string }

var fittedConsole = []fit{
	{"--cui-c-console-neutral-label-secondary", "oklch(35.9% 0.012 285.8 / 0.76)", "oklch(87.9% 0.015 286.1 / 0.63)", "the toolkit's 0.60 is 3.4:1 on white, 0.74 4.4:1 on a chip over a card; its dark 0.60 short of 4.5:1 on a success or warning banner's tint over a card (crashcart's demo has one)"},
	{"--cui-c-console-brand-accent-primary", "", "oklch(51.2% 0.203 260.5)", "white on the derived blue-600 is 2.7:1"},
	{"--cui-c-console-brand-fill-hover", "oklch(49.7% 0.209 260.5)", "oklch(43.6% 0.171 260.6)", "white on the toolkit's lighter blue is 3.9:1"},
	{"--cui-c-console-brand-fill-press", "", "oklch(43.6% 0.171 260.6)", "white on the derived blue-600 is 2.7:1"},
	{"--cui-c-console-brand-label-hover", "oklch(49.7% 0.209 260.5)", "oklch(68.3% 0.167 256.6)", "the toolkit's lighter blue is 3.5:1 on white, its derived dark 3.5:1 on black"},
	{"--cui-c-console-error-accent-primary", "oklch(49.5% 0.190 31.5)", "oklch(50.0% 0.186 31.4)", "white on the toolkit's red is 3.3:1"},
	{"--cui-c-console-field-border", "oklch(35.9% 0.012 285.8 / 0.58)", "oklch(87.9% 0.015 286.1 / 0.42)", "a field's edge, at the toolkit's 0.18, is 1.4:1 on white; a control's edge is held to 3:1"},
	{"--cui-c-console-brand-label-on-tint", "oklch(49.7% 0.209 260.5)", "oklch(75.6% 0.120 260.5)", "a link in a banner's sentence: the toolkit's blue 3.9:1 on the info tint over a card, its derived dark blue-600 4.47:1 on the warning tint over a card (crashcart's unsymbolicated notice)"},
	{"--cui-c-console-info-label-primary", "oklch(49.7% 0.209 260.5)", "oklch(68.3% 0.167 256.6)", "the toolkit's info blue is 4.0:1 on white; its derived dark blue-500 short of 4.5:1 on its tint over a card"},
	{"--cui-c-console-success-label-primary", "oklch(38.5% 0.094 150.1)", "", "the toolkit's green is 2.2:1 on white, green-800 short of 4.5:1 on the gray fill over a card"},
	{"--cui-c-console-warning-label-primary", "oklch(48.6% 0.112 61.8)", "", "the toolkit's yellow is 1.6:1 on white"},
	{"--cui-c-console-error-label-primary", "oklch(49.5% 0.190 31.5)", "oklch(70.6% 0.184 29.4)", "the toolkit's red is 3.6:1 on white, red-600 4.2:1 on a gray button; its derived dark red-500 short of 4.5:1 on its tint and a gray button over a card (crashcart: 4.09:1, 3.92:1)"},
	{"--cui-c-console-error-fill-hover", "oklch(37.8% 0.147 29.3)", "oklch(40.9% 0.152 29.5)", "white on the toolkit's lighter red is 3.0:1"},
	{"--cui-c-console-error-fill-press", "oklch(37.8% 0.147 29.3)", "oklch(40.9% 0.152 29.5)", "white on the toolkit's red is 3.5:1"},
	{"--cui-c-console-error-label-hover", "oklch(49.5% 0.190 31.5)", "oklch(70.6% 0.184 29.4)", "the toolkit's lighter red is 2.9:1 on white, its derived dark 4.3:1 on a gray button"},
	{"--cui-c-console-error-label-press", "oklch(49.5% 0.190 31.5)", "oklch(70.6% 0.184 29.4)", "the toolkit's red is 3.8:1 on white; its derived dark red-500 short of 4.5:1 on a gray button over a card"},
}

// @external: every named color CSS has; not transparent, currentColor or inherit, which
// are no color of a stylesheet's own
var named = regexp.MustCompile(`(?i)\b(aliceblue|antiquewhite|aqua|aquamarine|azure|beige|bisque|black|blanchedalmond|blue|blueviolet|brown|burlywood|cadetblue|chartreuse|chocolate|coral|cornflowerblue|cornsilk|crimson|cyan|darkblue|darkcyan|darkgoldenrod|darkgray|darkgreen|darkgrey|darkkhaki|darkmagenta|darkolivegreen|darkorange|darkorchid|darkred|darksalmon|darkseagreen|darkslateblue|darkslategray|darkslategrey|darkturquoise|darkviolet|deeppink|deepskyblue|dimgray|dimgrey|dodgerblue|firebrick|floralwhite|forestgreen|fuchsia|gainsboro|ghostwhite|gold|goldenrod|gray|green|greenyellow|grey|honeydew|hotpink|indianred|indigo|ivory|khaki|lavender|lavenderblush|lawngreen|lemonchiffon|lightblue|lightcoral|lightcyan|lightgoldenrodyellow|lightgray|lightgreen|lightgrey|lightpink|lightsalmon|lightseagreen|lightskyblue|lightslategray|lightslategrey|lightsteelblue|lightyellow|lime|limegreen|linen|magenta|maroon|mediumaquamarine|mediumblue|mediumorchid|mediumpurple|mediumseagreen|mediumslateblue|mediumspringgreen|mediumturquoise|mediumvioletred|midnightblue|mintcream|mistyrose|moccasin|navajowhite|navy|oldlace|olive|olivedrab|orange|orangered|orchid|palegoldenrod|palegreen|paleturquoise|palevioletred|papayawhip|peachpuff|peru|pink|plum|powderblue|purple|rebeccapurple|red|rosybrown|royalblue|saddlebrown|salmon|sandybrown|seagreen|seashell|sienna|silver|skyblue|slateblue|slategray|slategrey|snow|springgreen|steelblue|tan|teal|thistle|tomato|turquoise|violet|wheat|white|whitesmoke|yellow|yellowgreen)\b`)

var (
	varDecl     = regexp.MustCompile(`(?m)^\s*(--cui-[a-z0-9-]+):\s*([^;]+);`)
	hexColor    = regexp.MustCompile(`(?i)^#([0-9a-f]{6})$`)
	oklchColor  = regexp.MustCompile(`^oklch\(([\d.]+)% ([\d.]+) ([\d.]+)(?: / ([\d.]+))?\)$`)
	rgbColor    = regexp.MustCompile(`^rgb\((\d+) (\d+) (\d+) / ([\d.]+)\)$`)
	rgbaIn      = regexp.MustCompile(`rgba?\([^)]*\)`)
	zeroAlpha   = regexp.MustCompile(` 0\.0\)$`)
	rampOf      = regexp.MustCompile(`-palette-(?:light|dark)-([a-zA-Z]+)-`)
	jsFloat     = regexp.MustCompile(`^\s*[+-]?(\d+\.?\d*|\.\d+)([eE][+-]?\d+)?`)
	comment     = regexp.MustCompile(`(?s)/\*.*?\*/`)
	setSlot     = regexp.MustCompile(`--cui-set-[\w-]+`)
	declValue   = regexp.MustCompile(`:\s*([^;{}]+)[;}]`)
	customProp  = regexp.MustCompile(`--[\w-]+`)
	colorFunc   = regexp.MustCompile(`(?i)#[0-9a-f]{3,8}\b|\b(rgba?|hsla?|hwb|lab|lch|oklab|oklch|color)\(`)
	important   = regexp.MustCompile(`(?i)!\s*important`)
	reducedMove = regexp.MustCompile(`(?s)@media \(prefers-reduced-motion: reduce\) \{.*?\n\}`)
	colorRead   = regexp.MustCompile(`var\((--cui-c-[\w-]+)`)
	videoName   = regexp.MustCompile(`^--cui-c-(on-)?video`)
)

func jsRound(x float64) float64 { return math.Floor(x + 0.5) }

func parseFloat(s string) float64 {
	m := jsFloat.FindString(s)
	if m == "" {
		return math.NaN()
	}
	v, err := strconv.ParseFloat(strings.TrimSpace(m), 64)
	if err != nil {
		panic(err)
	}
	return v
}

func number(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}

// @external: CSS Color 4's matrices (OKLab to linear sRGB)
func oklchToSrgb(l, c, h float64) [3]float64 {
	a, b := c*math.Cos(h*math.Pi/180), c*math.Sin(h*math.Pi/180)
	l_ := math.Pow(l+0.3963377774*a+0.2158037573*b, 3)
	m_ := math.Pow(l-0.1055613458*a-0.0638541728*b, 3)
	s_ := math.Pow(l-0.0894841775*a-1.2914855480*b, 3)
	lin := [3]float64{4.0767416621*l_ - 3.3077115913*m_ + 0.2309699292*s_, -1.2684380046*l_ + 2.6097574011*m_ - 0.3413193965*s_, -0.0041960863*l_ - 0.7034186147*m_ + 1.7076147010*s_}
	var out [3]float64
	for i, v := range lin {
		v = math.Min(1, math.Max(0, v))
		if v <= 0.0031308 {
			out[i] = 12.92 * v * 255
		} else {
			out[i] = (1.055*math.Pow(v, 1/2.4) - 0.055) * 255
		}
	}
	return out
}

func TestTokens(t *testing.T) {
	read := func(rel string) string {
		b, err := os.ReadFile(filepath.Join(build.Root, rel))
		if err != nil {
			t.Fatal(err)
		}
		return string(b)
	}

	css := read("tokens.css")
	var tbTokens map[string]string
	var tbKeys []string
	tbTokens = map[string]string{}
	for _, token := range asc.Tokens.Tokens {
		tbKeys = append(tbKeys, token.Name)
		tbTokens[token.Name] = token.Value
	}
	tb := func(k string) string {
		v, ok := tbTokens[k]
		if !ok {
			return "undefined"
		}
		return v
	}

	// @invariant: the first block, :root's: the table a language's block overrides is not this one
	end := strings.Index(css, "\n}\n")
	if end < 0 {
		end = len(css) - 1
	}
	var varNames []string
	vars := map[string]string{}
	for _, m := range varDecl.FindAllStringSubmatch(css[:end], -1) {
		if _, seen := vars[m[1]]; !seen {
			varNames = append(varNames, m[1])
		}
		vars[m[1]] = strings.TrimSpace(m[2])
	}
	varOf := func(name string) string {
		v, ok := vars[name]
		if !ok {
			return "undefined"
		}
		return v
	}

	sides := func(name string) [2]string {
		v, ok := vars[name]
		if !ok {
			t.Fatalf("tokens.css has no %s", name)
		}
		if !strings.HasPrefix(v, "light-dark(") {
			return [2]string{v, v}
		}
		inner := v[len("light-dark(") : len(v)-1]
		depth := 0
		for i := 0; i < len(inner); i++ {
			switch inner[i] {
			case '(':
				depth++
			case ')':
				depth--
			case ',':
				if depth == 0 {
					return [2]string{strings.TrimSpace(inner[:i]), strings.TrimSpace(inner[i+1:])}
				}
			}
		}
		t.Fatalf("%s: no two sides in %s", name, v)
		return [2]string{}
	}
	channels := func(value string) [4]float64 {
		if h := hexColor.FindStringSubmatch(value); h != nil {
			n, err := strconv.ParseInt(h[1], 16, 64)
			if err != nil {
				t.Fatal(err)
			}
			return [4]float64{float64(n >> 16), float64((n >> 8) & 255), float64(n & 255), 1}
		}
		if o := oklchColor.FindStringSubmatch(value); o != nil {
			rgb := oklchToSrgb(number(o[1])/100, number(o[2]), number(o[3]))
			alpha := 1.0
			if o[4] != "" {
				alpha = number(o[4])
			}
			return [4]float64{rgb[0], rgb[1], rgb[2], alpha}
		}
		c := rgbColor.FindStringSubmatch(value)
		if c == nil {
			t.Fatalf("not a #rrggbb, rgb(r g b / a) or oklch(l%% c h / a) color: %s", value)
		}
		return [4]float64{number(c[1]), number(c[2]), number(c[3]), number(c[4])}
	}
	computedRgb := func(value string) string {
		c := channels(value)
		if c[3] == 1 {
			return fmt.Sprintf("rgb(%s, %s, %s)", check.Number(c[0]), check.Number(c[1]), check.Number(c[2]))
		}
		return fmt.Sprintf("rgba(%s, %s, %s, %s)", check.Number(c[0]), check.Number(c[1]), check.Number(c[2]), check.Number(c[3]))
	}
	ref := func(page reference.Page, probe string, width int, prop, state string) string {
		v, err := page.Value(probe, width, prop, state)
		if err != nil {
			t.Fatal(err)
		}
		return v
	}

	var colors []string
	for _, k := range varNames {
		if strings.HasPrefix(k, "--cui-c-") && !strings.HasPrefix(k, "--cui-c-console-") {
			colors = append(colors, k)
		}
	}
	var contractNames []string
	for _, c := range contract {
		contractNames = append(contractNames, c.name)
	}
	check.Equal(t, "the contract's colors are tokens.css's", strings.Join(colors, " "), strings.Join(contractNames, " "))

	for _, c := range contract {
		s := sides(c.name)
		for _, side := range []struct {
			side, value string
			source      source
		}{{"light", s[0], c.light}, {"dark", s[1], c.dark}} {
			src := side.source
			if src.fitted != "" {
				continue
			}
			// @external: a ring drawn as a box-shadow reads as its color first
			want := rgbaIn.FindString(ref(src.page, src.probe, src.width, src.prop, src.state))
			if want == "" {
				t.Fatalf("%s %s: no rgb() in apple.com's value", c.name, side.side)
			}
			state := ""
			if src.state != "" {
				state = ":" + src.state
			}
			check.Equal(t, fmt.Sprintf("%s %s is apple.com's %s:%s%s %s", c.name, side.side, src.page.Name, src.probe, state, src.prop), computedRgb(side.value), want)
		}
	}

	sourceOf := func(name string) string {
		if s, ok := sources[name]; ok {
			return "--tb-" + s
		}
		return "--tb-color-role-" + strings.Replace(name[len("--cui-c-console-"):], "on-accent", "onAccent", 1)
	}
	type lchValue struct {
		l, c, h  float64
		alpha    string
		hasAlpha bool
		opaque   string
	}
	lch := func(v string) lchValue {
		m := oklchColor.FindStringSubmatch(v)
		if m == nil {
			t.Fatalf("not an oklch() color: %s", v)
		}
		return lchValue{l: number(m[1]), c: number(m[2]), h: number(m[3]), alpha: m[4], hasAlpha: m[4] != "", opaque: fmt.Sprintf("oklch(%s%% %s %s)", m[1], m[2], m[3])}
	}
	type step struct{ key, value string }
	steps := func(side string) []step {
		var out []step
		for _, k := range tbKeys {
			if strings.HasPrefix(k, "--tb-color-palette-"+side+"-") {
				out = append(out, step{k, tbTokens[k]})
			}
		}
		return out
	}
	sameColor := func(a, b string) bool {
		return zeroAlpha.ReplaceAllString(a, " 0)") == zeroAlpha.ReplaceAllString(b, " 0)")
	}
	rampName := func(k string) string {
		m := rampOf.FindStringSubmatch(k)
		if m == nil {
			t.Fatalf("no ramp in %s", k)
		}
		return m[1]
	}
	stepOf := func(value, side string) step {
		v := lch(value)
		var exact []step
		for _, s := range steps(side) {
			if sameColor(s.value, v.opaque) {
				exact = append(exact, s)
			}
		}
		if len(exact) > 0 {
			var ramps []string
			for _, s := range exact {
				if r := rampName(s.key); !contains(ramps, r) {
					ramps = append(ramps, r)
				}
			}
			if len(ramps) != 1 {
				t.Fatalf("%s is a step of more than one ramp: %s", v.opaque, strings.Join(ramps, ", "))
			}
			return exact[0]
		}
		turn := func(a, b float64) float64 { return math.Min(math.Abs(a-b), 360-math.Abs(a-b)) }
		at500 := func(r string) lchValue { return lch(tbTokens["--tb-color-palette-"+side+"-"+r+"-500"]) }
		gray := v.c < 0.05
		var ramps []string
		for _, s := range steps(side) {
			if r := rampName(s.key); !contains(ramps, r) && (at500(r).c < 0.05) == gray {
				ramps = append(ramps, r)
			}
		}
		if len(ramps) == 0 {
			t.Fatalf("Reduce of empty array with no initial value")
		}
		ramp := ramps[0]
		for _, b := range ramps[1:] {
			if turn(at500(b).h, v.h) < turn(at500(ramp).h, v.h) {
				ramp = b
			}
		}
		var best step
		found := false
		for _, s := range steps(side) {
			if rampName(s.key) != ramp {
				continue
			}
			if !found || math.Abs(lch(s.value).l-v.l) < math.Abs(lch(best.value).l-v.l) {
				best, found = s, true
			}
		}
		return best
	}
	withAlpha := func(opaque string, v lchValue) string {
		if !v.hasAlpha {
			return opaque
		}
		return strings.TrimSuffix(opaque, ")") + " / " + v.alpha + ")"
	}
	darkOf := func(name, light string) string {
		v := lch(light)
		// @invariant: a label on a brand or error accent keeps its white, the accent a dark hue on
		// either side; the neutral's accent is black, white in dark, so its label turns with it
		if strings.Contains(name, "on-accent") && !strings.HasPrefix(name, "--cui-c-console-neutral-") {
			return light
		}
		if v.l == 0 && v.c == 0 {
			return withAlpha("oklch(100.0% 0.000 0)", v)
		}
		if v.l == 100 && v.c == 0 {
			return withAlpha("oklch(0.0% 0.000 0)", v)
		}
		return withAlpha(tb(strings.Replace(stepOf(light, "light").key, "-palette-light-", "-palette-dark-", 1)), v)
	}
	fitOf := map[string]fit{}
	for _, f := range fittedConsole {
		fitOf[f.name] = f
	}
	for _, name := range varNames {
		if !strings.HasPrefix(name, "--cui-c-console-") {
			continue
		}
		source := sourceOf(name)
		if _, ok := tbTokens[source]; !ok {
			t.Fatalf("%s: the toolkit has no %s", name, source)
		}
		s := sides(name)
		f := fitOf[name]
		derived := darkOf(name, tbTokens[source])
		if f.light != "" {
			check.Equal(t, fmt.Sprintf("%s light, fitted (%s), is a step of %s's ramp", name, f.why, source), rampName(stepOf(f.light, "light").key), rampName(stepOf(tbTokens[source], "light").key))
			check.Equal(t, name+" light is its fitted value", s[0], f.light)
		} else {
			check.Equal(t, fmt.Sprintf("%s light is App Store Connect's %s", name, source), s[0], tbTokens[source])
		}
		if f.dark != "" {
			check.Equal(t, fmt.Sprintf("%s dark, fitted (%s), is a step of its derived step's ramp", name, f.why), rampName(stepOf(f.dark, "dark").key), rampName(stepOf(derived, "dark").key))
			check.Equal(t, name+" dark is its fitted value", s[1], f.dark)
		} else {
			check.Equal(t, fmt.Sprintf("%s dark is %s's step of the dark palette", name, source), s[1], derived)
		}
	}
	for _, f := range fittedConsole {
		if _, ok := vars[f.name]; !ok {
			t.Fatalf("FITTED names %s, which tokens.css has not", f.name)
		}
	}

	// @external: WCAG 2 contrast. A translucent color is what it draws over the color
	// under it.
	over := func(layers ...string) string {
		under := layers[len(layers)-1]
		for i := len(layers) - 2; i >= 0; i-- {
			top, u := channels(layers[i]), channels(under)
			mix := func(tc, uc float64) float64 { return jsRound(tc*top[3] + uc*(1-top[3])) }
			under = fmt.Sprintf("#%02x%02x%02x", int(mix(top[0], u[0])), int(mix(top[1], u[1])), int(mix(top[2], u[2])))
		}
		return under
	}
	lum := func(value string) float64 {
		c := channels(value)
		weights := [3]float64{0.2126, 0.7152, 0.0722}
		s := 0.0
		for i := 0; i < 3; i++ {
			v := c[i] / 255
			if v <= 0.04045 {
				v = v / 12.92
			} else {
				v = math.Pow((v+0.055)/1.055, 2.4)
			}
			s += v * weights[i]
		}
		return s
	}
	ratio := func(a, b string) float64 {
		x, y := lum(a), lum(b)
		if y > x {
			x, y = y, x
		}
		return (x + 0.05) / (y + 0.05)
	}
	atLeast := func(name string, ok bool) {
		check.Equal(t, name, strconv.FormatBool(ok), "true")
	}
	sideNames := [2]string{"light", "dark"}

	type minimum struct {
		fg      string
		grounds []string
		min     float64
	}
	for _, m := range []minimum{
		{"--cui-c-text", []string{"--cui-c-bg", "--cui-c-bg-alt"}, 4.5},
		{"--cui-c-text-2", []string{"--cui-c-bg", "--cui-c-bg-alt"}, 4.5},
		{"--cui-c-link", []string{"--cui-c-bg", "--cui-c-bg-alt"}, 4.5},
		{"--cui-c-text-3", []string{"--cui-c-bg", "--cui-c-bg-alt"}, 3},
		// @invariant: a field's message on the page
		{"--cui-c-error", []string{"--cui-c-bg", "--cui-c-bg-alt"}, 4.5},
		{"--cui-c-border", []string{"--cui-c-bg", "--cui-c-bg-alt"}, 3},
		// @invariant: the player's text and controls on its video
		{"--cui-c-on-video-control", []string{"--cui-c-video"}, 4.5},
		{"--cui-c-on-video-time", []string{"--cui-c-video"}, 4.5},
	} {
		for _, bg := range m.grounds {
			for i := 0; i < 2; i++ {
				r := ratio(sides(m.fg)[i], sides(bg)[i])
				atLeast(fmt.Sprintf("%s on %s, %s, at least %s:1", m.fg, bg, sideNames[i], check.Number(m.min)), r >= m.min)
			}
		}
	}

	// @trap: each label on the fills it is drawn on (top first; [] the bare base), each
	// over both bases: pairs on a card were missed here and found by a site's
	// axe-core run
	const page = "--cui-c-console-neutral-background-primary"
	const card = "--cui-c-console-card"
	bases := [][]string{{page}, {card, page}}
	onFills := []string{"--cui-c-console-brand-accent-primary", "--cui-c-console-brand-fill-hover", "--cui-c-console-error-accent-primary", "--cui-c-console-error-fill-hover"}
	grayFills := []string{"--cui-c-console-neutral-fill-tertiary", "--cui-c-console-gray-fill-hover", "--cui-c-console-gray-fill-press"}
	roles := []string{"info", "success", "warning", "error"}
	each := func(names []string) [][]string {
		var out [][]string
		for _, n := range names {
			out = append(out, []string{n})
		}
		return out
	}
	cat := func(lists ...[][]string) [][]string {
		var out [][]string
		for _, l := range lists {
			out = append(out, l...)
		}
		return out
	}
	// @invariant: a banner's tints: the neutral's and each role's
	tints := [][]string{{"--cui-c-console-neutral-fill-quaternary"}}
	for _, role := range roles {
		tints = append(tints, []string{"--cui-c-console-" + role + "-fill-tertiary"})
	}
	type consoleMinimum struct {
		fg    string
		fills [][]string
		min   float64
	}
	consoleMin := []consoleMinimum{
		{"--cui-c-console-neutral-label-primary", cat([][]string{{}, {"--cui-c-console-neutral-accent-quaternary"}}, each(grayFills), tints), 4.5},
		// @invariant: the secondary label: a current menu item's, a banner's neutral icon, a
		// field's placeholder
		{"--cui-c-console-neutral-label-secondary", cat([][]string{{}}, each(grayFills), tints), 4.5},
		// @invariant: a button's label on its fill, at rest, hovered and pressed
		{"--cui-c-console-brand-label-on-accent-primary", cat(each(onFills), [][]string{{"--cui-c-console-brand-fill-press"}, {"--cui-c-console-error-fill-press"}}), 4.5},
		{"--cui-c-console-on-accent-press", [][]string{{"--cui-c-console-brand-fill-press"}, {"--cui-c-console-error-fill-press"}}, 4.5},
		{"--cui-c-console-neutral-label-hover", [][]string{{"--cui-c-console-gray-fill-hover"}}, 4.5},
		// @invariant: a picked segment's label on its black (white in dark), at rest, hovered and pressed
		{"--cui-c-console-neutral-label-on-accent-primary", [][]string{{"--cui-c-console-neutral-accent-primary"}, {"--cui-c-console-neutral-fill-hover"}}, 4.5},
		{"--cui-c-console-neutral-on-accent-press", [][]string{{"--cui-c-console-neutral-accent-primary"}}, 4.5},
	}
	// @invariant: a badge's label on the base, on its own tint and on the gray fill
	for _, role := range roles {
		consoleMin = append(consoleMin, consoleMinimum{"--cui-c-console-" + role + "-label-primary", [][]string{{}, {"--cui-c-console-" + role + "-accent-quaternary"}, {"--cui-c-console-neutral-fill-tertiary"}}, 4.5})
	}
	// @invariant: a banner's icon on its tint
	for _, role := range roles {
		consoleMin = append(consoleMin, consoleMinimum{"--cui-c-console-" + role + "-label-primary", [][]string{{"--cui-c-console-" + role + "-fill-tertiary"}}, 4.5})
	}
	// @invariant: a plain button's, a link's
	for _, fg := range []string{"--cui-c-console-brand-label-primary", "--cui-c-console-brand-label-hover", "--cui-c-console-brand-label-press"} {
		consoleMin = append(consoleMin, consoleMinimum{fg, [][]string{{}}, 4.5})
	}
	// @invariant: a link in a banner's sentence, on the banner's tint
	consoleMin = append(consoleMin, consoleMinimum{"--cui-c-console-brand-label-on-tint", tints, 4.5})
	// @invariant: a gray button's red label, on the base and on its gray
	for _, fg := range []string{"--cui-c-console-error-label-primary", "--cui-c-console-error-label-hover", "--cui-c-console-error-label-press"} {
		consoleMin = append(consoleMin, consoleMinimum{fg, cat([][]string{{}}, each(grayFills)), 4.5})
	}
	// @invariant: a field's edge and its red edge in error
	for _, fg := range []string{"--cui-c-console-field-border", "--cui-c-console-error-label-primary"} {
		consoleMin = append(consoleMin, consoleMinimum{fg, [][]string{{}}, 3})
	}
	for _, m := range consoleMin {
		for _, fill := range m.fills {
			for _, base := range bases {
				layers := append(append([]string{}, fill...), base...)
				for i := 0; i < 2; i++ {
					var drawn []string
					for _, l := range layers {
						drawn = append(drawn, sides(l)[i])
					}
					ground := over(append(drawn, [2]string{"#ffffff", "#000000"}[i])...)
					atLeast(fmt.Sprintf("%s on %s, %s, at least %s:1", m.fg, strings.Join(layers, " over "), sideNames[i], check.Number(m.min)), ratio(over(sides(m.fg)[i], ground), ground) >= m.min)
				}
			}
		}
	}
	for _, p := range []struct {
		on    string
		fills []string
	}{
		{"--cui-c-on-accent", []string{"--cui-c-accent", "--cui-c-accent-hover", "--cui-c-accent-active"}},
		{"--cui-c-on-neutral", []string{"--cui-c-neutral", "--cui-c-neutral-hover", "--cui-c-neutral-active"}},
		{"--cui-c-on-menu", []string{"--cui-c-menu"}},
		{"--cui-c-on-bar-hover", []string{"--cui-c-menu", "--cui-c-bg-alt"}},
		{"--cui-c-text-2", []string{"--cui-c-menu"}},
	} {
		for _, fill := range p.fills {
			for i := 0; i < 2; i++ {
				atLeast(fmt.Sprintf("%s on %s, %s, at least 4.5:1", p.on, fill, sideNames[i]), ratio(sides(p.on)[i], sides(fill)[i]) >= 4.5)
			}
		}
	}
	// @invariant: the nav's text on its bar, the bar over the page's grounds (what scrolls
	// under it is the page's)
	for _, p := range [][2]string{{"--cui-c-on-bar", "--cui-c-bar"}, {"--cui-c-on-bar-hover", "--cui-c-bar"}} {
		on, bar := p[0], p[1]
		for _, bg := range []string{"--cui-c-bg", "--cui-c-bg-alt"} {
			for i := 0; i < 2; i++ {
				ground := over(sides(bar)[i], sides(bg)[i])
				atLeast(fmt.Sprintf("%s on %s over %s, %s, at least 4.5:1", on, bar, bg, sideNames[i]), ratio(over(sides(on)[i], ground), ground) >= 4.5)
			}
		}
	}

	// @invariant: a field's text, its value and its label, on its ground over the field's
	// own (the page's ground, light: a field keeps its light look); in error,
	// its label on the error's ground
	for _, fg := range []string{"--cui-c-text", "--cui-c-text-2"} {
		ground := over(sides("--cui-c-field")[0], sides("--cui-c-bg")[0])
		atLeast(fg+" on --cui-c-field over --cui-c-bg, light, at least 4.5:1", ratio(sides(fg)[0], ground) >= 4.5)
	}
	atLeast("--cui-c-error on --cui-c-field-error, light, at least 4.5:1", ratio(sides("--cui-c-error")[0], sides("--cui-c-field-error")[0]) >= 4.5)

	// @invariant: the footer's text on its ground, the alternate (light: the footer keeps
	// its light look)
	for _, fg := range []string{"--cui-c-on-footer", "--cui-c-on-footer-link", "--cui-c-on-footer-title"} {
		ground := sides("--cui-c-bg-alt")[0]
		atLeast(fg+" on --cui-c-bg-alt, light, at least 4.5:1", ratio(over(sides(fg)[0], ground), ground) >= 4.5)
	}

	var tableKeys []string
	tableEm := map[string]float64{}
	for _, r := range table.CUITracking {
		tableKeys = append(tableKeys, strconv.Itoa(r.Size))
		tableEm[strconv.Itoa(r.Size)] = r.Em
	}
	var tracking []string
	for _, k := range varNames {
		if strings.HasPrefix(k, "--cui-tracking-") {
			tracking = append(tracking, k)
		}
	}
	// @invariant: the table is what was measured; a size no stylesheet reads has no token
	sheets := []string{"theme.css", "base.css", "type.css", "index.css", "site.css", "console.css"}
	sheets = append(sheets, nodeRecursiveCSS(t, filepath.Join(build.Root, "components"), "components")...)
	var texts []string
	uncommented := map[string]string{}
	for _, f := range sheets {
		text := read(f)
		texts = append(texts, text)
		uncommented[f] = comment.ReplaceAllString(text, "")
	}
	readAll := strings.Join(texts, "\n")
	var sizes, measured []string
	for _, k := range tracking {
		sizes = append(sizes, k[15:])
	}
	for _, size := range tableKeys {
		if strings.Contains(readAll, "--cui-tracking-"+size+")") {
			measured = append(measured, size)
		}
	}
	check.Equal(t, "a letter-spacing for each size of the table a stylesheet reads", strings.Join(sizes, " "), strings.Join(measured, " "))
	// @invariant: a part's removal leaves no slot behind
	var unread []string
	for _, name := range varNames {
		all := readAll + css
		if !strings.Contains(all, "var("+name+",") && !strings.Contains(all, "var("+name+")") {
			unread = append(unread, name)
		}
	}
	check.Equal(t, "every token a stylesheet reads", strings.Join(unread, " "), "")
	for _, k := range tracking {
		want := "undefined"
		if v, ok := tableEm[k[15:]]; ok {
			want = check.Number(v)
		}
		check.Equal(t, k, check.Number(parseFloat(vars[k])), want)
	}
	setSlots := func(files ...string) string {
		var slots []string
		for _, f := range files {
			for _, s := range setSlot.FindAllString(uncommented[f], -1) {
				if !contains(slots, s) {
					slots = append(slots, s)
				}
			}
		}
		sort.Strings(slots)
		return strings.Join(slots, " ")
	}
	slotsRead := setSlots("theme.css", "base.css")
	for _, sheet := range []string{"site.css", "console.css"} {
		check.Equal(t, sheet+" declares the --cui-set-* slots theme.css and base.css read, no more", setSlots(sheet), slotsRead)
	}

	check.Equal(t, "--cui-font", varOf("--cui-font"), `-apple-system, BlinkMacSystemFont, Inter, "Helvetica Neue", Helvetica, Arial, "PingFang TC", "Noto Sans TC", "Noto Sans CJK TC", "Microsoft JhengHei", sans-serif`)
	check.Equal(t, "--cui-font-mono", varOf("--cui-font-mono"), `ui-monospace, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace`)
	check.Equal(t, "--cui-focus-ring", varOf("--cui-focus-ring"), "2px solid var(--cui-c-accent)")
	check.Equal(t, "the focus ring is apple.com's width", ref(apple.Home, "button", W, "outline-width", "focus-visible"), "2px")
	check.Equal(t, "the focus ring is apple.com's color", computedRgb(sides("--cui-c-accent")[0]), ref(apple.Home, "button", W, "outline-color", "focus-visible"))
	check.Equal(t, "--cui-focus-ring-offset is apple.com's", varOf("--cui-focus-ring-offset"), ref(apple.Home, "button", W, "outline-offset", "focus-visible"))

	check.Equal(t, "--cui-console-shadow-sidebar is App Store Connect's --tb-component-Sidebar-variant-variant-opaque-boxShadow", varOf("--cui-console-shadow-sidebar"), tb("--tb-component-Sidebar-variant-variant-opaque-boxShadow"))
	check.Equal(t, "--cui-console-shadow-dialog is App Store Connect's --tb-component-Modal-variant-variant-modal-boxShadow", varOf("--cui-console-shadow-dialog"), tb("--tb-component-Modal-variant-variant-modal-boxShadow"))
	check.Equal(t, "--cui-console-shadow-popover is App Store Connect's --tb-component-Popover-variant-variant-stroked-boxShadow", varOf("--cui-console-shadow-popover"), tb("--tb-component-Popover-variant-variant-stroked-boxShadow"))
	check.Equal(t, "--cui-console-shadow-switch-knob is App Store Connect's --tb-component-Toggle-anatomy-knob-boxShadow", varOf("--cui-console-shadow-switch-knob"), tb("--tb-component-Toggle-anatomy-knob-boxShadow"))

	check.Equal(t, "--cui-video-scrim is tv.apple.com's tv-player:scrim.gradient background-color", varOf("--cui-video-scrim"), ref(apple.TvPlayer, "scrim.gradient", W, "background-color", ""))
	// @trap: golden: muse's subtitles, white under a soft shadow with no box (the owner, 2026-11-05, from Apple's view); tv.apple.com's
	// cues are its <video>'s own text track, none read. Two layers: a tight one that edges each letter on a white picture, a wide
	// faint one that darkens round the line without a box's edge
	check.Equal(t, "--cui-video-cue-shadow is muse's golden", varOf("--cui-video-cue-shadow"), "0 0.05em 0.1em rgb(0 0 0 / 0.6), 0 0 0.5em rgb(0 0 0 / 0.4)")
	// @trap: golden: muse's intro and end (the owner, 2026-11-08: 「修播放器還沒有開始播放的預設狀態是暗的」): tv.apple.com's scrim is
	// a flat 40% dim over the whole cover (read 2026-11-13). Dark only behind the words in the bottom-left corner, the cue's
	// shadow under them
	check.Equal(t, "--cui-video-title-scrim is muse's golden", varOf("--cui-video-title-scrim"), "radial-gradient(ellipse 75% 60% at 0% 100%, rgb(0 0 0 / 0.6), rgb(0 0 0 / 0.35) 40%, rgb(0 0 0 / 0))")
	// @trap: golden: muse's scrubber preview (the owner, 2026-11-05: a thumbnail over the time under the pointer); tv.apple.com's
	// trailers scrub with no preview. A 1px hairline edges it over a dark picture, a soft drop lifts it off a bright one
	check.Equal(t, "--cui-video-preview-shadow is muse's golden", varOf("--cui-video-preview-shadow"), "0 0 0 1px rgb(255 255 255 / 0.15), 0 4px 12px rgb(0 0 0 / 0.4)")
	// @trap: golden: muse's full-screen controls (回報問題, 下一個要看), a capsule of words on the picture; tv.apple.com's bar has none.
	// A faint white fill that reads on a dark or a bright picture, a little more hovered or pressed
	check.Equal(t, "--cui-video-control-fill is muse's golden", varOf("--cui-video-control-fill"), "rgb(255 255 255 / 0.16)")
	check.Equal(t, "--cui-video-control-fill-hover is muse's golden", varOf("--cui-video-control-fill-hover"), "rgb(255 255 255 / 0.24)")
	check.Equal(t, "--cui-video-thumb-shadow is tv.apple.com's tv-player:volume.range --volume-thumb-box-shadow", varOf("--cui-video-thumb-shadow"), ref(apple.TvPlayer, "volume.range", W, "--volume-thumb-box-shadow", ""))

	for _, file := range sheets {
		text := uncommented[file]
		// @trap: a custom property's name is no color (--cui-c-console-gray-fill-hover)
		var decls []string
		for _, m := range declValue.FindAllStringSubmatch(text, -1) {
			decls = append(decls, m[1])
		}
		values := customProp.ReplaceAllString(strings.Join(decls, " ; "), "")
		literal := colorFunc.FindString(values)
		if !colorFunc.MatchString(values) {
			literal = named.FindString(values)
			if !named.MatchString(values) {
				literal = "none"
			}
		}
		check.Equal(t, file+" writes no color of its own", literal, "none")
		// @invariant: rule 1: nothing !important, which a utility on the same element
		// would lose to. base.css's reduced-motion block is the exception: it must override
		// every transition a part declares, whatever its layer or specificity
		still := text
		if file == "base.css" {
			if loc := reducedMove.FindStringIndex(text); loc != nil {
				still = text[:loc[0]] + text[loc[1]:]
			}
		}
		has := "none"
		if important.MatchString(still) {
			has = "has !important"
		}
		check.Equal(t, file+" writes nothing !important (rule 1)", has, "none")
		// @invariant: a console part reads the console's slots alone; the dialog's backdrop the
		// scrim, which the toolkit names none of, and the player the video's
		// (decision 28)
		if strings.HasPrefix(file, filepath.Join("components", "console")) {
			var other []string
			for _, m := range colorRead.FindAllStringSubmatch(text, -1) {
				name := m[1]
				if !strings.HasPrefix(name, "--cui-c-console-") && name != "--cui-c-scrim" && !videoName.MatchString(name) {
					other = append(other, name)
				}
			}
			got := strings.Join(other, " ")
			if got == "" {
				got = "none"
			}
			check.Equal(t, file+" reads the console's colors alone", got, "none")
		}
	}
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// @external: Node's readdirSync(dir, { recursive: true }) lists breadth first: a
// directory's entries, then each subdirectory's in the order they were listed
func nodeRecursiveCSS(t *testing.T, dir, prefix string) []string {
	var out []string
	queue := []string{""}
	for len(queue) > 0 {
		rel := queue[0]
		queue = queue[1:]
		entries, err := os.ReadDir(filepath.Join(dir, rel))
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range entries {
			p := filepath.Join(rel, e.Name())
			if strings.HasSuffix(p, ".css") {
				out = append(out, filepath.Join(prefix, p))
			}
			if e.IsDir() {
				queue = append(queue, p)
			}
		}
	}
	return out
}

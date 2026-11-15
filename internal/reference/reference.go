// Package reference reads what was measured on Apple's pages (reference/apple.com/
// results/, reference/asc/results/), the values the tests hold ciderui to.
//
// @invariant: A probe, width or property the record lacks is an error: a check never
// compares with a value nobody measured.
package reference

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"sync"

	"github.com/at-least/ciderui/internal/build"
)

// Widths are the windows a case is read at unless it says otherwise.
//
// @trap: not apple.com's fourth, 1600: no part's rule starts above 1200, so only what
// the window sizes, up to its max-width, reads otherwise there, in the cases that say
// With1600
var Widths = []int{1200, 900, 390}

// With1600 is Widths with 1600 first.
var With1600 = []int{1600, 1200, 900, 390}

type box map[string]float64

type probe struct {
	Styles map[string]string            `json:"styles"`
	States map[string]map[string]string `json:"states"`
	Box    box                          `json:"box"`
}

type applePage struct {
	Widths map[string]map[string]probe `json:"widths"`
}

type ascPage struct {
	Probes map[string]probe  `json:"probes"`
	Tokens map[string]string `json:"tokens"`
}

var (
	mu         sync.Mutex
	applePages = map[string]*applePage{}
	ascPages   = map[string]*ascPage{}
)

func load(dir, name string, into any) error {
	b, err := os.ReadFile(filepath.Join(build.Root, dir, name+".json"))
	if err != nil {
		return fmt.Errorf("reference: %w", err)
	}
	if err := json.Unmarshal(b, into); err != nil {
		return fmt.Errorf("reference: %s/%s.json: %w", dir, name, err)
	}
	return nil
}

func px(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) + "px" }

// Apple is apple.com's measured value of prop on page's probe at width, in state if
// one is named.
func Apple(page, name string, width int, prop, state string) (string, error) {
	mu.Lock()
	p, ok := applePages[page]
	if !ok {
		p = &applePage{Widths: nil}
		if err := load("reference/apple.com/results", page, p); err != nil {
			mu.Unlock()
			return "", err
		}
		applePages[page] = p
	}
	mu.Unlock()
	at, ok := p.Widths[strconv.Itoa(width)]
	if !ok {
		return "", fmt.Errorf("reference %s has no width %d", page, width)
	}
	pr, ok := at[name]
	if !ok {
		return "", fmt.Errorf("reference %s:%s has nothing at %dpx", page, name, width)
	}
	// @external: the probe recorded the box at rest only
	if slices.Contains([]string{"box-height", "box-width", "box-top", "box-left", "box-right"}, prop) {
		var v float64
		var have bool
		if left, hasLeft := pr.Box["left"]; prop == "box-right" && hasLeft {
			v, have = left+pr.Box["width"], true
		} else {
			v, have = pr.Box[prop[len("box-"):]]
		}
		if state != "" || !have {
			s := ""
			if state != "" {
				s = state + " "
			}
			return "", fmt.Errorf("reference %s:%s has no %s%s at %dpx", page, name, s, prop, width)
		}
		return px(v), nil
	}
	styles := pr.Styles
	if state != "" {
		styles = pr.States[state]
	}
	v, ok := styles[prop]
	if !ok {
		s := ""
		if state != "" {
			s = ":" + state
		}
		return "", fmt.Errorf("reference %s:%s%s has no %s at %dpx", page, name, s, prop, width)
	}
	return v, nil
}

func asc(page string) (*ascPage, error) {
	mu.Lock()
	defer mu.Unlock()
	p, ok := ascPages[page]
	if !ok {
		p = &ascPage{Probes: nil, Tokens: nil}
		if err := load("reference/asc/results", page, p); err != nil {
			return nil, err
		}
		ascPages[page] = p
	}
	return p, nil
}

// ASC is App Store Connect's measured value of prop on page's probe.
//
// @external: App Store Connect's pages were read at one width, 1920px
func ASC(page, name, prop string) (string, error) {
	p, err := asc(page)
	if err != nil {
		return "", err
	}
	pr, ok := p.Probes[name]
	if !ok {
		return "", fmt.Errorf("reference asc %s has no probe %s", page, name)
	}
	if prop == "box-height" || prop == "box-width" {
		v, ok := pr.Box[prop[len("box-"):]]
		if !ok {
			return "", fmt.Errorf("reference asc %s:%s has no box", page, name)
		}
		return px(v), nil
	}
	v, ok := pr.Styles[prop]
	if !ok {
		return "", fmt.Errorf("reference asc %s:%s has no %s", page, name, prop)
	}
	return v, nil
}

// Token is a token of App Store Connect's toolkit, --tb-<name>, as its page
// computed it.
func Token(name string) (string, error) {
	p, err := asc("tokens")
	if err != nil {
		return "", err
	}
	v, ok := p.Tokens["--tb-"+name]
	if !ok {
		return "", fmt.Errorf("reference asc tokens has no --tb-%s", name)
	}
	return v, nil
}

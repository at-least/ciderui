//go:build browser

package pictures

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/at-least/ciderui/internal/reference"
	"github.com/mxschmitt/playwright-go"
)

// @invariant: a changed picture is written here for a reader to see; scripts/docker.mjs
// empties it before a run and copies it back after
const failedDir = "test/pictures-failed"

func TestPictures(t *testing.T) {
	update := os.Getenv("UPDATE_GOLDEN") == "1"
	goldenFile := filepath.Join(build.Root, "test/pictures.json")
	golden := map[string]string{}
	if !update {
		raw, err := os.ReadFile(goldenFile)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(raw, &golden); err != nil {
			t.Fatal(err)
		}
	}
	var seen [][2]string
	b := browser.Launch(t)
	pages := map[string]*browser.Page{}
	for _, set := range build.Sets {
		css, err := build.Shipped(set)
		if err != nil {
			t.Fatal(err)
		}
		p, err := b.Open(css, "", browser.Options{})
		if err != nil {
			t.Fatal(err)
		}
		defer p.Close()
		// @external: every face loaded before the first resize: one first asked for after a
		// resize would be drawn in its fallback
		if _, err := p.Page.Evaluate(`() => Promise.all([...document.fonts].map((face) => face.load()))`); err != nil {
			t.Fatal(err)
		}
		pages[set] = p
	}
	parts, err := build.Parts()
	if err != nil {
		t.Fatal(err)
	}
	for _, part := range parts {
		page := pages[part.Set]
		files, err := build.ExampleFiles(part)
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range files {
			html, err := build.ExampleHTML(part, file)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := page.Page.Evaluate(`(html) => { document.body.innerHTML = html; }`, html); err != nil {
				t.Fatal(err)
			}
			for _, width := range reference.Widths {
				if err := page.Resize(width); err != nil {
					t.Fatal(err)
				}
				picture, err := page.Page.Screenshot(playwright.PageScreenshotOptions{FullPage: playwright.Bool(true), Animations: playwright.ScreenshotAnimationsDisabled, Caret: playwright.ScreenshotCaretHide})
				if err != nil {
					t.Fatal(err)
				}
				sum := sha256.Sum256(picture)
				hash := hex.EncodeToString(sum[:])
				name := fmt.Sprintf("%s/%s at %d", part.Name, file, width)
				seen = append(seen, [2]string{name, hash})
				saved := fmt.Sprintf("%s/%s-%s-%d.png", failedDir, part.Name, file, width)
				if !update && !check.Hold(t, hash == golden[name], fmt.Sprintf("%spx: the picture changed, see %s (if meant: UPDATE_GOLDEN=1 node scripts/docker.mjs go test -count=1 -tags browser ./test/pictures)", name, saved)) {
					if err := os.MkdirAll(filepath.Join(build.Root, failedDir), 0o755); err != nil {
						t.Fatal(err)
					}
					if err := os.WriteFile(filepath.Join(build.Root, saved), picture, 0o644); err != nil {
						t.Fatal(err)
					}
				}
			}
		}
	}
	if update {
		var out bytes.Buffer
		out.WriteString("{\n")
		for i, kv := range seen {
			fmt.Fprintf(&out, " %q: %q", kv[0], kv[1])
			if i < len(seen)-1 {
				out.WriteString(",")
			}
			out.WriteString("\n")
		}
		out.WriteString("}\n")
		if err := os.WriteFile(goldenFile, out.Bytes(), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	var unseen []string
	for name := range golden {
		if !slices.ContainsFunc(seen, func(kv [2]string) bool { return kv[0] == name }) {
			unseen = append(unseen, name)
		}
	}
	slices.Sort(unseen)
	check.Equal(t, "every name in pictures.json has an example", strings.Join(unseen, " "), "")
}

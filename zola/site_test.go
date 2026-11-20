// Package zola holds the Zola theme beside it to the kit: its components write the
// parts' examples, every page of its test site passes ciderui.CheckMarkup, and, in
// the test image, no page scrolls sideways.
package zola

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/pinned"
)

// @invariant: the checkout reaches the site's copy as a site's submodule of ciderui
// does, at themes/ciderui (theme = "ciderui/zola"), so the theme's static/ links
// (dist/site.css, fonts/) resolve as they do on a site
func buildSite(t *testing.T, extra map[string]string) string {
	t.Helper()
	zola, err := pinned.Zola()
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	if err := os.CopyFS(root, os.DirFS(filepath.Join(build.Root, "zola/test/site"))); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "themes"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(build.Root, filepath.Join(root, "themes/ciderui")); err != nil {
		t.Fatal(err)
	}
	for path, content := range extra {
		if err := os.MkdirAll(filepath.Dir(filepath.Join(root, path)), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, path), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	out := filepath.Join(root, "public")
	if msg, err := exec.Command(zola, "--root", root, "build", "--output-dir", out).CombinedOutput(); err != nil {
		t.Fatalf("zola build: %v\n%s", err, msg)
	}
	return out
}

var redirect = regexp.MustCompile(`(?i)<meta http-equiv="refresh"`)

// @external: Zola writes a redirect page for each alias, no page of the site
func pages(t *testing.T, out string) map[string]string {
	t.Helper()
	found := map[string]string{}
	err := filepath.WalkDir(out, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".html" {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if !redirect.Match(b) {
			rel, err := filepath.Rel(out, path)
			if err != nil {
				return err
			}
			found[rel] = string(b)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return found
}

var (
	title       = regexp.MustCompile(`<title>[^<]+</title>`)
	description = regexp.MustCompile(`<meta name="description" content="[^"]+">`)
	stylesheets = regexp.MustCompile(`<link rel="stylesheet" href="[^"]*/ciderui\.css\?h=[0-9a-f]+">\n<link rel="stylesheet" href="[^"]*/ciderui-zola\.css\?h=[0-9a-f]+">`)
	listing     = regexp.MustCompile(`(?s)<pre.*?</pre>|<code.*?</code>`)
	tag         = regexp.MustCompile(`<[^>]+>`)
	tera        = regexp.MustCompile(`\{\{|\{%|\{#|%\}|#\}`)
	markdown    = regexp.MustCompile(`\[\^[^\]]*\]|\]\([^)]*\)|\*\*[^*]+\*\*`)
)

// @invariant: the test site holds every template (index, section with a heading,
// page, 404) and the footer's directory: fewer pages built means one of them went
const sitePages = 12

func TestEveryPageOfTheTestSiteIsWhole(t *testing.T) {
	found := pages(t, buildSite(t, nil))
	if len(found) != sitePages {
		t.Errorf("%d pages built, want %d", len(found), sitePages)
	}
	for path, html := range found {
		body := html[strings.Index(html, "<body"):]
		text := tag.ReplaceAllString(listing.ReplaceAllString(body, ""), " ")
		for _, want := range []struct {
			ok   bool
			what string
		}{
			{title.MatchString(html), "a title"},
			{description.MatchString(html), "a description"},
			{stylesheets.MatchString(html), "dist/site.css, then the theme's stylesheet"},
			{!tera.MatchString(text), "no Tera showing"},
			{!markdown.MatchString(text), "no markdown showing"},
		} {
			if !want.ok {
				t.Errorf("%s: wants %s", path, want.what)
			}
		}
		for _, problem := range ciderui.CheckMarkup(html) {
			t.Errorf("%s: %s", path, problem)
		}
	}
}

// @invariant: README's setup is the one this file builds: a submodule at
// themes/ciderui and theme = "ciderui/zola"
func TestReadmesSetupIsTheTestSites(t *testing.T) {
	readme, err := os.ReadFile(filepath.Join(build.Root, "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	config, err := os.ReadFile(filepath.Join(build.Root, "zola/test/site/config.toml"))
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range []string{"git submodule add https://github.com/at-least/ciderui themes/ciderui\n", "theme = \"ciderui/zola\"\n"} {
		if !strings.Contains(string(readme), line) {
			t.Errorf("README.md has no line %q", line)
		}
	}
	if !strings.Contains(string(config), "theme = \"ciderui/zola\"\n") {
		t.Errorf("zola/test/site/config.toml names another theme than README.md")
	}
}

func TestTheStylesheetsServedAreTheKitsAndTheThemes(t *testing.T) {
	out := buildSite(t, nil)
	for served, source := range map[string]string{
		"ciderui.css":                   "dist/site.css",
		"ciderui-zola.css":              "zola/static/ciderui-zola.css",
		"fonts/inter-roman-latin.woff2": "fonts/inter-roman-latin.woff2",
	} {
		got, err := os.ReadFile(filepath.Join(out, served))
		if err != nil {
			t.Fatal(err)
		}
		want, err := os.ReadFile(filepath.Join(build.Root, source))
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != string(want) {
			t.Errorf("the site's %s is not %s", served, source)
		}
	}
}

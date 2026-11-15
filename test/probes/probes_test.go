//go:build browser

// @invariant: Run in a copy of the repository: a probe writes results/<browser>.json, and
// the committed results stay the record of the runs each MEASURED.md names.
package probes

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/at-least/ciderui/internal/build"
)

type run struct {
	env map[string]string
}

// @external: each in Chromium; the image has no system font of Apple's, so stack's
// stack draws Inter
var probes = map[string]run{
	"console": {env: map[string]string{"SITE": "{copy}/console-site"}},
	"details": {},
	"dialog":  {},
	"flyout":  {},
	"forced":  {},
	"popover": {},
	"scheme":  {},
	"search":  {},
	"select":  {},
	"stack":   {env: map[string]string{"EXPECT": "inter"}},
	"tap":     {},
}

var notRun = map[string]string{
	"anchor":             "reads apple.com live",
	"apple-account":      "reads Apple Account live, signed in",
	"apple.com":          "reads apple.com live (its report step prints what it read)",
	"asc":                "reads App Store Connect live, signed in",
	"developer-account":  "reads the developer account live, signed in",
	"feedback-assistant": "reads Feedback Assistant live, signed in",
	"icons":              "reads apple.com live",
	"tracking":           "reads apple.com live (derive and tables read what it wrote)",
	"air":                "Safari on air, over ssh",
}

func TestProbes(t *testing.T) {
	entries, err := os.ReadDir(filepath.Join(build.Root, "reference"))
	if err != nil {
		t.Fatal(err)
	}
	var commands []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if gos, _ := filepath.Glob(filepath.Join(build.Root, "reference", e.Name(), "*.go")); len(gos) > 0 {
			commands = append(commands, e.Name())
		}
	}
	for _, c := range commands {
		_, run := probes[c]
		_, skipped := notRun[c]
		if !run && !skipped {
			t.Errorf("reference/%s: not run here and not in notRun", c)
		}
	}
	for c := range notRun {
		if !slices.Contains(commands, c) {
			t.Errorf("notRun names reference/%s, no such command", c)
		}
	}

	copy := t.TempDir()
	tree := filepath.Join(copy, "ciderui")
	if out, err := exec.Command("sh", "-c", fmt.Sprintf("mkdir -p %s && tar -c -C %s --exclude ./.git . | tar -x -C %s", tree, build.Root, tree)).CombinedOutput(); err != nil {
		t.Fatalf("copy: %v\n%s", err, out)
	}
	names := make([]string, 0, len(probes))
	for name := range probes {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		start := time.Now()
		cmd := exec.Command("go", "run", "./reference/"+name)
		cmd.Dir = tree
		cmd.Env = append(os.Environ(), "BROWSER=chromium")
		for k, v := range probes[name].env {
			cmd.Env = append(cmd.Env, k+"="+strings.ReplaceAll(v, "{copy}", copy))
		}
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Errorf("%s (%.1fs): %v\n%s", name, time.Since(start).Seconds(), err, out)
		}
	}
	t.Logf("probes: %d run, %d not run here", len(probes), len(notRun))
}

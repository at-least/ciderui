// Command docker runs a command in the test image (test/Dockerfile) on a copy of the
// repository, and copies back what the command wrote:
//
//	go run ./scripts/docker go test -count=1 -tags browser ./components/site/field
//
// @external: Nothing is bind-mounted: the Docker daemon may be on another machine
// (DOCKER_HOST=ssh://…), where a bind mount would read that machine's paths.
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/at-least/ciderui/internal/build"
)

const (
	image      = "ciderui-test"
	dockerfile = "test/Dockerfile"
)

// @invariant: the files the image's build reads: test/Dockerfile.dockerignore admits these
// alone, and their bytes name the image
var context = []string{"go.mod", "go.sum", "internal/pinned/pinned.go", "scripts/fetch/main.go"}

var imageID = regexp.MustCompile(`^sha256:[0-9a-f]{64}$`)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: go run ./scripts/docker <command> [argument…]")
		os.Exit(2)
	}
	status, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "docker:", err)
		os.Exit(1)
	}
	os.Exit(status)
}

func docker(args ...string) (string, error) {
	cmd := exec.Command("docker", args...)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("docker %s: %w", args[0], err)
	}
	return strings.TrimSpace(string(out)), nil
}

func run(command []string) (int, error) {
	root := build.Root
	ignore, err := os.ReadFile(filepath.Join(root, dockerfile+".dockerignore"))
	if err != nil {
		return 0, err
	}
	var admitted []string
	for _, line := range strings.Split(string(ignore), "\n") {
		if strings.HasPrefix(line, "!") {
			admitted = append(admitted, line[1:])
		}
	}
	// @trap: a file the image reads and the hash leaves out reuses a stale image
	if !slices.Equal(admitted, context) {
		return 0, fmt.Errorf("%s.dockerignore admits %v, the hash reads %v", dockerfile, admitted, context)
	}
	// @external: an image built from the same bytes carries their hash as a label and is
	// run by its ID, with no build (a cached build still took 2s over ssh). Not by the
	// tag: lab's daemon serves every Mac and checkout, and another's build can move the
	// tag meanwhile.
	inputs := sha256.New()
	for _, f := range append([]string{dockerfile}, context...) {
		b, err := os.ReadFile(filepath.Join(root, f))
		if err != nil {
			return 0, err
		}
		inputs.Write([]byte(f))
		inputs.Write(b)
	}
	label := "ciderui.inputs=" + hex.EncodeToString(inputs.Sum(nil))
	listed, err := docker("images", "-q", "--no-trunc", "--filter", "label="+label)
	if err != nil {
		return 0, err
	}
	id, _, _ := strings.Cut(listed, "\n")
	if id == "" {
		if id, err = docker("build", "-q", "-t", image, "--label", label, "-f", filepath.Join(root, dockerfile), root); err != nil {
			return 0, err
		}
	}
	if !imageID.MatchString(id) {
		return 0, fmt.Errorf("docker gave no image ID: %s", id)
	}
	// @external: CIDERUI_CHROME=1: the image with Google Chrome too
	// (reference/chrome.Dockerfile), for a probe of a page Playwright's Chromium cannot
	// play; the tests never run in it
	if os.Getenv("CIDERUI_CHROME") == "1" {
		// @external: FROM takes a name, not an ID: a tag of this image's own
		base := image + ":" + id[len("sha256:"):len("sha256:")+12]
		if _, err := docker("tag", id, base); err != nil {
			return 0, err
		}
		chrome, err := os.ReadFile(filepath.Join(root, "reference/chrome.Dockerfile"))
		if err != nil {
			return 0, err
		}
		cmd := exec.Command("docker", "build", "-q", "--build-arg", "BASE="+base, "-")
		cmd.Stdin = bytes.NewReader(chrome)
		cmd.Stderr = os.Stderr
		out, err := cmd.Output()
		if err != nil {
			return 0, fmt.Errorf("docker build chrome: %w", err)
		}
		if id = strings.TrimSpace(string(out)); !imageID.MatchString(id) {
			return 0, fmt.Errorf("docker gave no image ID for Chrome's: %s", id)
		}
	}

	args := []string{"create", "--init", "--ipc=host",
		// @external: Go's build cache a volume on the daemon's machine, no bind mount: a run
		// compiles what changed since the last, not the whole module
		"-v", "ciderui-gocache:/root/.cache/go-build"}
	for _, k := range []string{"UPDATE_GOLDEN"} {
		if v, ok := os.LookupEnv(k); ok {
			args = append(args, "-e", k+"="+v)
		}
	}
	// @invariant: The image's user (root): what it writes comes back through tar as the
	// caller's own, so no --user (the repository's top, the image's WORKDIR, is root's).
	// The stamp before the command: what is newer after it, the command wrote (the
	// copied files keep the checkout's mtimes).
	// @external: git init: the checkout's .git stays behind, and pointers_test.go reads
	// the repository's files from git (untracked ones, .gitignore's left out); the
	// container's own .git never comes back
	const wrap = `git init -q && touch /tmp/stamp && "$@"; status=$?; find . -path ./.git -prune -o -type f -newer /tmp/stamp -print > /tmp/written && tar -cf /tmp/written.tar -T /tmp/written; exit $status`
	args = append(append(args, id, "sh", "-c", wrap, "run"), command...)
	container, err := docker(args...)
	if err != nil {
		return 0, err
	}
	defer exec.Command("docker", "rm", "-f", container).Run()

	// @invariant: emptied here, not in the test: that runs in the container, which never
	// holds this directory (left out below), and the copy back only adds files
	if err := os.RemoveAll(filepath.Join(root, "test/pictures-failed")); err != nil {
		return 0, err
	}
	// @external: no extended attributes (docker cp refuses macOS's com.apple.provenance,
	// "operation not supported", on a file that carries it), no AppleDouble ._ files
	if err := shell(fmt.Sprintf(`tar -c --no-xattrs -C %s --exclude ./.git --exclude ./test/pictures-failed . | docker cp - %s:/src/ciderui`, quote(root), container)); err != nil {
		return 0, err
	}
	start := exec.Command("docker", "start", "-a", container)
	start.Stdout, start.Stderr = os.Stdout, os.Stderr
	status := 0
	if err := start.Run(); err != nil {
		exit, ok := err.(*exec.ExitError)
		if !ok {
			return 0, fmt.Errorf("docker start: %w", err)
		}
		status = exit.ExitCode()
	}
	out, err := os.MkdirTemp("", "ciderui-written-")
	if err != nil {
		return 0, err
	}
	defer os.RemoveAll(out)
	if err := shell(fmt.Sprintf(`docker cp %s:/tmp/written.tar - | tar -xOf - written.tar | tar -xf - -C %s`, container, quote(out))); err != nil {
		return 0, err
	}
	var written []string
	err = filepath.WalkDir(out, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(out, path)
		written = append(written, rel)
		return err
	})
	if err != nil {
		return 0, err
	}
	if len(written) > 0 {
		cmd := exec.Command("sh", "-c", fmt.Sprintf(`tar -c -C %s -T - | tar -xf - -C %s`, quote(out), quote(root)))
		cmd.Stdin = strings.NewReader(strings.Join(written, "\n"))
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return 0, fmt.Errorf("copy back: %w", err)
		}
		fmt.Printf("docker: %d written file(s) copied back\n", len(written))
	}
	return status, nil
}

func shell(script string) error {
	cmd := exec.Command("sh", "-c", script)
	cmd.Env = append(os.Environ(), "COPYFILE_DISABLE=1")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w", script, err)
	}
	return nil
}

func quote(s string) string { return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'" }

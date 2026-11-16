// Package pinned fetches the tools the repository's own checks and the kit's
// command run, each at one version and checksum, into the user's cache: Tailwind's
// standalone binary and the npm packages the checks read, unpacked. The repository
// holds no node_modules, and a site needs none. It also writes the entry of a site's
// own utilities' sheet, which the kit's command and the tests both build.
package pinned

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// TailwindVersion builds dist/ and, through `ciderui css`, a site's own utilities:
// one version, so the two sheets cannot part.
const TailwindVersion = "4.3.3"

type artifact struct {
	url    string
	sha256 string
}

// @external: the sums are the release's sha256sums.txt (Tailwind) and the registry's
// tarballs, whose sha512 matched package-lock.json's integrity when they were pinned
var tailwind = map[string]artifact{
	"linux/amd64":  {url: "https://github.com/tailwindlabs/tailwindcss/releases/download/v" + TailwindVersion + "/tailwindcss-linux-x64", sha256: "dc61b3ac6b8c9ca874c0cc4c57b2409791a64c5540404ca5f5367360babc313a"},
	"darwin/arm64": {url: "https://github.com/tailwindlabs/tailwindcss/releases/download/v" + TailwindVersion + "/tailwindcss-macos-arm64", sha256: "cdf646702987a743464dff4d9c60fd4480d1c1e73dd819a9a67f1078815dce9d"},
}

// LucideVersion is the lucide-static whose icons go/lucide/lucide.go holds.
const LucideVersion = "1.61.0"

var (
	lucide = artifact{url: "https://registry.npmjs.org/lucide-static/-/lucide-static-" + LucideVersion + ".tgz", sha256: "62773cd2ee411ad2e37c1973d1e4c4b49f9a6c57da6b5d5d0abca08756ba75f1"}
	axe    = artifact{url: "https://registry.npmjs.org/axe-core/-/axe-core-4.14.0.tgz", sha256: "2277b0a3ef47af9b52497faa3e71598e86e0a0e5e6eb295207a5657ede4709af"}
)

// Tailwind is the path of Tailwind's standalone binary for this machine.
func Tailwind() (string, error) {
	platform := runtime.GOOS + "/" + runtime.GOARCH
	a, ok := tailwind[platform]
	if !ok {
		return "", fmt.Errorf("pinned: no Tailwind binary pinned for %s", platform)
	}
	path, err := fetch(a)
	if err != nil {
		return "", err
	}
	if err := os.Chmod(path, 0o755); err != nil {
		return "", fmt.Errorf("pinned: %w", err)
	}
	return path, nil
}

// UtilitiesEntry is the Tailwind entry of a site's own utilities' sheet: the
// utilities.css at utilities, an absolute path, then one `@source <s>;` line for
// each of sources. `ciderui css` writes its @source of a site's packages, and the
// tests' pages their classes inline.
func UtilitiesEntry(utilities string, sources []string) []byte {
	entry := fmt.Appendf(nil, "@import %q;\n", utilities)
	for _, s := range sources {
		entry = fmt.Appendf(entry, "@source %s;\n", s)
	}
	return entry
}

// Lucide is the directory of the lucide-static package, unpacked.
func Lucide() (string, error) { return unpacked(lucide) }

// Axe is the directory of the axe-core package, unpacked.
func Axe() (string, error) { return unpacked(axe) }

// Fetch fetches every artifact this machine's checks use: the test image's build
// runs it, so a run in the image reaches no network.
func Fetch() error {
	if _, err := Tailwind(); err != nil {
		return err
	}
	if _, err := Lucide(); err != nil {
		return err
	}
	_, err := Axe()
	return err
}

func dir(a artifact) (string, error) {
	cache, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("pinned: %w", err)
	}
	return filepath.Join(cache, "ciderui", a.sha256), nil
}

func fetch(a artifact) (string, error) {
	d, err := dir(a)
	if err != nil {
		return "", err
	}
	path := filepath.Join(d, filepath.Base(a.url))
	if _, err := os.Stat(path); err == nil {
		return path, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("pinned: %w", err)
	}
	if err := os.MkdirAll(d, 0o755); err != nil {
		return "", fmt.Errorf("pinned: %w", err)
	}
	res, err := http.Get(a.url)
	if err != nil {
		return "", fmt.Errorf("pinned: fetch %s: %w", a.url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("pinned: fetch %s: %s", a.url, res.Status)
	}
	// @invariant: written beside its place and renamed once its sum holds, so a path
	// that exists is a checked file, whichever of two runs at once wrote it
	tmp, err := os.CreateTemp(d, "fetch-")
	if err != nil {
		return "", fmt.Errorf("pinned: %w", err)
	}
	defer os.Remove(tmp.Name())
	sum := sha256.New()
	if _, err := io.Copy(io.MultiWriter(tmp, sum), res.Body); err != nil {
		tmp.Close()
		return "", fmt.Errorf("pinned: fetch %s: %w", a.url, err)
	}
	if err := tmp.Close(); err != nil {
		return "", fmt.Errorf("pinned: %w", err)
	}
	if got := hex.EncodeToString(sum.Sum(nil)); got != a.sha256 {
		return "", fmt.Errorf("pinned: %s has sha256 %s, pinned %s", a.url, got, a.sha256)
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return "", fmt.Errorf("pinned: %w", err)
	}
	return path, nil
}

func unpacked(a artifact) (string, error) {
	tgz, err := fetch(a)
	if err != nil {
		return "", err
	}
	d := filepath.Dir(tgz)
	out := filepath.Join(d, "package")
	if _, err := os.Stat(out); err == nil {
		return out, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("pinned: %w", err)
	}
	tmp, err := os.MkdirTemp(d, "unpack-")
	if err != nil {
		return "", fmt.Errorf("pinned: %w", err)
	}
	defer os.RemoveAll(tmp)
	if err := untar(tgz, tmp); err != nil {
		return "", fmt.Errorf("pinned: unpack %s: %w", tgz, err)
	}
	// @external: a rename onto a run's that unpacked it meanwhile fails (directory not
	// empty), and that one stands
	if err := os.Rename(filepath.Join(tmp, "package"), out); err != nil {
		if _, statErr := os.Stat(out); statErr != nil {
			return "", fmt.Errorf("pinned: %w", err)
		}
	}
	return out, nil
}

func untar(tgz, into string) error {
	f, err := os.Open(tgz)
	if err != nil {
		return err
	}
	defer f.Close()
	z, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	r := tar.NewReader(z)
	for {
		h, err := r.Next()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		if h.Typeflag != tar.TypeReg {
			continue
		}
		name := filepath.Clean(h.Name)
		if !strings.HasPrefix(name, "package"+string(filepath.Separator)) {
			return fmt.Errorf("%s is outside package/", h.Name)
		}
		path := filepath.Join(into, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		w, err := os.Create(path)
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, r); err != nil {
			w.Close()
			return err
		}
		if err := w.Close(); err != nil {
			return err
		}
	}
}

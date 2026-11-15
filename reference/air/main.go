// Command air runs a probe on air (the owner's Mac that nobody sits at, `ssh air`),
// whose Safari is driven through safaridriver: never the Safari of the Mac in front
// of someone. The probe's command is built here for air (darwin/arm64); the
// repository goes to ~/ciderui-air on air (no .git or .claude), SITE and INTER to
// ~/ciderui-air-env/<name> where they name a path, the binary beside them; it runs in
// the probe's directory with BROWSER, SITE, MODE, EXPECT and INTER passed through
// (SITE and INTER as air's copies), and the directory's results/ comes back. air's Safari takes automation (Develop > Allow
// Remote Automation, on).
//
//	BROWSER=safari go run ./reference/air ./reference/scheme
//
// air has no Xcode: BROWSER=ios and the XCUITest probes stay on a Mac with it.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/at-least/ciderui/internal/build"
)

const host = "air"

// @trap: one run at a time: safaridriver drives one Safari, and each run quits
// Safari first, which would end another run's session. A lock left by a run cut
// off (its ssh dropped) names the run that made it; remove it with
// ssh air rmdir /tmp/ciderui-air.lock
const lock = "/tmp/ciderui-air.lock"

func main() {
	status, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(status)
}

func quote(s string) string { return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'" }

func ssh(script string) error {
	cmd := exec.Command("ssh", host, script)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}

// @trap: a directory's trailing slash: rsync copies its contents into the same path,
// not a copy of it inside; the parent made first, openrsync having no --mkpath
func sync(from, to string, dir bool, excludes ...string) error {
	args := []string{"-a", "--delete"}
	for _, e := range excludes {
		args = append(args, "--exclude", e)
	}
	if dir {
		from, to = from+"/", to+"/"
	}
	cmd := exec.Command("rsync", append(args, from, to)...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("rsync %s %s: %w", from, to, err)
	}
	return nil
}

func run(args []string) (int, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("usage: go run ./reference/air <probe's package> [args…]")
	}
	// @external: air has Safari alone: Playwright's browsers are not synced, and the
	// owner's call is that Chrome runs in Docker (scripts/docker)
	if b, set := os.LookupEnv("BROWSER"); set && b != "safari" {
		return 0, fmt.Errorf("BROWSER=%s does not run on air, which has Safari alone: Chrome, Chromium and Firefox run in Docker (scripts/docker), the iOS simulator on a Mac with Xcode", b)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return 0, err
	}
	pkg, rest := args[0], args[1:]
	list := exec.Command("go", "list", "-f", "{{.Dir}}", pkg)
	list.Stderr = os.Stderr
	out, err := list.Output()
	if err != nil {
		return 0, fmt.Errorf("go list %s: %w", pkg, err)
	}
	dir := strings.TrimSpace(string(out))
	if !strings.HasPrefix(dir, build.Root+string(filepath.Separator)) {
		return 0, fmt.Errorf("%s is not a package of the repository (%s)", pkg, build.Root)
	}
	// @trap: not this machine's path on air: lab's /home/newlix cannot be made on macOS
	// ("mkdir: /home/newlix: Operation not supported", 2026-11-15); the binary finds the
	// repository from the directory it runs in, so any place serves
	const remoteRoot, remoteEnv = "ciderui-air", "ciderui-air-env"
	rel, err := filepath.Rel(build.Root, dir)
	if err != nil {
		return 0, err
	}
	var pass []string
	for _, k := range []string{"BROWSER", "MODE", "EXPECT"} {
		if v, set := os.LookupEnv(k); set {
			pass = append(pass, k+"="+quote(v))
		}
	}
	for _, k := range []string{"SITE", "INTER"} {
		if _, set := os.LookupEnv(k); set {
			pass = append(pass, k+`="$HOME"/`+remoteEnv+"/"+k)
		}
	}
	paths := map[string]string{}
	for _, k := range []string{"SITE", "INTER"} {
		p, set := os.LookupEnv(k)
		if !set {
			continue
		}
		if _, err := os.Stat(p); err != nil || !filepath.IsAbs(p) {
			return 0, fmt.Errorf("%s must name an existing absolute path: %s", k, p)
		}
		paths[k] = p
	}

	// @external: built here for air: air needs no Go, and the binary finds the
	// repository from the directory it runs in (build.Root)
	local, err := os.MkdirTemp("", "ciderui-air-")
	if err != nil {
		return 0, err
	}
	defer os.RemoveAll(local)
	binary := filepath.Join(local, filepath.Base(dir))
	compile := exec.Command("go", "build", "-o", binary, pkg)
	compile.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=arm64", "CGO_ENABLED=0")
	compile.Stdout, compile.Stderr = os.Stdout, os.Stderr
	if err := compile.Run(); err != nil {
		return 0, fmt.Errorf("build %s for air: %w", pkg, err)
	}
	remoteBinary := "/tmp/ciderui-air-bin/" + filepath.Base(dir)

	hostname, err := os.Hostname()
	if err != nil {
		return 0, err
	}
	owner := fmt.Sprintf("%s pid %d %s", hostname, os.Getpid(), cwd)
	if err := ssh(fmt.Sprintf(`mkdir %s 2>/dev/null || { echo "air is busy with another probe (%s: $(cat %s/owner 2>/dev/null))" >&2; exit 1; }; echo %s > %s/owner`, lock, lock, lock, quote(owner), lock)); err != nil {
		return 0, fmt.Errorf("lock %s on air: %w", lock, err)
	}
	defer ssh("rm -rf " + lock)

	if err := ssh("mkdir -p " + strings.Join([]string{remoteRoot, remoteEnv, quote(filepath.Dir(remoteBinary))}, " ")); err != nil {
		return 0, err
	}
	if err := sync(build.Root, host+":"+remoteRoot, true, ".git", ".claude"); err != nil {
		return 0, err
	}
	for k, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			return 0, err
		}
		if err := sync(p, host+":"+remoteEnv+"/"+k, info.IsDir()); err != nil {
			return 0, err
		}
	}
	if err := sync(binary, host+":"+remoteBinary, false); err != nil {
		return 0, err
	}
	command := []string{quote(remoteBinary)}
	for _, a := range rest {
		command = append(command, quote(a))
	}
	// @external: Safari 27 takes no session in a Safari it did not launch ("Safari was
	// not launched for automation"): quit it first. caffeinate -du: with air's display
	// asleep, Safari's picture came back all black
	probe := exec.Command("ssh", host, strings.Join([]string{
		"pkill -x Safari; while pgrep -x Safari >/dev/null; do sleep 0.2; done;",
		"cd " + quote(remoteRoot+"/"+rel) + " &&",
		"exec env " + strings.Join(pass, " ") + " caffeinate -du " + strings.Join(command, " "),
	}, " "))
	probe.Stdin, probe.Stdout, probe.Stderr = os.Stdin, os.Stdout, os.Stderr
	status := 0
	if err := probe.Run(); err != nil {
		status = 1
		if exit, ok := err.(*exec.ExitError); ok && exit.ExitCode() > 0 {
			status = exit.ExitCode()
		}
	}
	results := filepath.Join(dir, "results")
	if exec.Command("ssh", host, "test -d "+quote(remoteRoot+"/"+rel+"/results")).Run() == nil {
		if err := os.MkdirAll(results, 0o755); err != nil {
			return status, err
		}
		if err := sync(host+":"+remoteRoot+"/"+rel+"/results", results, true); err != nil {
			return status, err
		}
	}
	return status, nil
}

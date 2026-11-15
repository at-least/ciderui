package ojson

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

var signedIn = regexp.MustCompile(`/reference/(asc|apple-account|developer-account|feedback-assistant)/`)

// @invariant: every record the Node probes wrote reads back and writes byte for byte:
// the writer is JSON.stringify(v, null, 1)'s
func TestRecordsRoundTrip(t *testing.T) {
	files, err := filepath.Glob("../../reference/*/results/*.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) < 50 {
		t.Fatalf("%d records: the glob missed them", len(files))
	}
	for _, f := range files {
		// @external: the signed-in sites' records are written in the reader's browser, not by
		// a probe here, some of their objects on one line
		if signedIn.MatchString(f) {
			continue
		}
		raw, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		v, err := Parse(raw)
		if err != nil {
			t.Fatalf("%s: %v", f, err)
		}
		got := Marshal(v)
		// @external: some writers add a newline after the value, some do not
		if string(got) != string(raw) && string(got)+"\n" != string(raw) {
			t.Errorf("%s: rewritten differs", f)
		}
	}
}

func TestFloatIsJavaScripts(t *testing.T) {
	tenth, fifth := 0.1, 0.2
	for v, want := range map[float64]Number{tenth + fifth: "0.30000000000000004", 1e21: "1e+21", 1e-7: "1e-7", 123.5: "123.5", -0.000001: "-0.000001", 44: "44"} {
		if got := Float(v); got != want {
			t.Errorf("Float(%v) = %s, want %s", v, got, want)
		}
	}
}

func TestWithKeepsPlace(t *testing.T) {
	o := Object{{"a", Number("1")}, {"b", Number("2")}}.With(Member{"a", Number("3")}, Member{"c", Number("4")})
	if got := string(Marshal(o)); got != "{\n \"a\": 3,\n \"b\": 2,\n \"c\": 4\n}" {
		t.Errorf("got %s", got)
	}
}

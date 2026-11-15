package reference

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/ojson"
)

const digest = "[\n  0,\n  1,\n  2,\n  3,\n  4,\n  5,\n  6,\n  7,\n  8,\n  9,\n  10,\n  11,\n  12,\n  13,\n  14,\n  15,\n  16,\n  17,\n  18,\n  19,\n  20,\n  21,\n  22,\n  23,\n  24,\n  25,\n  26,\n  27,\n  28,\n  29,\n  30,\n  255\n ]"

var signedInRecord = `{
 "path": "/apps/<app>/analytics/overview",
 "window": [
  1920,
  898,
  2
 ],
 "browser": "Chrome/154.0.0.0",
 "date": "2026-10-24T03:29:35.360Z",
 "probes": {
  "menu.item": {
   "html": "a {sc-FBvLw}",
   "styles": {
    "height": "36px",
    "color": "rgb(0, 0, 0)"
   },
   "box": {
    "width": 225.5,
    "height": 36
   }
  },
  "menu.label": {
   "html": "div",
   "styles": {}
  }
 },
 "probesDigest": ` + digest + `,
 "readerDigest": ` + digest + `,
 "sha256": ` + digest + `
}`

// TestSignedInRoundTrip holds a signed-in record read into Go and written back to the
// bytes it was read from, in the reader's order, and its Go file to gofmt.
func TestSignedInRoundTrip(t *testing.T) {
	for _, record := range []string{signedInRecord, strings.Replace(signedInRecord, `"probes": {
  "menu.item": {
   "html": "a {sc-FBvLw}",
   "styles": {
    "height": "36px",
    "color": "rgb(0, 0, 0)"
   },
   "box": {
    "width": 225.5,
    "height": 36
   }
  },
  "menu.label": {
   "html": "div",
   "styles": {}
  }
 }`, `"tokens": {
  "--tb-b": "1px",
  "--tb-a": "solid"
 }`, 1)} {
		v, err := ojson.Parse([]byte(record))
		if err != nil {
			t.Fatal(err)
		}
		page, err := SignedInFromJSON("hover-menu", v)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(ojson.Marshal(page.JSON())); got != record {
			t.Errorf("written back:\n%s", got)
		}
		if _, err := SignedInGoFile("asc", page); err != nil {
			t.Error(err)
		}
	}
}

func TestSignedInFails(t *testing.T) {
	for _, c := range []struct{ from, to, want string }{
		{`"html": "div"`, `"html": "div", "seen": 1`, `unknown key "seen"`},
		{`"height": "36px"`, `"height": 36`, `not a string`},
		{`"probes": {`, `"tokens": {}, "probes": {`, `probes or tokens`},
		{`"menu.label": {
   "html": "div",
   "styles": {}
  }`, `"menu.label": null`, `not an object`},
		{"  255\n ]\n}", "  256\n ]\n}", `byte 256`},
	} {
		v, err := ojson.Parse([]byte(strings.Replace(signedInRecord, c.from, c.to, 1)))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := SignedInFromJSON("hover-menu", v); err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%s: SignedInFromJSON = %v, want %s", c.to, err, c.want)
		}
	}
}

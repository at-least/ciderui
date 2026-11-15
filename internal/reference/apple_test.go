package reference

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/ojson"
)

const record = `{
 "url": "https://www.apple.com/",
 "finalUrl": "https://www.apple.com/",
 "browser": "chromium",
 "date": "2026-11-15T00:00:00.000Z",
 "widths": {
  "390": {
   "a": {
    "html": "<a>",
    "styles": {
     "color": ""
    },
    "keyframes": {},
    "box": {
     "width": 0.5,
     "height": 0,
     "top": 0,
     "left": 0
    },
    "animations": []
   },
   "b": null,
   "c": {
    "html": "<b>",
    "styles": {}
   }
  }
 }
}`

// TestRecordRoundTrip holds a record read into Go and written back to the bytes it was
// read from: an empty list stays apart from none, a zero from no position.
func TestRecordRoundTrip(t *testing.T) {
	v, err := ojson.Parse([]byte(record))
	if err != nil {
		t.Fatal(err)
	}
	page, err := FromJSON("home", v)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(ojson.Marshal(page.JSON())); got != record {
		t.Errorf("written back:\n%s", got)
	}
	if _, err := GoFile(page); err != nil {
		t.Error(err)
	}
}

func TestRecordUnknownKeyFails(t *testing.T) {
	v, err := ojson.Parse([]byte(strings.Replace(record, `"html": "<b>"`, `"html": "<b>", "seen": 1`, 1)))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := FromJSON("home", v); err == nil || !strings.Contains(err.Error(), `unknown key "seen"`) {
		t.Errorf("FromJSON = %v, want the unknown key", err)
	}
}

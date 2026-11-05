package glyph

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestEscapesItsClass(t *testing.T) {
	var b strings.Builder
	if err := More.Class(`a"b<c&d`).Render(&b); err != nil {
		t.Fatal(err)
	}
	if got := b.String(); !strings.HasPrefix(got, `<svg class="a&#34;b&lt;c&amp;d"`) {
		t.Errorf("got %s", got)
	}
}

// globalnav.css turns Menu's two lines into the close button's X, each about its
// own middle, onto one point: a move of a line must move those transforms too
func TestGlobalnavCSSTurnsMenusLinesIntoAnX(t *testing.T) {
	stylesheet, err := os.ReadFile("../../../components/site/globalnav/globalnav.css")
	if err != nil {
		t.Fatal(err)
	}
	lines := regexp.MustCompile(`M(\d+) (\d+)H(\d+)`).FindAllStringSubmatch(Menu.inner, -1)
	turns := regexp.MustCompile(`transform-origin: (\d+)px (\d+)px;\s+transform: translate\((-?[\d.]+)px, (-?[\d.]+)px\) rotate`).FindAllStringSubmatch(string(stylesheet), -1)
	if len(lines) != 2 || len(turns) != 2 {
		t.Fatalf("%d lines in Menu, %d transforms in globalnav.css", len(lines), len(turns))
	}
	number := func(s string) float64 {
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			t.Fatal(err)
		}
		return n
	}
	var crossings [2][2]float64
	for i := range lines {
		middle := (number(lines[i][1]) + number(lines[i][3])) / 2
		y := number(lines[i][2])
		if number(turns[i][1]) != middle || number(turns[i][2]) != y {
			t.Errorf("line %d's middle is %v %v, its transform-origin %s %s", i, middle, y, turns[i][1], turns[i][2])
		}
		crossings[i] = [2]float64{middle + number(turns[i][3]), y + number(turns[i][4])}
	}
	if crossings[0] != crossings[1] {
		t.Errorf("the X's lines cross at no one point: %v", crossings)
	}
}

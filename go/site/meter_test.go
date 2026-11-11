package site_test

import (
	"math"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestMeterRendersTheExamples(t *testing.T) {
	example.Check(t, site.Meter{Rows: []site.MeterRow{
		{Label: "5 stars", Share: 0.82, Value: "82%"},
		{Label: "4 stars", Share: 0.09, Value: "9%"},
		{Label: "1 star", Share: 0, Value: "0%"},
	}}, "site/meter/examples/meter.html")
}

func TestMeterPanics(t *testing.T) {
	example.Panics(t, func() site.Meter {
		return site.Meter{Rows: []site.MeterRow{{Label: "", Share: 0.5, Value: "50%"}}}
	}, []example.Break[site.Meter]{
		{Want: "site.Meter needs Rows", Change: func(m *site.Meter) { m.Rows = nil }},
		{Want: "site.MeterRow's Share is not between 0 and 1: 1.5", Change: func(m *site.Meter) { m.Rows[0].Share = 1.5 }},
		{Want: "site.MeterRow's Share is not between 0 and 1: -0.1", Change: func(m *site.Meter) { m.Rows[0].Share = -0.1 }},
		{Want: "site.MeterRow's Share is not between 0 and 1: NaN", Change: func(m *site.Meter) { m.Rows[0].Share = math.NaN() }},
		{Want: "site.MeterRow needs a Value: the track is hidden from assistive technology", Change: func(m *site.Meter) { m.Rows[0].Value = "" }},
	})
}

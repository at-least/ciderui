package console_test

import (
	"os"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleConsoleBadge() {
	_ = console.ConsoleBadge{Label: "Fatal", Color: console.ConsoleColorError, Class: ""}.Render(os.Stdout)
	// Output: <span class="cui-console-badge" data-color="error">Fatal</span>
}

func ExampleConsoleBanner() {
	_ = console.ConsoleBanner{Color: console.ConsoleColorInfo, Icon: lucide.Icon{}, Title: "Saved", Class: "", Children: nil}.Render(os.Stdout)
	// Output: <div class="cui-console-banner" data-color="info" role="status"><div><p class="cui-console-banner-title">Saved</p></div></div>
}

func ExampleConsoleChart() {
	_ = console.ConsoleChart{AriaLabel: "Events: 0", Scale: nil, Empty: "No data in this window", Axis: &console.ConsoleChartAxis{Start: "Oct 22", End: "Oct 28"}, Class: "",
		Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140"></svg>`)}.Render(os.Stdout)
	// Output: <figure class="cui-console-chart" aria-label="Events: 0"><div class="cui-console-chart-plot"><p class="cui-console-chart-empty">No data in this window</p><svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140"></svg></div><div class="cui-console-chart-axis" aria-hidden="true"><span>Oct 22</span><span>Oct 28</span></div></figure>
}

func ExampleConsoleCheckboxBox() {
	_ = h.Label(h.Class("cui-console-checkbox"),
		h.Input(h.Class("cui-console-checkbox-input"), h.Type("checkbox"), h.Name("mute")),
		console.ConsoleCheckboxBox{},
		g.Text("Mute"),
	).Render(os.Stdout)
	// Output: <label class="cui-console-checkbox"><input class="cui-console-checkbox-input" type="checkbox" name="mute"><span class="cui-console-checkbox-box" aria-hidden="true"><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 6 9 17l-5-5"/></svg><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/></svg></span>Mute</label>
}

func ExampleConsoleDialog() {
	_ = console.ConsoleDialog{ID: "ignore", Title: g.Text("Ignore?"), Attrs: nil, Class: "", Children: nil}.Render(os.Stdout)
	// Output: <dialog class="cui-console-dialog" id="ignore" aria-labelledby="ignore-title"><h2 class="cui-console-dialog-title" id="ignore-title">Ignore?</h2></dialog>
}

func ExampleConsoleDisclosure() {
	_ = console.ConsoleDisclosure{Open: false, Summary: "Filters", SummaryClass: "", Class: "", Children: h.P(g.Text("Level"))}.Render(os.Stdout)
	// Output: <details class="cui-console-disclosure"><summary class="cui-console-disclosure-summary"><svg class="cui-console-disclosure-chevron" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>Filters</summary><p>Level</p></details>
}

func ExampleConsoleEmpty() {
	_ = console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Backed: false, Icon: lucide.Icon{}, Heading: "No Issues", Caption: "", Class: "", Children: nil}.Render(os.Stdout)
	// Output: <div class="cui-console-empty" data-size="sm"><p class="cui-console-empty-heading">No Issues</p></div>
}

func ExampleConsoleSelect() {
	_ = console.ConsoleSelect{Class: "", Children: h.Select(h.Class("cui-console-field"), h.Name("level"), h.Option(g.Text("Any Level")))}.Render(os.Stdout)
	// Output: <span class="cui-console-select"><select class="cui-console-field" name="level"><option>Any Level</option></select><svg class="cui-console-select-chevron" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m7 15 5 5 5-5"/><path d="m7 9 5-5 5 5"/></svg></span>
}

func ExampleConsoleListItem() {
	_ = console.ConsoleListItem{Href: "/p/shop", ID: "", Label: "Shop", Secondary: "", Class: "", Children: nil}.Render(os.Stdout)
	// Output: <a class="cui-console-list-item" href="/p/shop"><span class="cui-console-list-label">Shop</span></a>
}

func ExampleConsolePlay() {
	_ = console.ConsolePlay{State: console.ConsolePlayStatePlaying, AriaLabel: "", Label: "Play This Line", PlayingLabel: "Stop", Size: console.ConsoleBtnSizeSmall, Class: "js-clip"}.Render(os.Stdout)
	// Output: <button class="cui-console-btn cui-console-play js-clip" data-size="sm" data-shape="capsule" data-state="playing" type="button"><span class="cui-console-play-glyph"><svg class="cui-console-play-paused" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 5a2 2 0 0 1 3.008-1.728l11.997 6.998a2 2 0 0 1 .003 3.458l-12 7A2 2 0 0 1 5 19z"/></svg><svg class="cui-console-play-playing" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/></svg></span><span class="cui-console-play-paused">Play This Line</span><span class="cui-console-play-playing">Stop</span></button>
}

func ExampleConsoleSidebar() {
	_ = console.ConsoleSidebar{ID: "sidebar", AriaLabel: "CrashCart", Home: "/", Wordmark: "CrashCart", CloseAriaLabel: "Close", Class: "", Footer: nil, Children: nil}.Render(os.Stdout)
	// Output: <nav class="cui-console-sidebar" id="sidebar" popover aria-label="CrashCart"><div class="cui-console-sidebar-header"><a class="cui-console-sidebar-wordmark" href="/">CrashCart</a><button class="cui-console-btn cui-console-btn-icon cui-console-sidebar-close" data-variant="plain" type="button" aria-label="Close" popovertarget="sidebar" popovertargetaction="hide" autofocus><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg></button></div><div class="cui-console-sidebar-body cui-console-menu" data-variant="sidebar"></div><div class="cui-console-sidebar-footer"></div></nav>
}

func ExampleConsoleStatus() {
	_ = console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Class: ""}.Render(os.Stdout)
	// Output: <span class="cui-console-status" data-color="success"><span class="cui-console-status-mark"></span>Delivered</span>
}

func ExampleConsoleSectionHead() {
	_ = console.ConsoleSectionHead{Title: h.H2(h.Class("cui-console-h4-emph"), g.Text("Findings")), Class: "", Children: nil}.Render(os.Stdout)
	// Output: <header class="cui-console-section-head"><div><h2 class="cui-console-h4-emph">Findings</h2></div></header>
}

func ExampleConsoleStrip() {
	item := console.ConsoleStripItem{Href: "#shot-1", Src: "/shot-1.png", Alt: "Shot 1", Lazy: true, Current: true, Variant: console.ConsoleStripItemVariantFrame, Caption: "S01-01", Detail: "", Unused: false, Tip: "Once upon a time", Mark: console.ConsoleStripMarkWaiting, MarkAriaLabel: "Waiting for you"}
	_ = console.ConsoleStrip{AriaLabel: "Shots", Groups: []console.ConsoleStripGroup{{Label: "Scene 1", Items: []console.ConsoleStripItem{item}}}, Class: ""}.Render(os.Stdout)
	// Output: <nav class="cui-console-strip" aria-label="Shots"><div class="cui-console-strip-group"><p class="cui-console-strip-head"><span class="cui-console-strip-label">Scene 1</span></p><a class="cui-console-strip-item" href="#shot-1" aria-current="true"><img src="/shot-1.png" alt="Shot 1" loading="lazy"><span class="cui-console-strip-caption">S01-01</span><span class="cui-console-strip-tip">Once upon a time</span><span class="cui-console-strip-mark" data-state="waiting" role="img" aria-label="Waiting for you"></span></a></div></nav>
}

func ExampleConsoleTable() {
	_ = console.ConsoleTable{AriaLabel: "Jobs", Columns: []console.ConsoleTableColumn{{Label: "Job", Figure: false}, {Label: "Runs", Figure: true}},
		Rows:      []console.ConsoleTableRow{{Selected: false, Attrs: nil, Cells: []console.ConsoleTableCell{{Content: g.Text("nightly"), Empty: false}, {Content: g.Text("142"), Empty: false}}}},
		BodyAttrs: nil, Class: ""}.Render(os.Stdout)
	// Output: <div class="cui-console-table-scroll" tabindex="0" role="region" aria-label="Jobs"><table class="cui-console-table" data-stack><thead><tr><th scope="col">Job</th><th scope="col" class="cui-console-table-figure">Runs</th></tr></thead><tbody><tr><td>nightly</td><td class="cui-console-table-figure" data-label="Runs">142</td></tr></tbody></table></div>
}

func ExampleConsoleTabs() {
	tabs := []console.ConsoleTab{
		{Label: "Unresolved", Href: "?status=unresolved", Current: true},
		{Label: "Resolved", Href: "?status=resolved", Current: false},
	}
	_ = console.ConsoleTabs{AriaLabel: "Status", Variant: console.ConsoleTabsVariantFilled, Tabs: tabs, Class: ""}.Render(os.Stdout)
	// Output: <div class="cui-console-tabs" data-variant="filled" role="group" aria-label="Status"><a class="cui-console-tab" href="?status=unresolved" aria-current="true">Unresolved</a><a class="cui-console-tab" href="?status=resolved">Resolved</a></div>
}

func ExampleConsoleTabButtons() {
	buttons := []console.ConsoleTabButton{
		{Label: "Keep", Value: "keep", Pressed: true},
		{Label: "Redraw", Value: "redraw", Pressed: false},
	}
	_ = console.ConsoleTabButtons{AriaLabel: "Mark", Name: "mark", Buttons: buttons, Class: ""}.Render(os.Stdout)
	// Output: <div class="cui-console-tabs" data-variant="filled" role="group" aria-label="Mark"><button class="cui-console-tab" type="submit" name="mark" value="keep" aria-pressed="true">Keep</button><button class="cui-console-tab" type="submit" name="mark" value="redraw" aria-pressed="false">Redraw</button></div>
}

func ExampleConsoleFormField() {
	page := console.ConsoleFormField{ID: "name", Label: "Name", Help: "Shown on every page.", Error: "", Control: func(attrs g.Node) g.Node {
		return h.Input(h.Class("cui-console-field"), attrs, h.Name("name"))
	}}
	_ = page.Render(os.Stdout)
	// Output: <div class="cui-console-form-field"><label class="cui-console-label" for="name">Name</label><input class="cui-console-field" id="name" aria-describedby="name-help" name="name"><p class="cui-console-help" id="name-help">Shown on every page.</p></div>
}

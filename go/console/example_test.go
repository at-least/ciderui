package console_test

import (
	"os"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/icon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleConsoleAppearance() {
	_ = console.ConsoleAppearance{Chosen: console.ConsoleAppearanceChoice{}, Class: ""}.Render(os.Stdout)
	// Output: <fieldset class="cui-console-tabs" data-variant="filled"><legend>Appearance</legend><label class="cui-console-tab"><input type="radio" name="appearance" value="light">Light</label><label class="cui-console-tab"><input type="radio" name="appearance" value="dark">Dark</label><label class="cui-console-tab"><input type="radio" name="appearance" value="auto">Auto</label></fieldset>
}

func ExampleConsoleBadge() {
	_ = console.ConsoleBadge{Label: "Fatal", Color: console.ConsoleColorError, Variant: console.ConsoleBadgeVariantTinted, Small: false, Class: ""}.Render(os.Stdout)
	// Output: <span class="cui-console-badge" data-color="error">Fatal</span>
}

func ExampleConsoleBanner() {
	_ = console.ConsoleBanner{Color: console.ConsoleColorInfo, Icon: icon.Icon{}, Title: "Saved", DismissLabel: "", Class: "", Children: nil}.Render(os.Stdout)
	// Output: <div class="cui-console-banner" data-color="info" role="status"><div><p class="cui-console-banner-title">Saved</p></div></div>
}

func ExampleConsoleMetric() {
	_ = console.ConsoleMetric{Title: "New Issues", Value: g.Text("8"), Change: "", ChangeColor: console.ConsoleColorNeutral, Caption: "", Href: "", Class: ""}.Render(os.Stdout)
	// Output: <div class="cui-console-card"><p class="cui-console-card-title">New Issues</p><p class="cui-console-card-figure"><span class="cui-console-card-value">8</span></p></div>
}

func ExampleConsoleChart() {
	_ = console.ConsoleChart{Peak: "", Half: "", Empty: "No data in this window", AxisStart: "Oct 22", AxisEnd: "Oct 28", Legend: nil, Class: "",
		Children: g.Raw(`<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" role="img" aria-label="Events: 0"></svg>`)}.Render(os.Stdout)
	// Output: <figure class="cui-console-chart"><div class="cui-console-chart-plot"><p class="cui-console-chart-empty">No data in this window</p><svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" role="img" aria-label="Events: 0"></svg></div><div class="cui-console-chart-axis" aria-hidden="true"><span>Oct 22</span><span>Oct 28</span></div></figure>
}

func ExampleConsoleSparkline() {
	_ = console.ConsoleSparkline{Label: "3 events over 2 days", Width: 40, Height: 28, Values: []float64{1, 2}, Class: ""}.Render(os.Stdout)
	// Output: <svg class="cui-console-chart-sparkline" viewBox="0 0 40 28" width="40" height="28" role="img" aria-label="3 events over 2 days"><polygon class="cui-console-chart-sparkline-area" points="0,28 0,15 40,2 40,28"/><polyline class="cui-console-chart-sparkline-line" points="0,15 40,2"/></svg>
}

func ExampleConsoleCheckboxBox() {
	_ = h.Label(h.Class("cui-console-checkbox"),
		h.Input(h.Class("cui-console-checkbox-input"), h.Type("checkbox"), h.Name("mute")),
		console.ConsoleCheckboxBox{},
		g.Text("Mute"),
	).Render(os.Stdout)
	// Output: <label class="cui-console-checkbox"><input class="cui-console-checkbox-input" type="checkbox" name="mute"><span class="cui-console-checkbox-box" aria-hidden="true"><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 6 9 17l-5-5"/></svg><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/></svg></span>Mute</label>
}

func ExampleConsoleChip() {
	_ = console.ConsoleChip{Key: "os", Value: "iOS", Href: "", RemoveHref: "", RemoveLabel: "", Class: ""}.Render(os.Stdout)
	// Output: <span class="cui-console-chip"><span class="cui-console-chip-key">os</span>iOS</span>
}

func ExampleConsoleCopyIcons() {
	_ = h.Button(h.Type("button"), h.Class("cui-console-copy-inline"), h.Aria("label", "Copy issue ID"), console.ConsoleCopyIcons{}).Render(os.Stdout)
	// Output: <button type="button" class="cui-console-copy-inline" aria-label="Copy issue ID"><svg class="cui-console-copy-idle" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></svg><svg class="cui-console-copy-done" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 6 9 17l-5-5"/></svg></button>
}

func ExampleConsoleDialog() {
	_ = console.ConsoleDialog{ID: "ignore", Title: g.Text("Ignore?"), Open: false, Class: "", Children: nil}.Render(os.Stdout)
	// Output: <dialog class="cui-console-dialog" id="ignore" aria-labelledby="ignore-title"><h2 class="cui-console-dialog-title" id="ignore-title">Ignore?</h2></dialog>
}

func ExampleConsoleDisclosure() {
	_ = console.ConsoleDisclosure{Variant: console.ConsoleDisclosureVariantButton, Open: false, Summary: "Filters", SummaryClass: "", Class: "", Children: h.P(g.Text("Level"))}.Render(os.Stdout)
	// Output: <details class="cui-console-disclosure" data-variant="button"><summary class="cui-console-btn">Filters<svg class="cui-console-disclosure-chevron" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg></summary><p>Level</p></details>
}

func ExampleConsoleEmpty() {
	_ = console.ConsoleEmpty{Size: console.ConsoleEmptySizeMedium, Backed: false, Icon: icon.Icon{}, Heading: "No Issues", Caption: "", Class: "", Children: nil}.Render(os.Stdout)
	// Output: <div class="cui-console-empty"><p class="cui-console-empty-heading">No Issues</p></div>
}

func ExampleConsoleSelect() {
	_ = console.ConsoleSelect{Class: "", Children: h.Select(h.Class("cui-console-field"), h.Name("level"), h.Option(g.Text("Any Level")))}.Render(os.Stdout)
	// Output: <span class="cui-console-select"><select class="cui-console-field" name="level"><option>Any Level</option></select><svg class="cui-console-select-chevron" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m7 15 5 5 5-5"/><path d="m7 9 5-5 5 5"/></svg></span>
}

func ExampleConsoleFormField() {
	_ = console.ConsoleFormField{ID: "name", Label: g.Text("Name"), Children: g.Group{
		h.Input(h.Class("cui-console-field"), h.ID("name"), h.Aria("describedby", console.ConsoleHelpID("name"))),
		console.ConsoleHelp{ControlID: "name", Text: "Shown on every page."},
	}}.Render(os.Stdout)
	// Output: <div class="cui-console-form-field"><label class="cui-console-label" for="name">Name</label><input class="cui-console-field" id="name" aria-describedby="name-help"><p class="cui-console-help" id="name-help">Shown on every page.</p></div>
}

func ExampleConsoleHead() {
	tabs := []console.ConsoleTab{{Label: "Overview", Href: "/p/shop", Current: true}}
	_ = console.ConsoleHead{Name: "Shop", Href: "/p/shop", Label: "Project", Tabs: tabs, Class: ""}.Render(os.Stdout)
	// Output: <div class="cui-console-head"><a class="cui-console-head-name" href="/p/shop">Shop</a><nav class="cui-console-tabs" aria-label="Project"><a class="cui-console-tab" href="/p/shop" aria-current="page">Overview</a></nav></div>
}

func ExampleConsoleHint() {
	_ = console.ConsoleHint{ID: "hint-a", Label: "What A means", Hint: "A is a."}.Render(os.Stdout)
	// Output: <button type="button" class="cui-console-hint" popovertarget="hint-a" aria-label="What A means">i</button><span class="cui-console-popover" id="hint-a" popover>A is a.</span>
}

func ExampleConsoleListItem() {
	_ = console.ConsoleListItem{Href: "/p/shop", ID: "", Label: "Shop", Secondary: "", Class: "", Children: nil}.Render(os.Stdout)
	// Output: <a class="cui-console-list-item" href="/p/shop"><span class="cui-console-list-label">Shop</span></a>
}

func ExampleConsoleListing() {
	_ = console.ConsoleListing{FirstLine: 41, Lines: []console.ConsoleListingLine{{Text: "let a = 1;", Highlighted: false}, {Text: "throw a;", Highlighted: true}}, Class: ""}.Render(os.Stdout)
	// Output:
	// <pre class="cui-console-listing"><code><span class="cui-console-listing-line" data-line="41">let a = 1;</span>
	// <span class="cui-console-listing-line" data-line="42" data-highlighted>throw a;</span></code></pre>
}

func ExampleConsoleMenu() {
	_ = console.ConsoleMenu{Label: "Analytics", Class: "", Children: console.ConsoleMenuItem{Href: "/overview", Label: "Overview", Icon: icon.Icon{}, Current: true, Disabled: false, Secondary: "", Class: ""}}.Render(os.Stdout)
	// Output: <nav class="cui-console-menu" aria-label="Analytics"><a class="cui-console-menu-item" href="/overview" aria-current="page"><span class="cui-console-menu-label">Overview</span></a></nav>
}

func ExampleConsolePager() {
	_ = console.ConsolePager{Label: "Issue pages", Summary: "1–50 of 230", First: "", Previous: "", Next: "?page=2", Class: ""}.Render(os.Stdout)
	// Output: <nav class="cui-console-pager" aria-label="Issue pages"><p class="cui-console-pager-summary">1–50 of 230</p><a class="cui-console-btn cui-console-btn-icon" href="?page=2" aria-label="Next page"><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg></a></nav>
}

func ExampleConsoleRadio() {
	_ = console.ConsoleRadio{Name: "kind", Value: "app", Checked: true, Required: false, Disabled: false, Caption: "", Class: "", Children: g.Text("App")}.Render(os.Stdout)
	// Output: <label class="cui-console-radio"><input class="cui-console-radio-input" type="radio" name="kind" value="app" checked>App</label>
}

func ExampleConsoleSidebar() {
	_ = console.ConsoleSidebar{ID: "sidebar", Label: "CrashCart", Home: "/", Wordmark: "CrashCart", CloseLabel: "Close", Class: "", Footer: nil, Children: nil}.Render(os.Stdout)
	// Output: <nav class="cui-console-sidebar" id="sidebar" popover aria-label="CrashCart"><div class="cui-console-sidebar-header"><a class="cui-console-sidebar-wordmark" href="/">CrashCart</a><button class="cui-console-btn cui-console-btn-icon cui-console-sidebar-close" data-variant="plain" type="button" popovertarget="sidebar" popovertargetaction="hide" aria-label="Close"><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg></button></div><div class="cui-console-sidebar-body cui-console-menu" data-variant="sidebar"></div><div class="cui-console-sidebar-footer"></div></nav>
}

func ExampleConsoleStatus() {
	_ = console.ConsoleStatus{Label: "Delivered", Color: console.ConsoleColorSuccess, Series: "", Class: ""}.Render(os.Stdout)
	// Output: <span class="cui-console-status" data-color="success"><span class="cui-console-status-mark"></span>Delivered</span>
}

func ExampleConsoleTabs() {
	tabs := []console.ConsoleTab{
		{Label: "Unresolved", Href: "?status=unresolved", Current: true},
		{Label: "Resolved", Href: "?status=resolved", Current: false},
	}
	_ = console.ConsoleTabs{Label: "Status", Variant: console.ConsoleTabsVariantFilled, Tabs: tabs, Class: ""}.Render(os.Stdout)
	// Output: <div class="cui-console-tabs" data-variant="filled" role="group" aria-label="Status"><a class="cui-console-tab" href="?status=unresolved" aria-current="true">Unresolved</a><a class="cui-console-tab" href="?status=resolved">Resolved</a></div>
}

func ExampleConsoleThumb() {
	_ = console.ConsoleThumb{Href: "/shot.png", Src: "/shot.png", Alt: "shot.png", Caption: "shot.png", Lazy: true, Class: ""}.Render(os.Stdout)
	// Output: <figure class="cui-console-thumb"><a href="/shot.png"><img src="/shot.png" alt="shot.png" loading="lazy"></a><figcaption class="cui-console-thumb-caption">shot.png</figcaption></figure>
}

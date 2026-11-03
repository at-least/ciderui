+++
title = "Console Tabs"
description = "App Store Connect’s tab bar: underlined for a page’s sections, filled for a choice of a few."

[extra]
category = "Console"
image = "pictures/console-tabs.png"
+++

App Store Connect’s tab bar, in its two forms: underlined, a page’s sections, the current one on a rule; filled, a choice of a few views of one thing (a chart’s period, a list’s status), each a pill, the chosen one filled, as Apple Developer Forums’ Topics and All Posts and the HIG’s segmented control. Each tab is the page’s own link. On a phone, an underlined bar too wide for the screen stays one row and scrolls sideways, the side still hiding content fading; a tab focused near its end is brought in clear of the fade.

```css
@import "ciderui/components/console/console-tabs/console-tabs.css";
```

{{ <cui_example file="console-tabs/tabs" title="Underlined tabs" pad={true} /> }}

{{ <cui_example file="console-tabs/many" title="A project's six sections: one scrolling row on a phone" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-tabs` | the bar: a `<nav>` with an `aria-label`, a group, or a `<fieldset>` of radios with its `<legend>` |
| `data-variant="filled"` | pills, for a choice of a few |
| `data-size="sm"`, `"lg"`, `"xl"` | 22px, 40px, 50px |
| `data-overflow`, `data-overflow-start`, `data-overflow-end` | the host’s script marks the bar’s scroll state; the side still hiding content fades (the phone row has no scrollbar) |
| `cui-console-tab` | a tab, a link; or a `<label>` round its radio, checked for the chosen one |
| `aria-current` | the current one (`"page"` for a page, `"true"` for a choice) |
| `aria-disabled="true"` | not available now |

## Filled

{{ <cui_example file="console-tabs/filled" title="Filled tabs" pad={true} /> }}

## A choice the form sends

Filled tabs of radios, where the choice is a setting rather than a page: Light, Dark and Auto, for a site that keeps its own appearance. Apple’s Human Interface Guidelines advise against an app-specific appearance setting; follow the system’s where you can. The radios are real, so the keyboard moves between them; the page’s script applies the choice and remembers it.

{{ <cui_example file="console-tabs/appearance" title="Light chosen" pad={true} /> }}

## Large

{{ <cui_example file="console-tabs/large" title="Large tabs" pad={true} /> }}

The Go module writes either with `console.ConsoleTabs` and its `console.ConsoleTab`s, and the appearance choice with `console.ConsoleAppearance` and `console.ConsoleAppearanceChoice`.

+++
title = "Console Tabs"
description = "App Store Connect’s tab bar: underlined for a page’s sections, filled for a choice of a few."

[extra]
category = "Console"
image = "pictures/console-tabs.png"
+++

App Store Connect’s tab bar, in its two forms: underlined, a page’s sections, the current one on a rule; filled, a choice of a few views of one thing (a chart’s period, a list’s status), each a pill, the chosen one filled, as Apple Developer Forums’ Topics and All Posts and the HIG’s segmented control. Each tab is the page’s own link. On a phone, an underlined bar too wide for the screen stays one row and scrolls sideways, the side still hiding content fading; a tab focused near its end is brought in clear of the fade.

```css
@import "ciderui/components/console-tabs/console-tabs.css";
```

{{ <cui_example file="console-tabs/tabs" title="Underlined tabs" pad={true} /> }}

{{ <cui_example file="console-tabs/many" title="A project's six sections: one scrolling row on a phone" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-tabs` | the bar: a `<nav>` with an `aria-label`, or a group |
| `data-variant="filled"` | pills, for a choice of a few |
| `data-size="sm"`, `"lg"`, `"xl"` | 22px, 40px, 50px |
| `data-overflow`, `data-overflow-start`, `data-overflow-end` | the host’s script marks the bar’s scroll state; the side still hiding content fades (the phone row has no scrollbar) |
| `cui-console-tab` | a tab, a link |
| `aria-current` | the current one (`"page"` for a page, `"true"` for a choice) |
| `aria-disabled="true"` | not available now |

## Filled

{{ <cui_example file="console-tabs/filled" title="Filled tabs" pad={true} /> }}

## Large

{{ <cui_example file="console-tabs/large" title="Large tabs" pad={true} /> }}

The crate writes either with `ConsoleTabs` and its `ConsoleTab`s.

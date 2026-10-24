+++
title = "Console Tabs"
description = "App Store Connect’s tab bar: underlined for a page’s sections, filled for a choice of a few."

[extra]
category = "Console"
image = "pictures/console-tabs.png"
+++

App Store Connect’s tab bar, in its two forms: underlined, a page’s sections, the current one on a rule; filled, a choice of a few (a chart’s period), each a pill, the chosen one filled. Each tab is the page’s own link.

```css
@import "ciderui/components/console-tabs/console-tabs.css";
```

{{ <cui_example file="console-tabs/tabs" title="Underlined tabs" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-tabs` | the bar: a `<nav>` with an `aria-label`, or a group |
| `data-variant="filled"` | pills, for a choice of a few |
| `data-size="sm"`, `"lg"`, `"xl"` | 22px, 40px, 50px |
| `cui-console-tab` | a tab, a link |
| `aria-current` | the current one (`"page"` for a page, `"true"` for a choice) |
| `aria-disabled="true"` | not available now |

## Filled

{{ <cui_example file="console-tabs/filled" title="Filled tabs" pad={true} /> }}

## Large

{{ <cui_example file="console-tabs/large" title="Large tabs" pad={true} /> }}

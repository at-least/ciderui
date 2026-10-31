+++
title = "Console Head"
description = "What a page belongs to by name, its tabs beside it."

[extra]
category = "Console"
image = "pictures/console-head.png"
+++

The top of a page that belongs to something — an app, a project: its name, large, and its tabs on the same row, as App Store Connect heads an app’s pages. The tabs say which page this is, so the page needs no second title for it. On a phone the tabs take a row of their own under the name.

```css
@import "ciderui/components/console/console-tabs/console-tabs.css";
@import "ciderui/components/console/console-head/console-head.css";
```

{{ <cui_example file="console-head/head" title="A project’s head, Issues current" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-head` | the row: the name, then the tabs |
| `cui-console-head-name` | the name, a link to the thing’s first page |

The tabs are `cui-console-tabs`, underlined; in the head they take App Store Connect’s app tabs’ size.

The crate writes it with `ConsoleHead`.

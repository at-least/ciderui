+++
title = "Console Hint"
description = "App Store Connect’s info button: a small “i” beside a word that opens what the word means."

[extra]
category = "Console"
image = "pictures/console-hint.png"
+++

A small “i” in a gray circle beside a word — a metric’s title, a column’s heading — that opens a popover saying what the word means. The popover is the button’s `popovertarget`: a click or a tap opens it, Escape or a click elsewhere closes it, with no script.

```css
@import "ciderui/components/console-hint/console-hint.css";
@import "ciderui/components/console-popover/console-popover.css";
```

{{ <cui_example file="console-hint/hint" title="Console hint" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-hint` | the `<button>` with `popovertarget` and an `aria-label` saying what it is about |
| `cui-console-popover` | the `popover` it opens (Console Popover): 16px in, 13px, 240px to 350px wide, under the button |

The crate’s `ConsoleHint` renders both.

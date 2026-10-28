+++
title = "Console Disclosure"
description = "A line or a button that shows what is hidden until it is needed."

[extra]
category = "Console"
image = "pictures/console-disclosure.png"
+++

A `<details>` that hides what is not needed yet — a stack’s system frames, a frame’s source, a phone’s filters — behind a line or a button that says what it hides. As the HIG’s disclosure controls: a line’s chevron sits at its leading edge, pointing at the label while closed and down while open; a button’s points down while closed and up while open. No script: the browser opens and closes it.

```css
@import "ciderui/components/console-disclosure/console-disclosure.css";
```

{{ <cui_example file="console-disclosure/disclosure" title="Disclosure lines, closed and open" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-disclosure` | the `<details>` |
| `cui-console-disclosure-summary` | the line: its chevron, then its label |
| `cui-console-disclosure-chevron` | the chevron: Lucide’s `chevron-right` before a line’s label, `chevron-down` after a button’s |
| `data-variant="button"` | its summary a `cui-console-btn` |

## A disclosure button

{{ <cui_example file="console-disclosure/button" title="A disclosure button, closed and open" pad={true} /> }}

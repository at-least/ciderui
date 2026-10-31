+++
title = "Console Checkbox"
description = "App Store Connect’s checkbox: a small rounded box before its label."

[extra]
category = "Console"
image = "pictures/console-checkbox.png"
+++

A 15px rounded box before its label, the input kept over it, transparent, so it takes the focus, the click and the space bar; checked, the brand’s blue with a white check. Where no label shows (a table row’s), the input carries an `aria-label`.

```css
@import "ciderui/components/console/console-checkbox/console-checkbox.css";
```

{{ <cui_example file="console-checkbox/checkbox" title="Console checkboxes" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-checkbox` | the checkbox and its label: a `<label>`, or a `<span>` |
| `cui-console-checkbox-input` | the page’s own `<input type="checkbox">`, first |
| `cui-console-checkbox-box` | the box after it, its check and its dash |

The crate writes the box with `ConsoleCheckboxBox`.

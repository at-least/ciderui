+++
title = "Field"
description = "apple.com’s 56px text field, its label the placeholder until it is filled."

[extra]
category = "Controls"
image = "pictures/field.png"
+++

apple.com’s own text field, the store’s: 56px, a 1px edge and 12px corners on a translucent white ground. Its label sits where a placeholder would, 17px in the secondary text, and moves to the field’s top at 12px as the field is focused or filled, in 0.125s, with no script.

```css
@import "ciderui/components/field/field.css";
```

{{ <cui_example file="field/field" title="An empty field" pad={true} /> }}

{{ <cui_example file="field/filled" title="A filled field" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-field` | the field |
| `cui-field-input` | the `<input>`; its placeholder a space, `placeholder=" "`, so the label can tell it empty |
| `cui-field-label` | its `<label>`, after the input |

Focused, the field shows a 2px accent edge in place of the focus ring: a text field is `:focus-visible` however it is focused, so a ring cannot show for the keyboard alone, as apple.com’s script’s does. In a dark block it keeps its light look on a ground of its own, as apple.com’s field does in a dark color scheme. The crate writes it with `Field`.

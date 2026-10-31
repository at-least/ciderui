+++
title = "Field"
description = "apple.com’s 56px text field, its label the placeholder until it is filled; in error, disabled, with a currency before it, a message of several lines, and the store’s dropdown."

[extra]
category = "Controls"
image = "pictures/field.png"
+++

apple.com’s own text field, the store’s: 56px, a 1px edge and 12px corners on a translucent white ground. Its label sits where a placeholder would, 17px in the secondary text, and moves to the field’s top at 12px as the field is focused or filled, in 0.125s, with no script.

```css
@import "ciderui/components/site/field/field.css";
```

{{ <cui_example file="field/field" title="An empty field" pad={true} /> }}

{{ <cui_example file="field/filled" title="A filled field" pad={true} /> }}

A field in error has a red edge and label on a pink ground, and its message under it, 12px in red after a mark. The error is the page’s to set, with its message, as its server answers the form: `aria-invalid="true"` on the input, and `aria-describedby` naming the message. A disabled field is grey, its edge faint. A currency or a country code can sit at the field’s start, on a faint ground of its own with a rule after it, its text and label moved past it.

{{ <cui_example file="field/error" title="A field in error and its message" pad={true} /> }}

{{ <cui_example file="field/disabled" title="A disabled field" pad={true} /> }}

{{ <cui_example file="field/prefix" title="A field with a currency before it" pad={true} /> }}

A message of several lines is a `<textarea>` in the field’s box, `cui-field-box`: three lines tall, the label in the room above them, taller or shorter by its handle, the box growing with it. Its states are the input’s: its placeholder a space, `aria-invalid="true"`, `disabled`.

{{ <cui_example file="field/textarea" title="A message of several lines" pad={true} /> }}

The store’s dropdown is a `<select>` with the same look and a chevron at its end. Its first choice is an empty one (`value=""`), chosen while nothing is: its label is then the placeholder and its own text is not shown; a choice made, the label moves to the top.

{{ <cui_example file="field/select" title="A dropdown, nothing chosen" pad={true} /> }}

{{ <cui_example file="field/selected" title="A dropdown with a choice made" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-field` | the field |
| `cui-field-input` | the `<input>` or the `<select>`; an input’s placeholder a space, `placeholder=" "`, so the label can tell it empty; a select’s empty choice first, `value=""`; `aria-invalid="true"` in error; `disabled` |
| `cui-field-label` | its `<label>`, after the input |
| `cui-field-box` | a textarea’s box, with the field’s edge, corners and ground, focused as the textarea is |
| `cui-field-textarea` | the `<textarea>`, in the box; its placeholder a space; `aria-invalid="true"`; `disabled` |
| `cui-field-prefix` | a currency or a country code, after the label; the input’s `aria-describedby` names it |
| `cui-field-chevron` | a select’s chevron, `icon::DROPDOWN`, after the label |
| `cui-field-message` | its error, after the field, its mark `icon::ALERT` |

Focused, the field shows a 2px accent edge in place of the focus ring: a text field is `:focus-visible` however it is focused, so a ring cannot show for the keyboard alone, as apple.com’s script’s does; in error too. In a dark block it keeps its light look on a ground of its own, as apple.com’s field does in a dark color scheme; its message there, on the block’s ground, is a lighter red, as apple.com’s is below 4.5:1 on black. The crate writes it with `Field`, the textarea with `TextArea` and the dropdown with `Select`.

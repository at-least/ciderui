+++
title = "Console Field"
description = "App Store Connect’s text field, dropdown and multi-line field, with their labels."

[extra]
category = "Console"
image = "pictures/console-field.png"
+++

App Store Connect’s fields: a text field, a dropdown and a multi-line field, each a rounded rectangle on white inside its edge; a label over one, help under it, an error in its place. Each is the page’s own `<input>`, `<select>` or `<textarea>` with the class.

```css
@import "ciderui/components/console/console-field/console-field.css";
```

{{ <cui_example file="console-field/fields" title="Console fields" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-field` | an `<input>`, `<select>` or `<textarea>`: 30px, 13px |
| `aria-invalid="true"` | in error: a red edge on a faint red |
| `disabled` | not available now |
| `data-size="sm"`, `"lg"`, `"xl"` | 22px, 40px, 50px |
| `cui-console-select` | a dropdown’s box, with its chevron (`cui-console-select-chevron`) |
| `cui-console-label` | a field’s label |
| `cui-console-help` | a line under a field, its `aria-describedby` |
| `cui-console-error` | a field’s error, its `aria-describedby` |
| `cui-console-form-field` | a label, its field and its lines |

## Dropdown

{{ <cui_example file="console-field/select" title="A console dropdown" pad={true} /> }}

## Sizes

{{ <cui_example file="console-field/sizes" title="Console field sizes" pad={true} /> }}

The crate writes a dropdown’s box with `ConsoleSelect`, and a field’s box, label, help and error with `ConsoleFormField`, `ConsoleHelp` and `ConsoleError`.

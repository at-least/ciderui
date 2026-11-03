+++
title = "Console Field"
description = "App Store Connect’s text field, dropdown, multi-line field and file field, with their labels."

[extra]
category = "Console"
image = "pictures/console-field.png"
+++

App Store Connect’s fields: a text field, a dropdown, a multi-line field and a file field, each a rounded rectangle on white inside its edge; a label over one, help under it, an error in its place. Each is the page’s own `<input>`, `<select>` or `<textarea>` with the class.

```css
@import "ciderui/components/console/console-field/console-field.css";
```

{{ <cui_example file="console-field/fields" title="Console fields" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-field` | an `<input>`, `<select>` or `<textarea>`: 30px, 13px; on an `<input type="file">`, its button the blue word |
| `aria-invalid="true"` | in error: a red edge on a faint red |
| `disabled` | not available now |
| `data-size="sm"`, `"lg"`, `"xl"` | 22px, 40px, 50px |
| `cui-console-select` | a dropdown’s box, with its chevron (`cui-console-select-chevron`) |
| `cui-console-label` | a field’s label |
| `cui-console-help` | a line under a field, its `aria-describedby` |
| `cui-console-error` | a field’s error, its `aria-describedby`; empty and hidden (`role="alert"`), the slot a script fills when a request fails |
| `cui-console-form-field` | a label, its field and its lines |

## Dropdown

{{ <cui_example file="console-field/select" title="A console dropdown" pad={true} /> }}

## File

A file to send: its button the blue word that chooses a file, the file’s name after it.

{{ <cui_example file="console-field/file" title="Console file fields" pad={true} /> }}

## Sizes

{{ <cui_example file="console-field/sizes" title="Console field sizes" pad={true} /> }}

The Go module writes a dropdown’s box with `console.ConsoleSelect`, and a field’s box, label, help and error with `console.ConsoleFormField`, `console.ConsoleHelp` and `console.ConsoleError`.

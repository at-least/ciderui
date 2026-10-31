+++
title = "Console Button"
description = "App Store Connect’s button: gray, filled, plain or red, in four sizes."

[extra]
category = "Console"
image = "pictures/console-btn.png"
+++

App Store Connect’s button: a rounded rectangle, gray by default, filled for the page’s one main action, plain as a link is, red where it destroys. It is a class on the page’s own `<button>` or `<a>`.

```css
@import "ciderui/components/console/console-btn/console-btn.css";
```

{{ <cui_example file="console-btn/variants" title="Console buttons" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-btn` | the button, gray: 30px, 13px at 500 |
| `data-variant="filled"` | the brand’s fill, for the page’s main action |
| `data-variant="plain"` | no fill, the brand’s label, as a link |
| `data-color="error"` | red, for an action that destroys |
| `disabled`, `aria-disabled="true"` | not available now |
| `cui-console-btn-icon` | an icon alone, square; its label its `aria-label` |

## Disabled

{{ <cui_example file="console-btn/disabled" title="Disabled console buttons" pad={true} /> }}

## Sizes

`data-size="sm"`, `"lg"` and `"xl"`: 22px, 40px and 50px.

{{ <cui_example file="console-btn/sizes" title="Console button sizes" pad={true} /> }}

## Icons

An icon, any of `ciderui::lucide`’s, before the label, or alone with `cui-console-btn-icon`.

{{ <cui_example file="console-btn/icon" title="Console buttons with icons" pad={true} /> }}

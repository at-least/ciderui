+++
title = "Console Badge"
description = "App Store Connect’s badge: a short word on a pill tinted in its color."

[extra]
category = "Console"
image = "pictures/console-badge.png"
+++

A short word on a pill tinted in its color: a level, a status, a count. Or on the gray fill, or plain, its color alone. A colored mark before a word is a [Console Status](@/components/console-status.md).

```css
@import "ciderui/components/console/console-badge/console-badge.css";
```

{{ <cui_example file="console-badge/badges" title="Console badges" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-badge` | the badge, a `<span>`: 20px, 13px at 600, tinted |
| `data-color` | `brand`, `info`, `success`, `warning`, `error`; none for the neutral |
| `data-variant="gray"`, `"plain"` | on the gray fill; or no fill |
| `data-size="sm"` | 18px, 11px |

## Gray

{{ <cui_example file="console-badge/gray" title="Gray console badges" pad={true} /> }}

## Plain

{{ <cui_example file="console-badge/plain" title="Plain console badges" pad={true} /> }}

## Small

{{ <cui_example file="console-badge/small" title="Small console badges" pad={true} /> }}

The Go module writes it with `console.ConsoleBadge`, its color a `console.ConsoleColor`.

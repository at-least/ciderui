+++
title = "Console Listing"
description = "Lines of code or of a log on a gray panel, as Apple’s developer documentation shows code."

[extra]
category = "Console"
image = "pictures/console-listing.png"
+++

Lines of code or of a log — a crash’s backtrace, a job’s output — on a gray rounded panel, as Apple’s developer documentation shows a code listing, at the console’s size. A line too long for the panel moves sideways inside it. On a `<pre>` the lines stay as written; on another element, they wrap.

```css
@import "ciderui/components/console/console-listing/console-listing.css";
```

{{ <cui_example file="console-listing/listing" title="A backtrace" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-listing` | the panel: 15px corners, 8px by 14px in, 13px on 22px in the mono face |
| `cui-console-listing-line` | a numbered line: its number from `data-line`, drawn before it, never selected or copied |
| `data-highlighted` | on a line: the line to look at, tinted, an edge at its number |

## Numbered lines

A frame’s source, its own line highlighted, as Apple’s tutorials mark the lines that changed.

{{ <cui_example file="console-listing/lines" title="Source lines, line 42 highlighted" pad={true} /> }}

## Inside a card

Inside a card, or inside another listing, whose gray is its own, a listing takes the page’s white.

{{ <cui_example file="console-listing/nested" title="A listing in a console card" pad={true} /> }}

Inline code in a line of text is `cui-console-code`.

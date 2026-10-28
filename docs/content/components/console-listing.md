+++
title = "Console Listing"
description = "Lines of code or of a log on a gray panel, as Apple’s developer documentation shows code."

[extra]
category = "Console"
image = "pictures/console-listing.png"
+++

Lines of code or of a log — a crash’s backtrace, a job’s output — on a gray rounded panel, as Apple’s developer documentation shows a code listing, at the console’s size. A line too long for the panel moves sideways inside it. On a `<pre>` the lines stay as written; on another element, they wrap.

```css
@import "ciderui/components/console-listing/console-listing.css";
```

{{ <cui_example file="console-listing/listing" title="A backtrace" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-listing` | the panel: 15px corners, 8px by 14px in, 13px on 22px in the mono face |

Inline code in a line of text is `cui-console-code`.

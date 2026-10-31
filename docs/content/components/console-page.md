+++
title = "Console Page"
description = "App Store Connect’s page: its ground, its text and its type, for a console."

[extra]
category = "Console"
image = "pictures/console-page.png"
+++

The page of a site that is a console, as App Store Connect’s newer look draws one: its ground, its text at 13px on 18px, and its type. A console’s parts (`cui-console-*`) are App Store Connect’s, not apple.com’s: their colors are the console’s own (`--cui-c-console-*`), and dark is derived from the toolkit’s dark palette.

```css
@import "ciderui/components/console/console-page/console-page.css";
```

{{ <cui_example file="console-page/page" title="A console page" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-page` | the page, on `<body>`: the console’s ground and text, 13px on 18px; 16px on 21px below 640px |

## Type

Each text style is a utility, App Store Connect’s of the same name: plain at 400, `-medium` at 500, `-emph` at its emphasis.

{{ <cui_example file="console-page/type" title="The console’s type" pad={true} /> }}

| utility | size |
| --- | --- |
| `cui-console-h1-plus`, `-emph` | 41px on 45px, a metric’s value |
| `cui-console-h1`, `-emph` | 34px on 41px |
| `cui-console-h2` | 26px on 30px at 700 |
| `cui-console-h3`, `-medium`, `-emph` | 20px on 24px |
| `cui-console-h4`, `-medium`, `-emph` | 18px on 22px, a section’s heading |
| `cui-console-h5`, `-medium`, `-emph` | 15px on 20px |
| `cui-console-h6`, `-medium`, `-emph` | 14px on 19px |
| `cui-console-body`, `-medium`, `-emph` | 13px on 18px |
| `cui-console-small`, `-medium`, `-emph` | 12px on 16px |
| `cui-console-xs`, `-medium`, `-emph` | 11px on 13px |
| `cui-console-mono` | code, 13px on 18px |

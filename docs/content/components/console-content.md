+++
title = "Console Content"
description = "App Store Connect’s page column: at most 1440px, centered, the sidebar inside it."

[extra]
category = "Console"
image = "pictures/console-content.png"
+++

The page’s one column, at most 1440px wide in the middle of the window, its sides 16px on a phone, 24px from 640px and 40px from 1024px. From 1024px a sidebar inside it stands in a 225px column at its start, 30px before the page, as App Store Connect’s menu does; below 1024px the sidebar is a sheet.

```css
@import "ciderui/components/console-content/console-content.css";
```

{{ <cui_example file="console-content/content" title="Console content" /> }}

| class | what |
| --- | --- |
| `cui-console-content` | the column; with a `cui-console-sidebar` child, two columns from 1024px |

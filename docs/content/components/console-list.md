+++
title = "Console List"
description = "App Store Connect’s grouped list: rows on a gray card, a hairline between them."

[extra]
category = "Console"
image = "pictures/console-list.png"
+++

Rows grouped on the card’s gray, a hairline between them from 16px in. A row that is a link fills when hovered; any row can hold the page’s own content (its controls, a chart).

```css
@import "ciderui/components/console-list/console-list.css";
```

{{ <cui_example file="console-list/list" title="Console list" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-list` | the group: the card’s gray, 10px corners |
| `cui-console-list-item` | a row (`<a>` or any element): at least 44px, 16px in |
| `cui-console-list-label` | its title, 13px at 500 (16px on a phone) |
| `cui-console-list-secondary` | what follows it, in the secondary label |

+++
title = "Console Bar"
description = "App Store Connect’s bar: the band across the top of a console’s page."

[extra]
category = "Console"
image = "pictures/console-bar.png"
+++

App Store Connect’s bar: the band across the top of a console’s page, 50px, on a light gray with no edge, holding the product’s wordmark and a few controls — the sidebar’s button on a narrow screen, a sign-in or sign-out button. Its place on the page is the page’s: a site adds `sticky top-0` to keep it in view.

```css
@import "ciderui/components/console-bar/console-bar.css";
```

{{ <cui_example file="console-bar/bar" title="A console bar" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-bar` | the bar: a row, 50px, 16px in (24px from 640px, 40px from 1024px), its parts 8px apart |

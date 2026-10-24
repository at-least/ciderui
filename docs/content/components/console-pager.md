+++
title = "Console Pager"
description = "Under a long list, what it shows and the buttons to the pages either side."

[extra]
category = "Console"
image = "pictures/console-pager.png"
+++

Under a long list: a line saying what it shows, and the buttons to the pages either side, a disabled one where there is no page that way.

```css
@import "ciderui/components/console-pager/console-pager.css";
```

{{ <cui_example file="console-pager/pager" title="Console pager" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-pager` | a `<nav>` named by its `aria-label`: the summary, then the buttons |
| `cui-console-pager-summary` | what the list shows, 13px, secondary |
| `cui-console-pager-actions` | the `cui-console-btn` buttons, 8px apart |

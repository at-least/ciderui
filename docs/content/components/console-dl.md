+++
title = "Console Description List"
description = "App Store Connect’s description list: an item’s facts, each a label and its value."

[extra]
category = "Console"
image = "pictures/console-dl.png"
+++

An item’s facts: each label beside its value, the labels in a column as wide as the longest, or each label over its value.

```css
@import "ciderui/components/console-dl/console-dl.css";
```

{{ <cui_example file="console-dl/dl" title="Console description list" pad={true} /> }}

{{ <cui_example file="console-dl/vertical" title="Labels over their values" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-dl` | the `<dl>`: a label (`dt`) beside its value (`dd`), 8px apart; 13px |
| `data-orientation="vertical"` | each label over its value, 24px under a value |
| `data-variant="strong"` | the labels in the label’s color at 600 |

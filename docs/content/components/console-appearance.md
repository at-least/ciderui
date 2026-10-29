+++
title = "Console Appearance"
description = "Light, Dark and Auto as one segmented control."

[extra]
category = "Console"
image = "pictures/console-appearance.png"
+++

Light, Dark and Auto as one segmented control, as Apple’s developer documentation offers in its footer, for a site that keeps its own appearance. Apple’s Human Interface Guidelines advise against an app-specific appearance setting; follow the system’s where you can. The radios are real, so the keyboard moves between them; the page’s script applies the choice and remembers it.

```css
@import "ciderui/components/console-appearance/console-appearance.css";
```

{{ <cui_example file="console-appearance/appearance" title="Light chosen" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-console-appearance` | the `<fieldset>`; its `<legend>` names it for a screen reader |
| `cui-console-appearance-choice` | a choice: the `<label>` around its radio and its word (a `<span>`) |

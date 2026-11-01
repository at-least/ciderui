+++
title = "Docs Content"
description = "An article’s prose as Apple’s developer documentation draws it, markdown’s elements styled as they come."

[extra]
category = "Documentation"
image = "pictures/docs-content.png"
+++

The prose of an article in Apple’s developer documentation: its headings, text, lists, links and inline code, its code listings and tables, each as DocC draws them. The plain elements markdown writes are styled as they come, with no class on them.

```css
@import "ciderui/components/docs/docs-content/docs-content.css";
```

{{ <cui_example file="docs-content/content" title="An article’s prose" pad={true} /> }}

| element | what |
| --- | --- |
| `h2`, `h3` | a section’s heading and a smaller one, at 600, larger from 736px and again from 1251px |
| `p`, `ul`, `ol` | the text, 17px on 25px, paragraphs and list items 13.6px apart |
| `a`, `code`, `em` | a link in the link’s blue, underlined; inline code in the mono face; emphasis |
| `pre` | a code listing on its gray, 15px corners, a long line moving sideways inside it |
| `table` | every cell edged in the divider’s color, the table moving sideways where the column is narrower |
| `img` | at its own size, centered |

A part inside it, a class named `cui-*`, keeps its own look: these rules match plain elements alone.

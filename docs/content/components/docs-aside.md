+++
title = "Docs Aside"
description = "A note set apart from an article’s text, as Apple’s developer documentation draws it: six kinds, each in its colors."

[extra]
category = "Documentation"
image = "pictures/docs-aside.png"
+++

A note set apart from an article’s text, its kind’s word over it: Apple’s developer documentation’s aside. Inside a [docs content](@/components/docs-content.md) it takes its room above.

```css
@import "ciderui/components/docs/docs-aside/docs-aside.css";
```

{{ <cui_example file="docs-aside/note" title="A note" pad={true} /> }}

## Its kinds

Important, warning, tip, deprecated and experiment, each a `data-variant`; a note has none.

{{ <cui_example file="docs-aside/kinds" title="The aside’s other kinds" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-docs-aside` | the aside: 16px in, a 1px edge, 15px corners, on its kind’s ground |
| `cui-docs-aside-label` | its kind’s word at 600, in the edge’s color |
| `data-variant` | `important`, `warning`, `tip`, `deprecated`, `experiment`; a note without one |

The crate writes an aside, its label and its accessible name with `DocsAside` and its `DocsAsideVariant`. On a page in ciderui-zola’s docs layout, markdown’s alert is an aside of its kind:

```markdown
> [!NOTE]
> A report from a beta build carries its beta identifier.
```

`[!IMPORTANT]`, `[!TIP]` and `[!WARNING]` likewise; deprecated and experiment have no alert of their own.

+++
title = "Docs Topics"
description = "Apple’s developer documentation’s Topics and See Also: groups of links to other pages, each with its abstract."

[extra]
category = "Documentation"
image = "pictures/docs-topics.png"
+++

After a page’s content, Apple’s developer documentation lists the pages under it (Topics) or beside it (See Also): a title over groups of links, each after its glyph, each with what the page is about under it.

```css
@import "ciderui/components/docs/docs-topics/docs-topics.css";
```

{{ <cui_example file="docs-topics/topics" title="A page’s See Also" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-docs-topics` | the section |
| `cui-docs-topics-title` | its title, as a section’s heading |
| `cui-docs-topics-group` | a group, 40px down |
| `cui-docs-topics-heading` | a group’s title |
| `cui-docs-topics-item` | a page: its link (`cui-docs-topics-link`), its glyph (`cui-docs-topics-icon`) and its abstract (`cui-docs-topics-abstract`) |

The crate writes the section, each link’s glyph (an article’s, or a collection’s for a page of pages) and its abstract with `DocsTopics`, its `DocsTopicsGroup`s and `DocsTopic`s.

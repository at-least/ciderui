+++
title = "Docs Hero"
description = "The top of a page in Apple’s developer documentation: its breadcrumbs, its eyebrow, its title and its abstract."

[extra]
category = "Documentation"
image = "pictures/docs-hero.png"
+++

The top of a page in Apple’s developer documentation: where the page is (its breadcrumbs), what kind of page it is (its eyebrow), its title and what it is about (its abstract). The title and the abstract grow from 736px and again from 1251px.

```css
@import "ciderui/components/docs/docs-hero/docs-hero.css";
```

{{ <cui_example file="docs-hero/hero" title="An article’s hero" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-docs-hero` | the hero: 40px above and below |
| `cui-docs-hero-crumbs` | the pages above this one, each a link, a “/” between them, the current page last (`aria-current="page"`), 20px over the eyebrow |
| `cui-docs-hero-eyebrow` | what kind of page it is, in the secondary text, 15px over the title |
| `cui-docs-hero-title` | the page’s title, 12px over the abstract |
| `cui-docs-hero-abstract` | what the page is about |

The Go module writes a hero, its crumbs ending with the page’s own title, with `docs.DocsHero` and its `docs.DocsCrumb`s.

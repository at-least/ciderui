+++
title = "Accordion"
description = "apple.com’s questions and answers, each opening below its question with no script."

[extra]
category = "Controls"
image = "pictures/accordion.png"
+++

apple.com’s FAQ, from its bag: questions under rules, each opening its answer below it, its chevron at its end turning up. It is HTML’s own `<details>`, so it opens with no script, from the keyboard too.

```css
@import "ciderui/components/accordion/accordion.css";
```

{{ <cui_example file="accordion/accordion" title="Two questions, the first open" pad={true} /> }}

| class | what |
| --- | --- |
| `cui-accordion` | the questions |
| `cui-accordion-item` | a `<details>`: a question and its answer, a rule above it |
| `cui-accordion-question` | its `<summary>`: 17px at 600 (14px below 735px), its chevron at its end in the muted text, the secondary hovered, turned up open |
| `cui-accordion-answer` | the answer, 14px, 28px above the next rule |

An item with `open` starts opened. Its height moves to the answer’s where the browser can animate to `auto` (Chrome), and opens at once where it cannot (Safari, Firefox).

## A large question

`data-size="lg"` is the bag’s outer question: 32px (24px below 735px), 32px from its rule.

{{ <cui_example file="accordion/large" title="A large question" pad={true} /> }}

The crate writes it with `Accordion` and `AccordionItem`, the chevron its `icon::DISCLOSURE`.

+++
title = "The Newsroom, in ciderui"
description = "Three parts from apple.com’s Newsroom: its story, its index and its archive."
date = 2026-10-19T20:49:07+08:00

[extra]
category = "Feature"
image = "pictures/stories.png"
+++

ciderui has three new components, read off apple.com’s Newsroom: a story, its index and its archive. This site’s news is made of them.

## The article

[Article](@/components/article.md), `cui-article`, is a story: an eyebrow of its category over its date, its headline at 700, a subhead, and its copy in a column narrower than its figures, 653px to their 980px on a wide window. Its copy styles the plain elements markdown writes, so a page’s markdown needs no class: the Newsroom’s paragraphs, subheads, lists and footnotes, and ciderui’s own code, quotations and tables, which the Newsroom has none of.

## The index

[Stories](@/components/stories.md), `cui-stories`, is the index: a grid of white tiles rounded 32px on the alternate ground, one, two or three across, the whole of each its link.

## The archive

[Posts](@/components/posts.md), `cui-posts`, is the archive’s month of ruled rows, and, smaller, the related stories under an article, like the ones below.

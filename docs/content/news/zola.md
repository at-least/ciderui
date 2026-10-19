+++
title = "ciderui comes to Zola"
description = "ciderui-zola, a Zola theme in apple.com’s look, for a static site’s pages and its news."
date = 2026-10-19T21:14:19+08:00

[extra]
category = "Update"
image = "pictures/article.png"
+++

ciderui-zola is a theme for Zola[^zola] built on ciderui: a site’s pages as articles, its news as the Newsroom’s index and archive, the global nav and the footer from its settings. This documentation is its first site.

## Nothing to build

A site needs Zola alone. The theme ships its stylesheet prebuilt (ciderui, every component and the type’s classes), ciderui’s Inter, and its one script, which moves a gallery’s and a shelf’s paddles and marks a gallery’s dot; everything else opens with HTML’s own elements.

## ciderui’s markup, held to it

Each part a page has is a Tera component named after the crate’s, with the crate’s fields, so the theme writes the markup ciderui’s examples have. Its checks hold it there:

- each component’s HTML, built by Zola, equals ciderui’s example;
- each example drawn by the theme’s stylesheet computes what it does with ciderui’s own, box by box and style by style, at six widths: 112,464 values, none apart;
- the script scrolls a row, disables the paddle at its end and marks the dot of the item in view;
- the sites built with it build, every page with its title, its description, the stylesheet and the script, and no template or markdown showing.

[^zola]: Zola 0.23.6, whose templates are Tera 2’s and whose shortcodes are Tera components.

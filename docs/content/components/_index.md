+++
title = "Components"
description = "apple.com’s parts as classes, each measured from the page it comes from."
sort_by = "title"

[extra]
heading = "All Components"
localnav = {}
+++

A component is a class, or a block of a few, that gives an element the look of one of apple.com’s parts, and nothing else: no script, no place on the page. Each is a stylesheet of its own, `components/<name>/<name>.css`, imported when a site uses it.

What varies is an attribute: a mutually exclusive choice is `data-variant` or `data-size`, a state the one the markup carries anyway (`aria-current`, `checked`, `disabled`, `[open]`). Each page shows the part live, in a frame of its own, over its markup.

# ciderui's documentation

ciderui's documentation, as a site built with ciderui-zola, the Zola theme on ciderui: getting started, the foundations (the colors, the type, the breakpoints), the Rust crate, a page for each component with its live examples, and the news, in the Newsroom's look. The pages are under `content/`.

From the repository's root, with ciderui-zola checked out beside this repository:

```sh
node scripts/build-docs.mjs   # the examples' copies and the frames' stylesheet
cd docs && zola serve
```

`themes/ciderui-zola` links to `../ciderui-zola`, so the site is drawn by the theme as it is there, with its prebuilt stylesheet, fonts and script. The site is not published: `base_url` names no host.

An example is a component's example file (`components/<set>/<name>/examples/<id>.html`, the markup the crate's tests render too), shown by the page's `cui_example` component (`templates/cui-example.html`) in a frame of its own, the article's figure, over the same markup as code. Zola reads files only inside the site, so `node scripts/build-docs.mjs` copies every example into `examples/` (a page's own examples live in `examples/pages/`) and builds the frames' stylesheet, `static/example.css`: ciderui with every component, its fonts and the type's classes the examples write. The frame is a page of its own, untouched by the article's styles; it runs the theme's script, so a gallery's paddles work, and `static/example.js`, which keeps it as tall as its content and its links where they are. Run the script after changing an example; `npm test` fails while the copies or the stylesheet are stale, and checks the site with `test/docs.mjs`.

The pictures (`static/pictures/`: a component's row, a story's, the home's tile) are ciderui drawn by itself, a scene of its examples in Chromium, by `node scripts/docs-pictures.mjs` after `build-docs`. Another Chromium draws other bytes, so what is held is what each was drawn from: `pictures.json` records its scene and the stylesheet's digest, and `npm test` fails (`docs-pictures.mjs --check`) once an example, the frames' stylesheet or a scene has changed since it was drawn: redraw them then.

The home is `templates/home.html`: a tile, the newest stories, and a section's header with its links, from `content/_index.md`'s front matter.

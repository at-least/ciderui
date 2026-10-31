# ciderui

apple.com's look for websites, built with Tailwind CSS v4: a small contract of variables (`--cui-*`: the colors, the sizes, the motion), the type as classes, and apple.com's parts as classes (`cui-btn`, `cui-globalnav`, `cui-tile`, …), each in a stylesheet of its own that a site imports when it uses the part; Inter for where the system font is not SF; and a Rust crate that writes the parts' markup with hypertext, carries the stylesheets and draws every [Lucide](https://lucide.dev) icon. It has no JavaScript. [ciderui-zola](../ciderui-zola), a Zola theme, is built on it. Why it is built as it is: [DESIGN.md](DESIGN.md).

## Use it

ciderui is not published: a site takes it as a checkout beside its own, `"ciderui": "file:../ciderui"`, with the Tailwind CLI, and imports it after Tailwind, its fonts first, then the parts it uses:

```css
@import "tailwindcss";
@import "ciderui/fonts.css";
@import "ciderui";
@import "ciderui/components/site/button/button.css";
```

The documentation is `docs/`, a site built with ciderui-zola (checked out beside this repository): getting started, the foundations, the crate, and each part live over its markup.

```sh
node scripts/build-docs.mjs   # the examples' copies and the frames' stylesheet
cd docs && zola serve
```

## The repository

| path | what |
| --- | --- |
| `index.css` | `theme.css`, `tokens.css`, `base.css` and `type.css`: the breakpoints and utilities, the variables, the page's base, the type |
| `fonts.css`, `fonts/` | Inter |
| `site.css`, `console.css` | a set's every part, one import each: the site set (apple.com's pages) and the console set (`cui-console-*`) |
| `components/<set>/<name>/` | a part, in its set (`site` or `console`): its stylesheet, its examples (the markup the crate's tests render too) and its test |
| `rust/` | the crate |
| `reference/` | the measurements, each with its probe, page, browser and date |
| `docs/` | the documentation's site |

`npm test` runs every test, the browser's in one Docker image (`test/Dockerfile`, `scripts/docker.mjs`): the variables, the type, the page, each part in headless Chromium against apple.com's measured values and its own goldens, the crate (its Lucide icons lucide-static's, formatted, clippy, its tests, its documentation's examples and its documentation), and the documentation (its examples' copies and pictures current, every page).

## License

MIT (`LICENSE`). Inter, in `fonts/`, is the Inter Project Authors', under the SIL Open Font License 1.1 (`LICENSE-Inter`). Lucide's icons, in `rust/src/lucide.rs`, are Lucide Icons and Contributors', under the ISC License, those derived from Feather also Cole Bemis's, under the MIT License (`LICENSE-Lucide`).

## Apple

What ciderui takes from apple.com is measured on its pages and recorded in `reference/`; no Apple stylesheet, font, image or logo is in this repository. Apple's names appear where they say where a value was read. Apple, iPhone and Mac are trademarks of Apple Inc. ciderui is not affiliated with Apple, and Apple does not endorse it.

# ciderui

Agents write this project entirely. It is fully optimized for agent vibe coding, debugging and RSI (recursive self-improvement): this README is the only file written for people; every other document, comment and line of code is written for the agents that read and change it.

apple.com's look for websites, built with Tailwind CSS v4: a small contract of variables (`--cui-*`: the colors, the type, the focus ring, the shadows), the type as classes, and apple.com's parts as classes (`cui-btn`, `cui-globalnav`, `cui-tile`, …), each in a stylesheet of its own that a site imports when it uses the part; Inter for where the system font is not SF; and a Go module that writes the parts' markup with gomponents, carries the stylesheets and draws every [Lucide](https://lucide.dev) icon. It has no JavaScript. [ciderui-zola](../ciderui-zola), a Zola theme, is built on it. Why it is built as it is: [DESIGN.md](DESIGN.md).

## Use it

ciderui is not on npm: a site's CSS build takes it as a checkout beside its own, `"ciderui": "file:../ciderui"`, with the Tailwind CLI, and imports it after Tailwind, its fonts first, then the parts it uses:

```css
@import "tailwindcss";
@import "ciderui/fonts.css";
@import "ciderui";
@import "ciderui/components/site/button/button.css";
```

Each part's markup is its examples, `components/<set>/<name>/examples/*.html`. A site that renders in Go writes the same markup with the Go module, pinned at a commit, which also carries the stylesheets (`go run github.com/at-least/ciderui/go/cmd/ciderui-css <dir>` writes them out for the Tailwind build), and calls `ciderui.CheckMarkup` on its pages in its tests:

```sh
go get github.com/at-least/ciderui@<commit>
```

## The repository

| path | what |
| --- | --- |
| `index.css` | `theme.css`, `tokens.css`, `base.css` and `type.css`: the breakpoints and utilities, the variables, the page's base, the type |
| `fonts.css`, `fonts/` | Inter |
| `site.css`, `console.css` | a set's every part, one import each: the site set (apple.com's pages) and the console set (`cui-console-*`) |
| `components/<set>/<name>/` | a part, in its set (`site` or `console`): its stylesheet, its examples (the markup the Go module's tests render too) and its test |
| `ciderui.go`, `markup.go`, `go/` | the Go module (`go.mod`): the root package embeds the stylesheets and checks a page's markup; `go/` holds the components (`site`, `console`), the icons (`icon`, `lucide`) and `ciderui-css` |
| `reference/` | the measurements, each with its probe, page, browser and date |

`npm test` runs every test, the browser's in one Docker image (`test/Dockerfile`, `scripts/docker.mjs`): the variables, the type, the page, each part in headless Chromium against apple.com's measured values and its own goldens, the Go module (its Lucide icons lucide-static's, gofmt, vet, exhaustruct, its tests and its examples), every control's touch target (`reference/touch/`) and the reference probes Chromium can run (`test/probes.mjs`).

## License

MIT (`LICENSE`). Inter, in `fonts/`, is the Inter Project Authors', under the SIL Open Font License 1.1 (`LICENSE-Inter`). Lucide's icons, in `go/lucide/lucide.go`, are Lucide Icons and Contributors', under the ISC License, those derived from Feather also Cole Bemis's, under the MIT License (`LICENSE-Lucide`).

## Apple

What ciderui takes from apple.com is measured on its pages and recorded in `reference/`; no Apple stylesheet, font, image or logo is in this repository. Apple's names appear where they say where a value was read. Apple, iPhone and Mac are trademarks of Apple Inc. ciderui is not affiliated with Apple, and Apple does not endorse it.

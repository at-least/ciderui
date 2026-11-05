# ciderui

Agents write this project entirely. It is fully optimized for agent vibe coding, debugging and RSI (recursive self-improvement): this README is the only file written for people; every other document, comment and line of code is written for the agents that read and change it.

apple.com's look for websites, built with Tailwind CSS v4: a small contract of variables (`--cui-*`: the colors, the type, the focus ring, the shadows), the type as classes, and apple.com's parts as classes (`cui-btn`, `cui-globalnav`, `cui-tile`, …), each in a stylesheet of its own, a set's every part one import; Inter for where the system font is not SF; and a Go module that writes the parts' markup with gomponents, carries the stylesheets and fonts and draws every [Lucide](https://lucide.dev) icon. It has no JavaScript. Why it is built as it is: [DESIGN.md](DESIGN.md).

## Use it

```sh
go get -tool github.com/at-least/ciderui/go/cmd/ciderui-css@<commit>
go tool ciderui-css target/ciderui
```

That writes the stylesheets and the fonts under `target/ciderui`. Your Tailwind entry imports one set, with one line:

```css
@import "tailwindcss";
@import "../target/ciderui/site.css";    /* or console.css */
```

To draw Inter where the system font is not SF, also add `@import "../target/ciderui/fonts.css";` and serve `target/ciderui/fonts/` as `fonts/` beside your built stylesheet: its URLs stay relative to that file. Without it, the text falls through to the platform's fonts.

Each part's markup is its examples, `components/<set>/<name>/examples/*.html`. Pages rendered in Go write the same markup with the Go module, which is the one pin for markup and CSS. Call `ciderui.CheckMarkup` on your pages in your tests.

## The repository

| path | what |
| --- | --- |
| `index.css` | `theme.css`, `tokens.css`, `base.css` and `type.css`: the breakpoints and utilities, the variables, the page's base, the type |
| `fonts.css`, `fonts/` | Inter |
| `site.css`, `console.css` | `index.css` and a set's every part, the one import a site writes: the site set (apple.com's pages) and the console set (`cui-console-*`) |
| `components/<set>/<name>/` | a part, in its set (`site` or `console`): its stylesheet, its examples (the markup the Go module's tests render too) and its test |
| `ciderui.go`, `markup.go`, `go/` | the Go module (`go.mod`): the root package embeds the stylesheets and fonts and checks a page's markup; `go/` holds the components (`site`, `console`), the icons (`lucide`, and the type `icon.Icon`) and `ciderui-css` |
| `reference/` | the measurements, each with its probe, page, browser and date |

`npm test` runs every test, the browser's in one Docker image (`test/Dockerfile`, `scripts/docker.mjs`): the variables, the type, the page, each part in headless Chromium against apple.com's measured values and its own goldens, the Go module (its Lucide icons lucide-static's, gofmt, vet, exhaustruct, its tests and its examples, each set built from what `ciderui-css` writes), every control's touch target (`reference/touch/`) and the reference probes Chromium can run (`test/probes.mjs`).

## License

MIT (`LICENSE`). Inter, in `fonts/`, is the Inter Project Authors', under the SIL Open Font License 1.1 (`LICENSE-Inter`). Lucide's icons, in `go/lucide/lucide.go`, are Lucide Icons and Contributors', under the ISC License, those derived from Feather also Cole Bemis's, under the MIT License (`LICENSE-Lucide`).

## Apple

What ciderui takes from apple.com is measured on its pages and recorded in `reference/`; no Apple stylesheet, font, image or logo is in this repository. Apple's names appear where they say where a value was read. Apple, iPhone and Mac are trademarks of Apple Inc. ciderui is not affiliated with Apple, and Apple does not endorse it.

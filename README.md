# ciderui

Agents write this project entirely. It is fully optimized for agent vibe coding, debugging and RSI (recursive self-improvement): this README is the only file written for people; every other document, comment and line of code is written for the agents that read and change it.

apple.com's look for websites, built with Tailwind CSS v4: a small contract of variables (`--cui-*`: the colors, the type, the focus ring, the shadows), the type as classes, and apple.com's parts as classes (`cui-btn`, `cui-globalnav`, `cui-field`, …); Inter for where the system font is not SF; and a Go module that writes the parts' markup with gomponents, serves the stylesheets and fonts and draws every [Lucide](https://lucide.dev) icon. It has no JavaScript. Why it is built as it is: [DESIGN.md](DESIGN.md).

## Use it

```go
mux.Handle(ciderui.AssetsPath, ciderui.Assets())

console.ConsolePage{Lang: "en", Title: "Jobs", Head: nil, Class: "", BodyClass: "", Children: body}
```

That is the whole setup: the page links the stylesheet and the module serves it, with its fonts. Use `site.Page` for a store or a marketing page.

Each part's markup is its examples, `components/<set>/<name>/examples/*.html`. Call `ciderui.CheckMarkup` on your pages in your tests.

### Your own utilities

A page that writes Tailwind utilities builds one more stylesheet and adds it in `Head`:

```sh
go get -tool github.com/at-least/ciderui/go/cmd/ciderui-css@<commit>
go tool ciderui-css target/ciderui
```

```css
@layer theme, base, components, utilities;
@import "tailwindcss/theme.css" layer(theme);
@import "tailwindcss/utilities.css" layer(utilities) source(none);
@import "../target/ciderui/utilities.css";
@source "../pages/*.go";
```

## The repository

| path | what |
| --- | --- |
| `index.css` | `theme.css`, `tokens.css`, `base.css` and `type.css`: the breakpoints and utilities, the variables, the page's base, the type |
| `fonts.css`, `fonts/` | Inter |
| `site.css`, `console.css` | `index.css` and a set's every part: the site set (apple.com's pages) and the console set (`cui-console-*`) |
| `dist/` | each set built, as the module serves it (`node scripts/build-css.mjs`) |
| `utilities.css` | what a site's own utilities' stylesheet imports |
| `components/<set>/<name>/` | a part, in its set (`site` or `console`): its stylesheet, its examples (the markup the Go module's tests render too) and its test |
| `ciderui.go`, `markup.go`, `go/` | the Go module (`go.mod`): the root package embeds and serves the stylesheets and fonts and checks a page's markup; `go/` holds the components (`site`, `console`), the icons (`lucide`, and the type `icon.Icon`) and `ciderui-css` |
| `reference/` | the measurements, each with its probe, page, browser and date |

`npm test` runs every test, the browser's in one Docker image (`test/Dockerfile`, `scripts/docker.mjs`): the variables, the type, the page, each part in headless Chromium against apple.com's measured values and its own goldens, the rules every part keeps (hover, the focus ring, color motion, letter-spacing) read off every example, a picture of every example at every width (when one changes on purpose, `UPDATE_GOLDEN=1 node scripts/docker.mjs node test/pictures.mjs`; a failure leaves the new picture in `test/pictures-failed/`), the Go module (its Lucide icons lucide-static's, gofmt, vet, exhaustruct, its tests and its examples, every statement of `site` and `console` run by a test, a site's utilities built from what `ciderui-css` writes), every control's touch target (`reference/touch/`) and the reference probes Chromium can run (`test/probes.mjs`).

## License

MIT (`LICENSE`). Inter, in `fonts/`, is the Inter Project Authors', under the SIL Open Font License 1.1 (`LICENSE-Inter`). Lucide's icons, in `go/lucide/lucide.go`, are Lucide Icons and Contributors', under the ISC License, those derived from Feather also Cole Bemis's, under the MIT License (`LICENSE-Lucide`).

## Apple

What ciderui takes from apple.com is measured on its pages and recorded in `reference/`; no Apple stylesheet, font, image or logo is in this repository. Apple's names appear where they say where a value was read. Apple, iPhone and Mac are trademarks of Apple Inc. ciderui is not affiliated with Apple, and Apple does not endorse it.

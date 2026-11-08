# ciderui

This project is agent-native: written and maintained entirely by AI agents, and fully optimized for autonomous development (vibe coding), deterministic debugging, and recursive self-improvement (RSI). This README is the only file written for people; every other document, comment and line of code is written for the agents that read and change them. What is decided, and why, is in [DESIGN.md](DESIGN.md).

## What it is

apple.com's look for websites, as one dependency: a small contract of CSS variables (`--cui-*`) — the colors, the type, the focus ring, the shadows —, the type as classes, apple.com's parts as classes (`cui-btn`, `cui-globalnav`, `cui-field`, …), Inter where the system font is not SF, and a Go module that writes the parts' markup with gomponents, serves the stylesheets and fonts, and draws every [Lucide](https://lucide.dev) icon. It has no JavaScript: behavior comes from HTML's own elements. Every value is measured on Apple's own pages.

## How it fits together

```
apple.com, measured (reference/: each value's probe, page, browser, date)
        │
        ▼
tokens.css · theme.css · base.css · type.css · components/<set>/<part>.css
        │
        │  scripts/build-css.mjs builds each set
        ▼
dist/site.css, dist/console.css — embedded in the Go module
        │  ciderui.Assets serves; site.Page, console.ConsolePage link;
        │  site.*, console.* write the markup; lucide.* the icons
        ▼
a site: its pages in Go, one dependency, no JavaScript of ciderui's
        │
        ▼
npm test: every part, in headless Chromium, against reference/'s values
```

## Use it

```go
mux.Handle(ciderui.AssetsPath, ciderui.Assets())

console.ConsolePage{Lang: "en", Title: "Jobs", Head: nil, Class: "", Skip: "", Children: body}
```

That is the whole setup: the page links the stylesheet and the module serves it, with its fonts. Use `site.Page` for a store or a marketing page. Each part's markup is its examples, `components/<set>/<name>/examples/*.html`. Under `go test`, every page `site.Page` or `console.ConsolePage` renders checks itself with `ciderui.CheckMarkup`; call it yourself on a fragment.

Add `go tool exhaustruct -i 'github.com/at-least/ciderui/.*' ./...` to your site's checks (after `go get -tool github.com/GaijinEntertainment/go-exhaustruct/v3/cmd/exhaustruct`), so a field the module adds fails your build instead of rendering a default nobody chose.

### Your own utilities

A set's stylesheet ships the utilities its examples write and no other: any other utility a page writes needs this second stylesheet. A page that writes Tailwind utilities builds it and adds it in `Head`. This writes what it imports, from the version your `go.mod` pins:

```sh
go run github.com/at-least/ciderui/go/cmd/ciderui-css target/ciderui
```

```css
@layer theme, base, components, utilities;
@import "tailwindcss/theme.css" layer(theme);
@import "tailwindcss/utilities.css" layer(utilities) source(none);
@import "../target/ciderui/utilities.css";
@source "../pages/*.go";
```

## Developing

`npm install`, then `npm test`: it checks each part in headless Chromium against apple.com's measured values, takes a picture of every example at every width, and tests the Go module. The browser's tests run in one Docker image, reached through `DOCKER_HOST`.

## License

MIT (`LICENSE`). Inter, in `fonts/`, is the Inter Project Authors', under the SIL Open Font License 1.1 (`LICENSE-Inter`). Lucide's icons, in `go/lucide/lucide.go`, are Lucide Icons and Contributors', under the ISC License, those derived from Feather also Cole Bemis's, under the MIT License (`LICENSE-Lucide`).

## Apple

What ciderui takes from apple.com is measured on its pages and recorded in `reference/`; no Apple stylesheet, font, image or logo is in this repository. Apple's names appear where they say where a value was read. Apple, iPhone and Mac are trademarks of Apple Inc. ciderui is not affiliated with Apple, and Apple does not endorse it.

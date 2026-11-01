+++
title = "Go"
description = "The Go module: the components’ markup as gomponents nodes, the icons, and the stylesheets."
template = "docs.html"
+++

A site that renders its pages in Go with [gomponents](https://www.gomponents.com) writes ciderui’s markup with the module `github.com/at-least/ciderui`, and takes its stylesheets from the same module, so its dependency, at one commit, pins both: the markup and the CSS cannot part. The repository is private, so the go command fetches it with your Git credentials, not through the public proxy:

```sh
export GOPRIVATE=github.com/at-least
go get github.com/at-least/ciderui@<commit>
```

## The components

The module holds a component for each part whose markup is more than a class on one element, in its set’s package: `site` for apple.com’s parts, `console` for the console’s (`console.ConsoleBadge`, its name keeping `Console`), `docs` for Apple’s developer documentation’s (`docs.DocsBar`); a button, a link in running text or a price is a class string on the page’s own element. Each component is a struct and a gomponents node; its fields are the part’s: its text, its links, its variant as a value of the package’s, and `Class`, the page’s utilities on its root (`""` for none). A call writes every field:

```go
page := h.Form(site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Error: "", Disabled: false, Prefix: "", Class: ""})
_ = page.Render(os.Stdout)
// Output: <form><div class="cui-field"><input class="cui-field-input" id="email" type="email" name="email" placeholder=" " autocomplete="email"><label class="cui-field-label" for="email">Email</label></div></form>
```

A field left out of a call would be a default nobody chose, and Go fills it in silently. So a site runs exhaustruct on its own code in its checks; a field added to a component then stops every call that does not set it. Add the tool to the site’s `go.mod` once:

```sh
go get -tool github.com/GaijinEntertainment/go-exhaustruct/v3/cmd/exhaustruct
```

and run it with the site’s tests:

```sh
go tool exhaustruct -i 'github.com/at-least/ciderui/.*' ./...
```

| component | the part |
| --- | --- |
| `MoreLink` | a more link and its chevron |
| `Section`, `SectionHeader` | a band and its header |
| `GlobalNav`, `NavLink`, `FlyoutGroup`, `LinkedSearch`, `FlyoutSearch`, `LinkedBag`, `FlyoutBag`, `BagLink`, `LocalNav` | the global nav, its links, their flyouts, its search and its bag, a product’s local nav |
| `Ribbon` | the message above a page |
| `Tile`, `Card` | a home page tile, a store card |
| `Shelf`, `Gallery` | a row of cards, a product page’s gallery |
| `Field`, `Select`, `TextArea`, `Checkbox`, `OptionTile`, `Swatch` | a form’s controls |
| `Modal`, `ModalButton` | the store’s overlay or a marketing page’s modal, and the button that opens it |
| `Accordion`, `AccordionItem` | questions and answers |
| `QuickLinks`, `ProductNav`, `ProductNavItem` | the store’s links and its products |
| `Footer`, `FooterSection` | the footer |
| `Stat`, `Footnote` | a number, a footnote’s mark |
| `Article`, `Stories`, `Story`, `Posts`, `Post` | the Newsroom’s story, index and archive |

Each one’s test renders its example, the markup this documentation shows, and holds the output equal to it.

## Icons

The package `icon` holds ciderui’s icons as inline SVG in `currentColor`, each drawn from the numbers of the glyph apple.com draws in its place: the more link’s chevron, the global nav’s search, bag and menu, the paddles’ chevrons, a checkbox’s check.

```go
page := h.A(h.Href("/iphone/"), g.Text("Learn more"), icon.More.Class("cui-link-icon"))
_ = page.Render(os.Stdout)
// Output: <a href="/iphone/">Learn more<svg class="cui-link-icon" aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540"><path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg></a>
```

Every other icon is [Lucide](https://lucide.dev/icons/)’s: the package `lucide` holds every one, named as Lucide names it, in Go’s capitals (`shopping-bag` is `lucide.ShoppingBag`). Each is 1em square in `currentColor`, so it takes its text’s size and color; a utility on it sizes it otherwise. In a line of text, add `inline`: Tailwind’s preflight makes an `<svg>` a block. A part apple.com draws an icon in draws ciderui’s own.

```go
page := h.A(h.Href("/saved"), lucide.Heart.Class("inline"), g.Text(" Saved"))
_ = page.Render(os.Stdout)
// Output: <a href="/saved"><svg class="inline" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 9.5a5.5 5.5 0 0 1 9.591-3.676.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5c0 2.29-1.5 4-3 5.5l-5.492 5.313a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5"/></svg> Saved</a>
```

## Checking the markup

A class misspelt styles nothing, and a `data-variant`, `data-size` or `data-color` value no rule selects draws the default, with no error. `ciderui.CheckMarkup` reads a rendered page and lists each word of ciderui’s that its stylesheets do not know; a site’s tests call it on every page they render:

```go
page := `<a class="cui-console-btn" data-variant="plian" href="/">Open</a>`
fmt.Println(ciderui.CheckMarkup(page))
// Output: [data-variant="plian" on <a class="cui-console-btn">: no rule of ciderui's selects it (cui-console-btn's: filled, plain)]
```

## The stylesheets

The module embeds every stylesheet, and writes them out in the kit’s layout before the site’s Tailwind build:

```sh
go run github.com/at-least/ciderui/go/cmd/ciderui-css target/ciderui
```

or, from the site’s own Go, `ciderui.WriteCSS("target/ciderui")`. The site’s entry then imports from there:

```css
@import "tailwindcss";
@import "../target/ciderui/fonts.css";
@import "../target/ciderui/index.css";
@import "../target/ciderui/site.css";
```

The Inter files are not embedded: copy ciderui’s `fonts/` beside the compiled stylesheet once.

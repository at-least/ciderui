use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::icon::{DOCS_ARTICLE, DOCS_CHEVRON, DOCS_COLLECTION, DOCS_FILTER};

/// A row of the navigator: a group's title over the pages after it, or a
/// page.
#[derive(Clone, Copy, Debug)]
pub enum DocsNavigatorItem<'a> {
    /// A group's title, the pages after it its own until the next.
    Group(&'a str),
    /// A page, and the pages under it.
    Page(DocsNavigatorPage<'a>),
}

/// A page in the navigator: a collection's glyph where it has pages under
/// it, which its chevron shows and hides, an article's where not.
#[derive(Clone, Copy, Debug)]
pub struct DocsNavigatorPage<'a> {
    /// Its name.
    pub label: &'a str,
    /// Its address. Its path unique among the navigator's pages: a page
    /// with pages under it names its chevron's checkbox after its path's
    /// words.
    pub href: &'a str,
    /// The page shown (`aria-current="page"`).
    pub current: bool,
    /// The pages under it shown at first (the current page's ancestors'
    /// are, as DocC's); `false` for a page with none.
    pub open: bool,
    /// The pages under it, and their groups' titles; `&[]` for none.
    pub children: &'a [DocsNavigatorItem<'a>],
}

/// `cui-docs-navigator` (`components/docs/docs-navigator`): the navigator
/// of Apple's developer documentation, its pages as a tree, a `<nav>` named
/// Navigator (the bar's button's name): the documentation's title, a link
/// to its first page, over its rows. A row
/// with pages under it opens and closes by its chevron, a checkbox and its
/// label (no script); the current page's row is marked. Under the rows, the
/// line saying the filter matches nothing, the filter at the navigator's
/// foot and a status read aloud, the first two hidden: the page's script
/// shows the filter, filters the rows and says what it found.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::docs::{DocsNavigator, DocsNavigatorItem, DocsNavigatorPage};
///
/// let items = [DocsNavigatorItem::Page(DocsNavigatorPage { label: "Get Started", href: "/get-started/", current: true, open: false, children: &[] })];
/// let html = rsx! { <DocsNavigator title="ciderui" href="/" items=(&items) class="" /> }.render();
/// assert!(html.as_inner().contains(r#"aria-current="page""#));
/// ```
pub struct DocsNavigator<'a> {
    /// The documentation's title, at the navigator's top. Required: an
    /// empty one panics.
    pub title: &'a str,
    /// Where the title links: the documentation's first page.
    pub href: &'a str,
    /// Its rows, from the top.
    pub items: &'a [DocsNavigatorItem<'a>],
    /// The page's utilities on the navigator; `""` for none.
    pub class: &'a str,
}

/// The rows of one level, written as a list.
struct Rows<'a>(&'a [DocsNavigatorItem<'a>]);

impl Renderable for Rows<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            <ul class="cui-docs-navigator-list">
                @for item in self.0 {
                    (Row(*item))
                }
            </ul>
        }
        .render_to(buffer);
    }
}

struct Row<'a>(DocsNavigatorItem<'a>);

impl Renderable for Row<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        match self.0 {
            DocsNavigatorItem::Group(title) => {
                rsx! { <li class="cui-docs-navigator-group">(title)</li> }.render_to(buffer);
            }
            DocsNavigatorItem::Page(page) => {
                assert!(
                    !page.label.is_empty(),
                    "ciderui::docs::DocsNavigator: a page needs a label"
                );
                assert!(
                    !page.open || !page.children.is_empty(),
                    "ciderui::docs::DocsNavigator: {} is open with no pages under it",
                    page.label
                );
                let parent = !page.children.is_empty();
                // a parent's chevron's checkbox id; a leaf has no chevron
                let id = parent.then(|| chevron_id(page.href));
                let icon = if parent {
                    DOCS_COLLECTION
                } else {
                    DOCS_ARTICLE
                };
                rsx! {
                    <li class="cui-docs-navigator-item">
                        @if let Some(id) = &id {
                            <input class="cui-docs-navigator-toggle" type="checkbox" id=(id.as_str()) aria-label=(page.label) checked[page.open]>
                        }
                        <div class="cui-docs-navigator-row">
                            @if let Some(id) = &id {
                                <label class="cui-docs-navigator-chevron" for=(id.as_str())>(DOCS_CHEVRON)</label>
                            }
                            (icon.class("cui-docs-navigator-icon"))
                            <a class="cui-docs-navigator-link" href=(page.href) aria-current=("page")[page.current]>(page.label)</a>
                        </div>
                        @if parent {
                            (Rows(page.children))
                        }
                    </li>
                }
                .render_to(buffer);
            }
        }
    }
}

/// The chevron's checkbox id, named after the page's address, its path's
/// words alone (`/components/` and `https://example.test/components/`
/// docs-navigator-components).
fn chevron_id(href: &str) -> String {
    let path = match href.split_once("://") {
        Some((_, rest)) => rest.find('/').map_or("", |i| &rest[i..]),
        None => href,
    };
    let mut id = String::from("docs-navigator");
    for word in path
        .split(|c: char| !c.is_ascii_alphanumeric())
        .filter(|w| !w.is_empty())
    {
        id.push('-');
        id.push_str(&word.to_ascii_lowercase());
    }
    id
}

impl Renderable for DocsNavigator<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.title.is_empty(),
            "ciderui::docs::DocsNavigator needs a title"
        );
        rsx! {
            <nav class=(Class("cui-docs-navigator", self.class)) aria-label="Navigator">
                <a class="cui-docs-navigator-title" href=(self.href)>(self.title)</a>
                (Rows(self.items))
                <p class="cui-docs-navigator-empty" hidden>"No results found."</p>
                <div class="cui-docs-navigator-filter" role="search" hidden>
                    <label class="cui-docs-navigator-field">(DOCS_FILTER.class("cui-docs-navigator-filter-icon"))<input class="cui-docs-navigator-input" type="search" placeholder="Filter" aria-label="Filter" autocomplete="off" spellcheck="false"></label>
                </div>
                <p class="cui-docs-navigator-status" aria-live="polite"></p>
            </nav>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    const fn page<'a>(
        label: &'a str,
        href: &'a str,
        current: bool,
        open: bool,
        children: &'a [DocsNavigatorItem<'a>],
    ) -> DocsNavigatorItem<'a> {
        DocsNavigatorItem::Page(DocsNavigatorPage {
            label,
            href,
            current,
            open,
            children,
        })
    }

    #[test]
    fn renders_the_example() {
        let reports = [
            DocsNavigatorItem::Group("Reports"),
            page("Reading a crash report", "#reading", true, false, &[]),
            page(
                "Symbolicating a backtrace",
                "#symbolicating",
                false,
                false,
                &[],
            ),
        ];
        let alerts = [page("Routing an alert", "#routing", false, false, &[])];
        let items = [
            page("Getting started", "#getting-started", false, false, &[]),
            page("Crash reports", "#crash-reports", false, true, &reports),
            page("Alerts", "#alerts", false, false, &alerts),
        ];
        let navigator = rsx! {
            <DocsNavigator title="Crashes" href="#crashes" items=(&items) class="" />
        }
        .render();
        assert_eq!(
            *navigator.as_inner(),
            example(include_str!(
                "../../../components/docs/docs-navigator/examples/navigator.html"
            ))
        );
    }

    #[test]
    fn a_chevron_is_named_after_its_pages_path() {
        let under = [page(
            "Button",
            "https://example.test/components/button/",
            false,
            false,
            &[],
        )];
        for href in ["/components/", "https://example.test/components/"] {
            let items = [page("Components", href, false, false, &under)];
            let html =
                rsx! { <DocsNavigator title="T" href="/" items=(&items) class="" /> }.render();
            assert!(
                html.as_inner()
                    .contains(r#"id="docs-navigator-components""#),
                "{href}"
            );
        }
    }

    #[test]
    #[should_panic(expected = "is open with no pages under it")]
    fn an_open_page_with_no_pages_panics() {
        let items = [page("Alone", "/alone/", false, true, &[])];
        let _ = rsx! { <DocsNavigator title="T" href="/" items=(&items) class="" /> }.render();
    }

    #[test]
    #[should_panic(expected = "needs a title")]
    fn an_empty_title_panics() {
        let _ = rsx! { <DocsNavigator title="" href="/" items=(&[]) class="" /> }.render();
    }
}

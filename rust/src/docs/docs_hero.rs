use hypertext::{Buffer, prelude::*};

use crate::Class;

/// A page above this one in the documentation, as a crumb: its name and
/// its address.
#[derive(Clone, Copy, Debug)]
pub struct DocsCrumb<'a> {
    /// Its name, the link's words.
    pub label: &'a str,
    /// Its address.
    pub href: &'a str,
}

/// `cui-docs-hero` (`components/docs/docs-hero`): the top of a page in
/// Apple's developer documentation: its breadcrumbs (the pages above it,
/// then its own title, the current crumb), its eyebrow, its title and its
/// abstract.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::docs::{DocsCrumb, DocsHero};
///
/// let crumbs = [DocsCrumb { label: "Components", href: "/components/" }];
/// let html = rsx! {
///     <DocsHero crumbs=(&crumbs) eyebrow="Component" title="Docs Hero" abstract_="The top of a page." class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<li aria-current="page">Docs Hero</li>"#));
/// ```
pub struct DocsHero<'a> {
    /// The pages above this one, from the top; `&[]` for none, a page at the
    /// top, which shows no breadcrumbs.
    pub crumbs: &'a [DocsCrumb<'a>],
    /// What kind of page it is (Article, Component); `""` for none.
    pub eyebrow: &'a str,
    /// Its title, the page's `<h1>` and its current crumb. Required: an
    /// empty one panics.
    pub title: &'a str,
    /// What it is about, under the title; `""` for none.
    pub abstract_: &'a str,
    /// The page's utilities on the hero; `""` for none.
    pub class: &'a str,
}

impl Renderable for DocsHero<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.title.is_empty(),
            "ciderui::docs::DocsHero needs a title"
        );
        rsx! {
            <header class=(Class("cui-docs-hero", self.class))>
                @if !self.crumbs.is_empty() {
                    <nav aria-label="Breadcrumbs">
                        <ol class="cui-docs-hero-crumbs">
                            @for crumb in self.crumbs {
                                <li><a href=(crumb.href)>(crumb.label)</a></li>
                            }
                            <li aria-current="page">(self.title)</li>
                        </ol>
                    </nav>
                }
                @if !self.eyebrow.is_empty() {
                    <p class="cui-docs-hero-eyebrow">(self.eyebrow)</p>
                }
                <h1 class="cui-docs-hero-title">(self.title)</h1>
                @if !self.abstract_.is_empty() {
                    <p class="cui-docs-hero-abstract">(self.abstract_)</p>
                }
            </header>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    #[test]
    fn renders_the_example() {
        let crumbs = [
            DocsCrumb {
                label: "Crash reports",
                href: "#crash-reports",
            },
            DocsCrumb {
                label: "Diagnosing crashes",
                href: "#diagnosing",
            },
        ];
        let hero = rsx! {
            <DocsHero crumbs=(&crumbs) eyebrow="Article" title="Reading a crash report" abstract_="Find the thread that crashed, the frame that failed and the images it had loaded." class="" />
        }
        .render();
        assert_eq!(
            *hero.as_inner(),
            example(include_str!(
                "../../../components/docs/docs-hero/examples/hero.html"
            ))
        );
    }

    #[test]
    fn a_page_at_the_top_has_no_crumbs() {
        let hero = rsx! {
            <DocsHero crumbs=(&[]) eyebrow="" title="Get Started" abstract_="" class="" />
        }
        .render();
        assert_eq!(
            *hero.as_inner(),
            r#"<header class="cui-docs-hero"><h1 class="cui-docs-hero-title">Get Started</h1></header>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a title")]
    fn an_empty_title_panics() {
        let _ = rsx! {
            <DocsHero crumbs=(&[]) eyebrow="" title="" abstract_="" class="" />
        }
        .render();
    }
}

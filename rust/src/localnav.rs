use hypertext::{Buffer, prelude::*};

use crate::{Class, icon::CHEVRON};

// the menu's checkbox: a page has one local nav
const TOGGLE_ID: &str = "cui-localnav-menu";

/// `cui-localnav` (`components/localnav`): the bar under the global nav on a
/// product's pages, stuck to the top as the page scrolls: the product's name,
/// its pages' links (the current one marked) and a button; to 833px the links
/// are a menu its chevron opens below the bar, with no script (a checkbox,
/// the chevron its label). No links, no menu.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::LocalNav;
///
/// let html = rsx! {
///     <LocalNav title="Cider Book 13" href="/cider-book/" links=(&[("Overview", "/cider-book/")]) current="/cider-book/" button=(Some(("Buy", "/buy/"))) dark=false class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<a class="cui-localnav-link" href="/cider-book/" aria-current="page">Overview</a>"#));
/// assert!(html.as_inner().contains(r#"<a class="cui-btn" data-size="sm" href="/buy/">Buy</a>"#));
/// ```
pub struct LocalNav<'a> {
    /// The product's name. Required: an empty one panics.
    pub title: &'a str,
    /// Where the name goes, the product's main page. Required.
    pub href: &'a str,
    /// The product's pages, (text, href) each, in order; none, no menu.
    pub links: &'a [(&'a str, &'a str)],
    /// The current page's href: its link is marked (`aria-current="page"`);
    /// `""` for none.
    pub current: &'a str,
    /// The button after the links, (text, href): a `cui-btn`,
    /// `data-size="sm"`.
    pub button: Option<(&'a str, &'a str)>,
    /// Whether the nav is dark (`cui-dark`).
    pub dark: bool,
    /// The page's utilities on the nav; `""` for none.
    pub class: &'a str,
}

impl Renderable for LocalNav<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.title.is_empty(), "ciderui::LocalNav needs a title");
        assert!(!self.href.is_empty(), "ciderui::LocalNav needs an href");
        let block = if self.dark {
            "cui-localnav cui-dark"
        } else {
            "cui-localnav"
        };
        let menu = !self.links.is_empty();
        rsx! {
            <nav class=(Class(block, self.class)) aria-label="Local">
                <div class="cui-localnav-content">
                    <a class="cui-localnav-title" href=(self.href)>(self.title)</a>
                    @if menu {
                        <input class="cui-localnav-toggle" type="checkbox" id=(TOGGLE_ID) aria-label="Menu">
                        <div class="cui-localnav-menu">
                            <ul class="cui-localnav-list">
                                @for (text, href) in self.links {
                                    <li><a class="cui-localnav-link" href=(href) aria-current="page"[*href == self.current]>(text)</a></li>
                                }
                            </ul>
                        </div>
                    }
                    <div class="cui-localnav-actions">
                        @if menu {
                            <label class="cui-localnav-menubutton" for=(TOGGLE_ID) aria-hidden="true">(CHEVRON)</label>
                        }
                        @if let Some((text, href)) = self.button {
                            <a class="cui-btn" data-size="sm" href=(href)>(text)</a>
                        }
                    </div>
                </div>
            </nav>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    const LINKS: &[(&str, &str)] = &[
        ("Overview", "/cider-book/"),
        ("Tech Specs", "/cider-book/specs/"),
        ("Compare", "/cider-book/compare/"),
        ("Switch to Cider Book", "/cider-book/switch/"),
    ];

    #[test]
    fn renders_the_example() {
        let html = rsx! {
            <LocalNav title="Cider Book 13" href="/cider-book/" links=(LINKS) current="/cider-book/" button=(Some(("Buy", "/buy/cider-book/"))) dark=false class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/localnav/examples/localnav.html"
            ))
        );
    }

    #[test]
    fn no_links_no_menu() {
        let html = rsx! {
            <LocalNav title="AirTag" href="/airtag/" links=(&[]) current="" button=None dark=true class="" />
        }
        .render();
        assert_eq!(
            html.as_inner(),
            concat!(
                r#"<nav class="cui-localnav cui-dark" aria-label="Local"><div class="cui-localnav-content">"#,
                r#"<a class="cui-localnav-title" href="/airtag/">AirTag</a><div class="cui-localnav-actions"></div></div></nav>"#,
            )
        );
    }

    #[test]
    fn marks_no_link_without_a_current_page() {
        let html = rsx! {
            <LocalNav title="Cider Book 13" href="/cider-book/" links=(LINKS) current="" button=None dark=false class="" />
        }
        .render();
        assert!(!html.as_inner().contains("aria-current"));
    }

    #[test]
    #[should_panic(expected = "needs a title")]
    fn an_empty_title_panics() {
        let _ = rsx! {
            <LocalNav title="" href="/x/" links=(&[]) current="" button=None dark=false class="" />
        }
        .render();
    }

    #[test]
    fn dark_then_the_pages_utilities() {
        let html = rsx! {
            <LocalNav title="x" href="/x/" links=(&[]) current="" button=None dark=true class="mt-0" />
        }
        .render();
        assert!(
            html.as_inner()
                .starts_with(r#"<nav class="cui-localnav cui-dark mt-0" aria-label="Local">"#)
        );
    }
}

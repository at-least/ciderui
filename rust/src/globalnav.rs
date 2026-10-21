use hypertext::{Buffer, prelude::*};

use crate::icon::{
    ARROW, ARROW_COMPACT, BAG, BAG_COMPACT, CLEAR, CLEAR_COMPACT, MENU, SEARCH, SEARCH_COMPACT,
    SEARCHFIELD, SEARCHFIELD_COMPACT,
};

// the list's id, which its button opens, and the search's: a page has one
// global nav
const MENU_ID: &str = "cui-globalnav-menu";
const SEARCH_ID: &str = "cui-globalnav-search";
const HEADING_ID: &str = "cui-globalnav-searchheading";

/// The global nav's search ([`GlobalNav`]'s `search`).
pub enum Search<'a> {
    /// A link to the site's search page.
    Link(&'a str),
    /// A button opening the search with no script: its field, focused, sent
    /// as `q` (GET) to `action`, the site's search page, its clear button
    /// shown while it holds text, and under it the
    /// quick links, (text, href) each, under the heading Quick Links (none,
    /// no heading); from 834px a flyout under the bar, to 833px a sheet over
    /// the window.
    Flyout {
        /// The site's search page. Required: an empty one panics.
        action: &'a str,
        /// The quick links, in order; `&[]` for none.
        links: &'a [(&'a str, &'a str)],
    },
}

/// `cui-globalnav` (`components/globalnav`): the bar across the top of the
/// page, the site's logo, its links and its search and bag; to 833px the
/// links are a menu its button opens over the window, with no script (the
/// list a popover beside the button, with a close button of its own that
/// takes the focus). Not placed on the page: a site adds `sticky top-0` (its
/// `class`) where the nav stays.
///
/// ```
/// use hypertext::{Raw, prelude::*};
/// use ciderui::GlobalNav;
///
/// let logo = Raw::<_>::dangerously_create(r#"<img src="/logo.svg" alt="" width="14" height="44">"#);
/// let html = rsx! {
///     <GlobalNav logo=(logo) home="/" name="Cider" links=(&[("Store", "/store")]) search=None bag=None dark=false clear=false class="sticky top-0" />
/// }
/// .render();
/// assert!(html.as_inner().starts_with(concat!(
///     r#"<nav class="cui-globalnav sticky top-0" aria-label="Global"><div class="cui-globalnav-content">"#,
///     r#"<ul class="cui-globalnav-list" id="cui-globalnav-menu" popover>"#,
/// )));
/// ```
pub struct GlobalNav<'a, L: Renderable> {
    /// The site's logo, its image (an `<svg>` in `currentColor` takes the
    /// bar's color), in the logo's link.
    pub logo: L,
    /// Where the logo goes, the site's home. Required: an empty one panics.
    pub home: &'a str,
    /// The site's name, the logo link's accessible name. Required.
    pub name: &'a str,
    /// The links, (text, href) each, in order. At least one: a menu of
    /// nothing panics.
    pub links: &'a [(&'a str, &'a str)],
    /// The site's search, if it has one ([`SEARCH`]): a link to its page,
    /// or the search opened over the page ([`Search`]).
    pub search: Option<Search<'a>>,
    /// The site's bag, if it has one ([`BAG`]).
    pub bag: Option<&'a str>,
    /// Whether the nav is dark (`cui-dark`): the dark bar and menu.
    pub dark: bool,
    /// Whether the bar is clear (`data-variant="clear"`): no ground and no
    /// blur, over a hero.
    pub clear: bool,
    /// The page's utilities on the nav; `""` for none.
    pub class: &'a str,
}

impl<L: Renderable> Renderable for GlobalNav<'_, L> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.home.is_empty(), "ciderui::GlobalNav needs a home");
        assert!(!self.name.is_empty(), "ciderui::GlobalNav needs a name");
        assert!(!self.links.is_empty(), "ciderui::GlobalNav needs links");
        if let Some(Search::Flyout { action, .. }) = self.search {
            assert!(
                !action.is_empty(),
                "ciderui::GlobalNav's search needs its page (action)"
            );
        }
        let mut class = String::from("cui-globalnav");
        if self.dark {
            class.push_str(" cui-dark");
        }
        if !self.class.is_empty() {
            class.push(' ');
            class.push_str(self.class);
        }
        rsx! {
            <nav class=(class) data-variant="clear"[self.clear] aria-label="Global">
                <div class="cui-globalnav-content">
                    <ul class="cui-globalnav-list" id=(MENU_ID) popover>
                        <li><a class="cui-globalnav-logo" href=(self.home) aria-label=(self.name)>(self.logo)</a></li>
                        <li><button class="cui-globalnav-closebutton" popovertarget=(MENU_ID) popovertargetaction="hide" aria-label="Close" autofocus>(MENU)</button></li>
                        @for (text, href) in self.links {
                            <li><a class="cui-globalnav-link" href=(href)>(text)</a></li>
                        }
                        @if let Some(Search::Link(href)) = self.search {
                            <li><a class="cui-globalnav-icon" href=(href) aria-label="Search">(SEARCH.class("cui-globalnav-image"))(SEARCH_COMPACT.class("cui-globalnav-image-compact"))</a></li>
                        }
                        @if let Some(Search::Flyout { .. }) = self.search {
                            <li><button class="cui-globalnav-icon" popovertarget=(SEARCH_ID) aria-label="Search">(SEARCH.class("cui-globalnav-image"))(SEARCH_COMPACT.class("cui-globalnav-image-compact"))</button></li>
                        }
                        @if let Some(href) = self.bag {
                            <li><a class="cui-globalnav-icon" href=(href) aria-label="Bag">(BAG.class("cui-globalnav-image"))(BAG_COMPACT.class("cui-globalnav-image-compact"))</a></li>
                        }
                    </ul>
                    <button class="cui-globalnav-menubutton" popovertarget=(MENU_ID) aria-label="Menu">(MENU)</button>
                    @if let Some(Search::Flyout { action, links }) = self.search {
                        <div class="cui-globalnav-search" id=(SEARCH_ID) popover>
                            <div class="cui-globalnav-searchpanel">
                                <form class="cui-globalnav-searchfield" action=(action) method="get" role="search">
                                    (SEARCHFIELD.class("cui-globalnav-image"))(SEARCHFIELD_COMPACT.class("cui-globalnav-image-compact"))
                                    <input class="cui-globalnav-searchinput" type="text" name="q" placeholder="Search" aria-label="Search" autocomplete="off" autocapitalize="off" autofocus>
                                    <button class="cui-globalnav-searchreset" type="reset" aria-label="Clear search">(CLEAR.class("cui-globalnav-image"))(CLEAR_COMPACT.class("cui-globalnav-image-compact"))</button>
                                </form>
                                @if !links.is_empty() {
                                    <h2 class="cui-globalnav-searchheading" id=(HEADING_ID)>"Quick Links"</h2>
                                    <ul class="cui-globalnav-searchlinks" aria-labelledby=(HEADING_ID)>
                                        @for (text, href) in links {
                                            <li><a class="cui-globalnav-searchlink" href=(href)>(ARROW.class("cui-globalnav-image"))(ARROW_COMPACT.class("cui-globalnav-image-compact"))(text)</a></li>
                                        }
                                    </ul>
                                }
                                <button class="cui-globalnav-closebutton" popovertarget=(SEARCH_ID) popovertargetaction="hide" aria-label="Close">(MENU)</button>
                            </div>
                            <button class="cui-globalnav-curtain" popovertarget=(SEARCH_ID) popovertargetaction="hide" tabindex="-1" aria-hidden="true"></button>
                        </div>
                    }
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
    use crate::icon::{Icon, SUGGESTION, SUGGESTION_COMPACT};
    use hypertext::Raw;

    const LINKS: &[(&str, &str)] = &[
        ("Store", "/store"),
        ("Mac", "/mac"),
        ("iPad", "/ipad"),
        ("Support", "/support"),
    ];

    fn logo() -> Raw<&'static str> {
        Raw::<_>::dangerously_create(
            r#"<svg aria-hidden="true" width="14" height="44" viewBox="0 0 14 44"><circle cx="7" cy="22" r="6" fill="currentColor"/></svg>"#,
        )
    }

    #[test]
    fn renders_the_example() {
        let html = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(LINKS) search=(Some(Search::Link("/search"))) bag=(Some("/bag")) dark=false clear=false class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/globalnav/examples/globalnav.html"
            ))
        );
    }

    const QUICK: &[(&str, &str)] = &[
        ("Find a Store", "/retail"),
        ("Mac", "/mac"),
        ("iPad", "/ipad"),
        ("Gift Cards", "/gift-cards"),
        ("Support", "/support"),
    ];

    #[test]
    fn renders_the_search_example() {
        let html = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(LINKS) search=(Some(Search::Flyout { action: "/search", links: QUICK })) bag=(Some("/bag")) dark=false clear=false class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/globalnav/examples/search.html"
            ))
        );
    }

    // the suggestions' example is the search's as a site's script leaves it,
    // its sections in the quick links' place, each row after the crate's
    // icons: a link's arrow, a suggested search's magnifier
    #[test]
    fn the_suggestions_example_is_the_search_with_its_sections() {
        let search = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(LINKS) search=(Some(Search::Flyout { action: "/search", links: QUICK })) bag=(Some("/bag")) dark=false clear=false class="" />
        }
        .render();
        let search = search.as_inner();
        let suggestions = example(include_str!(
            "../../components/globalnav/examples/suggestions.html"
        ));
        // from the first heading to the last list's end
        let sections = |html: &str| {
            let start = html
                .find(r#"<h2 class="cui-globalnav-searchheading""#)
                .unwrap();
            (start, html.rfind("</ul>").unwrap() + "</ul>".len())
        };
        let ((a, b), (c, d)) = (sections(search), sections(&suggestions));
        assert_eq!(
            format!(
                "{}{}{}",
                &suggestions[..c],
                &search[a..b],
                &suggestions[d..]
            ),
            *search
        );
        let images = |regular: Icon, compact: Icon| {
            rsx! { (regular.class("cui-globalnav-image"))(compact.class("cui-globalnav-image-compact")) }
                .render()
                .into_inner()
        };
        assert!(suggestions.contains(&format!(
            r#"href="/mac">{}<span><mark>Ma</mark>c</span>"#,
            images(ARROW, ARROW_COMPACT)
        )));
        assert!(suggestions.contains(&format!(
            r#"href="/search?q=mac">{}<span><mark>ma</mark>c</span>"#,
            images(SUGGESTION, SUGGESTION_COMPACT)
        )));
    }

    #[test]
    fn a_search_with_no_quick_links_has_no_heading() {
        let html = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(LINKS) search=(Some(Search::Flyout { action: "/find?a&b", links: &[] })) bag=None dark=false clear=false class="" />
        }
        .render();
        assert!(html.as_inner().contains(
            r#"<form class="cui-globalnav-searchfield" action="/find?a&amp;b" method="get" role="search">"#
        ));
        assert!(!html.as_inner().contains("cui-globalnav-searchheading"));
        assert!(!html.as_inner().contains("cui-globalnav-searchlinks"));
    }

    #[test]
    #[should_panic(expected = "needs its page")]
    fn a_search_with_no_page_panics() {
        let _ = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(LINKS) search=(Some(Search::Flyout { action: "", links: &[] })) bag=None dark=false clear=false class="" />
        }
        .render();
    }

    #[test]
    fn dark_clear_and_the_pages_utilities() {
        let html = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(LINKS) search=None bag=None dark=true clear=true class="sticky top-0" />
        }
        .render();
        assert!(html.as_inner().starts_with(
            r#"<nav class="cui-globalnav cui-dark sticky top-0" data-variant="clear" aria-label="Global">"#
        ));
        assert!(!html.as_inner().contains("cui-globalnav-icon"));
    }

    #[test]
    fn escapes_its_text() {
        let html = rsx! {
            <GlobalNav logo=(logo()) home="/?a&b" name="A \"B\"" links=(&[("<b>", "/x")]) search=None bag=None dark=false clear=false class="" />
        }
        .render();
        assert!(
            html.as_inner()
                .contains(r#"href="/?a&amp;b" aria-label="A &quot;B&quot;""#)
        );
        assert!(
            html.as_inner()
                .contains(r#"<a class="cui-globalnav-link" href="/x">&lt;b&gt;</a>"#)
        );
    }

    #[test]
    #[should_panic(expected = "needs links")]
    fn no_links_panics() {
        let _ = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(&[]) search=None bag=None dark=false clear=false class="" />
        }
        .render();
    }
}

use hypertext::{Buffer, prelude::*};

use crate::icon::{
    ARROW, ARROW_COMPACT, BAG, BAG_COMPACT, CLEAR, CLEAR_COMPACT, FLYOUT, FLYOUT_BACK,
    FLYOUT_COMPACT, MENU, SEARCH, SEARCH_COMPACT, SEARCHFIELD, SEARCHFIELD_COMPACT,
};

// the list's id, which its button opens, and the search's: a page has one
// global nav; a flyout's is `cui-globalnav-flyout-<n>`, n its link's place
const MENU_ID: &str = "cui-globalnav-menu";
const SEARCH_ID: &str = "cui-globalnav-search";
const HEADING_ID: &str = "cui-globalnav-searchheading";

/// A link of the global nav ([`GlobalNav`]'s `links`): its text, where it
/// goes, and the groups of the flyout it opens, if it opens one: from 834px
/// a menu under the bar, which the pointer resting on the link shows, its
/// button (the keyboard's, and on a screen with no hover a tap on the link)
/// opens; to 833px the menu's row opens it over the menu.
///
/// ```
/// use ciderui::{FlyoutGroup, NavLink};
///
/// const MAC: &[FlyoutGroup] = &[FlyoutGroup { heading: "Explore Mac", large: &[("Explore All Mac", "/mac")], links: &[("Compare Mac", "/mac/compare")] }];
/// let links = [NavLink::new("Store", "/store"), NavLink::new("Mac", "/mac").flyout(MAC)];
/// assert!(links[1].flyout.len() == 1);
/// ```
#[derive(Clone, Copy, Debug)]
pub struct NavLink<'a> {
    /// Its text. Required.
    pub text: &'a str,
    /// Where it goes. Required.
    pub href: &'a str,
    /// Its flyout's groups, in order; `&[]` for no flyout.
    pub flyout: &'a [FlyoutGroup<'a>],
}

impl<'a> NavLink<'a> {
    /// A link with no flyout.
    pub const fn new(text: &'a str, href: &'a str) -> Self {
        NavLink {
            text,
            href,
            flyout: &[],
        }
    }

    /// The same link opening a flyout of these groups.
    pub const fn flyout(self, groups: &'a [FlyoutGroup<'a>]) -> Self {
        NavLink {
            flyout: groups,
            ..self
        }
    }
}

/// A group of a flyout ([`NavLink`]'s `flyout`): a heading over its links,
/// (text, href) each; its large links first (apple.com's elevated group, the
/// flyout's first: 24px from 834px, 28px to 833px, where its heading is not
/// shown), then its small ones.
#[derive(Clone, Copy, Debug)]
pub struct FlyoutGroup<'a> {
    /// Its heading, which names its list. Required.
    pub heading: &'a str,
    /// Its large links; `&[]` for none.
    pub large: &'a [(&'a str, &'a str)],
    /// Its small links, after the large ones; `&[]` for none.
    pub links: &'a [(&'a str, &'a str)],
}

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
/// takes the focus); a link may open a flyout ([`NavLink`]). Not placed on
/// the page: a site adds `sticky top-0` and a `z-index` over its page's own
/// (its `class`) where the nav stays.
///
/// ```
/// use hypertext::{Raw, prelude::*};
/// use ciderui::{GlobalNav, NavLink};
///
/// let logo = Raw::<_>::dangerously_create(r#"<img src="/logo.svg" alt="" width="14" height="44">"#);
/// let html = rsx! {
///     <GlobalNav logo=(logo) home="/" name="Cider" links=(&[NavLink::new("Store", "/store")]) search=None bag=None dark=false clear=false class="sticky top-0 z-10" />
/// }
/// .render();
/// assert!(html.as_inner().starts_with(concat!(
///     r#"<nav class="cui-globalnav sticky top-0 z-10" aria-label="Global"><div class="cui-globalnav-content">"#,
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
    /// The links, in order, each perhaps opening a flyout. At least one: a
    /// menu of nothing panics.
    pub links: &'a [NavLink<'a>],
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
        for link in self.links {
            assert!(
                !link.text.is_empty() && !link.href.is_empty(),
                "ciderui::GlobalNav's links need their text and href"
            );
            for group in link.flyout {
                assert!(
                    !group.heading.is_empty(),
                    "ciderui::GlobalNav's flyout groups need their heading"
                );
            }
        }
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
                        @for (n, link) in self.links.iter().enumerate() {
                            @if link.flyout.is_empty() {
                                <li><a class="cui-globalnav-link" href=(link.href)>(link.text)</a></li>
                            } @else {
                                <li>(Flyout { link, n: n + 1 })</li>
                            }
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

// a link with its flyout, the `n`th link: the link, the flyout's button, the
// flyout (a popover) with its panel (the sheet's back button and X, its
// groups) and its curtain
struct Flyout<'a> {
    link: &'a NavLink<'a>,
    n: usize,
}

impl Renderable for Flyout<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        let id = format!("cui-globalnav-flyout-{}", self.n);
        let label = format!("{} menu", self.link.text);
        rsx! {
            <a class="cui-globalnav-link" href=(self.link.href)>(self.link.text)</a>
            <button class="cui-globalnav-flyoutbutton" popovertarget=(id) aria-label=(label)>(FLYOUT.class("cui-globalnav-image"))(FLYOUT_COMPACT.class("cui-globalnav-image-compact"))</button>
            <div class="cui-globalnav-flyout" id=(id) popover>
                <div class="cui-globalnav-flyoutpanel">
                    <button class="cui-globalnav-backbutton" popovertarget=(id) popovertargetaction="hide" aria-label="Main menu">(FLYOUT_BACK)</button>
                    <button class="cui-globalnav-closebutton" popovertarget=(MENU_ID) popovertargetaction="hide" aria-label="Close">(MENU)</button>
                    @for (g, group) in self.link.flyout.iter().enumerate() {
                        @let heading = format!("{id}-{}", g + 1);
                        <div class="cui-globalnav-flyoutgroup">
                            <h2 class="cui-globalnav-flyoutheading" id=(heading)>(group.heading)</h2>
                            <ul class="cui-globalnav-flyoutlinks" aria-labelledby=(heading)>
                                @for (text, href) in group.large {
                                    <li><a class="cui-globalnav-flyoutlink" data-size="lg" href=(href)>(text)</a></li>
                                }
                                @for (text, href) in group.links {
                                    <li><a class="cui-globalnav-flyoutlink" href=(href)>(text)</a></li>
                                }
                            </ul>
                        </div>
                    }
                </div>
                <button class="cui-globalnav-curtain" popovertarget=(id) popovertargetaction="hide" tabindex="-1" aria-hidden="true"></button>
            </div>
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

    const LINKS: &[NavLink] = &[
        NavLink::new("Store", "/store"),
        NavLink::new("Mac", "/mac"),
        NavLink::new("iPad", "/ipad"),
        NavLink::new("Support", "/support"),
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

    const STORE: &[FlyoutGroup] = &[
        FlyoutGroup {
            heading: "Shop",
            large: &[
                ("Shop the Latest", "/store"),
                ("Mac", "/shop/mac"),
                ("iPad", "/shop/ipad"),
                ("iPhone", "/shop/iphone"),
            ],
            links: &[("Accessories", "/shop/accessories")],
        },
        FlyoutGroup {
            heading: "Quick Links",
            large: &[],
            links: &[
                ("Find a Store", "/retail"),
                ("Order Status", "/shop/orders"),
                ("Apple Trade In", "/shop/trade-in"),
            ],
        },
    ];

    const MAC: &[FlyoutGroup] = &[
        FlyoutGroup {
            heading: "Explore Mac",
            large: &[
                ("Explore All Mac", "/mac"),
                ("MacBook Air", "/macbook-air"),
                ("MacBook Pro", "/macbook-pro"),
                ("iMac", "/imac"),
            ],
            links: &[("Compare Mac", "/mac/compare")],
        },
        FlyoutGroup {
            heading: "Shop Mac",
            large: &[],
            links: &[
                ("Shop Mac", "/shop/mac"),
                ("Mac Accessories", "/shop/mac/accessories"),
                ("Financing", "/shop/financing"),
            ],
        },
        FlyoutGroup {
            heading: "More from Mac",
            large: &[],
            links: &[
                ("Mac Support", "/support/mac"),
                ("macOS", "/macos"),
                ("Continuity", "/continuity"),
            ],
        },
    ];

    const FLYOUTS: &[NavLink] = &[
        NavLink::new("Store", "/store").flyout(STORE),
        NavLink::new("Mac", "/mac").flyout(MAC),
        NavLink::new("iPad", "/ipad"),
        NavLink::new("Support", "/support"),
    ];

    #[test]
    fn renders_the_flyout_example() {
        let html = rsx! {
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(FLYOUTS) search=(Some(Search::Link("/search"))) bag=(Some("/bag")) dark=false clear=false class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/globalnav/examples/flyout.html"
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
            <GlobalNav logo=(logo()) home="/?a&b" name="A \"B\"" links=(&[NavLink::new("<b>", "/x")]) search=None bag=None dark=false clear=false class="" />
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

use hypertext::{Buffer, prelude::*};

use crate::icon::{BAG, BAG_COMPACT, MENU, SEARCH, SEARCH_COMPACT};

// the list's id, which its button opens: a page has one global nav
const MENU_ID: &str = "cui-globalnav-menu";

/// `cui-globalnav` (`components/globalnav`): the bar across the top of the
/// page, the site's logo, its links and its search and bag; to 833px the
/// links are a menu its button opens over the window, with no script (the
/// list a popover, the button inside it). Not placed on the page: a site
/// adds `sticky top-0` (its `class`) where the nav stays.
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
    /// The site's search page, if it has one ([`SEARCH`]).
    pub search: Option<&'a str>,
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
                        <li><button class="cui-globalnav-menubutton" popovertarget=(MENU_ID) aria-label="Menu">(MENU)</button></li>
                        @for (text, href) in self.links {
                            <li><a class="cui-globalnav-link" href=(href)>(text)</a></li>
                        }
                        @if let Some(href) = self.search {
                            <li><a class="cui-globalnav-icon" href=(href) aria-label="Search">(SEARCH.class("cui-globalnav-image"))(SEARCH_COMPACT.class("cui-globalnav-image-compact"))</a></li>
                        }
                        @if let Some(href) = self.bag {
                            <li><a class="cui-globalnav-icon" href=(href) aria-label="Bag">(BAG.class("cui-globalnav-image"))(BAG_COMPACT.class("cui-globalnav-image-compact"))</a></li>
                        }
                    </ul>
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
            <GlobalNav logo=(logo()) home="/" name="Cider" links=(LINKS) search=(Some("/search")) bag=(Some("/bag")) dark=false clear=false class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/globalnav/examples/globalnav.html"
            ))
        );
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

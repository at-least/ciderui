use hypertext::{Buffer, prelude::*};

use crate::{Class, icon};

/// A section of a [`Footer`]'s directory: its title and its links.
#[derive(Clone, Copy, Debug)]
pub struct FooterSection<'a> {
    /// Its toggle's id (to 833px it opens by its title). Required: an
    /// empty one panics.
    pub id: &'a str,
    /// Its title. Required: an empty one panics.
    pub title: &'a str,
    /// Its links, (text, href).
    pub links: &'a [(&'a str, &'a str)],
}

/// `cui-footer` (`components/footer`): apple.com's footer on the alternate
/// ground, light on a dark page too: its footnotes, its directory of
/// columns of [`FooterSection`]s (to 833px a column of sections that open
/// one by one, with no script), the shop line and the legal row.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{Footer, FooterSection};
///
/// let none = ();
/// let directory: &[&[FooterSection]] = &[&[FooterSection { id: "f-shop", title: "Shop", links: &[("Store", "/store/")] }]];
/// let html = rsx! {
///     <Footer notes=none directory=(directory) shop=none copyright="Copyright © 2026 Example Inc." links=(&[("Privacy Policy", "/privacy/")]) locale=None class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<input class="cui-footer-toggle" type="checkbox" id="f-shop">"#));
/// ```
pub struct Footer<'a, N: Renderable, S: Renderable> {
    /// Its footnotes (`<p>`s, `<ol>`), or `()` for none.
    pub notes: N,
    /// Its directory: columns of sections.
    pub directory: &'a [&'a [FooterSection<'a>]],
    /// The shop line's content (its text and links), or `()` for none.
    pub shop: S,
    /// The copyright. Required: an empty one panics.
    pub copyright: &'a str,
    /// The legal links, (text, href).
    pub links: &'a [(&'a str, &'a str)],
    /// The locale's link, (text, href); `None` for none.
    pub locale: Option<(&'a str, &'a str)>,
    /// The page's utilities on the footer; `""` for none.
    pub class: &'a str,
}

/// Whether a part was given: `()` renders nothing.
fn given(part: &impl Renderable) -> bool {
    !rsx! { (part) }.render().as_inner().is_empty()
}

impl<N: Renderable, S: Renderable> Renderable for Footer<'_, N, S> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.copyright.is_empty(),
            "ciderui::Footer needs a copyright"
        );
        for section in self.directory.iter().flat_map(|c| c.iter()) {
            assert!(!section.id.is_empty(), "ciderui::FooterSection needs an id");
            assert!(
                !section.title.is_empty(),
                "ciderui::FooterSection needs a title"
            );
        }
        rsx! {
            <footer class=(Class("cui-footer", self.class))>
                <div class="cui-footer-content">
                    @if given(&self.notes) {
                        <section class="cui-footer-notes" aria-label="Footnotes">(self.notes)</section>
                    }
                    @if !self.directory.is_empty() {
                        <nav class="cui-footer-directory" aria-label="Directory">
                            @for column in self.directory {
                                <div class="cui-footer-column">
                                    @for section in column.iter() {
                                        <div class="cui-footer-section">
                                            <input class="cui-footer-toggle" type="checkbox" id=(section.id)>
                                            <h3 class="cui-footer-title"><label for=(section.id)>(section.title)(icon::FOOTER_CHEVRON)</label></h3>
                                            <ul class="cui-footer-list">
                                                @for (text, href) in section.links {
                                                    <li><a href=(href)>(text)</a></li>
                                                }
                                            </ul>
                                        </div>
                                    }
                                </div>
                            }
                        </nav>
                    }
                    <div class="cui-footer-legal">
                        @if given(&self.shop) {
                            <p class="cui-footer-shop">(self.shop)</p>
                        }
                        <div class="cui-footer-row">
                            <p class="cui-footer-copyright">(self.copyright)</p>
                            @if !self.links.is_empty() {
                                <ul class="cui-footer-links">
                                    @for (text, href) in self.links {
                                        <li><a href=(href)>(text)</a></li>
                                    }
                                </ul>
                            }
                            @if let Some((text, href)) = self.locale {
                                <a class="cui-footer-locale" href=(href)>(text)</a>
                            }
                        </div>
                    </div>
                </div>
            </footer>
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
        let notes = rsx! {
            <p>"Trade-in values vary. See "<a href="/shop/trade-in/">"apple.com/shop/trade-in"</a>" for details."</p>
        };
        let shop = rsx! { "More ways to shop: "<a href="/retail/">"Find a store"</a>" near you." };
        let directory: &[&[FooterSection]] = &[
            &[
                FooterSection {
                    id: "footer-shop",
                    title: "Shop and Learn",
                    links: &[
                        ("Store", "/store/"),
                        ("Mac", "/mac/"),
                        ("iPhone", "/iphone/"),
                    ],
                },
                FooterSection {
                    id: "footer-wallet",
                    title: "Apple Wallet",
                    links: &[("Wallet", "/wallet/"), ("Apple Card", "/apple-card/")],
                },
            ],
            &[FooterSection {
                id: "footer-account",
                title: "Account",
                links: &[
                    ("Manage Your Account", "/account/"),
                    ("Store Account", "/shop/account/"),
                ],
            }],
        ];
        let links = &[
            ("Privacy Policy", "/privacy/"),
            ("Terms of Use", "/terms/"),
            ("Site Map", "/sitemap/"),
        ];
        let html = rsx! {
            <Footer notes=(notes) directory=(directory) shop=(shop) copyright="Copyright © 2026 Example Inc. All rights reserved." links=(links) locale=(Some(("United States", "/choose-country/"))) class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!("../../components/footer/examples/footer.html"))
        );
    }

    #[test]
    fn no_notes_no_directory_no_shop() {
        let none = ();
        let html = rsx! {
            <Footer notes=none directory=(&[]) shop=none copyright="© 2026" links=(&[]) locale=None class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<footer class="cui-footer"><div class="cui-footer-content"><div class="cui-footer-legal"><div class="cui-footer-row"><p class="cui-footer-copyright">© 2026</p></div></div></div></footer>"#
        );
    }

    #[test]
    #[should_panic(expected = "FooterSection needs an id")]
    fn a_section_with_no_id_panics() {
        let none = ();
        let directory: &[&[FooterSection]] = &[&[FooterSection {
            id: "",
            title: "x",
            links: &[],
        }]];
        let _ = rsx! {
            <Footer notes=none directory=(directory) shop=none copyright="©" links=(&[]) locale=None class="" />
        }
        .render();
    }
}

use hypertext::{Buffer, RenderableExt, prelude::*};

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

/// `cui-footer` (`components/site/footer`): apple.com's footer on the alternate
/// ground, light on a dark page too: its footnotes, its directory of
/// columns of [`FooterSection`]s (to 833px a column of sections that open
/// one by one, with no script), the shop line and the legal row.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::site::{Footer, FooterSection};
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

impl<N: Renderable, S: Renderable> Renderable for Footer<'_, N, S> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.copyright.is_empty(),
            "ciderui::site::Footer needs a copyright"
        );
        for section in self.directory.iter().flat_map(|c| c.iter()) {
            assert!(
                !section.id.is_empty(),
                "ciderui::site::FooterSection needs an id"
            );
            assert!(
                !section.title.is_empty(),
                "ciderui::site::FooterSection needs a title"
            );
        }
        // rendered once: whether a part was given is whether it renders
        // anything (`()` renders nothing)
        let notes = self.notes.memoize();
        let shop = self.shop.memoize();
        rsx! {
            <footer class=(Class("cui-footer", self.class))>
                <div class="cui-footer-content">
                    @if !notes.as_str().is_empty() {
                        <section class="cui-footer-notes" aria-label="Footnotes">(notes)</section>
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
                        @if !shop.as_str().is_empty() {
                            <p class="cui-footer-shop">(shop)</p>
                        }
                        <div class="cui-footer-row">
                            <p class="cui-footer-copyright">(self.copyright)</p>
                            @if !self.links.is_empty() {
                                <ul class="cui-footer-links">
                                    @for (n, (text, href)) in self.links.iter().enumerate() {
                                        // a space between them, as apple.com's: its
                                        // inline items' gap is the space's too
                                        @if n > 0 {
                                            " "
                                        }
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
            <p>"Credit for a trade-in depends on the device. Details at "<a href="/shop/trade-in/">"example.com/shop/trade-in"</a>"."</p>
        };
        let shop =
            rsx! { "Want to see it first? "<a href="/retail/">"Visit a store"</a>" near you." };
        let directory: &[&[FooterSection]] = &[
            &[
                FooterSection {
                    id: "footer-shop",
                    title: "Products",
                    links: &[
                        ("Store", "/store/"),
                        ("Laptops", "/laptops/"),
                        ("Phones", "/phones/"),
                    ],
                },
                FooterSection {
                    id: "footer-wallet",
                    title: "Payments",
                    links: &[("Wallet", "/wallet/"), ("Cider Card", "/cider-card/")],
                },
            ],
            &[FooterSection {
                id: "footer-account",
                title: "Account",
                links: &[
                    ("Your Account", "/account/"),
                    ("Order History", "/shop/account/"),
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
            example(include_str!(
                "../../../components/site/footer/examples/footer.html"
            ))
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

    // a part is a page's: it may cost to render, or render otherwise a
    // second time
    #[test]
    fn its_notes_and_shop_rendered_once() {
        struct Counted<'a>(&'a std::cell::Cell<u32>);
        impl Renderable for Counted<'_> {
            fn render_to(&self, buffer: &mut Buffer) {
                self.0.set(self.0.get() + 1);
                "x".render_to(buffer);
            }
        }
        let notes = std::cell::Cell::new(0);
        let shop = std::cell::Cell::new(0);
        let (n, s) = (&notes, &shop);
        let _ = rsx! {
            <Footer notes=(Counted(n)) directory=(&[]) shop=(Counted(s)) copyright="©" links=(&[]) locale=None class="" />
        }
        .render();
        assert_eq!((notes.get(), shop.get()), (1, 1));
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

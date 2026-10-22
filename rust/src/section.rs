use hypertext::{Buffer, prelude::*};

use crate::MoreLink;

/// A [`Section`]'s ground (`data-variant`).
#[derive(Clone, Copy, Debug, Default, PartialEq, Eq)]
pub enum Ground {
    /// No ground of its own: the page's.
    #[default]
    Page,
    /// The alternate ground (`data-variant="alt"`), dark in a dark block.
    Alt,
}

/// A [`Section`]'s rhythm, its space above and below (`data-size`).
#[derive(Clone, Copy, Debug, Default, PartialEq, Eq)]
pub enum Rhythm {
    /// The iPhone page's: 56px, 96px from 735px, 112px from 1069px.
    #[default]
    Default,
    /// iPhone 18 Pro's (`data-size="lg"`): 112px, 120px, 144px.
    Large,
}

/// `cui-section` (`components/section`): a band of the page, its ground, its
/// rhythm and, with `dark`, a dark block (`cui-dark`), its content inside
/// `cui-section-content` (87.5% of the window, centered).
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{Ground, Rhythm, Section};
///
/// let html = rsx! {
///     <Section ground=(Ground::Alt) rhythm=(Rhythm::Large) dark=true class="">
///         <p>"Content"</p>
///     </Section>
/// }
/// .render();
/// assert_eq!(
///     html.as_inner(),
///     concat!(
///         r#"<section class="cui-section cui-dark" data-variant="alt" data-size="lg">"#,
///         r#"<div class="cui-section-content"><p>Content</p></div></section>"#,
///     ),
/// );
/// ```
pub struct Section<'a, R: Renderable> {
    /// Its ground.
    pub ground: Ground,
    /// Its rhythm.
    pub rhythm: Rhythm,
    /// Whether it is a dark block.
    pub dark: bool,
    /// The page's utilities on the section; `""` for none.
    pub class: &'a str,
    /// The content.
    pub children: R,
}

impl<R: Renderable> Renderable for Section<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        let mut class = String::from("cui-section");
        if self.dark {
            class.push_str(" cui-dark");
        }
        if !self.class.is_empty() {
            class.push(' ');
            class.push_str(self.class);
        }
        let alt = self.ground == Ground::Alt;
        let large = self.rhythm == Rhythm::Large;
        rsx! {
            <section class=(class) data-variant="alt"[alt] data-size="lg"[large]>
                <div class="cui-section-content">(self.children)</div>
            </section>
        }
        .render_to(buffer);
    }
}

/// `cui-section-header`: a section's headline (`cui-section-headline`, an
/// `<h2>`) and the more links beside it (`cui-section-links`), each a
/// [`MoreLink`] given as (text, href).
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::SectionHeader;
///
/// let html = rsx! { <SectionHeader headline="Get the highlights." links=(&[("Watch the film", "/film")]) class="" /> }.render();
/// assert!(html.as_inner().starts_with(
///     r#"<header class="cui-section-header"><h2 class="cui-section-headline">Get the highlights.</h2><ul class="cui-section-links"><li><a class="cui-link" data-variant="more" href="/film">"#,
/// ));
/// ```
pub struct SectionHeader<'a> {
    /// The headline. Required: an empty one panics.
    pub headline: &'a str,
    /// The links beside it, (text, href) each; none, no list.
    pub links: &'a [(&'a str, &'a str)],
    /// The page's utilities on its root; `""` for none.
    pub class: &'a str,
}

impl Renderable for SectionHeader<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.headline.is_empty(),
            "ciderui::SectionHeader needs a headline"
        );
        rsx! {
            <header class=(crate::Class("cui-section-header", self.class))>
                <h2 class="cui-section-headline">(self.headline)</h2>
                @if !self.links.is_empty() {
                    <ul class="cui-section-links">
                        @for (text, href) in self.links {
                            <li><MoreLink href=(href) text=(text) class="" /></li>
                        }
                    </ul>
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
    fn renders_the_section_example() {
        let html = rsx! {
            <Section ground=(Ground::Page) rhythm=(Rhythm::Default) dark=false class="">
                <p>"Content"</p>
            </Section>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/section/examples/section.html"
            ))
        );
    }

    #[test]
    fn renders_the_header_example() {
        let html = rsx! {
            <SectionHeader headline="Get the highlights." links=(&[("Watch the film", "#film"), ("Watch the event", "#event")]) class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/section/examples/header.html"
            ))
        );
    }

    #[test]
    fn a_header_without_links_has_no_list() {
        let html = rsx! { <SectionHeader headline="Values" links=(&[]) class="" /> }.render();
        assert_eq!(
            html.as_inner(),
            r#"<header class="cui-section-header"><h2 class="cui-section-headline">Values</h2></header>"#
        );
    }

    #[test]
    fn takes_the_pages_utilities() {
        let html =
            rsx! { <Section ground=(Ground::Page) rhythm=(Rhythm::Default) dark=false class="pt-0">"x"</Section> }
                .render();
        assert!(
            html.as_inner()
                .starts_with(r#"<section class="cui-section pt-0">"#)
        );
    }

    #[test]
    #[should_panic(expected = "needs a headline")]
    fn an_empty_headline_panics() {
        let _ = rsx! { <SectionHeader headline="" links=(&[]) class="" /> }.render();
    }

    #[test]
    fn a_header_takes_the_pages_utilities() {
        let html = rsx! { <SectionHeader headline="x" links=(&[]) class="mb-0" /> }.render();
        assert!(
            html.as_inner()
                .starts_with(r#"<header class="cui-section-header mb-0">"#)
        );
    }
}

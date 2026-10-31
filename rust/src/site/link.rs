use hypertext::{Buffer, prelude::*};

use crate::{Class, icon::MORE};

/// `cui-link`'s more link (`components/site/link`, `data-variant="more"`): a
/// standalone link, its text, and the chevron after it ([`MORE`]), not
/// underlined until hovered, then its text alone. One line of markup: a
/// space between the text and the chevron would widen the gap.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::site::MoreLink;
///
/// let html = rsx! { <MoreLink href="#learn" text="Learn more" .. /> }.render();
/// assert!(html.as_inner().starts_with(
///     r##"<a class="cui-link" data-variant="more" href="#learn"><span class="cui-link-text">Learn more</span><svg class="cui-link-icon""##,
/// ));
/// ```
#[derive(Default)]
pub struct MoreLink<'a> {
    /// Where it goes. Required: a link nowhere is not a link, so an empty
    /// `href` panics.
    pub href: &'a str,
    /// Its text (`cui-link-text`). Required, so an empty text panics.
    pub text: &'a str,
    /// The page's utilities on the link; `""` for none.
    pub class: &'a str,
}

impl Renderable for MoreLink<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        // `..` fills a missing field with "", which would draw an empty link
        assert!(
            !self.href.is_empty(),
            "ciderui::site::MoreLink needs an href"
        );
        assert!(
            !self.text.is_empty(),
            "ciderui::site::MoreLink needs a text"
        );
        rsx! {
            <a class=(Class("cui-link", self.class)) data-variant="more" href=(self.href)>
                <span class="cui-link-text">(self.text)</span>
                (MORE.class("cui-link-icon"))
            </a>
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
        let html = rsx! { <MoreLink href="#learn" text="Learn more" .. /> }.render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../../components/site/link/examples/more.html"
            ))
        );
    }

    #[test]
    fn takes_the_pages_utilities() {
        let html = rsx! { <MoreLink href="/x" text="More" class="mt-4" /> }.render();
        assert!(
            html.as_inner()
                .starts_with(r#"<a class="cui-link mt-4" data-variant="more""#)
        );
    }

    #[test]
    fn escapes_its_text() {
        let html = rsx! { <MoreLink href="/a?b&c" text="<b>" .. /> }.render();
        assert!(html.as_inner().contains(r#"href="/a?b&amp;c""#));
        assert!(
            html.as_inner()
                .contains(r#"<span class="cui-link-text">&lt;b&gt;</span>"#)
        );
    }

    #[test]
    #[should_panic(expected = "needs a text")]
    fn an_empty_text_panics() {
        let _ = rsx! { <MoreLink href="/x" .. /> }.render();
    }
}

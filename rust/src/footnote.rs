use hypertext::{Buffer, prelude::*};

/// `cui-footnote` (`components/footnote`): a footnote's mark, a link to its
/// note at the text's size and on its baseline, its digits the font's
/// numerators.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Footnote;
///
/// let html = rsx! { <p>"A full day"<Footnote href="#footnote-1" mark="1" class="" /></p> }.render();
/// assert_eq!(html.as_inner(), r##"<p>A full day<sup class="cui-footnote"><a href="#footnote-1" aria-label="Footnote 1">1</a></sup></p>"##);
/// ```
pub struct Footnote<'a> {
    /// The note's link (`#footnote-1`).
    pub href: &'a str,
    /// The mark: a number, or a symbol (◊, ※). Required: an empty one
    /// panics.
    pub mark: &'a str,
    /// The page's utilities on its root; `""` for none.
    pub class: &'a str,
}

impl Renderable for Footnote<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.mark.is_empty(), "ciderui::Footnote needs a mark");
        let label = format!("Footnote {}", self.mark);
        rsx! {
            <sup class=(crate::Class("cui-footnote", self.class))><a href=(self.href) aria-label=(label)>(self.mark)</a></sup>
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
        let html = rsx! {
            <p class="cui-body">"A full day of battery life"<Footnote href="#footnote-1" mark="1" class="" /></p>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/footnote/examples/footnote.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a mark")]
    fn an_empty_mark_panics() {
        let _ = rsx! { <Footnote href="#f" mark="" class="" /> }.render();
    }

    #[test]
    fn takes_the_pages_utilities() {
        let html = rsx! { <Footnote href="#f" mark="1" class="ml-1" /> }.render();
        assert!(
            html.as_inner()
                .starts_with(r#"<sup class="cui-footnote ml-1">"#)
        );
    }
}

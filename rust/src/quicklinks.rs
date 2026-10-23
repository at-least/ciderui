use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-quicklinks` (`components/quicklinks`): the store's row of pill links
/// (`cui-btn`, `data-variant="secondary-neutral"`) under a title, as far in
/// as a shelf's header.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::QuickLinks;
///
/// let html = rsx! {
///     <QuickLinks id="quicklinks" title="Shortcuts" links=(&[("Visit a Store", "/retail/")]) class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<li><a class="cui-btn" data-variant="secondary-neutral" href="/retail/">Visit a Store</a></li>"#));
/// ```
pub struct QuickLinks<'a> {
    /// The title's id, the nav's `aria-labelledby` (`{id}-title`).
    /// Required: an empty one panics.
    pub id: &'a str,
    /// The title. Required: an empty one panics.
    pub title: &'a str,
    /// The links, (text, href).
    pub links: &'a [(&'a str, &'a str)],
    /// The page's utilities on the nav; `""` for none.
    pub class: &'a str,
}

impl Renderable for QuickLinks<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::QuickLinks needs an id");
        assert!(!self.title.is_empty(), "ciderui::QuickLinks needs a title");
        let title_id = format!("{}-title", self.id);
        rsx! {
            <nav class=(Class("cui-quicklinks", self.class)) aria-labelledby=(title_id)>
                <h2 class="cui-quicklinks-title" id=(title_id)>(self.title)</h2>
                <ul class="cui-quicklinks-list">
                    @for (text, href) in self.links {
                        <li><a class="cui-btn" data-variant="secondary-neutral" href=(href)>(text)</a></li>
                    }
                </ul>
            </nav>
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
        let links = &[
            ("Visit a Store", "/retail/"),
            ("Track an Order", "/shop/order/list"),
            ("Help with Buying", "/shop/help"),
        ];
        let html = rsx! {
            <QuickLinks id="quicklinks" title="Shortcuts" links=(links) class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/quicklinks/examples/quicklinks.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a title")]
    fn an_empty_title_panics() {
        let _ = rsx! { <QuickLinks id="q" title="" links=(&[]) class="" /> }.render();
    }
}

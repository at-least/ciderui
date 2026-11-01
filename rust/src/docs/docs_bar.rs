use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::icon::DOCS_SIDEBAR;

/// `cui-docs-bar` (`components/docs/docs-bar`): the documentation's bar in
/// Apple's developer documentation, under the global nav: the button that
/// opens the navigator below 1024px (its `popovertarget`, the navigator's
/// id), then the documentation's title, a link home.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::docs::DocsBar;
///
/// let html = rsx! { <DocsBar title="Documentation" href="/" navigator="docs-navigator" class="" /> }.render();
/// assert!(html.as_inner().contains(r#"popovertarget="docs-navigator""#));
/// ```
pub struct DocsBar<'a> {
    /// The documentation's title, its link's words and the bar's name.
    /// Required: an empty one panics.
    pub title: &'a str,
    /// Where the title links: the documentation's first page.
    pub href: &'a str,
    /// The navigator's id, which the button opens; `""` for no navigator,
    /// and no button.
    pub navigator: &'a str,
    /// The page's utilities on the bar; `""` for none.
    pub class: &'a str,
}

impl Renderable for DocsBar<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.title.is_empty(),
            "ciderui::docs::DocsBar needs a title"
        );
        rsx! {
            <nav class=(Class("cui-docs-bar", self.class)) aria-label=(self.title)>
                @if !self.navigator.is_empty() {
                    <button class="cui-docs-bar-toggle" type="button" "popovertarget"=(self.navigator) aria-label="Navigator">(DOCS_SIDEBAR)</button>
                }
                <a class="cui-docs-bar-title" href=(self.href)>(self.title)</a>
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
        let bar = rsx! {
            <DocsBar title="Documentation" href="#documentation" navigator="docs-navigator" class="" />
        }
        .render();
        assert_eq!(
            *bar.as_inner(),
            example(include_str!(
                "../../../components/docs/docs-bar/examples/bar.html"
            ))
        );
    }

    #[test]
    fn no_navigator_no_button() {
        let bar = rsx! {
            <DocsBar title="Documentation" href="/" navigator="" class="" />
        }
        .render();
        assert_eq!(
            *bar.as_inner(),
            r#"<nav class="cui-docs-bar" aria-label="Documentation"><a class="cui-docs-bar-title" href="/">Documentation</a></nav>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a title")]
    fn an_empty_title_panics() {
        let _ = rsx! { <DocsBar title="" href="/" navigator="" class="" /> }.render();
    }
}

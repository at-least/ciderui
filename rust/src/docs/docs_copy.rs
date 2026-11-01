use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::icon::{DOCS_COPIED, DOCS_COPY};

/// `cui-docs-copy` (`components/docs/docs-copy`): a code listing with
/// Apple's developer documentation's copy button at its corner, hidden
/// until the page's script shows it: the script copies the listing's code
/// and sets the button's `data-state="copied"` for a second. The listing
/// is the page's (a `<pre>` the docs content draws).
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::docs::DocsCopy;
///
/// let html = rsx! { <DocsCopy class=""><pre><code>"cargo build"</code></pre></DocsCopy> }.render();
/// assert!(html.as_inner().contains(r#"<button class="cui-docs-copy-button" type="button" aria-label="Copy" hidden>"#));
/// ```
pub struct DocsCopy<'a, R: Renderable> {
    /// The page's utilities on the listing's box; `""` for none.
    pub class: &'a str,
    /// The listing: the page's `<pre>`.
    pub children: R,
}

impl<R: Renderable> Renderable for DocsCopy<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            <div class=(Class("cui-docs-copy", self.class))>
                (self.children)
                <button class="cui-docs-copy-button" type="button" aria-label="Copy" hidden>(DOCS_COPY.class("cui-docs-copy-idle"))(DOCS_COPIED.class("cui-docs-copy-done"))</button>
            </div>
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
            <div class="cui-docs-content">
                <DocsCopy class=""><pre><code>"@import \"ciderui/docs.css\";"</code></pre></DocsCopy>
            </div>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../../components/docs/docs-copy/examples/copy.html"
            ))
        );
    }
}

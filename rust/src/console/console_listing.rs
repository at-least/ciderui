use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-console-listing` (`components/console/console-listing`) with its
/// lines numbered (`cui-console-listing-line`): a frame's source, one line
/// highlighted, as Apple's developer documentation's tutorials mark the
/// lines that changed. A listing without numbers is one class on the
/// page's own `<pre>`.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleListing;
///
/// let html = rsx! { <ConsoleListing first_line=41 lines=(&["let a = 1;", "throw a;"]) highlighted=(Some(42)) class="" /> }.render();
/// assert_eq!(
///     html.as_inner(),
///     "<pre class=\"cui-console-listing\"><code><span class=\"cui-console-listing-line\" data-line=\"41\">let a = 1;</span>\n<span class=\"cui-console-listing-line\" data-line=\"42\" data-highlighted>throw a;</span></code></pre>"
/// );
/// ```
pub struct ConsoleListing<'a, S: AsRef<str>> {
    /// The first line's number.
    pub first_line: u32,
    /// The lines, each without its newline. Required: none panics.
    pub lines: &'a [S],
    /// The highlighted line's number; `None` for none.
    pub highlighted: Option<u32>,
    /// The page's utilities on the `<pre>`; `""` for none.
    pub class: &'a str,
}

impl<S: AsRef<str>> Renderable for ConsoleListing<'_, S> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.lines.is_empty(),
            "ciderui::console::ConsoleListing needs a line"
        );
        rsx! {
            <pre class=(Class("cui-console-listing", self.class))><code>
                @for (i, line) in self.lines.iter().enumerate() {
                    @let number = self.first_line + u32::try_from(i).expect("a listing's line count fits in u32");
                    @if i > 0 {
                        "\n"
                    }
                    <span class="cui-console-listing-line" data-line=(number) data-highlighted[self.highlighted == Some(number)]>(line.as_ref())</span>
                }
            </code></pre>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    /// An example whose listing's lines are kept as written: the harness's
    /// `example` trims each line, which would drop the code's indent and
    /// the newlines between its lines. Each line outside the `<pre>` is
    /// trimmed as `example` does.
    fn example_with_pre(html: &str) -> String {
        let start = html.find("<pre").expect("the example has a <pre>");
        let end = html.find("</pre>").expect("the example closes its <pre>") + "</pre>".len();
        let trim = |s: &str| s.lines().map(str::trim).collect::<String>();
        format!(
            "{}{}{}",
            trim(&html[..start]),
            &html[start..end],
            trim(&html[end..])
        )
    }

    #[test]
    fn renders_the_examples() {
        let lines = rsx! {
            <div class="cui-console-page max-w-xl">
                <ConsoleListing first_line=40 lines=(&["func total(_ cart: Cart) -> Int {", "    let items = cart.items", "    return items.reduce(0) { $0 + $1.price }", "}"]) highlighted=(Some(42)) class="" />
            </div>
        }
        .render();
        assert_eq!(
            *lines.as_inner(),
            example_with_pre(include_str!(
                "../../../components/console/console-listing/examples/lines.html"
            ))
        );
        let nested = rsx! {
            <div class="cui-console-page max-w-xl">
                <div class="cui-console-card">
                    <ConsoleListing first_line=41 lines=(&["    let items = cart.items", "    return items.reduce(0) { $0 + $1.price }"]) highlighted=(Some(42)) class="" />
                </div>
            </div>
        }
        .render();
        assert_eq!(
            *nested.as_inner(),
            example_with_pre(include_str!(
                "../../../components/console/console-listing/examples/nested.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a line")]
    fn no_lines_panic() {
        let none: [&str; 0] = [];
        let _ = rsx! { <ConsoleListing first_line=1 lines=(&none) highlighted=None class="" /> }
            .render();
    }
}

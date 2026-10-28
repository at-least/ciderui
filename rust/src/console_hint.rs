use hypertext::{Buffer, prelude::*};

/// `cui-console-hint` (`components/console-hint`): App Store Connect's
/// newer look's info button, a small "i" beside a word that opens a
/// popover saying what the word means. The popover is the button's
/// `popovertarget`: a click or a tap opens it, Escape or a click elsewhere
/// closes it, no script.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::ConsoleHint;
///
/// let html = rsx! { <ConsoleHint id="hint-a" label="What A means" hint="A is a." /> }.render();
/// assert!(html.as_inner().contains(r#"popovertarget="hint-a""#));
/// ```
pub struct ConsoleHint<'a> {
    /// The popover's id, unique on the page. Required: an empty one panics.
    pub id: &'a str,
    /// The button's accessible name (`aria-label`), what it is about.
    /// Required: an empty one panics.
    pub label: &'a str,
    /// What the word means: the popover's text. Required: an empty one
    /// panics.
    pub hint: &'a str,
}

impl Renderable for ConsoleHint<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::ConsoleHint needs an id");
        assert!(!self.label.is_empty(), "ciderui::ConsoleHint needs a label");
        assert!(!self.hint.is_empty(), "ciderui::ConsoleHint needs a hint");
        // the popover a <span>: a hint sits in a heading or a <p>, where a
        // <div> would close the <p>
        rsx! {
            <button type="button" class="cui-console-hint" "popovertarget"=(self.id) aria-label=(self.label)>"i"</button>
            <span class="cui-console-popover" id=(self.id) "popover">(self.hint)</span>
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
        let hint = rsx! {
            <ConsoleHint id="hint-trend" label="What Trend (7d) means" hint="Events per day over the last 7 days, whatever the window above." />
        }
        .render()
        .into_inner();
        assert_eq!(
            format!(r#"<p class="cui-console-card-title">Trend (7d){hint}</p>"#),
            example(include_str!(
                "../../components/console-hint/examples/hint.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs an id")]
    fn an_empty_id_panics() {
        let _ = rsx! { <ConsoleHint id="" label="x" hint="x" /> }.render();
    }
}

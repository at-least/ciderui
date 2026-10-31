use hypertext::{Buffer, prelude::*};

use crate::lucide;

/// `cui-console-checkbox-box` (`components/console/console-checkbox`): the box of
/// App Store Connect's newer look's checkbox, its check and its dash, which
/// goes after the page's own
/// `<input class="cui-console-checkbox-input" type="checkbox">` in a
/// `<label class="cui-console-checkbox">` (or a `<span>`, the input then
/// carrying an `aria-label`), so the input keeps whatever attributes the
/// page gives it.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleCheckboxBox;
///
/// let html = rsx! {
///     <label class="cui-console-checkbox">
///         <input class="cui-console-checkbox-input" type="checkbox" name="mute">
///         <ConsoleCheckboxBox />
///         "Mute"
///     </label>
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<span class="cui-console-checkbox-box" aria-hidden="true"><svg"#));
/// ```
pub struct ConsoleCheckboxBox;

impl Renderable for ConsoleCheckboxBox {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            <span class="cui-console-checkbox-box" aria-hidden="true">(lucide::CHECK)(lucide::MINUS)</span>
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
        let checkboxes = rsx! {
            <div class="grid gap-3">
                <label class="cui-console-checkbox">
                    <input class="cui-console-checkbox-input" type="checkbox" name="mute">
                    <ConsoleCheckboxBox />
                    "Mute this issue"
                </label>
                <label class="cui-console-checkbox">
                    <input class="cui-console-checkbox-input" type="checkbox" name="notify" checked>
                    <ConsoleCheckboxBox />
                    "Notify me when it comes back"
                </label>
                <label class="cui-console-checkbox">
                    <input class="cui-console-checkbox-input" type="checkbox" name="share" disabled>
                    <ConsoleCheckboxBox />
                    "Share with every project"
                </label>
                <span class="cui-console-checkbox">
                    <input class="cui-console-checkbox-input" type="checkbox" name="all" aria-label="Select all">
                    <ConsoleCheckboxBox />
                </span>
            </div>
        }
        .render();
        assert_eq!(
            *checkboxes.as_inner(),
            example(include_str!(
                "../../../components/console/console-checkbox/examples/checkbox.html"
            ))
        );
    }
}

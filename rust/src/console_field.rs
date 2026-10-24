use hypertext::{Buffer, prelude::*};

use crate::{Class, lucide};

/// `cui-console-select` (`components/console-field`): App Store Connect's
/// newer look's dropdown, the page's own `<select class="cui-console-field">`
/// (its children) in its box, with its chevron. A text field or a
/// multi-line one is the page's own `<input>` or `<textarea>` with the
/// class.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::ConsoleSelect;
///
/// let html = rsx! {
///     <ConsoleSelect class="">
///         <select class="cui-console-field" name="level"><option>"Any Level"</option></select>
///     </ConsoleSelect>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<span class="cui-console-select"><select class="cui-console-field" name="level">"#));
/// ```
pub struct ConsoleSelect<'a, R: Renderable> {
    /// The page's utilities on its box; `""` for none.
    pub class: &'a str,
    /// Its `<select>`.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleSelect<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            <span class=(Class("cui-console-select", self.class))>
                (self.children)
                (lucide::CHEVRONS_UP_DOWN.class("cui-console-select-chevron"))
            </span>
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
        let select = rsx! {
            <div class="cui-console-form-field">
                <label class="cui-console-label" for="level">"Level"</label>
                <ConsoleSelect class="">
                    <select class="cui-console-field" id="level" name="level">
                        <option>"Any Level"</option>
                        <option>"Fatal"</option>
                        <option>"Error"</option>
                    </select>
                </ConsoleSelect>
            </div>
        }
        .render();
        assert_eq!(
            *select.as_inner(),
            example(include_str!(
                "../../components/console-field/examples/select.html"
            ))
        );
    }
}

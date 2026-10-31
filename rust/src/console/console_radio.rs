use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-console-radio` (`components/console/console-radio`): App Store
/// Connect's newer look's radio, the input itself drawn, in its
/// `<label>` before its label (`children`), and the line under it that
/// says what the choice means. Its group's `<fieldset>` and legend are the
/// page's own.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleRadio;
///
/// let html = rsx! {
///     <ConsoleRadio name="kind" value="app" checked=true required=false disabled=false caption="" class="">"App"</ConsoleRadio>
/// }
/// .render();
/// assert_eq!(html.as_inner(), r#"<label class="cui-console-radio"><input class="cui-console-radio-input" type="radio" name="kind" value="app" checked>App</label>"#);
/// ```
pub struct ConsoleRadio<'a, R: Renderable> {
    /// Its group's name. Required: an empty one panics.
    pub name: &'a str,
    /// Its value.
    pub value: &'a str,
    /// Whether it is the chosen one.
    pub checked: bool,
    /// Whether its group needs a choice (`required`, on one radio of it).
    pub required: bool,
    /// Whether it cannot be chosen.
    pub disabled: bool,
    /// The line under its label; `""` for none.
    pub caption: &'a str,
    /// The page's utilities on its `<label>`; `""` for none.
    pub class: &'a str,
    /// Its label.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleRadio<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.name.is_empty(),
            "ciderui::console::ConsoleRadio needs a name"
        );
        rsx! {
            <label class=(Class("cui-console-radio", self.class))>
                <input class="cui-console-radio-input" type="radio" name=(self.name) value=(self.value) checked[self.checked] required[self.required] disabled[self.disabled]>
                (self.children)
                @if !self.caption.is_empty() {
                    <span class="cui-console-radio-caption">(self.caption)</span>
                }
            </label>
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
        let radios = rsx! {
            <fieldset class="grid gap-3">
                <legend class="cui-console-label">"Kind"</legend>
                <ConsoleRadio name="kind" value="service" checked=true required=false disabled=false caption="A backend or a server. Listed under Services." class="">"Service"</ConsoleRadio>
                <ConsoleRadio name="kind" value="app" checked=false required=false disabled=false caption="A mobile or desktop app. Listed under Apps." class="">"App"</ConsoleRadio>
                <ConsoleRadio name="kind" value="job" checked=false required=false disabled=true caption="" class="">"Job"</ConsoleRadio>
            </fieldset>
        }
        .render();
        assert_eq!(
            *radios.as_inner(),
            example(include_str!(
                "../../../components/console/console-radio/examples/radios.html"
            ))
        );
    }

    #[test]
    fn a_required_one_says_so() {
        let html = rsx! {
            <ConsoleRadio name="kind" value="service" checked=false required=true disabled=false caption="" class="">"Service"</ConsoleRadio>
        }
        .render();
        assert_eq!(
            html.as_inner(),
            r#"<label class="cui-console-radio"><input class="cui-console-radio-input" type="radio" name="kind" value="service" required>Service</label>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a name")]
    fn an_empty_name_panics() {
        let _ = rsx! { <ConsoleRadio name="" value="a" checked=false required=false disabled=false caption="" class="">"A"</ConsoleRadio> }.render();
    }
}

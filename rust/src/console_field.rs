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

/// `cui-console-form-field` (`components/console-field`): one field's box,
/// its label (`cui-console-label`, 5px over it) first, then the body: the
/// control the page writes, then a [`ConsoleHelp`] or a [`ConsoleError`]
/// where there is one.
///
/// The label's `for` is `id`. The control is the page's, so the page gives
/// it `id`, `aria-describedby` naming the help's and the error's ids
/// ([`console_help_id`], [`console_error_id`]) where they are written, and
/// `aria-invalid="true"` while there is an error.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{ConsoleFormField, ConsoleHelp};
///
/// let html = rsx! {
///     <ConsoleFormField id="name" label="Name">
///         <input class="cui-console-field" id="name" aria-describedby="name-help">
///         <ConsoleHelp id="name" text="Shown on every page." />
///     </ConsoleFormField>
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<p class="cui-console-help" id="name-help">"#));
/// ```
pub struct ConsoleFormField<'a, R: Renderable, L: Renderable = &'a str> {
    /// The control's id: the label's `for`.
    pub id: &'a str,
    /// The label's content, text or markup.
    pub label: L,
    /// The control, then its help or its error.
    pub children: R,
}

impl<R: Renderable, L: Renderable> Renderable for ConsoleFormField<'_, R, L> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::ConsoleFormField needs an id");
        rsx! {
            <div class="cui-console-form-field">
                <label class="cui-console-label" for=(self.id)>(self.label)</label>
                (self.children)
            </div>
        }
        .render_to(buffer);
    }
}

/// `cui-console-help`: the line under a control saying what it takes, its
/// id [`console_help_id`] for the control's `aria-describedby`.
pub struct ConsoleHelp<'a> {
    /// The control's id.
    pub id: &'a str,
    /// The help. Required: an empty one panics.
    pub text: &'a str,
}

impl Renderable for ConsoleHelp<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.text.is_empty(), "ciderui::ConsoleHelp needs a text");
        rsx! { <p class="cui-console-help" id=(console_help_id(self.id))>(self.text)</p> }
            .render_to(buffer);
    }
}

/// `cui-console-error`: the line under a control saying what is wrong, its
/// id [`console_error_id`] for the control's `aria-describedby`. Written
/// only while there is an error.
pub struct ConsoleError<'a> {
    /// The control's id.
    pub id: &'a str,
    /// The error. Required: an empty one panics.
    pub text: &'a str,
}

impl Renderable for ConsoleError<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.text.is_empty(), "ciderui::ConsoleError needs a text");
        rsx! { <p class="cui-console-error" id=(console_error_id(self.id))>(self.text)</p> }
            .render_to(buffer);
    }
}

/// The help line's id for the control `id`.
pub fn console_help_id(id: &str) -> String {
    format!("{id}-help")
}

/// The error line's id for the control `id`.
pub fn console_error_id(id: &str) -> String {
    format!("{id}-error")
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

    #[test]
    fn renders_the_fields_example() {
        let fields = rsx! {
            <div class="grid max-w-80 gap-4">
                <ConsoleFormField id="project" label="Project Name">
                    <input class="cui-console-field" id="project" value="Shop" aria-describedby=(console_help_id("project"))>
                    <ConsoleHelp id="project" text="Shown in alerts and on every page." />
                </ConsoleFormField>
                <ConsoleFormField id="search" label="Search">
                    <input class="cui-console-field" id="search" placeholder="is:unresolved level:error">
                </ConsoleFormField>
                <ConsoleFormField id="url" label="Webhook URL">
                    <input class="cui-console-field" id="url" value="htp:/example" aria-invalid="true" aria-describedby=(console_error_id("url"))>
                    <ConsoleError id="url" text="Enter a URL that starts with https://." />
                </ConsoleFormField>
                <ConsoleFormField id="dsn" label="DSN">
                    <input class="cui-console-field" id="dsn" value="https://key@crashcart.test/1" disabled>
                </ConsoleFormField>
                <ConsoleFormField id="note" label="Note">
                    <textarea class="cui-console-field" id="note" rows="3" placeholder="What did you decide?"></textarea>
                </ConsoleFormField>
            </div>
        }
        .render();
        assert_eq!(
            *fields.as_inner(),
            example(include_str!(
                "../../components/console-field/examples/fields.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "ConsoleError needs a text")]
    fn an_empty_error_panics() {
        let _ = rsx! { <ConsoleError id="a" text="" /> }.render();
    }
}

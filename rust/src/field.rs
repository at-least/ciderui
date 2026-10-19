use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-field` (`components/field`): apple.com's 56px text field, its label
/// the placeholder until the field is focused or filled, then small at its
/// top (read off the input: its placeholder is a space, so the label can
/// tell it empty).
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Field;
///
/// let html = rsx! {
///     <Field id="email" kind="email" name="email" label="Email" autocomplete="email" value="" class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<label class="cui-field-label" for="email">Email</label>"#));
/// ```
pub struct Field<'a> {
    /// The input's id, its label's `for`. Required: an empty one panics.
    pub id: &'a str,
    /// The input's type (`text`, `email`, `tel`, `password`, …).
    pub kind: &'a str,
    /// The input's name.
    pub name: &'a str,
    /// The label. Required: an empty one panics.
    pub label: &'a str,
    /// The input's `autocomplete`; `""` for none.
    pub autocomplete: &'a str,
    /// Its value; `""` for an empty field.
    pub value: &'a str,
    /// The page's utilities on the field; `""` for none.
    pub class: &'a str,
}

impl Renderable for Field<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::Field needs an id");
        assert!(!self.label.is_empty(), "ciderui::Field needs a label");
        rsx! {
            <div class=(Class("cui-field", self.class))>
                <input class="cui-field-input" id=(self.id) type=(self.kind) name=(self.name) placeholder=" " autocomplete=(self.autocomplete)[!self.autocomplete.is_empty()] value=(self.value)[!self.value.is_empty()]>
                <label class="cui-field-label" for=(self.id)>(self.label)</label>
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
    fn renders_the_examples() {
        let field = rsx! {
            <Field id="email" kind="email" name="email" label="Email" autocomplete="email" value="" class="" />
        }
        .render();
        assert_eq!(
            *field.as_inner(),
            example(include_str!("../../components/field/examples/field.html"))
        );
        let filled = rsx! {
            <Field id="zip" kind="text" name="zip" label="ZIP Code" autocomplete="postal-code" value="95014" class="" />
        }
        .render();
        assert_eq!(
            *filled.as_inner(),
            example(include_str!("../../components/field/examples/filled.html"))
        );
    }

    #[test]
    fn no_autocomplete_no_attribute() {
        let html = rsx! {
            <Field id="code" kind="text" name="code" label="Promo Code" autocomplete="" value="" class="mt-4" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<div class="cui-field mt-4"><input class="cui-field-input" id="code" type="text" name="code" placeholder=" "><label class="cui-field-label" for="code">Promo Code</label></div>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn an_empty_label_panics() {
        let _ = rsx! {
            <Field id="x" kind="text" name="x" label="" autocomplete="" value="" class="" />
        }
        .render();
    }
}

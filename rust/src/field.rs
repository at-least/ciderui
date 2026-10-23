use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::icon::{ALERT, DROPDOWN};

/// `cui-field` (`components/field`): apple.com's 56px text field, its label
/// the placeholder until the field is focused or filled, then small at its
/// top (read off the input: its placeholder is a space, so the label can
/// tell it empty); in error with its message under it, disabled, or with a
/// currency or a country code before it.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Field;
///
/// let html = rsx! {
///     <Field id="email" kind="email" name="email" label="Email" autocomplete="email" value="" error="" disabled=false prefix="" class="" />
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
    /// Its error, the message under it (`aria-invalid`, `aria-describedby`,
    /// its id the field's with `-message`); `""` for none.
    pub error: &'a str,
    /// Whether it is disabled.
    pub disabled: bool,
    /// A currency or a country code before it (`$`, `+1`), which describes
    /// it (its id the field's with `-prefix`); `""` for none.
    pub prefix: &'a str,
    /// The page's utilities on the field; `""` for none.
    pub class: &'a str,
}

impl Renderable for Field<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::Field needs an id");
        assert!(!self.label.is_empty(), "ciderui::Field needs a label");
        let message = format!("{}-message", self.id);
        let prefix = format!("{}-prefix", self.id);
        let described = described(
            (!self.prefix.is_empty()).then_some(prefix.as_str()),
            (!self.error.is_empty()).then_some(message.as_str()),
        );
        rsx! {
            <div class=(Class("cui-field", self.class))>
                <input class="cui-field-input" id=(self.id) type=(self.kind) name=(self.name) placeholder=" " autocomplete=(self.autocomplete)[!self.autocomplete.is_empty()] value=(self.value)[!self.value.is_empty()] aria-invalid="true"[!self.error.is_empty()] aria-describedby=(described)[!described.is_empty()] disabled[self.disabled]>
                <label class="cui-field-label" for=(self.id)>(self.label)</label>
                @if !self.prefix.is_empty() {
                    <span class="cui-field-prefix" id=(prefix)>(self.prefix)</span>
                }
            </div>
            @if !self.error.is_empty() {
                <p class="cui-field-message" id=(message)>(ALERT)(self.error)</p>
            }
        }
        .render_to(buffer);
    }
}

/// `cui-field` on a `select` (`components/field`): the store's dropdown, the
/// field's look with a chevron at its end; its empty choice first, chosen
/// while nothing is, its label the placeholder until a choice is made.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Select;
///
/// let html = rsx! {
///     <Select id="model" name="model" label="Choose your model" options=(&[("phone-3", "Cider Phone 3")]) selected="" error="" disabled=false class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<option value="" selected>Select</option><option value="phone-3">Cider Phone 3</option>"#));
/// ```
pub struct Select<'a> {
    /// The select's id, its label's `for`. Required: an empty one panics.
    pub id: &'a str,
    /// The select's name.
    pub name: &'a str,
    /// The label. Required: an empty one panics.
    pub label: &'a str,
    /// The choices, (value, text) each, after the empty one. At least one;
    /// no value empty, which is the empty choice's.
    pub options: &'a [(&'a str, &'a str)],
    /// The chosen choice's value; `""` for none (the empty choice). One the
    /// options do not have panics.
    pub selected: &'a str,
    /// Its error, the message under it; `""` for none.
    pub error: &'a str,
    /// Whether it is disabled.
    pub disabled: bool,
    /// The page's utilities on the field; `""` for none.
    pub class: &'a str,
}

impl Renderable for Select<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::Select needs an id");
        assert!(!self.label.is_empty(), "ciderui::Select needs a label");
        assert!(!self.options.is_empty(), "ciderui::Select needs options");
        assert!(
            self.options.iter().all(|(value, _)| !value.is_empty()),
            "ciderui::Select's options need values: the empty one is its empty choice's"
        );
        assert!(
            self.selected.is_empty()
                || self
                    .options
                    .iter()
                    .any(|(value, _)| *value == self.selected),
            "ciderui::Select's selected value is none of its options'"
        );
        let message = format!("{}-message", self.id);
        rsx! {
            <div class=(Class("cui-field", self.class))>
                <select class="cui-field-input" id=(self.id) name=(self.name) aria-invalid="true"[!self.error.is_empty()] aria-describedby=(&message)[!self.error.is_empty()] disabled[self.disabled]>
                    <option value="" selected[self.selected.is_empty()]>"Select"</option>
                    @for (value, text) in self.options {
                        <option value=(value) selected[*value == self.selected]>(text)</option>
                    }
                </select>
                <label class="cui-field-label" for=(self.id)>(self.label)</label>
                (DROPDOWN.class("cui-field-chevron"))
            </div>
            @if !self.error.is_empty() {
                <p class="cui-field-message" id=(message)>(ALERT)(self.error)</p>
            }
        }
        .render_to(buffer);
    }
}

/// `cui-field` on a `textarea` (`components/field`): a message of several
/// lines, its textarea in the field's box (`cui-field-box`), 3 lines tall
/// and taller by its handle, its label the placeholder until it is focused
/// or filled.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::TextArea;
///
/// let html = rsx! {
///     <TextArea id="comments" name="comments" label="Comments" value="" error="" disabled=false class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<div class="cui-field-box"><textarea class="cui-field-textarea" id="comments" name="comments" placeholder=" "></textarea></div>"#));
/// ```
pub struct TextArea<'a> {
    /// The textarea's id, its label's `for`. Required: an empty one panics.
    pub id: &'a str,
    /// The textarea's name.
    pub name: &'a str,
    /// The label. Required: an empty one panics.
    pub label: &'a str,
    /// Its text; `""` for an empty one.
    pub value: &'a str,
    /// Its error, the message under it; `""` for none.
    pub error: &'a str,
    /// Whether it is disabled.
    pub disabled: bool,
    /// The page's utilities on the field; `""` for none.
    pub class: &'a str,
}

impl Renderable for TextArea<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::TextArea needs an id");
        assert!(!self.label.is_empty(), "ciderui::TextArea needs a label");
        let message = format!("{}-message", self.id);
        rsx! {
            <div class=(Class("cui-field", self.class))>
                <div class="cui-field-box"><textarea class="cui-field-textarea" id=(self.id) name=(self.name) placeholder=" " aria-invalid="true"[!self.error.is_empty()] aria-describedby=(&message)[!self.error.is_empty()] disabled[self.disabled]>(self.value)</textarea></div>
                <label class="cui-field-label" for=(self.id)>(self.label)</label>
            </div>
            @if !self.error.is_empty() {
                <p class="cui-field-message" id=(message)>(ALERT)(self.error)</p>
            }
        }
        .render_to(buffer);
    }
}

// the ids that describe a field, in order, space-separated
fn described(prefix: Option<&str>, message: Option<&str>) -> String {
    [prefix, message]
        .into_iter()
        .flatten()
        .collect::<Vec<_>>()
        .join(" ")
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    const MODELS: &[(&str, &str)] = &[
        ("phone-3-plus", "Cider Phone 3 Plus"),
        ("phone-3", "Cider Phone 3"),
        ("phone-2", "Cider Phone 2"),
    ];

    #[test]
    fn renders_the_examples() {
        let cases = [
            (
                rsx! { <Field id="email" kind="email" name="email" label="Email" autocomplete="email" value="" error="" disabled=false prefix="" class="" /> }.render(),
                include_str!("../../components/field/examples/field.html"),
            ),
            (
                rsx! { <Field id="zip" kind="text" name="zip" label="ZIP Code" autocomplete="postal-code" value="10001" error="" disabled=false prefix="" class="" /> }.render(),
                include_str!("../../components/field/examples/filled.html"),
            ),
            (
                rsx! { <Field id="recipient" kind="email" name="recipient" label="Recipient Email Address" autocomplete="email" value="" error="Please enter a valid email address." disabled=false prefix="" class="" /> }.render(),
                include_str!("../../components/field/examples/error.html"),
            ),
            (
                rsx! { <Field id="sender" kind="text" name="sender" label="Sender Name" autocomplete="name" value="" error="" disabled=true prefix="" class="" /> }.render(),
                include_str!("../../components/field/examples/disabled.html"),
            ),
            (
                rsx! { <Field id="amount" kind="text" name="amount" label="Other Amount" autocomplete="" value="" error="" disabled=false prefix="$" class="" /> }.render(),
                include_str!("../../components/field/examples/prefix.html"),
            ),
            (
                rsx! { <Select id="model" name="model" label="Choose your model" options=(MODELS) selected="" error="" disabled=false class="" /> }.render(),
                include_str!("../../components/field/examples/select.html"),
            ),
            (
                rsx! { <Select id="model" name="model" label="Choose your model" options=(MODELS) selected="phone-3" error="" disabled=false class="" /> }.render(),
                include_str!("../../components/field/examples/selected.html"),
            ),
        ];
        let textarea = rsx! { <TextArea id="comments" name="comments" label="Comments" value="" error="" disabled=false class="" /> }.render();
        assert_eq!(
            *textarea.as_inner(),
            example(include_str!(
                "../../components/field/examples/textarea.html"
            ))
        );
        for (html, file) in cases {
            assert_eq!(*html.as_inner(), example(file));
        }
    }

    #[test]
    fn no_autocomplete_no_attribute() {
        let html = rsx! {
            <Field id="code" kind="text" name="code" label="Promo Code" autocomplete="" value="" error="" disabled=false prefix="" class="mt-4" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<div class="cui-field mt-4"><input class="cui-field-input" id="code" type="text" name="code" placeholder=" "><label class="cui-field-label" for="code">Promo Code</label></div>"#
        );
    }

    #[test]
    fn a_prefix_and_an_error_both_describe_it() {
        let html = rsx! {
            <Field id="phone" kind="tel" name="phone" label="Phone Number" autocomplete="tel" value="" error="Please enter a valid phone number." disabled=false prefix="+1" class="" />
        }
        .render();
        assert!(
            html.as_inner()
                .contains(r#"aria-invalid="true" aria-describedby="phone-prefix phone-message">"#)
        );
        assert!(
            html.as_inner()
                .contains(r#"<span class="cui-field-prefix" id="phone-prefix">+1</span></div><p class="cui-field-message" id="phone-message">"#)
        );
    }

    #[test]
    fn a_selects_error_and_disabled() {
        let html = rsx! {
            <Select id="state" name="state" label="State" options=(&[("CA", "California")]) selected="" error="Please choose a state." disabled=true class="" />
        }
        .render();
        assert!(html.as_inner().contains(
            r#"<select class="cui-field-input" id="state" name="state" aria-invalid="true" aria-describedby="state-message" disabled>"#
        ));
        assert!(html.as_inner().ends_with(r#"Please choose a state.</p>"#));
    }

    #[test]
    fn a_textareas_text_error_and_disabled() {
        let html = rsx! {
            <TextArea id="note" name="note" label="Message" value="Happy <birthday> & more" error="Please enter a message." disabled=true class="mt-4" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            concat!(
                r#"<div class="cui-field mt-4"><div class="cui-field-box"><textarea class="cui-field-textarea" id="note" name="note" placeholder=" " aria-invalid="true" aria-describedby="note-message" disabled>Happy &lt;birthday&gt; &amp; more</textarea></div>"#,
                r#"<label class="cui-field-label" for="note">Message</label></div><p class="cui-field-message" id="note-message">"#,
            )
            .to_owned()
                + &ALERT.render().into_inner()
                + "Please enter a message.</p>"
        );
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn an_empty_label_panics() {
        let _ = rsx! {
            <Field id="x" kind="text" name="x" label="" autocomplete="" value="" error="" disabled=false prefix="" class="" />
        }
        .render();
    }

    #[test]
    #[should_panic(expected = "none of its options")]
    fn a_selected_value_no_option_has_panics() {
        let _ = rsx! {
            <Select id="m" name="m" label="Model" options=(MODELS) selected="phone-1" error="" disabled=false class="" />
        }
        .render();
    }

    #[test]
    #[should_panic(expected = "need values")]
    fn an_option_with_no_value_panics() {
        let _ = rsx! {
            <Select id="m" name="m" label="Model" options=(&[("", "None")]) selected="" error="" disabled=false class="" />
        }
        .render();
    }
}

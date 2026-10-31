use hypertext::{Buffer, prelude::*};

use crate::{Class, icon};

/// `cui-checkbox` (`components/site/checkbox`): apple.com's 16px box before its
/// label, the input itself kept under the box (it takes the focus, the
/// click and the space bar); checked, the accent with a white check.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::site::Checkbox;
///
/// let html = rsx! {
///     <Checkbox id="remember" name="remember" label="Remember me" checked=false class="" />
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-checkbox"><input class="cui-checkbox-input" type="checkbox" id="remember" name="remember">"#));
/// ```
pub struct Checkbox<'a> {
    /// The input's id, its label's `for`. Required: an empty one panics.
    pub id: &'a str,
    /// The input's name.
    pub name: &'a str,
    /// The label. Required: an empty one panics.
    pub label: &'a str,
    /// Whether it is checked.
    pub checked: bool,
    /// The page's utilities on the checkbox; `""` for none.
    pub class: &'a str,
}

impl Renderable for Checkbox<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::site::Checkbox needs an id");
        assert!(
            !self.label.is_empty(),
            "ciderui::site::Checkbox needs a label"
        );
        rsx! {
            <div class=(Class("cui-checkbox", self.class))>
                <input class="cui-checkbox-input" type="checkbox" id=(self.id) name=(self.name) checked[self.checked]>
                <label class="cui-checkbox-label" for=(self.id)><span class="cui-checkbox-indicator" aria-hidden="true">(icon::CHECK)</span>(self.label)</label>
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
        let unchecked = rsx! {
            <Checkbox id="remember" name="remember" label="Remember me" checked=false class="" />
        }
        .render();
        assert_eq!(
            *unchecked.as_inner(),
            example(include_str!(
                "../../../components/site/checkbox/examples/checkbox.html"
            ))
        );
        let checked = rsx! {
            <Checkbox id="terms" name="terms" label="I agree to the terms" checked=true class="" />
        }
        .render();
        assert_eq!(
            *checked.as_inner(),
            example(include_str!(
                "../../../components/site/checkbox/examples/checked.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn an_empty_label_panics() {
        let _ = rsx! { <Checkbox id="x" name="x" label="" checked=false class="" /> }.render();
    }
}

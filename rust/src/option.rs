use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-option` (`components/option`): apple.com's buy page's choice, a tile
/// on a radio, its title (a detail under it) at one side and a price at the
/// other; checked, its edge the accent. (`OptionTile`: `Option` is Rust's.)
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::OptionTile;
///
/// let html = rsx! {
///     <OptionTile id="s-256" name="storage" value="256" title="256GB" detail="" price="From $1099" checked=false class="" />
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<span class="cui-option-price">From $1099</span>"#));
/// ```
pub struct OptionTile<'a> {
    /// The radio's id, its label's `for`. Required: an empty one panics.
    pub id: &'a str,
    /// The radio's name, its group's.
    pub name: &'a str,
    /// The radio's value.
    pub value: &'a str,
    /// The title. Required: an empty one panics.
    pub title: &'a str,
    /// A detail under the title; `""` for none.
    pub detail: &'a str,
    /// The price at its end; `""` for none.
    pub price: &'a str,
    /// Whether it is checked.
    pub checked: bool,
    /// The page's utilities on the option; `""` for none.
    pub class: &'a str,
}

impl Renderable for OptionTile<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::OptionTile needs an id");
        assert!(!self.title.is_empty(), "ciderui::OptionTile needs a title");
        rsx! {
            <div class=(Class("cui-option", self.class))>
                <input class="cui-option-input" type="radio" id=(self.id) name=(self.name) value=(self.value) checked[self.checked]>
                <label class="cui-option-label" for=(self.id)>
                    <span class="cui-option-title">
                        (self.title)
                        @if !self.detail.is_empty() {
                            <span class="cui-option-detail">(self.detail)</span>
                        }
                    </span>
                    @if !self.price.is_empty() {
                        <span class="cui-option-price">(self.price)</span>
                    }
                </label>
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
        let option = rsx! {
            <OptionTile id="storage-256" name="storage" value="256" title="256GB" detail="" price="From $1099" checked=false class="" />
        }
        .render();
        assert_eq!(
            *option.as_inner(),
            example(include_str!("../../components/option/examples/option.html"))
        );
        let checked = rsx! {
            <OptionTile id="model-pro" name="model" value="pro" title="iPhone 18 Pro" detail="6.3-inch display" price="" checked=true class="" />
        }
        .render();
        assert_eq!(
            *checked.as_inner(),
            example(include_str!(
                "../../components/option/examples/checked.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a title")]
    fn an_empty_title_panics() {
        let _ = rsx! {
            <OptionTile id="x" name="x" value="x" title="" detail="" price="" checked=false class="" />
        }
        .render();
    }
}

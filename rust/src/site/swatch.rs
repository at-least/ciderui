use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-swatch` (`components/site/swatch`): apple.com's buy page's color choice,
/// an `<li>` of a `<ul class="cui-swatches">`: a 28px circle of the color
/// (its children: the page's `<img class="cui-swatch-color">`, or an element
/// its style colors) in a 36px ring, on a radio; checked, the ring the
/// accent; the color's name hidden but read.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::site::Swatch;
///
/// let html = rsx! {
///     <ul class="cui-swatches">
///         <Swatch id="c-silver" name="color" value="silver" label="Silver" checked=true class="">
///             <img class="cui-swatch-color" src="/silver.png" alt="">
///         </Swatch>
///     </ul>
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<span class="cui-swatch-name">Silver</span>"#));
/// ```
pub struct Swatch<'a, R: Renderable> {
    /// The radio's id, its ring's `for`. Required: an empty one panics.
    pub id: &'a str,
    /// The radio's name, its group's.
    pub name: &'a str,
    /// The radio's value.
    pub value: &'a str,
    /// The color's name, read but not shown. Required: an empty one panics.
    pub label: &'a str,
    /// Whether it is checked.
    pub checked: bool,
    /// The page's utilities on the item; `""` for none.
    pub class: &'a str,
    /// The color: `cui-swatch-color` on an image, or on an element its
    /// style colors.
    pub children: R,
}

impl<R: Renderable> Renderable for Swatch<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::site::Swatch needs an id");
        assert!(
            !self.label.is_empty(),
            "ciderui::site::Swatch needs a label"
        );
        rsx! {
            <li class=(Class("cui-swatch", self.class))>
                <input class="cui-swatch-input" type="radio" id=(self.id) name=(self.name) value=(self.value) checked[self.checked]>
                <label class="cui-swatch-ring" for=(self.id)>
                    (self.children)
                    <span class="cui-swatch-name">(self.label)</span>
                </label>
            </li>
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
        let html = rsx! {
            <ul class="cui-swatches">
                <Swatch id="color-burgundy" name="color" value="burgundy" label="Burgundy" checked=true class="">
                    <span class="cui-swatch-color" style="background-color: #5c2a2e"></span>
                </Swatch>
                <Swatch id="color-silver" name="color" value="silver" label="Silver" checked=false class="">
                    <span class="cui-swatch-color" style="background-color: #e3e4e5"></span>
                </Swatch>
                <Swatch id="color-sky" name="color" value="sky" label="Sky Blue" checked=false class="">
                    <span class="cui-swatch-color"></span>
                </Swatch>
            </ul>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../../components/site/swatch/examples/swatches.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn an_empty_label_panics() {
        let _ = rsx! {
            <Swatch id="x" name="x" value="x" label="" checked=false class="">"x"</Swatch>
        }
        .render();
    }
}

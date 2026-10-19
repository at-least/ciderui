use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-stat` (`components/stat`): a product page's number, a caption above
/// it and one below in the secondary text; with `large`, AirPods Pro's.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Stat;
///
/// let html = rsx! { <Stat large=false before="Up to" value="25 more hours" after="of video playback" class="" /> }.render();
/// assert!(html.as_inner().contains(r#"<span class="cui-stat-value">25 more hours</span>"#));
/// ```
pub struct Stat<'a> {
    /// Whether it is a large one (`data-size="lg"`).
    pub large: bool,
    /// The caption above the value; `""` for none.
    pub before: &'a str,
    /// The value. Required: an empty one panics.
    pub value: &'a str,
    /// The caption below the value; `""` for none.
    pub after: &'a str,
    /// The page's utilities on the stat; `""` for none.
    pub class: &'a str,
}

impl Renderable for Stat<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.value.is_empty(), "ciderui::Stat needs a value");
        rsx! {
            <p class=(Class("cui-stat", self.class)) data-size="lg"[self.large]>
                @if !self.before.is_empty() {
                    <span class="cui-stat-caption">(self.before)</span>
                }
                <span class="cui-stat-value">(self.value)</span>
                @if !self.after.is_empty() {
                    <span class="cui-stat-caption">(self.after)</span>
                }
            </p>
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
        let stat = rsx! { <Stat large=false before="Up to" value="25 more hours" after="of video playback" class="" /> }.render();
        assert_eq!(
            *stat.as_inner(),
            example(include_str!("../../components/stat/examples/stat.html"))
        );
        let large = rsx! { <Stat large=true before="Removes up to" value="2x more" after="unwanted noise" class="" /> }.render();
        assert_eq!(
            *large.as_inner(),
            example(include_str!("../../components/stat/examples/large.html"))
        );
    }

    #[test]
    #[should_panic(expected = "needs a value")]
    fn an_empty_value_panics() {
        let _ = rsx! { <Stat large=false before="" value="" after="" class="" /> }.render();
    }
}

use hypertext::{Buffer, Raw, prelude::*};

use crate::{Class, console::ConsoleColor};

/// A metric's `cui-console-card` (`components/console/console-card`): App Store
/// Connect's newer look's analytics card, its title, its value, a change up
/// or down and a caption; with an `href`, all of it a link to what it counts
/// (`a.cui-console-card`). A card holding anything else (a list, a chart) is
/// the page's own `<div class="cui-console-card">`.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::{ConsoleColor, ConsoleMetric};
///
/// let html = rsx! { <ConsoleMetric title="New Issues" value="8" change="" change_color=(ConsoleColor::Neutral) caption="" href=None class="" /> }.render();
/// assert!(html.as_inner().contains(r#"<span class="cui-console-card-value">8</span>"#));
/// ```
pub struct ConsoleMetric<'a, V: Renderable = &'a str> {
    /// Its title. Required: an empty one panics.
    pub title: &'a str,
    /// Its value, text or markup (a figure in a color, a badge beside it).
    /// Required: one that renders empty panics.
    pub value: V,
    /// Its change beside the value; `""` for none.
    pub change: &'a str,
    /// The change's color: the neutral (no change), success (up, for the
    /// better) or error (for the worse); any other panics.
    pub change_color: ConsoleColor,
    /// A line under it; `""` for none.
    pub caption: &'a str,
    /// Where it goes, the whole card a link; `None` for a card that is no
    /// link.
    pub href: Option<&'a str>,
    /// The page's utilities on the card; `""` for none.
    pub class: &'a str,
}

impl<V: Renderable> Renderable for ConsoleMetric<'_, V> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.title.is_empty(),
            "ciderui::console::ConsoleMetric needs a title"
        );
        // rendered once to check it, then written as the markup it is
        let value = self.value.render().into_inner();
        assert!(
            !value.is_empty(),
            "ciderui::console::ConsoleMetric needs a value"
        );
        let value = Raw::<_>::dangerously_create(value);
        assert!(
            matches!(
                self.change_color,
                ConsoleColor::Neutral | ConsoleColor::Success | ConsoleColor::Error
            ),
            "ciderui::console::ConsoleMetric's change is neutral, success or error"
        );
        let color = self.change_color.attribute();
        let content = rsx! {
            <p class="cui-console-card-title">(self.title)</p>
            <p class="cui-console-card-figure">
                <span class="cui-console-card-value">(value)</span>
                @if !self.change.is_empty() {
                    <span class="cui-console-card-change" data-color=(color.unwrap_or(""))[color.is_some()]>(self.change)</span>
                }
            </p>
            @if !self.caption.is_empty() {
                <p class="cui-console-card-caption">(self.caption)</p>
            }
        };
        match self.href {
            Some(href) => rsx! {
                <a class=(Class("cui-console-card", self.class)) href=(href)>(content)</a>
            }
            .render_to(buffer),
            None => rsx! {
                <div class=(Class("cui-console-card", self.class))>(content)</div>
            }
            .render_to(buffer),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    #[test]
    fn renders_the_example() {
        let cards: String = [
            ("Crash-Free Sessions", "99.2%", "+0.4%", ConsoleColor::Success, "Latest release, 2.4.1"),
            ("Unhandled Errors", "1.4K", "+44%", ConsoleColor::Error, "In the last 7 days"),
            ("New Issues", "8", "No change", ConsoleColor::Neutral, "First seen in the last 7 days"),
        ]
        .iter()
        .map(|(t, v, ch, c, cap)| {
            rsx! { <ConsoleMetric title=t value=v change=ch change_color=(*c) caption=cap href=None class="" /> }
                .render()
                .into_inner()
        })
        .collect();
        assert_eq!(
            format!(r#"<div class="grid grid-cols-3 gap-5">{cards}</div>"#),
            example(include_str!(
                "../../../components/console/console-card/examples/metrics.html"
            ))
        );
    }

    #[test]
    fn renders_the_link_example() {
        let card = rsx! {
            <div class="grid grid-cols-2 gap-5">
                <ConsoleMetric title="New Issues" value="8" change="" change_color=(ConsoleColor::Neutral) caption="First seen in the last 7 days" href=(Some("#new-issues")) class="" />
            </div>
        }
        .render();
        assert_eq!(
            *card.as_inner(),
            example(include_str!(
                "../../../components/console/console-card/examples/link.html"
            ))
        );
    }

    #[test]
    fn a_value_of_markup_is_written_as_it_is() {
        let html = rsx! {
            <ConsoleMetric title="Crash-Free" value=(rsx! { <span class="text-danger">"0%"</span> }) change="" change_color=(ConsoleColor::Neutral) caption="" href=None class="" />
        }
        .render();
        assert!(
            html.as_inner().contains(
                r#"<span class="cui-console-card-value"><span class="text-danger">0%</span></span>"#
            ),
            "{}",
            html.as_inner()
        );
    }

    #[test]
    #[should_panic(expected = "needs a value")]
    fn a_value_of_empty_markup_panics() {
        let _ = rsx! { <ConsoleMetric title="x" value=(rsx! { "" }) change="" change_color=(ConsoleColor::Neutral) caption="" href=None class="" /> }.render();
    }

    #[test]
    #[should_panic(expected = "is neutral, success or error")]
    fn a_warning_change_panics() {
        let _ = rsx! { <ConsoleMetric title="x" value="1" change="x" change_color=(ConsoleColor::Warning) caption="" href=None class="" /> }.render();
    }
}

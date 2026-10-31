use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::console::ConsoleColor;

/// `cui-console-status` (`components/console/console-status`): App Store
/// Connect's status, a mark in its color before its word. The line under it
/// (`<span class="cui-console-status-detail">`) is the page's own.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::{ConsoleColor, ConsoleStatus};
///
/// let html = rsx! { <ConsoleStatus label="Delivered" color=(ConsoleColor::Success) series="" class="" /> }.render();
/// assert_eq!(html.as_inner(), r#"<span class="cui-console-status" data-color="success"><span class="cui-console-status-mark"></span>Delivered</span>"#);
/// ```
pub struct ConsoleStatus<'a> {
    /// Its word. Required: an empty one panics.
    pub label: &'a str,
    /// Its mark's color.
    pub color: ConsoleColor,
    /// A series' color for its mark, a CSS color (a chart's key,
    /// `--cui-console-status-color`), in place of `color`'s; `""` for none.
    pub series: &'a str,
    /// The page's utilities on it; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleStatus<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsoleStatus needs a label"
        );
        assert!(
            self.series.is_empty() || self.color == ConsoleColor::Neutral,
            "ciderui::console::ConsoleStatus takes a color or a series, not both"
        );
        rsx! {
            <span class=(Class("cui-console-status", self.class)) data-color=(self.color.attribute().unwrap_or(""))[self.color.attribute().is_some()]>
                @if self.series.is_empty() {
                    <span class="cui-console-status-mark"></span>
                } @else {
                    <span class="cui-console-status-mark" style={ "--cui-console-status-color: " (self.series) }></span>
                }
                (self.label)
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
    fn renders_the_examples() {
        let statuses = rsx! {
            <div class="cui-console-page grid gap-3">
                <ConsoleStatus label="Muted" color=(ConsoleColor::Neutral) series="" class="" />
                <ConsoleStatus label="New" color=(ConsoleColor::Brand) series="" class="" />
                <ConsoleStatus label="Processing" color=(ConsoleColor::Info) series="" class="" />
                <ConsoleStatus label="Delivered" color=(ConsoleColor::Success) series="" class="" />
                <ConsoleStatus label="No response" color=(ConsoleColor::Warning) series="" class="" />
                <ConsoleStatus label="Failed" color=(ConsoleColor::Error) series="" class="" />
            </div>
        }
        .render();
        assert_eq!(
            *statuses.as_inner(),
            example(include_str!(
                "../../../components/console/console-status/examples/status.html"
            ))
        );
        let detail = rsx! {
            <div class="cui-console-page">
                <div>
                    <ConsoleStatus label="Delivered" color=(ConsoleColor::Success) series="" class="" />
                    <span class="cui-console-status-detail">"200 · 410 ms"</span>
                </div>
            </div>
        }
        .render();
        assert_eq!(
            *detail.as_inner(),
            example(include_str!(
                "../../../components/console/console-status/examples/detail.html"
            ))
        );
        let key = rsx! {
            <div class="cui-console-page flex flex-wrap gap-4">
                <ConsoleStatus label="iOS" color=(ConsoleColor::Neutral) series="oklch(54.2% 0.162 251.9)" class="" />
                <ConsoleStatus label="Android" color=(ConsoleColor::Neutral) series="oklch(66.6% 0.179 58.3)" class="" />
            </div>
        }
        .render();
        assert_eq!(
            *key.as_inner(),
            example(include_str!(
                "../../../components/console/console-status/examples/key.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn an_empty_label_panics() {
        let _ =
            rsx! { <ConsoleStatus label="" color=(ConsoleColor::Neutral) series="" class="" /> }
                .render();
    }

    #[test]
    #[should_panic(expected = "a color or a series")]
    fn a_color_and_a_series_panic() {
        let _ = rsx! { <ConsoleStatus label="iOS" color=(ConsoleColor::Brand) series="red" class="" /> }
            .render();
    }
}

use hypertext::{Buffer, prelude::*};

use crate::Class;

/// A console part's color (`data-color`): the toolkit's color roles.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ConsoleColor {
    /// The neutral, the default: no `data-color`.
    Neutral,
    /// The brand's blue.
    Brand,
    /// Information's blue.
    Info,
    /// Success's green.
    Success,
    /// A warning's yellow.
    Warning,
    /// An error's red.
    Error,
}

impl ConsoleColor {
    /// Its `data-color` value; `None` for the neutral, which writes none.
    pub fn attribute(self) -> Option<&'static str> {
        match self {
            ConsoleColor::Neutral => None,
            ConsoleColor::Brand => Some("brand"),
            ConsoleColor::Info => Some("info"),
            ConsoleColor::Success => Some("success"),
            ConsoleColor::Warning => Some("warning"),
            ConsoleColor::Error => Some("error"),
        }
    }
}

/// A [`ConsoleBadge`]'s ground (`data-variant`).
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ConsoleBadgeVariant {
    /// Its color at 0.15, the default: no `data-variant`.
    Tinted,
    /// The gray fill under its colored label.
    Gray,
    /// No fill, its color alone.
    Plain,
}

/// `cui-console-badge` (`components/console-badge`): App Store Connect's
/// newer look's badge, a short word on a pill tinted in its color, on the
/// gray, or plain; with `dot`, a dot before it.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{ConsoleBadge, ConsoleBadgeVariant, ConsoleColor};
///
/// let html = rsx! { <ConsoleBadge label="Fatal" color=(ConsoleColor::Error) variant=(ConsoleBadgeVariant::Tinted) small=false dot=false class="" /> }.render();
/// assert_eq!(html.as_inner(), r#"<span class="cui-console-badge" data-color="error">Fatal</span>"#);
/// ```
pub struct ConsoleBadge<'a> {
    /// Its word. Required: an empty one panics.
    pub label: &'a str,
    /// Its color.
    pub color: ConsoleColor,
    /// Its ground.
    pub variant: ConsoleBadgeVariant,
    /// Whether it is small (`data-size="sm"`).
    pub small: bool,
    /// Whether a dot in its color goes before its word.
    pub dot: bool,
    /// The page's utilities on the badge; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleBadge<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::ConsoleBadge needs a label"
        );
        let variant = match self.variant {
            ConsoleBadgeVariant::Tinted => None,
            ConsoleBadgeVariant::Gray => Some("gray"),
            ConsoleBadgeVariant::Plain => Some("plain"),
        };
        rsx! {
            <span class=(Class("cui-console-badge", self.class)) data-color=(self.color.attribute().unwrap_or(""))[self.color.attribute().is_some()] data-variant=(variant.unwrap_or(""))[variant.is_some()] data-size="sm"[self.small]>
                @if self.dot {
                    <span class="cui-console-badge-dot"></span>
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

    const COLORS: [(ConsoleColor, &str); 6] = [
        (ConsoleColor::Neutral, "Unhandled"),
        (ConsoleColor::Brand, "New"),
        (ConsoleColor::Info, "Info"),
        (ConsoleColor::Success, "Resolved"),
        (ConsoleColor::Warning, "Warning"),
        (ConsoleColor::Error, "Fatal"),
    ];

    fn row(
        class: &str,
        badges: &[(ConsoleColor, &str)],
        ground: ConsoleBadgeVariant,
        is_small: bool,
        with_dot: bool,
    ) -> String {
        let inner: String = badges
            .iter()
            .map(|(c, word)| {
                rsx! { <ConsoleBadge label=word color=(*c) variant=ground small=is_small dot=with_dot class="" /> }
                    .render()
                    .into_inner()
            })
            .collect();
        format!(r#"<p class="{class}">{inner}</p>"#)
    }

    #[test]
    fn renders_the_examples() {
        let tinted = row(
            "flex flex-wrap gap-2",
            &COLORS,
            ConsoleBadgeVariant::Tinted,
            false,
            false,
        );
        assert_eq!(
            tinted,
            example(include_str!(
                "../../components/console-badge/examples/badges.html"
            ))
        );
        let gray = row(
            "flex flex-wrap gap-2",
            &COLORS,
            ConsoleBadgeVariant::Gray,
            false,
            false,
        );
        assert_eq!(
            gray,
            example(include_str!(
                "../../components/console-badge/examples/gray.html"
            ))
        );
        let plain = row(
            "flex flex-wrap gap-4",
            &COLORS,
            ConsoleBadgeVariant::Plain,
            false,
            true,
        );
        assert_eq!(
            plain,
            example(include_str!(
                "../../components/console-badge/examples/plain.html"
            ))
        );
        let small = row(
            "flex flex-wrap gap-2",
            &COLORS[..3],
            ConsoleBadgeVariant::Tinted,
            true,
            false,
        );
        assert_eq!(
            small,
            example(include_str!(
                "../../components/console-badge/examples/small.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn an_empty_label_panics() {
        let _ = rsx! { <ConsoleBadge label="" color=(ConsoleColor::Neutral) variant=(ConsoleBadgeVariant::Tinted) small=false dot=false class="" /> }.render();
    }
}

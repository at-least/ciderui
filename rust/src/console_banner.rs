use hypertext::{Buffer, prelude::*};

use crate::icon::Icon;
use crate::{Class, ConsoleColor, lucide};

/// `cui-console-banner` (`components/console-banner`): App Store Connect's
/// newer look's banner, a message on a rounded panel tinted in its color:
/// its icon, its title, what the page gives it under the title (a
/// `<p class="cui-console-banner-body">`, a caption, its actions), and a
/// button that closes it, which the page's script wires. An error banner is
/// `role="alert"`, any other `role="status"`.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{ConsoleBanner, ConsoleColor};
///
/// let html = rsx! {
///     <ConsoleBanner color=(ConsoleColor::Info) icon=None title="Saved" dismiss="" class="">""</ConsoleBanner>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-console-banner" data-color="info" role="status">"#));
/// ```
pub struct ConsoleBanner<'a, R: Renderable> {
    /// Its color: the neutral, info, success, warning or error. The brand
    /// has no banner: it panics.
    pub color: ConsoleColor,
    /// Its icon at its start, a [`crate::lucide`] icon; `None` for none.
    pub icon: Option<Icon>,
    /// Its title; `""` for none.
    pub title: &'a str,
    /// Its close button's label (`aria-label`); `""` for no close button.
    pub dismiss: &'a str,
    /// The page's utilities on the banner; `""` for none.
    pub class: &'a str,
    /// What goes under its title.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleBanner<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            self.color != ConsoleColor::Brand,
            "ciderui::ConsoleBanner has no brand color"
        );
        let color = self.color.attribute();
        // an error interrupts: a screen reader announces it at once
        let role = if self.color == ConsoleColor::Error {
            "alert"
        } else {
            "status"
        };
        rsx! {
            <div class=(Class("cui-console-banner", self.class)) data-color=(color.unwrap_or(""))[color.is_some()] role=(role)>
                @if let Some(icon) = self.icon {
                    (icon.class("cui-console-banner-icon"))
                }
                <div>
                    @if !self.title.is_empty() {
                        <p class="cui-console-banner-title">(self.title)</p>
                    }
                    (self.children)
                </div>
                @if !self.dismiss.is_empty() {
                    <button class="cui-console-btn cui-console-btn-icon cui-console-banner-dismiss" data-variant="plain" data-size="sm" type="button" aria-label=(self.dismiss)>(lucide::X)</button>
                }
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
        let warning = rsx! {
            <ConsoleBanner color=(ConsoleColor::Warning) icon=(Some(lucide::TRIANGLE_ALERT)) title="Symbols are missing" dismiss="Close" class="">
                <p class="cui-console-banner-body">"Upload a dSYM to see function names in this crash."</p>
                <p class="cui-console-banner-actions"><a class="cui-console-btn" data-size="sm" href="#upload">"Upload dSYM"</a></p>
            </ConsoleBanner>
        }
        .render();
        assert_eq!(
            *warning.as_inner(),
            example(include_str!(
                "../../components/console-banner/examples/banner.html"
            ))
        );
        let rows: String = [
            (
                ConsoleColor::Neutral,
                Some(lucide::INFO),
                "Data is up to date as of 9:41.",
            ),
            (
                ConsoleColor::Info,
                Some(lucide::INFO),
                "A new version of the SDK is out.",
            ),
            (
                ConsoleColor::Success,
                Some(lucide::CIRCLE_CHECK),
                "The alert channel works.",
            ),
            (
                ConsoleColor::Error,
                Some(lucide::CIRCLE_X),
                "The upload failed. Try again.",
            ),
        ]
        .iter()
        .map(|(c, i, text)| {
            rsx! {
                <ConsoleBanner color=(*c) icon=(*i) title="" dismiss="" class="">
                    <p class="cui-console-banner-body">(text)</p>
                </ConsoleBanner>
            }
            .render()
            .into_inner()
        })
        .collect();
        assert_eq!(
            format!(r#"<div class="grid gap-3">{rows}</div>"#),
            example(include_str!(
                "../../components/console-banner/examples/colors.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "has no brand color")]
    fn a_brand_banner_panics() {
        let _ = rsx! { <ConsoleBanner color=(ConsoleColor::Brand) icon=None title="" dismiss="" class="">""</ConsoleBanner> }.render();
    }
}

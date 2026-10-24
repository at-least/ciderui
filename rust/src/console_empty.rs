use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::icon::Icon;

/// A [`ConsoleEmpty`]'s size (`data-size`).
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ConsoleEmptySize {
    /// In a small place (a card's): `data-size="sm"`.
    Small,
    /// The default: no `data-size`.
    Medium,
    /// A whole page's: `data-size="lg"`.
    Large,
}

/// `cui-console-empty` (`components/console-empty`): App Store Connect's
/// newer look's empty state, centered where a list or a chart has nothing
/// to show: its icon, heading, caption, and what the page gives it, its
/// actions (`<p class="cui-console-empty-actions">`).
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{ConsoleEmpty, ConsoleEmptySize};
///
/// let html = rsx! {
///     <ConsoleEmpty size=(ConsoleEmptySize::Medium) backed=false icon=None heading="No Issues" caption="" class="">""</ConsoleEmpty>
/// }
/// .render();
/// assert_eq!(html.as_inner(), r#"<div class="cui-console-empty"><p class="cui-console-empty-heading">No Issues</p></div>"#);
/// ```
pub struct ConsoleEmpty<'a, R: Renderable> {
    /// Its size.
    pub size: ConsoleEmptySize,
    /// Whether it is on its own faint ground (`data-variant="backed"`).
    pub backed: bool,
    /// Its icon, a [`crate::lucide`] icon; `None` for none.
    pub icon: Option<Icon>,
    /// Its heading. Required: an empty one panics.
    pub heading: &'a str,
    /// The line under it, what to do next; `""` for none.
    pub caption: &'a str,
    /// The page's utilities on it; `""` for none.
    pub class: &'a str,
    /// Its actions.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleEmpty<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.heading.is_empty(),
            "ciderui::ConsoleEmpty needs a heading"
        );
        let size = match self.size {
            ConsoleEmptySize::Small => Some("sm"),
            ConsoleEmptySize::Medium => None,
            ConsoleEmptySize::Large => Some("lg"),
        };
        rsx! {
            <div class=(Class("cui-console-empty", self.class)) data-size=(size.unwrap_or(""))[size.is_some()] data-variant="backed"[self.backed]>
                @if let Some(icon) = self.icon {
                    (icon.class("cui-console-empty-icon"))
                }
                <p class="cui-console-empty-heading">(self.heading)</p>
                @if !self.caption.is_empty() {
                    <p class="cui-console-empty-caption">(self.caption)</p>
                }
                (self.children)
            </div>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;
    use crate::lucide;

    #[test]
    fn renders_the_examples() {
        let empty = rsx! {
            <ConsoleEmpty size=(ConsoleEmptySize::Medium) backed=false icon=(Some(lucide::INBOX)) heading="No Issues" caption="When your app reports a crash, it shows here." class="">
                <p class="cui-console-empty-actions"><a class="cui-console-btn" href="#setup">"Set Up Reporting"</a></p>
            </ConsoleEmpty>
        }
        .render();
        assert_eq!(
            *empty.as_inner(),
            example(include_str!(
                "../../components/console-empty/examples/empty.html"
            ))
        );
        let small = rsx! {
            <ConsoleEmpty size=(ConsoleEmptySize::Small) backed=true icon=(Some(lucide::CHART_LINE)) heading="No Data" caption="Nothing was reported in this period." class="">""</ConsoleEmpty>
        }
        .render();
        assert_eq!(
            *small.as_inner(),
            example(include_str!(
                "../../components/console-empty/examples/small.html"
            ))
        );
        let large = rsx! {
            <ConsoleEmpty size=(ConsoleEmptySize::Large) backed=false icon=(Some(lucide::FOLDER_PLUS)) heading="No Projects" caption="Create a project to get a DSN for your app." class="">
                <p class="cui-console-empty-actions"><a class="cui-console-btn" data-variant="filled" href="#new">"New Project"</a></p>
            </ConsoleEmpty>
        }
        .render();
        assert_eq!(
            *large.as_inner(),
            example(include_str!(
                "../../components/console-empty/examples/large.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a heading")]
    fn an_empty_heading_panics() {
        let _ = rsx! { <ConsoleEmpty size=(ConsoleEmptySize::Medium) backed=false icon=None heading="" caption="" class="">""</ConsoleEmpty> }.render();
    }
}

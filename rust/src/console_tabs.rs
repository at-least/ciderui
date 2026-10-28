use hypertext::{Buffer, prelude::*};

/// `cui-console-tabs` (`components/console-tabs`): a row of tabs, each a
/// link to a page.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{ConsoleTab, ConsoleTabs, ConsoleTabsVariant};
///
/// let tabs = [
///     ConsoleTab { label: "Unresolved", href: Some("?status=unresolved"), current: true },
///     ConsoleTab { label: "Resolved", href: Some("?status=resolved"), current: false },
/// ];
/// let html = rsx! { <ConsoleTabs label="Status" variant=(ConsoleTabsVariant::Filled) tabs=(&tabs) /> }.render();
/// assert!(html.as_inner().contains(r#"aria-current="true""#));
/// ```
pub struct ConsoleTabs<'a> {
    /// What the row is called (its `aria-label`). Required: an empty one
    /// panics.
    pub label: &'a str,
    /// Underlined or filled.
    pub variant: ConsoleTabsVariant,
    /// The tabs, in order.
    pub tabs: &'a [ConsoleTab<'a>],
}

/// What a row of tabs is for, which sets its look and its markup.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ConsoleTabsVariant {
    /// A page's sections, each its own page: App Store Connect's app tabs.
    /// A `<nav>`; the current tab `aria-current="page"`.
    Underlined,
    /// A choice of a few views of one thing (a list's status, a chart's
    /// period): the HIG's segmented control, Apple Developer Forums'
    /// Topics and All Posts. A group; the chosen tab `aria-current="true"`.
    Filled,
}

/// One tab of a [`ConsoleTabs`].
pub struct ConsoleTab<'a> {
    /// Its text. Required: an empty one panics.
    pub label: &'a str,
    /// The page it opens; `None` for a tab not available now
    /// (`aria-disabled`).
    pub href: Option<&'a str>,
    /// Whether it is the page or the choice shown.
    pub current: bool,
}

impl Renderable for ConsoleTabs<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.label.is_empty(), "ciderui::ConsoleTabs needs a label");
        assert!(
            self.tabs.iter().filter(|t| t.current).count() <= 1,
            "ciderui::ConsoleTabs has one current tab at most"
        );
        let current = match self.variant {
            ConsoleTabsVariant::Underlined => "page",
            ConsoleTabsVariant::Filled => "true",
        };
        let tabs = rsx! {
            @for tab in self.tabs {
                (Tab { tab, current })
            }
        };
        match self.variant {
            ConsoleTabsVariant::Underlined => rsx! {
                <nav class="cui-console-tabs" aria-label=(self.label)>(tabs)</nav>
            }
            .render_to(buffer),
            ConsoleTabsVariant::Filled => rsx! {
                <div class="cui-console-tabs" data-variant="filled" role="group" aria-label=(self.label)>(tabs)</div>
            }
            .render_to(buffer),
        }
    }
}

struct Tab<'a> {
    tab: &'a ConsoleTab<'a>,
    current: &'static str,
}

impl Renderable for Tab<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        let tab = self.tab;
        assert!(!tab.label.is_empty(), "ciderui::ConsoleTab needs a label");
        rsx! {
            @if let Some(href) = tab.href {
                <a class="cui-console-tab" href=(href) aria-current=(self.current)[tab.current]>(tab.label)</a>
            } @else {
                <a class="cui-console-tab" aria-disabled="true">(tab.label)</a>
            }
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    fn tab<'a>(label: &'a str, href: &'a str, current: bool) -> ConsoleTab<'a> {
        ConsoleTab {
            label,
            href: Some(href),
            current,
        }
    }

    #[test]
    fn renders_the_examples() {
        let underlined = [
            tab("Overview", "#overview", true),
            tab("Issues", "#issues", false),
            tab("Releases", "#releases", false),
            ConsoleTab {
                label: "Jobs",
                href: None,
                current: false,
            },
        ];
        let html = rsx! { <ConsoleTabs label="Project" variant=(ConsoleTabsVariant::Underlined) tabs=(&underlined) /> }.render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/console-tabs/examples/tabs.html"
            ))
        );
        let filled = [
            tab("24h", "#24h", false),
            tab("7d", "#7d", true),
            tab("30d", "#30d", false),
            tab("90d", "#90d", false),
        ];
        let html = rsx! { <ConsoleTabs label="Period" variant=(ConsoleTabsVariant::Filled) tabs=(&filled) /> }.render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/console-tabs/examples/filled.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "one current tab at most")]
    fn two_current_tabs_panic() {
        let tabs = [tab("A", "#a", true), tab("B", "#b", true)];
        let _ = rsx! { <ConsoleTabs label="X" variant=(ConsoleTabsVariant::Underlined) tabs=(&tabs) /> }.render();
    }
}

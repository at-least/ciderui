use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::console::{ConsoleTab, ConsoleTabs, ConsoleTabsVariant};

/// `cui-console-head` (`components/console/console-head`): what a page
/// belongs to (an app, a project) by name, a link to it, with its
/// underlined tabs beside the name, App Store Connect's app head.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::{ConsoleHead, ConsoleTab};
///
/// let tabs = [ConsoleTab { label: "Overview", href: Some("/p/shop"), current: true }];
/// let html = rsx! { <ConsoleHead name="Shop" href="/p/shop" label="Project" tabs=(&tabs) class="" /> }.render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-console-head"><a class="cui-console-head-name" href="/p/shop">Shop</a><nav class="cui-console-tabs" aria-label="Project">"#));
/// ```
pub struct ConsoleHead<'a> {
    /// What the page belongs to. Required: an empty one panics.
    pub name: &'a str,
    /// Its page.
    pub href: &'a str,
    /// What its tabs are called (their `aria-label`).
    pub label: &'a str,
    /// Its tabs, in order.
    pub tabs: &'a [ConsoleTab<'a>],
    /// The page's utilities on the row; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleHead<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.name.is_empty(),
            "ciderui::console::ConsoleHead needs a name"
        );
        rsx! {
            <div class=(Class("cui-console-head", self.class))>
                <a class="cui-console-head-name" href=(self.href)>(self.name)</a>
                <ConsoleTabs label=(self.label) variant=(ConsoleTabsVariant::Underlined) tabs=(self.tabs) class="" />
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
    fn renders_the_example() {
        let tabs = [
            ConsoleTab {
                label: "Overview",
                href: Some("#overview"),
                current: false,
            },
            ConsoleTab {
                label: "Issues",
                href: Some("#issues"),
                current: true,
            },
            ConsoleTab {
                label: "Releases",
                href: Some("#releases"),
                current: false,
            },
            ConsoleTab {
                label: "Jobs",
                href: Some("#jobs"),
                current: false,
            },
        ];
        let head = rsx! {
            <div class="cui-console-page max-w-4xl">
                <ConsoleHead name="Shop" href="#overview" label="Project" tabs=(&tabs) class="" />
                <p>"The page’s content."</p>
            </div>
        }
        .render();
        assert_eq!(
            *head.as_inner(),
            example(include_str!(
                "../../../components/console/console-head/examples/head.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a name")]
    fn an_empty_name_panics() {
        let _ =
            rsx! { <ConsoleHead name="" href="/" label="Project" tabs=(&[]) class="" /> }.render();
    }
}

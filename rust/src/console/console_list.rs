use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-console-list-item` (`components/console/console-list`): a row of
/// App Store Connect's newer look's grouped list, its label and what
/// follows it, a link where it has an `href`. The group,
/// `<div class="cui-console-list">`, is the page's own.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleListItem;
///
/// let html = rsx! { <ConsoleListItem href=(Some("/p/shop")) label="Shop" secondary="" class="" /> }.render();
/// assert_eq!(html.as_inner(), r#"<a class="cui-console-list-item" href="/p/shop"><span class="cui-console-list-label">Shop</span></a>"#);
/// ```
pub struct ConsoleListItem<'a> {
    /// Where it goes; `None` for a row that is no link.
    pub href: Option<&'a str>,
    /// Its title. Required: an empty one panics.
    pub label: &'a str,
    /// What follows it, in the secondary label; `""` for none.
    pub secondary: &'a str,
    /// The page's utilities on the row; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleListItem<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsoleListItem needs a label"
        );
        let content = rsx! {
            <span class="cui-console-list-label">(self.label)</span>
            @if !self.secondary.is_empty() {
                <span class="cui-console-list-secondary">(self.secondary)</span>
            }
        };
        match self.href {
            Some(href) => rsx! {
                <a class=(Class("cui-console-list-item", self.class)) href=(href)>(content)</a>
            }
            .render_to(buffer),
            None => rsx! {
                <div class=(Class("cui-console-list-item", self.class))>(content)</div>
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
        let list = rsx! {
            <div class="cui-console-list max-w-md">
                <ConsoleListItem href=(Some("#shop")) label="Shop App" secondary="8 open issues · crash-free 99.2%" class="" />
                <ConsoleListItem href=(Some("#api")) label="Orders API" secondary="2 open issues" class="" />
                <ConsoleListItem href=None label="nightly-backup" secondary="Last run 4 hours ago" class="" />
            </div>
        }
        .render();
        assert_eq!(
            *list.as_inner(),
            example(include_str!(
                "../../../components/console/console-list/examples/list.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn an_empty_label_panics() {
        let _ = rsx! { <ConsoleListItem href=None label="" secondary="" class="" /> }.render();
    }
}

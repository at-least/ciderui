use hypertext::{Buffer, Raw, prelude::*};

use crate::Class;

/// `cui-console-list-item` (`components/console/console-list`): a row of
/// App Store Connect's newer look's grouped list, its label and what
/// follows it, then what the page gives it (its controls, its lines), a
/// link where it has an `href`. The group,
/// `<div class="cui-console-list">`, is the page's own.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleListItem;
///
/// let html = rsx! { <ConsoleListItem href=(Some("/p/shop")) id="" label="Shop" secondary="" class="">""</ConsoleListItem> }.render();
/// assert_eq!(html.as_inner(), r#"<a class="cui-console-list-item" href="/p/shop"><span class="cui-console-list-label">Shop</span></a>"#);
/// ```
pub struct ConsoleListItem<'a, R: Renderable> {
    /// Where it goes; `None` for a row that is no link.
    pub href: Option<&'a str>,
    /// Its id; `""` for none.
    pub id: &'a str,
    /// Its title; `""` for a row whose content is all the page's own.
    pub label: &'a str,
    /// What follows it, in the secondary label; `""` for none.
    pub secondary: &'a str,
    /// The page's utilities on the row; `""` for none.
    pub class: &'a str,
    /// What the page gives it after its label; `""` for nothing. A row of
    /// no label and nothing panics.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleListItem<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        // rendered once to check it, then written as the markup it is
        let children = self.children.render().into_inner();
        assert!(
            !self.label.is_empty() || !children.is_empty(),
            "ciderui::console::ConsoleListItem needs a label or its own content"
        );
        let children = Raw::<_>::dangerously_create(children);
        let content = rsx! {
            @if !self.label.is_empty() {
                <span class="cui-console-list-label">(self.label)</span>
            }
            @if !self.secondary.is_empty() {
                <span class="cui-console-list-secondary">(self.secondary)</span>
            }
            (children)
        };
        match self.href {
            Some(href) => rsx! {
                <a class=(Class("cui-console-list-item", self.class)) id=(self.id)[!self.id.is_empty()] href=(href)>(content)</a>
            }
            .render_to(buffer),
            None => rsx! {
                <div class=(Class("cui-console-list-item", self.class)) id=(self.id)[!self.id.is_empty()]>(content)</div>
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
                <ConsoleListItem href=(Some("#shop")) id="" label="Shop App" secondary="8 open issues · crash-free 99.2%" class="">""</ConsoleListItem>
                <ConsoleListItem href=(Some("#api")) id="" label="Orders API" secondary="2 open issues" class="">""</ConsoleListItem>
                <ConsoleListItem href=None id="" label="nightly-backup" secondary="Last run 4 hours ago" class="">""</ConsoleListItem>
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
    fn renders_the_content_example() {
        let list = rsx! {
            <div class="cui-console-list max-w-md">
                <ConsoleListItem href=None id="nightly-backup" label="nightly-backup" secondary="Last run 4 hours ago" class="">
                    <p class="mt-2"><button class="cui-console-btn" type="button">"Pause"</button></p>
                </ConsoleListItem>
            </div>
        }
        .render();
        assert_eq!(
            *list.as_inner(),
            example(include_str!(
                "../../../components/console/console-list/examples/content.html"
            ))
        );
    }

    #[test]
    fn a_row_of_its_own_content_needs_no_label() {
        let html = rsx! {
            <ConsoleListItem href=None id="" label="" secondary="" class="py-3"><p>"Shop App"</p></ConsoleListItem>
        }
        .render();
        assert_eq!(
            html.as_inner(),
            r#"<div class="cui-console-list-item py-3"><p>Shop App</p></div>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a label or its own content")]
    fn a_row_of_nothing_panics() {
        let _ = rsx! { <ConsoleListItem href=None id="" label="" secondary="" class="">""</ConsoleListItem> }.render();
    }
}

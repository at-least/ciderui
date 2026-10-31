use hypertext::{Buffer, prelude::*};

use crate::icon::Icon;
use crate::{Class, lucide};

/// `cui-console-sidebar` (`components/console/console-sidebar`): App Store Connect's
/// newer look's app sidebar: its wordmark (a link home), its list (its
/// children: [`ConsoleSidebarItem`]s, headings, rules) and its foot (an
/// account); a popover, so below 1024px it is a sheet the page's
/// `cui-console-sidebar-open` button opens (`popovertarget` its `id`) and
/// its own close button closes.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleSidebar;
///
/// let html = rsx! {
///     <ConsoleSidebar id="sidebar" label="CrashCart" home="/" wordmark="CrashCart" close="Close" class="" footer="">""</ConsoleSidebar>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<nav class="cui-console-sidebar" id="sidebar" popover aria-label="CrashCart">"#));
/// ```
pub struct ConsoleSidebar<'a, R: Renderable, F: Renderable> {
    /// Its id, which the page's open button names. Required: an empty one
    /// panics.
    pub id: &'a str,
    /// What it is, for a screen reader (`aria-label`). Required.
    pub label: &'a str,
    /// Where its wordmark goes.
    pub home: &'a str,
    /// Its wordmark's text, the product's name. Required.
    pub wordmark: &'a str,
    /// Its close button's label (`aria-label`). Required.
    pub close: &'a str,
    /// The page's utilities on it; `""` for none.
    pub class: &'a str,
    /// What goes at its foot (an account); `""` for nothing.
    pub footer: F,
    /// Its list.
    pub children: R,
}

impl<R: Renderable, F: Renderable> Renderable for ConsoleSidebar<'_, R, F> {
    fn render_to(&self, buffer: &mut Buffer) {
        for (field, value) in [
            ("id", self.id),
            ("label", self.label),
            ("wordmark", self.wordmark),
            ("close", self.close),
        ] {
            assert!(
                !value.is_empty(),
                "ciderui::console::ConsoleSidebar needs a {field}"
            );
        }
        rsx! {
            <nav class=(Class("cui-console-sidebar", self.class)) id=(self.id) popover aria-label=(self.label)>
                <div class="cui-console-sidebar-header">
                    <a class="cui-console-sidebar-wordmark" href=(self.home)>(self.wordmark)</a>
                    <button class="cui-console-btn cui-console-btn-icon cui-console-sidebar-close" data-variant="plain" type="button" popovertarget=(self.id) popovertargetaction="hide" aria-label=(self.close)>(lucide::X)</button>
                </div>
                <div class="cui-console-sidebar-body">(self.children)</div>
                <div class="cui-console-sidebar-footer">(self.footer)</div>
            </nav>
        }
        .render_to(buffer);
    }
}

/// A page in a [`ConsoleSidebar`], a link.
pub struct ConsoleSidebarItem<'a> {
    /// Where it goes.
    pub href: &'a str,
    /// Its label. Required: an empty one panics.
    pub label: &'a str,
    /// Its icon, a [`crate::lucide`] icon; `None` for none.
    pub icon: Option<Icon>,
    /// Whether it is the current page (`aria-current="page"`).
    pub current: bool,
    /// A count at its end; `""` for none.
    pub secondary: &'a str,
    /// The page's utilities on it; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleSidebarItem<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsoleSidebarItem needs a label"
        );
        rsx! {
            <a class=(Class("cui-console-sidebar-item", self.class)) href=(self.href) aria-current="page"[self.current]>
                @if let Some(icon) = self.icon {
                    (icon.class("cui-console-sidebar-icon"))
                }
                <span class="cui-console-sidebar-label">(self.label)</span>
                @if !self.secondary.is_empty() {
                    <span class="cui-console-sidebar-secondary">(self.secondary)</span>
                }
            </a>
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
        let account = rsx! {
            <a class="cui-console-sidebar-account" href="#account">(lucide::CIRCLE_USER.class("cui-console-sidebar-avatar"))"Alex Morgan"</a>
        };
        let sidebar = rsx! {
            <ConsoleSidebar id="sidebar" label="CrashCart" home="#home" wordmark="CrashCart" close="Close" class="" footer=account>
                <ConsoleSidebarItem href="#projects" label="All Projects" icon=(Some(lucide::LAYOUT_GRID)) current=false secondary="" class="" />
                <ConsoleSidebarItem href="#issues" label="All Issues" icon=(Some(lucide::BUG)) current=true secondary="" class="" />
                <p class="cui-console-sidebar-heading">"Projects"</p>
                <ConsoleSidebarItem href="#shop" label="Shop" icon=None current=false secondary="8" class="" />
                <ConsoleSidebarItem href="#demo" label="Demo Shop" icon=None current=false secondary="12" class="" />
                <hr class="cui-console-sidebar-rule">
                <ConsoleSidebarItem href="#alerts" label="Alerts" icon=(Some(lucide::BELL)) current=false secondary="" class="" />
            </ConsoleSidebar>
        }
        .render();
        assert_eq!(
            *sidebar.as_inner(),
            example(include_str!(
                "../../../components/console/console-sidebar/examples/sidebar.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a close")]
    fn a_sidebar_with_no_close_label_panics() {
        let _ = rsx! { <ConsoleSidebar id="s" label="x" home="/" wordmark="x" close="" class="" footer="">""</ConsoleSidebar> }.render();
    }
}

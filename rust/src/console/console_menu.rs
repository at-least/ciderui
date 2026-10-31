use hypertext::{Buffer, prelude::*};

use crate::icon::Icon;
use crate::{Class, lucide};

/// `cui-console-menu` (`components/console/console-menu`): App Store Connect's
/// newer look's list of a section's pages down its left side; its
/// [`ConsoleMenuItem`]s, [`ConsoleMenuGroup`]s, and a heading
/// (`<p class="cui-console-menu-heading">`) or a rule
/// (`<hr class="cui-console-menu-rule">`) between runs of them.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::{ConsoleMenu, ConsoleMenuItem};
///
/// let html = rsx! {
///     <ConsoleMenu label="Analytics" class="">
///         <ConsoleMenuItem href="/overview" label="Overview" icon=None current=true disabled=false secondary="" class="" />
///     </ConsoleMenu>
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<a class="cui-console-menu-item" href="/overview" aria-current="page">"#));
/// ```
pub struct ConsoleMenu<'a, R: Renderable> {
    /// What the list is, for a screen reader (`aria-label`). Required: an
    /// empty one panics.
    pub label: &'a str,
    /// The page's utilities on the list; `""` for none.
    pub class: &'a str,
    /// Its items, groups, headings and rules.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleMenu<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsoleMenu needs a label"
        );
        rsx! {
            <nav class=(Class("cui-console-menu", self.class)) aria-label=(self.label)>(self.children)</nav>
        }
        .render_to(buffer);
    }
}

/// A page in a [`ConsoleMenu`] or a [`ConsoleMenuGroup`], a link.
pub struct ConsoleMenuItem<'a> {
    /// Where it goes. Required unless it is disabled, when it is left out.
    pub href: &'a str,
    /// Its label. Required: an empty one panics.
    pub label: &'a str,
    /// Its icon before the label, a [`crate::lucide`] icon; `None` for none.
    pub icon: Option<Icon>,
    /// Whether it is the current page (`aria-current="page"`).
    pub current: bool,
    /// Whether it goes nowhere now (`aria-disabled`, no `href`).
    pub disabled: bool,
    /// What follows it at its end (a count); `""` for nothing.
    pub secondary: &'a str,
    /// The page's utilities on the item; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleMenuItem<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsoleMenuItem needs a label"
        );
        assert!(
            self.disabled == self.href.is_empty(),
            "ciderui::console::ConsoleMenuItem needs an href, unless it is disabled"
        );
        rsx! {
            <a class=(Class("cui-console-menu-item", self.class)) href=(self.href)[!self.disabled] aria-current="page"[self.current] aria-disabled="true"[self.disabled]>
                @if let Some(icon) = self.icon {
                    (icon.class("cui-console-menu-icon"))
                }
                <span class="cui-console-menu-label">(self.label)</span>
                @if !self.secondary.is_empty() {
                    <span class="cui-console-menu-secondary">(self.secondary)</span>
                }
            </a>
        }
        .render_to(buffer);
    }
}

/// A group of a [`ConsoleMenu`] (`<details>`): its label opening its items,
/// set in from the left, with no script.
pub struct ConsoleMenuGroup<'a, R: Renderable> {
    /// Its label. Required: an empty one panics.
    pub label: &'a str,
    /// Its icon before the label; `None` for none.
    pub icon: Option<Icon>,
    /// Whether it is open.
    pub open: bool,
    /// The page's utilities on the group; `""` for none.
    pub class: &'a str,
    /// Its items, [`ConsoleMenuItem`]s.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleMenuGroup<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsoleMenuGroup needs a label"
        );
        rsx! {
            <details class=(Class("cui-console-menu-group", self.class)) open[self.open]>
                <summary class="cui-console-menu-item">
                    @if let Some(icon) = self.icon {
                        (icon.class("cui-console-menu-icon"))
                    }
                    <span class="cui-console-menu-label">(self.label)</span>
                    (lucide::CHEVRON_RIGHT.class("cui-console-menu-chevron"))
                </summary>
                (self.children)
            </details>
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
        let menu = rsx! {
            <ConsoleMenu label="Analytics" class="">
                <ConsoleMenuItem href="#overview" label="Overview" icon=(Some(lucide::LAYOUT_GRID)) current=true disabled=false secondary="" class="" />
                <ConsoleMenuGroup label="Acquisition" icon=(Some(lucide::USER_PLUS)) open=true class="">
                    <ConsoleMenuItem href="#sources" label="Sources" icon=None current=false disabled=false secondary="" class="" />
                    <ConsoleMenuItem href="" label="Product Pages" icon=None current=false disabled=true secondary="" class="" />
                </ConsoleMenuGroup>
                <ConsoleMenuGroup label="Monetization" icon=(Some(lucide::CHART_LINE)) open=false class="">
                    <ConsoleMenuItem href="#sales" label="Sales" icon=None current=false disabled=false secondary="" class="" />
                </ConsoleMenuGroup>
                <hr class="cui-console-menu-rule">
                <p class="cui-console-menu-heading">"Projects"</p>
                <ConsoleMenuItem href="#shop" label="Shop" icon=None current=false disabled=false secondary="8" class="" />
                <ConsoleMenuItem href="#demo" label="Demo Shop" icon=None current=false disabled=false secondary="12" class="" />
            </ConsoleMenu>
        }
        .render();
        assert_eq!(
            *menu.as_inner(),
            example(include_str!(
                "../../../components/console/console-menu/examples/menu.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs an href, unless it is disabled")]
    fn an_item_with_no_href_panics() {
        let _ = rsx! { <ConsoleMenuItem href="" label="x" icon=None current=false disabled=false secondary="" class="" /> }.render();
    }

    #[test]
    #[should_panic(expected = "needs an href, unless it is disabled")]
    fn a_disabled_item_with_an_href_panics() {
        let _ = rsx! { <ConsoleMenuItem href="#x" label="x" icon=None current=false disabled=true secondary="" class="" /> }.render();
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn a_menu_with_no_label_panics() {
        let _ = rsx! { <ConsoleMenu label="" class="">"x"</ConsoleMenu> }.render();
    }
}

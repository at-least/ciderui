use hypertext::{Buffer, prelude::*};

use crate::{Class, lucide};

/// `cui-console-chip` (`components/console-chip`): App Store Connect's
/// newer look's chip, a key and its value on the gray fill. A filter the
/// page applies has a control that removes it (`remove_href`, a link with
/// a Lucide x); a value that opens a list filtered by it is a link as a
/// whole (`href`).
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::ConsoleChip;
///
/// let html = rsx! { <ConsoleChip key="os" value="iOS" href=None remove_href=None remove_label="" class="" /> }.render();
/// assert_eq!(html.as_inner(), r#"<span class="cui-console-chip"><span class="cui-console-chip-key">os</span>iOS</span>"#);
/// ```
pub struct ConsoleChip<'a> {
    /// Its key, as the page words it. Required: an empty one panics.
    pub key: &'a str,
    /// Its value.
    pub value: &'a str,
    /// `Some`: the chip is a link to this URL. Not with `remove_href`: a
    /// control inside a link is no valid HTML, so both panic.
    pub href: Option<&'a str>,
    /// `Some`: the chip has a remove control, a link to this URL.
    pub remove_href: Option<&'a str>,
    /// The remove control's `aria-label`. Required with `remove_href`: the
    /// control is an icon alone.
    pub remove_label: &'a str,
    /// The page's utilities on the chip; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleChip<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.key.is_empty(), "ciderui::ConsoleChip needs a key");
        assert!(
            self.href.is_none() || self.remove_href.is_none(),
            "ciderui::ConsoleChip is a link or has a remove control, not both"
        );
        assert!(
            self.remove_href.is_none() || !self.remove_label.is_empty(),
            "ciderui::ConsoleChip needs a remove_label with a remove_href"
        );
        let parts = rsx! { <span class="cui-console-chip-key">(self.key)</span>(self.value) };
        rsx! {
            @if let Some(href) = self.href {
                <a class=(Class("cui-console-chip", self.class)) href=(href)>(parts)</a>
            } @else {
                <span class=(Class("cui-console-chip", self.class))>
                    (parts)
                    @if let Some(remove) = self.remove_href {
                        <a class="cui-console-chip-remove" href=(remove) aria-label=(self.remove_label)>(lucide::X)</a>
                    }
                </span>
            }
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
        let chips = rsx! {
            <ConsoleChip key="level" value="error" href=None remove_href=(Some("#all")) remove_label="Remove level: error" class="" />
            <ConsoleChip key="release" value="2.4.1" href=None remove_href=(Some("#all")) remove_label="Remove release: 2.4.1" class="" />
            <ConsoleChip key="os" value="iOS 18.1" href=(Some("#ios")) remove_href=None remove_label="" class="" />
            <ConsoleChip key="device" value="iPhone 16" href=(Some("#device")) remove_href=None remove_label="" class="" />
        }
        .render();
        assert_eq!(
            format!(
                r#"<p class="flex flex-wrap gap-2">{}</p>"#,
                chips.as_inner()
            ),
            example(include_str!(
                "../../components/console-chip/examples/chips.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "is a link or has a remove control, not both")]
    fn a_link_with_a_remove_control_panics() {
        let _ = rsx! { <ConsoleChip key="env" value="prod" href=(Some("/a")) remove_href=(Some("/b")) remove_label="Remove" class="" /> }.render();
    }

    #[test]
    #[should_panic(expected = "needs a remove_label with a remove_href")]
    fn a_remove_control_without_a_label_panics() {
        let _ = rsx! { <ConsoleChip key="env" value="prod" href=None remove_href=(Some("/b")) remove_label="" class="" /> }.render();
    }
}

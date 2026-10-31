use hypertext::{Buffer, prelude::*};

use crate::Class;

/// A choice of a [`ConsoleAppearance`].
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ConsoleAppearanceChoice {
    /// Light, `value="light"`.
    Light,
    /// Dark, `value="dark"`.
    Dark,
    /// The system's, `value="auto"`.
    Auto,
}

/// `cui-console-appearance` (`components/console/console-appearance`):
/// Light, Dark and Auto as one segmented control of radios
/// (`name="appearance"`), Apple's developer documentation's footer control,
/// for a site that offers its own appearance; the page's script applies the
/// choice.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleAppearance;
///
/// let html = rsx! { <ConsoleAppearance chosen=None class="" /> }.render();
/// assert!(html.as_inner().starts_with(r#"<fieldset class="cui-console-appearance"><legend>Appearance</legend><label class="cui-console-appearance-choice"><input type="radio" name="appearance" value="light"><span>Light</span></label>"#));
/// ```
pub struct ConsoleAppearance<'a> {
    /// The chosen one; `None` for none yet (the page's script chooses it
    /// from what it stored).
    pub chosen: Option<ConsoleAppearanceChoice>,
    /// The page's utilities on the `<fieldset>`; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleAppearance<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        let choices = [
            (ConsoleAppearanceChoice::Light, "light", "Light"),
            (ConsoleAppearanceChoice::Dark, "dark", "Dark"),
            (ConsoleAppearanceChoice::Auto, "auto", "Auto"),
        ];
        rsx! {
            <fieldset class=(Class("cui-console-appearance", self.class))>
                <legend>"Appearance"</legend>
                @for (choice, value, word) in choices {
                    <label class="cui-console-appearance-choice"><input type="radio" name="appearance" value=(value) checked[self.chosen == Some(choice)]><span>(word)</span></label>
                }
            </fieldset>
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
        let appearance = rsx! {
            <div class="cui-console-page">
                <ConsoleAppearance chosen=(Some(ConsoleAppearanceChoice::Light)) class="" />
            </div>
        }
        .render();
        assert_eq!(
            *appearance.as_inner(),
            example(include_str!(
                "../../../components/console/console-appearance/examples/appearance.html"
            ))
        );
    }

    #[test]
    fn none_chosen_checks_none() {
        let html = rsx! { <ConsoleAppearance chosen=None class="hx-only" /> }.render();
        assert!(
            html.as_inner()
                .starts_with(r#"<fieldset class="cui-console-appearance hx-only">"#)
        );
        assert!(!html.as_inner().contains("checked"), "{}", html.as_inner());
    }
}

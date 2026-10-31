use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-console-dialog` (`components/console/console-dialog`): App Store Connect's
/// newer look's modal, a `<dialog>` the page opens with `showModal()`: its
/// title (the dialog's name, `aria-labelledby`), and what the page gives
/// it — what it asks, its buttons in a `<form method="dialog"
/// class="cui-console-dialog-actions">` or the page's own form.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleDialog;
///
/// let html = rsx! { <ConsoleDialog id="ignore" title="Ignore?" open=false class="">""</ConsoleDialog> }.render();
/// assert!(html.as_inner().starts_with(r#"<dialog class="cui-console-dialog" id="ignore" aria-labelledby="ignore-title">"#));
/// ```
pub struct ConsoleDialog<'a, R: Renderable, T: Renderable = &'a str> {
    /// Its id, which the page's script opens; its title's is `{id}-title`.
    /// Required: an empty one panics.
    pub id: &'a str,
    /// Its title, text or markup.
    pub title: T,
    /// Rendered open (`open`): a page shown with it open, no script.
    pub open: bool,
    /// The page's utilities on it; `""` for none.
    pub class: &'a str,
    /// What it asks and its buttons.
    pub children: R,
}

impl<R: Renderable, T: Renderable> Renderable for ConsoleDialog<'_, R, T> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.id.is_empty(),
            "ciderui::console::ConsoleDialog needs an id"
        );
        let title_id = format!("{}-title", self.id);
        rsx! {
            <dialog class=(Class("cui-console-dialog", self.class)) id=(self.id) aria-labelledby=(&title_id) open[self.open]>
                <h2 class="cui-console-dialog-title" id=(&title_id)>(self.title)</h2>
                (self.children)
            </dialog>
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
        let dialog = rsx! {
            <button class="cui-console-btn" type="button" "commandfor"="ignore" "command"="show-modal">"Ignore…"</button>
            <ConsoleDialog id="ignore" title="Ignore 3 issues?" open=false class="">
                <p>"Ignored issues come back when the condition you choose is met."</p>
                <form method="dialog" class="cui-console-dialog-actions">
                    <button class="cui-console-btn">"Cancel"</button>
                    <button class="cui-console-btn" data-variant="filled">"Ignore"</button>
                </form>
            </ConsoleDialog>
        }
        .render();
        assert_eq!(
            *dialog.as_inner(),
            example(include_str!(
                "../../../components/console/console-dialog/examples/dialog.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs an id")]
    fn a_dialog_without_an_id_panics() {
        let _ =
            rsx! { <ConsoleDialog id="" title="x" open=false class="">""</ConsoleDialog> }.render();
    }
}

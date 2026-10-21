use hypertext::{Buffer, Raw, prelude::*};

use crate::Class;
use crate::icon::{CLOSE, CLOSE_SMALL};

/// `cui-modal` (`components/modal`): apple.com's store overlay on a
/// `<dialog>`, the page dimmed and held under it, its content in a white
/// panel with a close button at its corner (a form's, `method="dialog"`),
/// which and Escape close it with no script; or, `page`, the marketing
/// pages' modal, centered over the page dimmed and blurred, its close
/// button the neutral fill; `cui-dark` in its `class`, its dark one. A
/// [`ModalButton`] opens it.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{Modal, ModalButton};
///
/// let html = rsx! {
///     <ModalButton modal="models" class="cui-btn">"Show more"</ModalButton>
///     <Modal id="models" labelledby="models-title" page=false class="">
///         <h2 id="models-title">"Which model is right for you?"</h2>
///     </Modal>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<button class="cui-btn" commandfor="models" command="show-modal">Show more</button>"#));
/// assert!(html.as_inner().contains(r#"<dialog class="cui-modal" id="models" aria-labelledby="models-title"><div class="cui-modal-panel"><form method="dialog"><button class="cui-modal-close" aria-label="Close">"#));
/// ```
pub struct Modal<'a, R: Renderable> {
    /// The dialog's id, which its [`ModalButton`] names. Required: an empty
    /// one panics.
    pub id: &'a str,
    /// The id of its headline, which names it (`aria-labelledby`).
    /// Required: an empty one panics.
    pub labelledby: &'a str,
    /// Whether it is the marketing pages' modal (`data-variant="page"`), not
    /// the store's.
    pub page: bool,
    /// The page's utilities on the dialog (`cui-dark` with `page`, the dark
    /// one); `""` for none.
    pub class: &'a str,
    /// The panel's content, after its close button.
    pub children: R,
}

impl<R: Renderable> Renderable for Modal<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::Modal needs an id");
        assert!(
            !self.labelledby.is_empty(),
            "ciderui::Modal needs the id of its headline (labelledby)"
        );
        rsx! {
            <dialog class=(Class("cui-modal", self.class)) data-variant="page"[self.page] id=(self.id) aria-labelledby=(self.labelledby)>
                <div class="cui-modal-panel">
                    <form method="dialog"><button class="cui-modal-close" aria-label="Close">(if self.page { CLOSE_SMALL } else { CLOSE })</button></form>
                    (self.children)
                </div>
            </dialog>
        }
        .render_to(buffer);
    }
}

/// The button that opens a [`Modal`] with no script: `commandfor` naming
/// the dialog, `command="show-modal"` (attributes hypertext's `<button>` does
/// not know, so written here), its class the page's (a `cui-btn`).
pub struct ModalButton<'a, R: Renderable> {
    /// The id of the [`Modal`] it opens. Required: an empty one panics.
    pub modal: &'a str,
    /// Its class: the page's (`cui-btn`, a link's).
    pub class: &'a str,
    /// Its text.
    pub children: R,
}

impl<R: Renderable> Renderable for ModalButton<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.modal.is_empty(),
            "ciderui::ModalButton needs its modal's id"
        );
        Raw::<_>::dangerously_create(format!(
            r#"<button class="{}" commandfor="{}" command="show-modal">"#,
            attribute(self.class),
            attribute(self.modal)
        ))
        .render_to(buffer);
        self.children.render_to(buffer);
        Raw::<_>::dangerously_create("</button>").render_to(buffer);
    }
}

// a value inside double quotes, escaped as hypertext escapes an attribute's
fn attribute(value: &str) -> String {
    let mut out = String::with_capacity(value.len());
    for ch in value.chars() {
        match ch {
            '&' => out.push_str("&amp;"),
            '"' => out.push_str("&quot;"),
            '<' => out.push_str("&lt;"),
            _ => out.push(ch),
        }
    }
    out
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    #[test]
    fn renders_the_example() {
        let html = rsx! {
            <ModalButton modal="models" class="cui-btn">"Show more"</ModalButton>
            <Modal id="models" labelledby="models-title" page=false class="">
                <h2 class="cui-headline" id="models-title">"Which model is right for you?"</h2>
                <p class="cui-body">"Both have the same chip and the same cameras. The larger one has a larger display and the longest battery life."</p>
            </Modal>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!("../../components/modal/examples/modal.html"))
        );
    }

    #[test]
    fn renders_the_page_example() {
        let html = rsx! {
            <ModalButton modal="trade-in" class="cui-btn">"Read more about Apple Trade In"</ModalButton>
            <Modal id="trade-in" labelledby="trade-in-title" page=true class="">
                <p class="cui-eyebrow">"Apple Trade In"</p>
                <h2 class="cui-headline" id="trade-in-title">"Trade in your eligible device for credit."</h2>
                <p class="cui-body">"Just add a trade-in when you choose a new product. Once your device has been received and verified, we credit its value to your payment method."</p>
            </Modal>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!("../../components/modal/examples/page.html"))
        );
    }

    #[test]
    fn its_dark_one_takes_cui_dark() {
        let html =
            rsx! { <Modal id="m5" labelledby="m5-title" page=true class="cui-dark">"x"</Modal> }
                .render();
        assert!(html.as_inner().starts_with(
            r#"<dialog class="cui-modal cui-dark" data-variant="page" id="m5" aria-labelledby="m5-title">"#
        ));
    }

    #[test]
    fn escapes_the_buttons_attributes() {
        let html = rsx! { <ModalButton modal=r#"a"b"# class="x<y&z">"Open"</ModalButton> }.render();
        assert_eq!(
            html.as_inner(),
            r#"<button class="x&lt;y&amp;z" commandfor="a&quot;b" command="show-modal">Open</button>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs an id")]
    fn an_empty_id_panics() {
        let _ = rsx! { <Modal id="" labelledby="t" page=false class="">"x"</Modal> }.render();
    }

    #[test]
    #[should_panic(expected = "labelledby")]
    fn no_headline_panics() {
        let _ = rsx! { <Modal id="m" labelledby="" page=false class="">"x"</Modal> }.render();
    }

    #[test]
    #[should_panic(expected = "needs its modal's id")]
    fn a_button_with_no_modal_panics() {
        let _ = rsx! { <ModalButton modal="" class="cui-btn">"x"</ModalButton> }.render();
    }
}

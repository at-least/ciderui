use hypertext::{Buffer, prelude::*};

use crate::{Class, icon};

/// `cui-accordion` (`components/site/accordion`): apple.com's FAQ, questions under
/// rules, each opening its answer below it with no script (`<details>`),
/// its chevron turning up; with `large`, the bag's large question.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::site::{Accordion, AccordionItem};
///
/// let html = rsx! {
///     <Accordion large=false class="">
///         <AccordionItem question="How long does delivery take?" open=false class="">
///             <p>"Add your address at checkout to see a delivery date."</p>
///         </AccordionItem>
///     </Accordion>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-accordion"><details class="cui-accordion-item"><summary class="cui-accordion-question">How long does delivery take?<svg"#));
/// ```
pub struct Accordion<'a, R: Renderable> {
    /// Whether its questions are large (`data-size="lg"`, the bag's outer one).
    pub large: bool,
    /// The page's utilities on the accordion; `""` for none.
    pub class: &'a str,
    /// Its items, [`AccordionItem`]s.
    pub children: R,
}

impl<R: Renderable> Renderable for Accordion<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            <div class=(Class("cui-accordion", self.class)) data-size="lg"[self.large]>(self.children)</div>
        }
        .render_to(buffer);
    }
}

/// A question of an [`Accordion`] and its answer, a `<details>`.
pub struct AccordionItem<'a, R: Renderable> {
    /// The question. Required: an empty one panics.
    pub question: &'a str,
    /// Whether its answer is open.
    pub open: bool,
    /// The page's utilities on its root; `""` for none.
    pub class: &'a str,
    /// The answer.
    pub children: R,
}

impl<R: Renderable> Renderable for AccordionItem<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.question.is_empty(),
            "ciderui::site::AccordionItem needs a question"
        );
        rsx! {
            <details class=(crate::Class("cui-accordion-item", self.class)) open[self.open]>
                <summary class="cui-accordion-question">(self.question)(icon::DISCLOSURE)</summary>
                <div class="cui-accordion-answer">(self.children)</div>
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
    fn renders_the_examples() {
        let accordion = rsx! {
            <Accordion large=false class="">
                <AccordionItem question="How long does delivery take?" open=true class="">
                    <p>"Add your address at checkout to see a delivery date."</p>
                </AccordionItem>
                <AccordionItem question="Can I collect my order in person?" open=false class="">
                    <p>"Yes. Pick a store when you check out."</p>
                </AccordionItem>
            </Accordion>
        }
        .render();
        assert_eq!(
            *accordion.as_inner(),
            example(include_str!(
                "../../../components/site/accordion/examples/accordion.html"
            ))
        );
        let large = rsx! {
            <Accordion large=true class="">
                <AccordionItem question="Before You Buy" open=false class="">
                    <p>"What people ask most, answered."</p>
                </AccordionItem>
            </Accordion>
        }
        .render();
        assert_eq!(
            *large.as_inner(),
            example(include_str!(
                "../../../components/site/accordion/examples/large.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a question")]
    fn an_empty_question_panics() {
        let _ =
            rsx! { <AccordionItem question="" open=false class="">"x"</AccordionItem> }.render();
    }

    #[test]
    fn an_item_takes_the_pages_utilities() {
        let html =
            rsx! { <AccordionItem question="x" open=false class="hidden">"x"</AccordionItem> }
                .render();
        assert!(
            html.as_inner()
                .starts_with(r#"<details class="cui-accordion-item hidden">"#)
        );
    }
}

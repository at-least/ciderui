use hypertext::{Buffer, prelude::*};

use crate::{Class, lucide};

/// A [`ConsoleDisclosure`]'s look (`data-variant`).
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ConsoleDisclosureVariant {
    /// A line, the chevron before its label, the default: no `data-variant`.
    Line,
    /// A console button, the chevron after its label.
    Button,
}

/// `cui-console-disclosure` (`components/console/console-disclosure`): a
/// `<details>` hiding what is not needed yet (`children`) behind a line or
/// a console button, the HIG's disclosure control.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::{ConsoleDisclosure, ConsoleDisclosureVariant};
///
/// let html = rsx! {
///     <ConsoleDisclosure variant=(ConsoleDisclosureVariant::Button) open=false summary="Filters" summary_class="" class="">
///         <p>"Level"</p>
///     </ConsoleDisclosure>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<details class="cui-console-disclosure" data-variant="button"><summary class="cui-console-btn">Filters<svg class="cui-console-disclosure-chevron""#));
/// ```
pub struct ConsoleDisclosure<'a, R: Renderable> {
    /// Its look.
    pub variant: ConsoleDisclosureVariant,
    /// Whether it starts open.
    pub open: bool,
    /// Its label. Required: an empty one panics.
    pub summary: &'a str,
    /// The page's utilities on its `<summary>`; `""` for none.
    pub summary_class: &'a str,
    /// The page's utilities on its `<details>`; `""` for none.
    pub class: &'a str,
    /// What it hides.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleDisclosure<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.summary.is_empty(),
            "ciderui::console::ConsoleDisclosure needs a summary"
        );
        let chevron = "cui-console-disclosure-chevron";
        match self.variant {
            ConsoleDisclosureVariant::Line => rsx! {
                <details class=(Class("cui-console-disclosure", self.class)) open[self.open]>
                    <summary class=(Class("cui-console-disclosure-summary", self.summary_class))>(lucide::CHEVRON_RIGHT.class(chevron))(self.summary)</summary>
                    (self.children)
                </details>
            }
            .render_to(buffer),
            ConsoleDisclosureVariant::Button => rsx! {
                <details class=(Class("cui-console-disclosure", self.class)) data-variant="button" open[self.open]>
                    <summary class=(Class("cui-console-btn", self.summary_class))>(self.summary)(lucide::CHEVRON_DOWN.class(chevron))</summary>
                    (self.children)
                </details>
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
    fn renders_the_examples() {
        let lines = rsx! {
            <div class="cui-console-page grid gap-3">
                <ConsoleDisclosure variant=(ConsoleDisclosureVariant::Line) open=false summary="12 system frames" summary_class="" class="">
                    <p>"libsystem_kernel.dylib __pthread_kill + 8"</p>
                </ConsoleDisclosure>
                <ConsoleDisclosure variant=(ConsoleDisclosureVariant::Line) open=true summary="Source" summary_class="" class="">
                    <p>"let total = items.reduce(0, +)"</p>
                </ConsoleDisclosure>
            </div>
        }
        .render();
        assert_eq!(
            *lines.as_inner(),
            example(include_str!(
                "../../../components/console/console-disclosure/examples/disclosure.html"
            ))
        );
        let buttons = rsx! {
            <div class="cui-console-page grid gap-3">
                <ConsoleDisclosure variant=(ConsoleDisclosureVariant::Button) open=false summary="Filters" summary_class="" class="">
                    <p>"Level, release and the search box."</p>
                </ConsoleDisclosure>
                <ConsoleDisclosure variant=(ConsoleDisclosureVariant::Button) open=true summary="Filters" summary_class="" class="">
                    <p>"Level, release and the search box."</p>
                </ConsoleDisclosure>
            </div>
        }
        .render();
        assert_eq!(
            *buttons.as_inner(),
            example(include_str!(
                "../../../components/console/console-disclosure/examples/button.html"
            ))
        );
    }

    #[test]
    fn the_pages_utilities_go_on_each_part() {
        let html = rsx! {
            <ConsoleDisclosure variant=(ConsoleDisclosureVariant::Line) open=false summary="Raw" summary_class="text-xs" class="mt-6">""</ConsoleDisclosure>
        }
        .render();
        assert!(
            html.as_inner().starts_with(
                r#"<details class="cui-console-disclosure mt-6"><summary class="cui-console-disclosure-summary text-xs"><svg"#
            ),
            "{}",
            html.as_inner()
        );
    }

    #[test]
    #[should_panic(expected = "needs a summary")]
    fn an_empty_summary_panics() {
        let _ = rsx! { <ConsoleDisclosure variant=(ConsoleDisclosureVariant::Line) open=false summary="" summary_class="" class="">""</ConsoleDisclosure> }.render();
    }
}

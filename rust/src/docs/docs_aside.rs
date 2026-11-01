use hypertext::{Buffer, prelude::*};

use crate::Class;

/// An aside's kind: DocC's six, each its ground and its edge's color, and the
/// word its label says.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum DocsAsideVariant {
    /// A note, the gray one.
    Note,
    /// Something the reader must not miss.
    Important,
    /// A danger: data lost, a build broken.
    Warning,
    /// A way to do it better.
    Tip,
    /// What not to use any more.
    Deprecated,
    /// Something to try.
    Experiment,
}

impl DocsAsideVariant {
    /// Its `data-variant`, none for a note.
    fn attribute(self) -> Option<&'static str> {
        match self {
            Self::Note => None,
            Self::Important => Some("important"),
            Self::Warning => Some("warning"),
            Self::Tip => Some("tip"),
            Self::Deprecated => Some("deprecated"),
            Self::Experiment => Some("experiment"),
        }
    }

    /// The word its label says, and the aside's accessible name.
    fn label(self) -> &'static str {
        match self {
            Self::Note => "Note",
            Self::Important => "Important",
            Self::Warning => "Warning",
            Self::Tip => "Tip",
            Self::Deprecated => "Deprecated",
            Self::Experiment => "Experiment",
        }
    }
}

/// `cui-docs-aside` (`components/docs/docs-aside`): Apple's developer
/// documentation's aside, a note set apart from an article's text, its
/// kind's word over what the page gives it, in its kind's colors.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::docs::{DocsAside, DocsAsideVariant};
///
/// let html = rsx! {
///     <DocsAside variant=(DocsAsideVariant::Tip) class="">
///         <p>"Mark a deploy."</p>
///     </DocsAside>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<aside class="cui-docs-aside" data-variant="tip" aria-label="Tip">"#));
/// ```
pub struct DocsAside<'a, R: Renderable> {
    /// Its kind.
    pub variant: DocsAsideVariant,
    /// The page's utilities on the aside; `""` for none.
    pub class: &'a str,
    /// What it says, under its label: the page's paragraphs.
    pub children: R,
}

impl<R: Renderable> Renderable for DocsAside<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        let variant = self.variant.attribute();
        let label = self.variant.label();
        rsx! {
            <aside class=(Class("cui-docs-aside", self.class)) data-variant=(variant.unwrap_or(""))[variant.is_some()] aria-label=(label)>
                <p class="cui-docs-aside-label">(label)</p>
                (self.children)
            </aside>
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
        let note = rsx! {
            <DocsAside variant=(DocsAsideVariant::Note) class="">
                <p>"A report from a beta build carries its beta identifier; one from the App Store does not."</p>
            </DocsAside>
        }
        .render();
        assert_eq!(
            *note.as_inner(),
            example(include_str!(
                "../../../components/docs/docs-aside/examples/note.html"
            ))
        );
        let kinds: String = [
            (
                DocsAsideVariant::Important,
                "Upload the symbol files of every build you ship.",
            ),
            (
                DocsAsideVariant::Warning,
                "Deleting a project deletes its reports.",
            ),
            (
                DocsAsideVariant::Tip,
                "Mark a deploy to compare its crashes with the release before it.",
            ),
            (
                DocsAsideVariant::Deprecated,
                "Use the store’s key in its place.",
            ),
            (
                DocsAsideVariant::Experiment,
                "Try grouping by the crashed frame alone.",
            ),
        ]
        .iter()
        .map(|(variant, text)| {
            rsx! {
                <DocsAside variant=(*variant) class=""><p>(text)</p></DocsAside>
            }
            .render()
            .into_inner()
        })
        .collect();
        assert_eq!(
            format!(r#"<div class="grid gap-4">{kinds}</div>"#),
            example(include_str!(
                "../../../components/docs/docs-aside/examples/kinds.html"
            ))
        );
    }

    #[test]
    fn takes_the_pages_utilities() {
        let html = rsx! {
            <DocsAside variant=(DocsAsideVariant::Note) class="mt-8"><p>"A."</p></DocsAside>
        }
        .render();
        assert!(
            html.as_inner()
                .starts_with(r#"<aside class="cui-docs-aside mt-8" aria-label="Note">"#)
        );
    }
}

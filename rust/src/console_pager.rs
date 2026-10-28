use hypertext::{Buffer, prelude::*};

use crate::lucide;

/// `cui-console-pager` (`components/console-pager`): under a long list,
/// the buttons to the pages either side of a line saying what it shows,
/// as Apple Developer Forums' pager draws them: a chevron alone, named
/// "Previous page" or "Next page", and no button at all for a direction
/// with no page.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::ConsolePager;
///
/// let html = rsx! { <ConsolePager label="Issue pages" summary="1–50 of 230" previous=None next=(Some("?page=2")) /> }.render();
/// assert!(html.as_inner().contains(r#"aria-label="Next page""#));
/// assert!(!html.as_inner().contains("Previous page"));
/// ```
pub struct ConsolePager<'a> {
    /// The `<nav>`'s `aria-label`, what the list's pages are called ("Issue
    /// pages"), so a page with two lists names each. Required: an empty one
    /// panics.
    pub label: &'a str,
    /// What the list shows ("51–100 of 230"). Required: an empty one panics.
    pub summary: &'a str,
    /// The previous page's URL; `None` where there is none.
    pub previous: Option<&'a str>,
    /// The next page's URL; `None` where there is none.
    pub next: Option<&'a str>,
}

impl Renderable for ConsolePager<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::ConsolePager needs a label"
        );
        assert!(
            !self.summary.is_empty(),
            "ciderui::ConsolePager needs a summary"
        );
        rsx! {
            <nav class="cui-console-pager" aria-label=(self.label)>
                @if let Some(href) = self.previous {
                    <a class="cui-console-btn cui-console-btn-icon" href=(href) aria-label="Previous page">(lucide::CHEVRON_LEFT)</a>
                }
                <p class="cui-console-pager-summary">(self.summary)</p>
                @if let Some(href) = self.next {
                    <a class="cui-console-btn cui-console-btn-icon" href=(href) aria-label="Next page">(lucide::CHEVRON_RIGHT)</a>
                }
            </nav>
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
        let pager = rsx! {
            <div class="cui-console-page">
                <ConsolePager label="Issue pages" summary="51–100 of 230" previous=(Some("#page-1")) next=(Some("#page-3")) />
            </div>
        }
        .render();
        assert_eq!(
            *pager.as_inner(),
            example(include_str!(
                "../../components/console-pager/examples/pager.html"
            ))
        );
        let first = rsx! {
            <div class="cui-console-page">
                <ConsolePager label="Issue pages" summary="1–50 of 230" previous=None next=(Some("#page-2")) />
            </div>
        }
        .render();
        assert_eq!(
            *first.as_inner(),
            example(include_str!(
                "../../components/console-pager/examples/first.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a summary")]
    fn an_empty_summary_panics() {
        let _ = rsx! { <ConsolePager label="Pages" summary="" previous=None next=None /> }.render();
    }
}

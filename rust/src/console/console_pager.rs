use hypertext::{Buffer, prelude::*};

use crate::{Class, lucide};

/// `cui-console-pager` (`components/console/console-pager`): under a long list,
/// the buttons to the pages either side of a line saying what it shows,
/// as Apple Developer Forums' pager draws them: a chevron alone, named
/// "Previous page" or "Next page", a "First" button before them, and no
/// button at all for a direction with no page.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsolePager;
///
/// let html = rsx! { <ConsolePager label="Issue pages" summary="1–50 of 230" first=None previous=None next=(Some("?page=2")) class="" /> }.render();
/// assert!(html.as_inner().contains(r#"aria-label="Next page""#));
/// assert!(!html.as_inner().contains("Previous page"));
/// ```
pub struct ConsolePager<'a> {
    /// The `<nav>`'s `aria-label`, what the list's pages are called ("Issue
    /// pages"), so a page with two lists names each. Required: an empty one
    /// panics.
    pub label: &'a str,
    /// What the list shows ("51–100 of 230"); `""` for none (a list paged
    /// by a cursor knows nothing to say past its first page).
    pub summary: &'a str,
    /// The first page's URL, a "First" button; `None` where there is none
    /// or where the previous page is the first. A list paged by a cursor,
    /// which cannot step back one page, gives this and no `previous`.
    pub first: Option<&'a str>,
    /// The previous page's URL; `None` where there is none.
    pub previous: Option<&'a str>,
    /// The next page's URL; `None` where there is none.
    pub next: Option<&'a str>,
    /// The page's utilities on the `<nav>`; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsolePager<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsolePager needs a label"
        );
        assert!(
            !self.summary.is_empty()
                || self.first.is_some()
                || self.previous.is_some()
                || self.next.is_some(),
            "ciderui::console::ConsolePager needs a summary or a page to go to"
        );
        rsx! {
            <nav class=(Class("cui-console-pager", self.class)) aria-label=(self.label)>
                @if let Some(href) = self.first {
                    <a class="cui-console-btn" href=(href) aria-label="First page">"First"</a>
                }
                @if let Some(href) = self.previous {
                    <a class="cui-console-btn cui-console-btn-icon" href=(href) aria-label="Previous page">(lucide::CHEVRON_LEFT)</a>
                }
                @if !self.summary.is_empty() {
                    <p class="cui-console-pager-summary">(self.summary)</p>
                }
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
                <ConsolePager label="Issue pages" summary="51–100 of 230" first=None previous=(Some("#page-1")) next=(Some("#page-3")) class="" />
            </div>
        }
        .render();
        assert_eq!(
            *pager.as_inner(),
            example(include_str!(
                "../../../components/console/console-pager/examples/pager.html"
            ))
        );
        let first = rsx! {
            <div class="cui-console-page">
                <ConsolePager label="Issue pages" summary="1–50 of 230" first=None previous=None next=(Some("#page-2")) class="" />
            </div>
        }
        .render();
        assert_eq!(
            *first.as_inner(),
            example(include_str!(
                "../../../components/console/console-pager/examples/first.html"
            ))
        );
    }

    #[test]
    fn a_later_page_of_a_cursor_list() {
        let later = rsx! {
            <div class="cui-console-page">
                <ConsolePager label="Timeline pages" summary="" first=(Some("#newest")) previous=None next=(Some("#older")) class="" />
            </div>
        }
        .render();
        assert_eq!(
            *later.as_inner(),
            example(include_str!(
                "../../../components/console/console-pager/examples/later.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a summary or a page to go to")]
    fn nothing_to_show_panics() {
        let _ = rsx! { <ConsolePager label="Pages" summary="" first=None previous=None next=None class="" /> }.render();
    }
}

use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::docs::{DocsBar, DocsNavigator, DocsNavigatorItem};
use crate::icon::DOCS_CLOSE;

/// The navigator's id: the bar's button and the sheet's close button name
/// it (`popovertarget`). One docs page to a document.
pub const DOCS_NAVIGATOR_ID: &str = "docs-navigator";

/// `cui-docs-page` (`components/docs/docs-page`): a page of Apple's
/// developer documentation, under the site's global nav: the
/// documentation's bar ([`DocsBar`]) across it, the navigator
/// ([`DocsNavigator`]) beside the content from 1024px and a sheet the bar's
/// button opens below that, with its close button, and the page's `<main>`,
/// whose children are the page's: its columns (`cui-docs-page-column`),
/// each holding a hero, the content or the topics.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::docs::DocsPage;
///
/// let html = rsx! {
///     <DocsPage title="Documentation" href="/" navigator=(&[]) class="">
///         <div class="cui-docs-page-column">"…"</div>
///     </DocsPage>
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<div class="cui-docs-page-sidebar" id="docs-navigator" popover>"#));
/// ```
pub struct DocsPage<'a, R: Renderable> {
    /// The documentation's title, the bar's and the navigator's. Required:
    /// an empty one panics.
    pub title: &'a str,
    /// Where the title links: the documentation's first page.
    pub href: &'a str,
    /// The navigator's rows.
    pub navigator: &'a [DocsNavigatorItem<'a>],
    /// The page's utilities on the page; `""` for none.
    pub class: &'a str,
    /// The page's `<main>`: its columns.
    pub children: R,
}

impl<R: Renderable> Renderable for DocsPage<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            <div class=(Class("cui-docs-page", self.class))>
                <DocsBar title=(self.title) href=(self.href) navigator=(DOCS_NAVIGATOR_ID) class="" />
                <div class="cui-docs-page-sidebar" id=(DOCS_NAVIGATOR_ID) "popover">
                    <button class="cui-docs-page-close" type="button" "popovertarget"=(DOCS_NAVIGATOR_ID) "popovertargetaction"="hide" aria-label="Close Navigator">(DOCS_CLOSE)</button>
                    <DocsNavigator title=(self.title) href=(self.href) items=(self.navigator) class="" />
                </div>
                <main class="cui-docs-page-main">(self.children)</main>
            </div>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::docs::{
        DocsCrumb, DocsHero, DocsNavigatorPage, DocsTopic, DocsTopics, DocsTopicsGroup,
    };
    use crate::example::example;

    #[test]
    fn renders_the_example() {
        let reports = [
            DocsNavigatorItem::Page(DocsNavigatorPage {
                label: "Reading a crash report",
                href: "#reading",
                current: true,
                open: false,
                children: &[],
            }),
            DocsNavigatorItem::Page(DocsNavigatorPage {
                label: "Symbolicating a backtrace",
                href: "#symbolicating",
                current: false,
                open: false,
                children: &[],
            }),
        ];
        let navigator = [
            DocsNavigatorItem::Page(DocsNavigatorPage {
                label: "Getting started",
                href: "#getting-started",
                current: false,
                open: false,
                children: &[],
            }),
            DocsNavigatorItem::Page(DocsNavigatorPage {
                label: "Crash reports",
                href: "#crash-reports",
                current: false,
                open: true,
                children: &reports,
            }),
        ];
        let crumbs = [DocsCrumb {
            label: "Crash reports",
            href: "#crash-reports",
        }];
        let topics = [DocsTopic {
            label: "Symbolicating a backtrace",
            href: "#symbolicating",
            abstract_: "Replace the addresses in a backtrace with the names of the functions they fall in.",
            collection: false,
        }];
        let groups = [DocsTopicsGroup {
            heading: "",
            topics: &topics,
        }];
        let page = rsx! {
            <DocsPage title="Crashes" href="#crashes" navigator=(&navigator) class="">
                <div class="cui-docs-page-column">
                    <DocsHero crumbs=(&crumbs) eyebrow="Article" title="Reading a crash report" abstract_="Find the thread that crashed, the frame that failed and the images it had loaded." class="" />
                </div>
                <div class="cui-docs-page-column">
                    <div class="cui-docs-content">
                        <h2>"Overview"</h2>
                        <p>"A crash report records what an app was doing when it stopped."</p>
                    </div>
                </div>
                <div class="cui-docs-page-column">
                    <DocsTopics title="See Also" groups=(&groups) class="" />
                </div>
            </DocsPage>
        }
        .render();
        assert_eq!(
            *page.as_inner(),
            example(include_str!(
                "../../../components/docs/docs-page/examples/page.html"
            ))
        );
    }
}

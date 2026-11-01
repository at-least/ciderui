use hypertext::{Buffer, prelude::*};

use crate::Class;
use crate::icon::{DOCS_ARTICLE, DOCS_COLLECTION};

/// A page the topics link to.
#[derive(Clone, Copy, Debug)]
pub struct DocsTopic<'a> {
    /// Its name, the link's words. Required: an empty one panics.
    pub label: &'a str,
    /// Its address.
    pub href: &'a str,
    /// What it is about, under its link; `""` for none.
    pub abstract_: &'a str,
    /// A page of pages: a collection's glyph before it, an article's where
    /// not.
    pub collection: bool,
}

/// A group of the topics: its title over its pages.
#[derive(Clone, Copy, Debug)]
pub struct DocsTopicsGroup<'a> {
    /// Its title; `""` for none.
    pub heading: &'a str,
    /// Its pages.
    pub topics: &'a [DocsTopic<'a>],
}

/// `cui-docs-topics` (`components/docs/docs-topics`): Apple's developer
/// documentation's Topics and See Also, after the content: its title over
/// groups of links, each after its glyph, each with its abstract under it.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::docs::{DocsTopic, DocsTopics, DocsTopicsGroup};
///
/// let topics = [DocsTopic { label: "Docs Bar", href: "/components/docs-bar/", abstract_: "", collection: false }];
/// let groups = [DocsTopicsGroup { heading: "", topics: &topics }];
/// let html = rsx! { <DocsTopics title="See Also" groups=(&groups) class="" /> }.render();
/// assert!(html.as_inner().contains(r#"<h2 class="cui-docs-topics-title">See Also</h2>"#));
/// ```
pub struct DocsTopics<'a> {
    /// Its title (Topics, See Also). Required: an empty one panics.
    pub title: &'a str,
    /// Its groups.
    pub groups: &'a [DocsTopicsGroup<'a>],
    /// The page's utilities on the section; `""` for none.
    pub class: &'a str,
}

impl Renderable for DocsTopics<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.title.is_empty(),
            "ciderui::docs::DocsTopics needs a title"
        );
        for topic in self.groups.iter().flat_map(|g| g.topics) {
            assert!(
                !topic.label.is_empty(),
                "ciderui::docs::DocsTopics: a topic needs a label"
            );
        }
        rsx! {
            <section class=(Class("cui-docs-topics", self.class))>
                <h2 class="cui-docs-topics-title">(self.title)</h2>
                @for group in self.groups {
                    <div class="cui-docs-topics-group">
                        @if !group.heading.is_empty() {
                            <h3 class="cui-docs-topics-heading">(group.heading)</h3>
                        }
                        @for topic in group.topics {
                            <div class="cui-docs-topics-item">
                                <a class="cui-docs-topics-link" href=(topic.href)>
                                    ((if topic.collection { DOCS_COLLECTION } else { DOCS_ARTICLE }).class("cui-docs-topics-icon"))
                                    <span>(topic.label)</span>
                                </a>
                                @if !topic.abstract_.is_empty() {
                                    <p class="cui-docs-topics-abstract">(topic.abstract_)</p>
                                }
                            </div>
                        }
                    </div>
                }
            </section>
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
        let topics = [
            DocsTopic {
                label: "Symbolicating a backtrace",
                href: "#symbolicating",
                abstract_: "Replace the addresses in a backtrace with the names of the functions they fall in.",
                collection: false,
            },
            DocsTopic {
                label: "Identifying the cause of common crashes",
                href: "#common-crashes",
                abstract_: "Match a crash report to a pattern and find what to fix.",
                collection: false,
            },
            DocsTopic {
                label: "Alerts",
                href: "#alerts",
                abstract_: "Hear about a new crash where your team already is.",
                collection: true,
            },
        ];
        let groups = [DocsTopicsGroup {
            heading: "Crash reports",
            topics: &topics,
        }];
        let html = rsx! {
            <DocsTopics title="See Also" groups=(&groups) class="" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../../components/docs/docs-topics/examples/topics.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a title")]
    fn an_empty_title_panics() {
        let _ = rsx! { <DocsTopics title="" groups=(&[]) class="" /> }.render();
    }

    #[test]
    #[should_panic(expected = "a topic needs a label")]
    fn an_empty_label_panics() {
        let topics = [DocsTopic {
            label: "",
            href: "/",
            abstract_: "",
            collection: false,
        }];
        let groups = [DocsTopicsGroup {
            heading: "",
            topics: &topics,
        }];
        let _ = rsx! { <DocsTopics title="See Also" groups=(&groups) class="" /> }.render();
    }
}

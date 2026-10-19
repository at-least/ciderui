use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-article` (`components/article`): apple.com's Newsroom story, its
/// header (an eyebrow of its category and its date, the headline, a
/// subhead) over its copy (`body`, the plain elements markdown writes, its
/// footnotes a `cui-article-footnotes` section at its end).
///
/// ```
/// use hypertext::{Raw, prelude::*};
/// use ciderui::Article;
///
/// let body = Raw::<_>::dangerously_create("<p>The copy.</p>");
/// let html = rsx! {
///     <Article category="Update" datetime="2026-10-19" date="October 19, 2026" headline="ciderui comes to Zola" subhead="" body=(body) class="" />
/// }
/// .render();
/// assert!(html.as_inner().ends_with(r#"<div class="cui-article-body"><p>The copy.</p></div></article>"#));
/// ```
pub struct Article<'a, B: Renderable> {
    /// Its category, over its date (`cui-article-category`); `""` for none.
    pub category: &'a str,
    /// Its date as a machine reads it (`2026-10-19`), the `<time>`'s
    /// `datetime`; `""` with no date.
    pub datetime: &'a str,
    /// Its date as the page shows it (`October 19, 2026`); `""` for none. A
    /// date with no `datetime` panics.
    pub date: &'a str,
    /// The headline, the page's `<h1>`. Required: an empty one panics.
    pub headline: &'a str,
    /// The subhead under it; `""` for none.
    pub subhead: &'a str,
    /// The copy: paragraphs, subheads, lists, figures, its footnotes.
    pub body: B,
    /// The page's utilities on the article; `""` for none.
    pub class: &'a str,
}

impl<B: Renderable> Renderable for Article<'_, B> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.headline.is_empty(),
            "ciderui::Article needs a headline"
        );
        assert!(
            self.date.is_empty() || !self.datetime.is_empty(),
            "ciderui::Article's date needs its datetime"
        );
        rsx! {
            <article class=(Class("cui-article", self.class))>
                <header class="cui-article-header">
                    @if !self.category.is_empty() || !self.date.is_empty() {
                        <p class="cui-article-eyebrow">
                            @if !self.category.is_empty() {
                                <span class="cui-article-category">(self.category)</span>
                            }
                            @if !self.date.is_empty() {
                                <time class="cui-article-date" datetime=(self.datetime)>(self.date)</time>
                            }
                        </p>
                    }
                    <h1 class="cui-article-headline">(self.headline)</h1>
                    @if !self.subhead.is_empty() {
                        <p class="cui-article-subhead">(self.subhead)</p>
                    }
                </header>
                <div class="cui-article-body">(self.body)</div>
            </article>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;
    use hypertext::Raw;

    const BODY: &str = r#"<div class="cui-article-body">"#;

    #[test]
    fn renders_the_example() {
        let want = example(include_str!(
            "../../components/article/examples/article.html"
        ));
        // the example's copy, as the page's markdown would give it
        let start = want.find(BODY).unwrap() + BODY.len();
        let end = want.rfind("</div></article>").unwrap();
        let body = Raw::<_>::dangerously_create(&want[start..end]);
        let article = rsx! {
            <Article category="Update" datetime="2026-10-19" date="October 19, 2026" headline="ciderui comes to Zola" subhead="The kit's look, now for a static site's pages and its news." body=(body) class="" />
        }
        .render();
        assert_eq!(*article.as_inner(), want);
    }

    #[test]
    fn no_eyebrow_and_no_subhead_write_nothing() {
        let body = Raw::<_>::dangerously_create("<p>x</p>");
        let html = rsx! {
            <Article category="" datetime="" date="" headline="Privacy" subhead="" body=(body) class="pb-24" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<article class="cui-article pb-24"><header class="cui-article-header"><h1 class="cui-article-headline">Privacy</h1></header><div class="cui-article-body"><p>x</p></div></article>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a headline")]
    fn an_empty_headline_panics() {
        let _ = rsx! {
            <Article category="" datetime="" date="" headline="" subhead="" body=() class="" />
        }
        .render();
    }

    #[test]
    #[should_panic(expected = "date needs its datetime")]
    fn a_date_with_no_datetime_panics() {
        let _ = rsx! {
            <Article category="" datetime="" date="Today" headline="x" subhead="" body=() class="" />
        }
        .render();
    }
}

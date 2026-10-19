use hypertext::{Buffer, prelude::*};

/// A [`Posts`] list's size (`data-size`).
#[derive(Clone, Copy, Debug, Default, PartialEq, Eq)]
pub enum PostSize {
    /// The archive's: 980px, a month's heading, rows ruled above.
    #[default]
    Default,
    /// The related stories under an article (`data-size="sm"`): the
    /// article's text column, rows ruled below.
    Small,
}

/// `cui-posts` (`components/posts`): apple.com's Newsroom archive, a
/// heading over ruled rows, each a [`Post`]; [`PostSize::Small`], the
/// related stories under an article.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{PostSize, Posts};
///
/// let html = rsx! {
///     <Posts size=(PostSize::Default) heading="October 2026" class="">
///         <li>"…"</li>
///     </Posts>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<section class="cui-posts"><h2 class="cui-posts-heading">October 2026</h2><ul class="cui-posts-list">"#));
/// ```
pub struct Posts<'a, R: Renderable> {
    /// The archive's, or an article's related stories: [`PostSize`].
    pub size: PostSize,
    /// Its heading, an `<h2>`. Required: an empty one panics.
    pub heading: &'a str,
    /// The page's utilities on it; `""` for none.
    pub class: &'a str,
    /// Its rows, each an `<li>` with a [`Post`] in it.
    pub children: R,
}

impl<R: Renderable> Renderable for Posts<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.heading.is_empty(), "ciderui::Posts needs a heading");
        let small = self.size == PostSize::Small;
        rsx! {
            <section class=(crate::Class("cui-posts", self.class)) data-size="sm"[small]>
                <h2 class="cui-posts-heading">(self.heading)</h2>
                <ul class="cui-posts-list">(self.children)</ul>
            </section>
        }
        .render_to(buffer);
    }
}

/// `cui-post` (`components/posts`): a row of a [`Posts`] list, a link: its
/// picture before its category, its headline and its date.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Post;
///
/// let none = ();
/// let html = rsx! {
///     <Post href="/news/zola/" media=none category="Update" datetime="2026-10-19" date="October 19, 2026" headline="ciderui comes to Zola" />
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<a class="cui-post" href="/news/zola/"><div class="cui-post-content"><p class="cui-post-category">Update</p>"#));
/// ```
pub struct Post<'a, M: Renderable> {
    /// Where it goes, the row's link. Required: an empty one panics.
    pub href: &'a str,
    /// Its picture (`cui-post-media` on an `<img>`), or `()` for none.
    pub media: M,
    /// Its category, over its headline; `""` for none.
    pub category: &'a str,
    /// Its date as a machine reads it, the `<time>`'s `datetime`; `""` with
    /// no date.
    pub datetime: &'a str,
    /// Its date as the page shows it, under its headline; `""` for none. A
    /// date with no `datetime` panics.
    pub date: &'a str,
    /// Its headline, an `<h3>`. Required: an empty one panics.
    pub headline: &'a str,
}

impl<M: Renderable> Renderable for Post<'_, M> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.href.is_empty(), "ciderui::Post needs an href");
        assert!(!self.headline.is_empty(), "ciderui::Post needs a headline");
        assert!(
            self.date.is_empty() || !self.datetime.is_empty(),
            "ciderui::Post's date needs its datetime"
        );
        rsx! {
            <a class="cui-post" href=(self.href)>
                (self.media)
                <div class="cui-post-content">
                    @if !self.category.is_empty() {
                        <p class="cui-post-category">(self.category)</p>
                    }
                    <h3 class="cui-post-headline">(self.headline)</h3>
                    @if !self.date.is_empty() {
                        <time class="cui-post-date" datetime=(self.datetime)>(self.date)</time>
                    }
                </div>
            </a>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    const GIF: &str = "data:image/gif;base64,R0lGODlhAQABAAAAACw=";

    #[test]
    fn renders_the_examples() {
        let archive = [
            (
                "/news/zola/",
                "Update",
                "2026-10-19",
                "October 19, 2026",
                "ciderui comes to Zola",
            ),
            (
                "/news/article/",
                "Update",
                "2026-10-18",
                "October 18, 2026",
                "An article in the Newsroom's look",
            ),
            (
                "/news/stories/",
                "Feature",
                "2026-10-17",
                "October 17, 2026",
                "A grid of stories",
            ),
        ];
        let related = [
            (
                "/news/archive/",
                "Update",
                "2026-10-16",
                "October 16, 2026",
                "The archive's rows",
            ),
            (
                "/news/fonts/",
                "Update",
                "2026-10-15",
                "October 15, 2026",
                "Inter where SF is not",
            ),
        ];
        for (size, heading, rows, file) in [
            (
                PostSize::Default,
                "October 2026",
                &archive[..],
                include_str!("../../components/posts/examples/archive.html"),
            ),
            (
                PostSize::Small,
                "More from the Newsroom",
                &related[..],
                include_str!("../../components/posts/examples/related.html"),
            ),
        ] {
            let html = rsx! {
                <Posts size=(size) heading=(heading) class="">
                    @for (href, category, datetime, date, headline) in rows {
                        <li><Post href=(href) media=(rsx! { <img class="cui-post-media" src=(GIF) alt="" width="1280" height="720"> }) category=(category) datetime=(datetime) date=(date) headline=(headline) /></li>
                    }
                </Posts>
            }
            .render();
            assert_eq!(*html.as_inner(), example(file));
        }
    }

    #[test]
    #[should_panic(expected = "date needs its datetime")]
    fn a_date_with_no_datetime_panics() {
        let none = ();
        let _ = rsx! {
            <Post href="/x/" media=none category="" datetime="" date="Today" headline="x" />
        }
        .render();
    }

    #[test]
    #[should_panic(expected = "Posts needs a heading")]
    fn an_empty_heading_panics() {
        let _ = rsx! { <Posts size=(PostSize::Default) heading="" class="">""</Posts> }.render();
    }
}

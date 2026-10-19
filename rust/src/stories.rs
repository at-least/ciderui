use hypertext::{Buffer, prelude::*};

use crate::Class;

/// A [`Story`]'s size (`data-size`): how many of it go across the grid.
#[derive(Clone, Copy, Debug, Default, PartialEq, Eq)]
pub enum StorySize {
    /// Three across, four of the grid's twelve columns.
    #[default]
    Default,
    /// Two across (`data-size="md"`), six columns.
    Medium,
    /// One across (`data-size="lg"`), its picture beside its text from 1069px.
    Large,
}

/// `cui-stories` (`components/stories`): apple.com's Newsroom index, a band
/// on the alternate ground, its heading over a grid of [`Story`] tiles.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Stories;
///
/// let html = rsx! {
///     <Stories heading="Latest News" class="">
///         <li>"…"</li>
///     </Stories>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<section class="cui-stories"><div class="cui-stories-content"><h2 class="cui-stories-heading">Latest News</h2><ul class="cui-stories-list">"#));
/// ```
pub struct Stories<'a, R: Renderable> {
    /// Its heading, an `<h2>`. Required: an empty one panics.
    pub heading: &'a str,
    /// The page's utilities on the band; `""` for none.
    pub class: &'a str,
    /// Its stories, each an `<li>` with a [`Story`] in it.
    pub children: R,
}

impl<R: Renderable> Renderable for Stories<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.heading.is_empty(), "ciderui::Stories needs a heading");
        rsx! {
            <section class=(Class("cui-stories", self.class))>
                <div class="cui-stories-content">
                    <h2 class="cui-stories-heading">(self.heading)</h2>
                    <ul class="cui-stories-list">(self.children)</ul>
                </div>
            </section>
        }
        .render_to(buffer);
    }
}

/// `cui-story` (`components/stories`): a story in a [`Stories`] grid, a
/// white rounded tile, the whole of it its link: its picture, its category,
/// its headline, its date.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{Story, StorySize};
///
/// let media = rsx! { <img class="cui-story-media" src="/zola.jpg" alt=""> };
/// let html = rsx! {
///     <Story size=(StorySize::Large) href="/news/zola/" media=(media) category="Update" datetime="2026-10-19" date="October 19, 2026" headline="ciderui comes to Zola" />
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<a class="cui-story" data-size="lg" href="/news/zola/"><img class="cui-story-media""#));
/// ```
pub struct Story<'a, M: Renderable> {
    /// How many go across: [`StorySize`].
    pub size: StorySize,
    /// Where it goes, the whole tile's link. Required: an empty one panics.
    pub href: &'a str,
    /// Its picture (`cui-story-media` on an `<img>`), or `()` for none.
    pub media: M,
    /// Its category, over its headline; `""` for none.
    pub category: &'a str,
    /// Its date as a machine reads it, the `<time>`'s `datetime`; `""` with
    /// no date.
    pub datetime: &'a str,
    /// Its date as the page shows it, at the tile's foot; `""` for none. A
    /// date with no `datetime` panics.
    pub date: &'a str,
    /// Its headline, an `<h3>`. Required: an empty one panics.
    pub headline: &'a str,
}

impl<M: Renderable> Renderable for Story<'_, M> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.href.is_empty(), "ciderui::Story needs an href");
        assert!(!self.headline.is_empty(), "ciderui::Story needs a headline");
        assert!(
            self.date.is_empty() || !self.datetime.is_empty(),
            "ciderui::Story's date needs its datetime"
        );
        let size = match self.size {
            StorySize::Default => None,
            StorySize::Medium => Some("md"),
            StorySize::Large => Some("lg"),
        };
        rsx! {
            <a class="cui-story" data-size=(size.unwrap_or_default())[size.is_some()] href=(self.href)>
                (self.media)
                <div class="cui-story-content">
                    <div class="cui-story-head">
                        @if !self.category.is_empty() {
                            <p class="cui-story-category">(self.category)</p>
                        }
                        <h3 class="cui-story-headline">(self.headline)</h3>
                    </div>
                    @if !self.date.is_empty() {
                        <time class="cui-story-date" datetime=(self.datetime)>(self.date)</time>
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
    fn renders_the_example() {
        let stories = [
            (
                StorySize::Large,
                "/news/zola/",
                "Update",
                "2026-10-19",
                "October 19, 2026",
                "ciderui comes to Zola",
            ),
            (
                StorySize::Medium,
                "/news/article/",
                "Update",
                "2026-10-18",
                "October 18, 2026",
                "An article in the Newsroom's look",
            ),
            (
                StorySize::Medium,
                "/news/stories/",
                "Feature",
                "2026-10-17",
                "October 17, 2026",
                "A grid of stories",
            ),
            (
                StorySize::Default,
                "/news/archive/",
                "Update",
                "2026-10-16",
                "October 16, 2026",
                "The archive's rows",
            ),
            (
                StorySize::Default,
                "/news/fonts/",
                "Update",
                "2026-10-15",
                "October 15, 2026",
                "Inter where SF is not",
            ),
            (
                StorySize::Default,
                "/news/tokens/",
                "Update",
                "2026-10-14",
                "October 14, 2026",
                "A fixed palette",
            ),
        ];
        let html = rsx! {
            <Stories heading="Latest News" class="">
                @for (size, href, category, datetime, date, headline) in stories {
                    <li><Story size=(size) href=(href) media=(rsx! { <img class="cui-story-media" src=(GIF) alt="" width="1280" height="720"> }) category=(category) datetime=(datetime) date=(date) headline=(headline) /></li>
                }
            </Stories>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/stories/examples/stories.html"
            ))
        );
    }

    #[test]
    fn no_category_no_date_no_picture() {
        let none = ();
        let html = rsx! {
            <Story size=(StorySize::Default) href="/x/" media=none category="" datetime="" date="" headline="x" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<a class="cui-story" href="/x/"><div class="cui-story-content"><div class="cui-story-head"><h3 class="cui-story-headline">x</h3></div></div></a>"#
        );
    }

    #[test]
    #[should_panic(expected = "date needs its datetime")]
    fn a_date_with_no_datetime_panics() {
        let none = ();
        let _ = rsx! {
            <Story size=(StorySize::Default) href="/x/" media=none category="" datetime="" date="Today" headline="x" />
        }
        .render();
    }

    #[test]
    #[should_panic(expected = "Stories needs a heading")]
    fn an_empty_heading_panics() {
        let _ = rsx! { <Stories heading="" class="">""</Stories> }.render();
    }
}

use hypertext::{Buffer, prelude::*};

use crate::{Class, icon};

/// `cui-shelf` (`components/shelf`): the store's row of cards, a header (a
/// headline and its secondary text after it) over a scroller across the
/// window, its items snapping into place; with `paddles`, a large paddle at
/// each side, the previous disabled (the row at its start): the page's
/// script scrolls the row and disables the one at an end.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{Card, Shelf};
///
/// let none = ();
/// let html = rsx! {
///     <Shelf id="latest" headline="New arrivals." subhead="See what just came in." paddles=false class="">
///         <li><Card dark=false href="/phone/" media=none eyebrow="" tag=false headline="Cider Phone" description="" price="From $799" class="" /></li>
///     </Shelf>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-shelf"><div class="cui-shelf-header"><h2 class="cui-shelf-headline">New arrivals.</h2> <span class="cui-shelf-subhead">"#));
/// ```
pub struct Shelf<'a, R: Renderable> {
    /// The scroller's id, the paddles' `aria-controls`. Required: an empty
    /// one panics.
    pub id: &'a str,
    /// The headline, an `<h2>`. Required: an empty one panics.
    pub headline: &'a str,
    /// The secondary text after it; `""` for none.
    pub subhead: &'a str,
    /// Whether it shows its paddles (a page with a script to scroll the row).
    pub paddles: bool,
    /// The page's utilities on the shelf; `""` for none.
    pub class: &'a str,
    /// Its items, each an `<li>` (a [`Card`](crate::Card) in it).
    pub children: R,
}

impl<R: Renderable> Renderable for Shelf<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::Shelf needs an id");
        assert!(!self.headline.is_empty(), "ciderui::Shelf needs a headline");
        rsx! {
            <div class=(Class("cui-shelf", self.class))>
                <div class="cui-shelf-header">
                    <h2 class="cui-shelf-headline">(self.headline)</h2>
                    @if !self.subhead.is_empty() {
                        " "<span class="cui-shelf-subhead">(self.subhead)</span>
                    }
                </div>
                <div class="cui-shelf-row">
                    <ul class="cui-shelf-scroller" id=(self.id)>(self.children)</ul>
                    @if self.paddles {
                        <button class="cui-gallery-paddle cui-shelf-paddle" data-size="lg" type="button" aria-controls=(self.id) aria-label="Previous" disabled>(icon::PADDLE_PREVIOUS)</button>
                        <button class="cui-gallery-paddle cui-shelf-paddle" data-size="lg" type="button" aria-controls=(self.id) aria-label="Next">(icon::PADDLE_NEXT)</button>
                    }
                </div>
            </div>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::{Card, example::example};

    #[test]
    fn renders_the_examples() {
        for (paddles, file) in [
            (
                false,
                include_str!("../../components/shelf/examples/shelf.html"),
            ),
            (
                true,
                include_str!("../../components/shelf/examples/paddles.html"),
            ),
        ] {
            let none = ();
            let html = rsx! {
                <Shelf id="shelf" headline="New arrivals." subhead="See what just came in." paddles=(paddles) class="">
                    <li><Card dark=false href="/phone/" media=none eyebrow="NEW" tag=true headline="Cider Phone 3" description="" price="From $1199" class="" /></li>
                    <li><Card dark=false href="/phone-fold/" media=none eyebrow="NEW" tag=true headline="Cider Phone Fold" description="" price="From $899" class="" /></li>
                    <li><Card dark=false href="/watch/" media=none eyebrow="" tag=false headline="Cider Watch" description="" price="From $399" class="" /></li>
                    <li><Card dark=false href="/buds/" media=none eyebrow="" tag=false headline="Cider Buds" description="" price="From $129" class="" /></li>
                </Shelf>
            }
            .render();
            assert_eq!(*html.as_inner(), example(file));
        }
    }

    #[test]
    fn no_subhead_no_space() {
        let html = rsx! {
            <Shelf id="accessories" headline="Accessories." subhead="" paddles=false class="mt-8">"x"</Shelf>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<div class="cui-shelf mt-8"><div class="cui-shelf-header"><h2 class="cui-shelf-headline">Accessories.</h2></div><div class="cui-shelf-row"><ul class="cui-shelf-scroller" id="accessories">x</ul></div></div>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs an id")]
    fn an_empty_id_panics() {
        let _ = rsx! { <Shelf id="" headline="x" subhead="" paddles=false class="">"x"</Shelf> }
            .render();
    }

    #[test]
    #[should_panic(expected = "needs a headline")]
    fn an_empty_headline_panics() {
        let _ = rsx! { <Shelf id="x" headline="" subhead="" paddles=false class="">"x"</Shelf> }
            .render();
    }
}

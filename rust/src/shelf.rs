use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-shelf` (`components/shelf`): the store's row of cards, a header (a
/// headline and its secondary text after it) over a scroller across the
/// window, its items snapping into place.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{Card, Shelf};
///
/// let none = ();
/// let html = rsx! {
///     <Shelf headline="The latest." subhead="Take a look at what's new." class="">
///         <li><Card dark=false href="/iphone/" media=none eyebrow="" tag=false headline="iPhone" description="" price="From $799" class="" /></li>
///     </Shelf>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-shelf"><div class="cui-shelf-header"><h2 class="cui-shelf-headline">The latest.</h2> <span class="cui-shelf-subhead">"#));
/// ```
pub struct Shelf<'a, R: Renderable> {
    /// The headline, an `<h2>`. Required: an empty one panics.
    pub headline: &'a str,
    /// The secondary text after it; `""` for none.
    pub subhead: &'a str,
    /// The page's utilities on the shelf; `""` for none.
    pub class: &'a str,
    /// Its items, each an `<li>` (a [`Card`](crate::Card) in it).
    pub children: R,
}

impl<R: Renderable> Renderable for Shelf<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.headline.is_empty(), "ciderui::Shelf needs a headline");
        rsx! {
            <div class=(Class("cui-shelf", self.class))>
                <div class="cui-shelf-header">
                    <h2 class="cui-shelf-headline">(self.headline)</h2>
                    @if !self.subhead.is_empty() {
                        " "<span class="cui-shelf-subhead">(self.subhead)</span>
                    }
                </div>
                <ul class="cui-shelf-scroller">(self.children)</ul>
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
    fn renders_the_example() {
        let none = ();
        let html = rsx! {
            <Shelf headline="The latest." subhead="Take a look at what's new." class="">
                <li><Card dark=false href="/iphone-18-pro/" media=none eyebrow="NEW" tag=true headline="iPhone 18 Pro" description="" price="From $1199" class="" /></li>
                <li><Card dark=false href="/iphone-duo/" media=none eyebrow="NEW" tag=true headline="iPhone Duo" description="" price="From $899" class="" /></li>
                <li><Card dark=false href="/watch/" media=none eyebrow="" tag=false headline="Apple Watch" description="" price="From $399" class="" /></li>
                <li><Card dark=false href="/airpods/" media=none eyebrow="" tag=false headline="AirPods" description="" price="From $129" class="" /></li>
            </Shelf>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!("../../components/shelf/examples/shelf.html"))
        );
    }

    #[test]
    fn no_subhead_no_space() {
        let html = rsx! {
            <Shelf headline="Accessories." subhead="" class="mt-8">"x"</Shelf>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<div class="cui-shelf mt-8"><div class="cui-shelf-header"><h2 class="cui-shelf-headline">Accessories.</h2></div><ul class="cui-shelf-scroller">x</ul></div>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a headline")]
    fn an_empty_headline_panics() {
        let _ = rsx! { <Shelf headline="" subhead="" class="">"x"</Shelf> }.render();
    }
}

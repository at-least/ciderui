use hypertext::{Buffer, prelude::*};

use crate::{Class, icon};

/// `cui-gallery` (`components/gallery`): a row of items across the window,
/// each snapping into place where the content starts; with `paddles`, its
/// paddles' row, the previous disabled (the row at its start); with `dots`,
/// a dot for each item, a link to it, the current one marked. The page's
/// script scrolls the row from the paddles, disables the one at an end and
/// moves the mark: the gallery shows the state the page sets.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::Gallery;
///
/// let html = rsx! {
///     <Gallery id="camera" paddles=false dots=2 current=1 class="">
///         <li id="camera-1">"One"</li>
///         <li id="camera-2">"Two"</li>
///     </Gallery>
/// }
/// .render();
/// assert!(html.as_inner().contains(r##"<a class="cui-gallery-dot" href="#camera-1" aria-label="Item 1" aria-current="true"></a>"##));
/// ```
pub struct Gallery<'a, R: Renderable> {
    /// The row's id: the paddles' `aria-controls`, and the dots' links to
    /// its items, `#{id}-1`, `#{id}-2`, …. Required: an empty one panics.
    pub id: &'a str,
    /// Whether it shows its paddles (a page with a script to scroll the row).
    pub paddles: bool,
    /// How many dots, one for each item (`0` for none).
    pub dots: usize,
    /// The current dot, counted from 1: with dots, one outside them panics.
    pub current: usize,
    /// The page's utilities on the gallery (`cui-dark`); `""` for none.
    pub class: &'a str,
    /// Its items, each an `<li>`, with the id its dot links to.
    pub children: R,
}

impl<R: Renderable> Renderable for Gallery<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::Gallery needs an id");
        assert!(
            self.dots == 0 || (1..=self.dots).contains(&self.current),
            "ciderui::Gallery's current dot is outside its dots"
        );
        rsx! {
            <div class=(Class("cui-gallery", self.class))>
                <ul class="cui-gallery-scroller" id=(self.id)>(self.children)</ul>
                @if self.paddles {
                    <div class="cui-gallery-nav">
                        <button class="cui-gallery-paddle" type="button" aria-controls=(self.id) aria-label="Previous" disabled>(icon::PADDLE_PREVIOUS)</button>
                        <button class="cui-gallery-paddle" type="button" aria-controls=(self.id) aria-label="Next">(icon::PADDLE_NEXT)</button>
                    </div>
                }
                @if self.dots > 0 {
                    <ul class="cui-gallery-dots">
                        @for n in 1..=self.dots {
                            <li><a class="cui-gallery-dot" href=(format!("#{}-{n}", self.id)) aria-label=(format!("Item {n}")) aria-current="true"[n == self.current]></a></li>
                        }
                    </ul>
                }
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
            <Gallery id="gallery" paddles=true dots=4 current=1 class="">
                <li id="gallery-1"><Card dark=false href="/phone/" media=none eyebrow="" tag=false headline="Cider Phone 3" description="" price="From $1199" class="" /></li>
                <li id="gallery-2"><Card dark=false href="/phone-fold/" media=none eyebrow="" tag=false headline="Cider Phone Fold" description="" price="From $899" class="" /></li>
                <li id="gallery-3"><Card dark=false href="/watch/" media=none eyebrow="" tag=false headline="Cider Watch" description="" price="From $399" class="" /></li>
                <li id="gallery-4"><Card dark=false href="/buds/" media=none eyebrow="" tag=false headline="Cider Buds" description="" price="From $129" class="" /></li>
            </Gallery>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/gallery/examples/gallery.html"
            ))
        );
    }

    #[test]
    fn no_paddles_no_dots() {
        let html = rsx! {
            <Gallery id="g" paddles=false dots=0 current=0 class="cui-dark">"x"</Gallery>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<div class="cui-gallery cui-dark"><ul class="cui-gallery-scroller" id="g">x</ul></div>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs an id")]
    fn an_empty_id_panics() {
        let _ =
            rsx! { <Gallery id="" paddles=false dots=0 current=0 class="">"x"</Gallery> }.render();
    }

    #[test]
    #[should_panic(expected = "outside its dots")]
    fn a_current_dot_outside_panics() {
        let _ =
            rsx! { <Gallery id="g" paddles=false dots=3 current=4 class="">"x"</Gallery> }.render();
    }
}

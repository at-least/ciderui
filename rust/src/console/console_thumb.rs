use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-console-thumb` (`components/console/console-thumb`): a screenshot
/// shown small, a link to itself whole, its name under it, App Store
/// Connect's App Previews and Screenshots'. Their row,
/// `<div class="cui-console-thumbs">`, is the page's own.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleThumb;
///
/// let html = rsx! { <ConsoleThumb href="/shot.png" src="/shot.png" alt="shot.png" caption="shot.png" lazy=true class="" /> }.render();
/// assert_eq!(html.as_inner(), r#"<figure class="cui-console-thumb"><a href="/shot.png"><img src="/shot.png" alt="shot.png" loading="lazy"></a><figcaption class="cui-console-thumb-caption">shot.png</figcaption></figure>"#);
/// ```
pub struct ConsoleThumb<'a> {
    /// The image whole, what it links to.
    pub href: &'a str,
    /// The image shown.
    pub src: &'a str,
    /// What the image shows, its `alt`.
    pub alt: &'a str,
    /// Its name under it. Required: an empty one panics.
    pub caption: &'a str,
    /// Whether it loads only near the window (`loading="lazy"`), for a row
    /// of many.
    pub lazy: bool,
    /// The page's utilities on the `<figure>`; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleThumb<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.caption.is_empty(),
            "ciderui::console::ConsoleThumb needs a caption"
        );
        rsx! {
            <figure class=(Class("cui-console-thumb", self.class))>
                <a href=(self.href)><img src=(self.src) alt=(self.alt) loading="lazy"[self.lazy]></a>
                <figcaption class="cui-console-thumb-caption">(self.caption)</figcaption>
            </figure>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn renders_the_example() {
        let html = include_str!("../../../components/console/console-thumb/examples/thumbs.html");
        let src = |alt: &str| {
            let at = html.find(&format!(r#"" alt="{alt}""#)).unwrap();
            let start = html[..at].rfind(r#"src=""#).unwrap() + r#"src=""#.len();
            &html[start..at]
        };
        let thumbs = rsx! {
            <div class="cui-console-page">
                <div class="cui-console-thumbs">
                    <ConsoleThumb href="#checkout" src=(src("checkout.png")) alt="checkout.png" caption="checkout.png" lazy=false class="" />
                    <ConsoleThumb href="#cart" src=(src("cart.png")) alt="cart.png" caption="cart.png" lazy=false class="" />
                </div>
            </div>
        }
        .render();
        assert_eq!(*thumbs.as_inner(), crate::example::example(html));
    }

    #[test]
    #[should_panic(expected = "needs a caption")]
    fn an_empty_caption_panics() {
        let _ = rsx! { <ConsoleThumb href="/a.png" src="/a.png" alt="" caption="" lazy=false class="" /> }.render();
    }
}

use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-card` (`components/site/card`): the store's card, its picture filling it
/// and its text at its top, an eyebrow (or a tag), its headline, a
/// description and a price; the whole card its headline's link.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::site::Card;
///
/// let none = ();
/// let html = rsx! {
///     <Card dark=true href="/phone/" media=none eyebrow="NEW" tag=true headline="Cider Phone" description="" price="From $799" class="" />
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-card cui-dark"><div class="cui-card-content"><p class="cui-card-eyebrow" data-variant="tag">NEW</p>"#));
/// ```
pub struct Card<'a, M: Renderable> {
    /// Whether it is a dark card (`cui-dark`): black, its text white.
    pub dark: bool,
    /// Its headline's link, the whole card's.
    pub href: &'a str,
    /// Its picture (`cui-card-media` on an `<img>`), or `()` for none.
    pub media: M,
    /// The eyebrow above the headline; `""` for none.
    pub eyebrow: &'a str,
    /// Whether the eyebrow is a tag ("NEW"), in the tag's color. A tag with
    /// no eyebrow panics.
    pub tag: bool,
    /// The headline, an `<h3>`, its text the link's. Required: an empty one
    /// panics.
    pub headline: &'a str,
    /// The description; `""` for none.
    pub description: &'a str,
    /// The price; `""` for none.
    pub price: &'a str,
    /// The page's utilities on the card (another width: `w-full`); `""` for
    /// none.
    pub class: &'a str,
}

impl<M: Renderable> Renderable for Card<'_, M> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.headline.is_empty(),
            "ciderui::site::Card needs a headline"
        );
        assert!(
            !self.tag || !self.eyebrow.is_empty(),
            "ciderui::site::Card's tag needs an eyebrow"
        );
        let block = if self.dark {
            "cui-card cui-dark"
        } else {
            "cui-card"
        };
        rsx! {
            <div class=(Class(block, self.class))>
                (self.media)
                <div class="cui-card-content">
                    @if !self.eyebrow.is_empty() {
                        <p class="cui-card-eyebrow" data-variant="tag"[self.tag]>(self.eyebrow)</p>
                    }
                    <h3 class="cui-card-headline"><a class="cui-card-link" href=(self.href)>(self.headline)</a></h3>
                    @if !self.description.is_empty() {
                        <p class="cui-card-description">(self.description)</p>
                    }
                    @if !self.price.is_empty() {
                        <p class="cui-card-price">(self.price)</p>
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
    use crate::example::example;

    #[test]
    fn renders_the_examples() {
        let none = ();
        let card = rsx! {
            <Card dark=false href="/phone/" media=none eyebrow="NEW" tag=true headline="Cider Phone 3" description="Our most capable phone." price="From $1199" class="" />
        }
        .render();
        assert_eq!(
            *card.as_inner(),
            example(include_str!(
                "../../../components/site/card/examples/card.html"
            ))
        );
        let eyebrow = rsx! {
            <Card dark=false href="/shop/specialist/" media=none eyebrow="EXPERT HELP" tag=false headline="Talk it through with an expert." description="" price="" class="" />
        }
        .render();
        assert_eq!(
            *eyebrow.as_inner(),
            example(include_str!(
                "../../../components/site/card/examples/eyebrow.html"
            ))
        );
    }

    #[test]
    fn writes_its_picture_and_the_pages_utilities() {
        let media = rsx! { <img class="cui-card-media" src="/phone.jpg" alt=""> };
        let html = rsx! {
            <Card dark=false href="/phone/" media=(media) eyebrow="" tag=false headline="Cider Phone" description="" price="" class="w-full" />
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            r#"<div class="cui-card w-full"><img class="cui-card-media" src="/phone.jpg" alt=""><div class="cui-card-content"><h3 class="cui-card-headline"><a class="cui-card-link" href="/phone/">Cider Phone</a></h3></div></div>"#
        );
    }

    #[test]
    #[should_panic(expected = "needs a headline")]
    fn an_empty_headline_panics() {
        let none = ();
        let _ = rsx! {
            <Card dark=false href="/x/" media=none eyebrow="" tag=false headline="" description="" price="" class="" />
        }
        .render();
    }

    #[test]
    #[should_panic(expected = "tag needs an eyebrow")]
    fn a_tag_with_no_eyebrow_panics() {
        let none = ();
        let _ = rsx! {
            <Card dark=false href="/x/" media=none eyebrow="" tag=true headline="x" description="" price="" class="" />
        }
        .render();
    }

    #[test]
    fn dark_then_the_pages_utilities() {
        let none = ();
        let html = rsx! {
            <Card dark=true href="/x/" media=none eyebrow="" tag=false headline="x" description="" price="" class="w-full" />
        }
        .render();
        assert!(
            html.as_inner()
                .starts_with(r#"<div class="cui-card cui-dark w-full">"#)
        );
    }
}

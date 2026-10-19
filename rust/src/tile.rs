use hypertext::{Buffer, prelude::*};

/// A [`Tile`]'s size (`data-size`).
#[derive(Clone, Copy, Debug, Default, PartialEq, Eq)]
pub enum TileSize {
    /// A promo tile, two to a row from 735px: 580px, 490px, 500px tall.
    #[default]
    Default,
    /// A full-width tile (`data-size="lg"`): 692px, 650px, 500px, its type a
    /// step larger.
    Large,
}

/// `cui-tile` (`components/tile`): a block of its own ground, its headline,
/// subhead, callout and buttons centered at its top, its picture filling it;
/// with `link`, the whole tile a link, its buttons above it.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{Tile, TileSize};
///
/// let none = ();
/// let html = rsx! {
///     <Tile size=(TileSize::Large) rounded=false dark=true link=(Some(("/iphone/", "iPhone"))) media=none headline="iPhone" subhead="Hello." callout="">
///         <a class="cui-btn" data-size="cta" href="/iphone/">"Learn more"</a>
///     </Tile>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-tile cui-dark" data-size="lg"><a class="cui-tile-link" href="/iphone/" aria-label="iPhone"></a>"#));
/// ```
pub struct Tile<'a, M: Renderable, C: Renderable> {
    /// Its size.
    pub size: TileSize,
    /// Whether its corners are rounded (`cui-tile-rounded`, 28px).
    pub rounded: bool,
    /// Whether it is a dark block (`cui-dark`): black, its text light.
    pub dark: bool,
    /// The whole tile's link, (href, its name); `None` for none.
    pub link: Option<(&'a str, &'a str)>,
    /// Its picture (`cui-tile-media` on an `<img>`), or `()` for none.
    pub media: M,
    /// The headline, an `<h2>`. Required: an empty one panics.
    pub headline: &'a str,
    /// The subhead; `""` for none.
    pub subhead: &'a str,
    /// The callout, in the secondary text; `""` for none.
    pub callout: &'a str,
    /// The buttons (`cui-btn`, `data-size="cta"` in a large tile).
    pub children: C,
}

impl<M: Renderable, C: Renderable> Renderable for Tile<'_, M, C> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.headline.is_empty(), "ciderui::Tile needs a headline");
        let mut class = String::from("cui-tile");
        if self.rounded {
            class.push_str(" cui-tile-rounded");
        }
        if self.dark {
            class.push_str(" cui-dark");
        }
        let large = self.size == TileSize::Large;
        rsx! {
            <div class=(class) data-size="lg"[large]>
                (self.media)
                @if let Some((href, name)) = self.link {
                    <a class="cui-tile-link" href=(href) aria-label=(name)></a>
                }
                <div class="cui-tile-content">
                    <h2 class="cui-tile-headline">(self.headline)</h2>
                    @if !self.subhead.is_empty() {
                        <p class="cui-tile-subhead">(self.subhead)</p>
                    }
                    @if !self.callout.is_empty() {
                        <p class="cui-tile-callout">(self.callout)</p>
                    }
                    <div class="cui-tile-ctas">(self.children)</div>
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
        let tile = rsx! {
            <Tile size=(TileSize::Default) rounded=false dark=false link=(Some(("/upgrade/", "Apple Upgrade"))) media=none headline="Apple Upgrade" subhead="Get the latest iPhone." callout="">
                <a class="cui-btn" href="/upgrade/">"Learn more"</a>
            </Tile>
        }
        .render();
        assert_eq!(
            *tile.as_inner(),
            example(include_str!("../../components/tile/examples/tile.html"))
        );
        let hero = rsx! {
            <Tile size=(TileSize::Large) rounded=false dark=false link=(Some(("/iphone/", "iPhone Duo"))) media=none headline="iPhone Duo" subhead="Hello, hello." callout="Available starting 10.23">
                <a class="cui-btn" data-size="cta" href="/iphone/">"Learn more"</a>
                <a class="cui-btn" data-size="cta" data-variant="secondary" href="/buy/">"Pre-order"</a>
            </Tile>
        }
        .render();
        assert_eq!(
            *hero.as_inner(),
            example(include_str!("../../components/tile/examples/hero.html"))
        );
    }

    #[test]
    #[should_panic(expected = "needs a headline")]
    fn an_empty_headline_panics() {
        let none = ();
        let _ = rsx! {
            <Tile size=(TileSize::Default) rounded=true dark=false link=None media=none headline="" subhead="" callout="">"x"</Tile>
        }
        .render();
    }
}

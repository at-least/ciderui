use hypertext::{Buffer, prelude::*};

use crate::Class;

/// `cui-ribbon` (`components/site/ribbon`): the message above a page's content, a
/// band on the alternate ground, its text centered; with `announce`, it drops
/// in on the accent as the page shows it and turns to its ground.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::site::{MoreLink, Ribbon};
///
/// let html = rsx! {
///     <Ribbon announce=true class="">
///         <p>"New this week. "<MoreLink href="/new" text="See what's new" .. /></p>
///     </Ribbon>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<div class="cui-ribbon" data-variant="announce"><div class="cui-ribbon-content"><p>New this week. "#));
/// ```
pub struct Ribbon<'a, R: Renderable> {
    /// Whether it drops in on the accent and turns (`data-variant="announce"`).
    pub announce: bool,
    /// The page's utilities on the ribbon; `""` for none.
    pub class: &'a str,
    /// The message: a paragraph, a more link in it.
    pub children: R,
}

impl<R: Renderable> Renderable for Ribbon<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            <div class=(Class("cui-ribbon", self.class)) data-variant="announce"[self.announce]>
                <div class="cui-ribbon-content">(self.children)</div>
            </div>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::{example::example, site::MoreLink};

    #[test]
    fn renders_the_examples() {
        for (announce, file) in [
            (
                false,
                include_str!("../../../components/site/ribbon/examples/ribbon.html"),
            ),
            (
                true,
                include_str!("../../../components/site/ribbon/examples/announce.html"),
            ),
        ] {
            let html = rsx! {
                <Ribbon announce=(announce) class="">
                    <p>"Pre-order Cider Phone Fold now. Already have one saved for pre-order? "<MoreLink href="#checkout" text="Check out now" .. /></p>
                </Ribbon>
            }
            .render();
            assert_eq!(*html.as_inner(), example(file));
        }
    }
}

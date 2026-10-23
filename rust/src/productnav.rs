use hypertext::{Buffer, prelude::*};

use crate::{Class, icon};

/// `cui-productnav` (`components/productnav`): the store's shelf of small
/// cards ([`ProductNavItem`]s), each a product's image above its name, on a
/// `cui-shelf` with no header; with `paddles`, the shelf's large paddles.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::{ProductNav, ProductNavItem};
///
/// let html = rsx! {
///     <ProductNav id="products" label="Products" paddles=false class="">
///         <ProductNavItem href="/shop/laptops" name="Laptops" class="">
///             <img class="cui-productnav-image" src="/mac.png" width="200" height="130" alt="">
///         </ProductNavItem>
///     </ProductNav>
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"<span class="cui-productnav-name">Laptops</span>"#));
/// ```
pub struct ProductNav<'a, R: Renderable> {
    /// The scroller's id, the paddles' `aria-controls`. Required: an empty
    /// one panics.
    pub id: &'a str,
    /// The nav's name (`aria-label`). Required: an empty one panics.
    pub label: &'a str,
    /// Whether it shows its paddles (a page with a script to scroll the row).
    pub paddles: bool,
    /// The page's utilities on the nav; `""` for none.
    pub class: &'a str,
    /// Its cards, [`ProductNavItem`]s.
    pub children: R,
}

impl<R: Renderable> Renderable for ProductNav<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(!self.id.is_empty(), "ciderui::ProductNav needs an id");
        assert!(!self.label.is_empty(), "ciderui::ProductNav needs a label");
        rsx! {
            <nav class=(Class("cui-shelf cui-productnav", self.class)) aria-label=(self.label)>
                <div class="cui-shelf-row">
                    <ul class="cui-shelf-scroller" id=(self.id)>(self.children)</ul>
                    @if self.paddles {
                        <button class="cui-gallery-paddle cui-shelf-paddle" data-size="lg" type="button" aria-controls=(self.id) aria-label="Previous" disabled>(icon::PADDLE_PREVIOUS)</button>
                        <button class="cui-gallery-paddle cui-shelf-paddle" data-size="lg" type="button" aria-controls=(self.id) aria-label="Next">(icon::PADDLE_NEXT)</button>
                    }
                </div>
            </nav>
        }
        .render_to(buffer);
    }
}

/// A card of a [`ProductNav`]: a link, its image above its name.
pub struct ProductNavItem<'a, M: Renderable> {
    /// The link.
    pub href: &'a str,
    /// The product's name. Required: an empty one panics.
    pub name: &'a str,
    /// The page's utilities on its root, the `<li>`; `""` for none.
    pub class: &'a str,
    /// Its image, `cui-productnav-image` on an `<img>` (with its width and
    /// height: its shape sets its width).
    pub children: M,
}

impl<M: Renderable> Renderable for ProductNavItem<'_, M> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.name.is_empty(),
            "ciderui::ProductNavItem needs a name"
        );
        rsx! {
            <li class=(self.class)[!self.class.is_empty()]><a class="cui-productnav-item" href=(self.href)>(self.children)<span class="cui-productnav-name">(self.name)</span></a></li>
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    #[test]
    fn renders_the_example() {
        let items = [
            ("/shop/laptops", "Laptops"),
            ("/shop/phones", "Phones"),
            ("/shop/tablets", "Tablets"),
            ("/shop/watches", "Watches"),
            ("/shop/audio", "Audio"),
            ("/shop/accessories", "Accessories"),
        ];
        let html = rsx! {
            <ProductNav id="products" label="Products" paddles=false class="">
                @for (href, name) in items {
                    <ProductNavItem href=(href) name=(name) class="">
                        <img class="cui-productnav-image" src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyMDAiIGhlaWdodD0iMTMwIi8+" width="200" height="130" alt="">
                    </ProductNavItem>
                }
            </ProductNav>
        }
        .render();
        assert_eq!(
            *html.as_inner(),
            example(include_str!(
                "../../components/productnav/examples/productnav.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "needs a name")]
    fn an_empty_name_panics() {
        let _ = rsx! { <ProductNavItem href="/x" name="" class="">"x"</ProductNavItem> }.render();
    }

    #[test]
    fn an_item_takes_the_pages_utilities() {
        let html = rsx! { <ProductNavItem href="/x" name="x" class="hidden">"x"</ProductNavItem> }
            .render();
        assert!(
            html.as_inner()
                .starts_with(r#"<li class="hidden"><a class="cui-productnav-item""#)
        );
    }
}

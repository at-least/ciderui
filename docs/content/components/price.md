+++
title = "Price"
description = "The buy page’s price under its title, its amounts kept whole."

[extra]
category = "Store"
image = "pictures/price.png"
+++

The price under a product’s title on apple.com’s buy page: 17px on 25px at 600 (12px on 16px below 735px), in the text’s color. Each amount is a `cui-price-amount`, which never breaks across lines, so “$45.79/mo.” stays one.

```css
@import "ciderui/components/site/price/price.css";
```

{{ <cui_example file="price/price" title="A price" pad={true} /> }}

It is a class on the page’s own sentence, so the crate has no component for it.

+++
title = "Footer"
description = "apple.com’s footer: its notes, its directory, the shop line and the legal row."

[extra]
category = "Navigation"
image = "pictures/footer.png"
+++

apple.com’s footer, on the alternate ground: its footnotes over a rule, its directory of links in columns, the shop line, and the legal row (the copyright, the legal links, the locale). It stays light on a dark page, as apple.com’s does. Below 834px the directory becomes a column of sections, each opened by its title with no script: a checkbox and its label, as the local nav’s menu. This page’s own footer is one.

```css
@import "ciderui/components/footer/footer.css";
```

{{ <cui_example file="footer/footer" title="A footer" /> }}

| class | what |
| --- | --- |
| `cui-footer` | the footer, 12px on 16px |
| `cui-footer-content` | 980px at most, 22px in (16px to 833px) |
| `cui-footer-notes` | its footnotes, a rule under them |
| `cui-footer-directory` | the columns (`cui-footer-column`, a quarter of the row at most) |
| `cui-footer-section` | a title (`cui-footer-title`) over its links (`cui-footer-list`); to 833px a row over a rule, its checkbox `cui-footer-toggle` before it |
| `cui-footer-legal` | the shop line (`cui-footer-shop`) and the legal row (`cui-footer-row`: `cui-footer-copyright`, `cui-footer-links`, `cui-footer-locale`) |

Every part but the copyright is optional. The crate writes it with `Footer` and `FooterSection`, the chevron its `icon::FOOTER_CHEVRON`.

// The account's page as it opens (its Sign-In and Security section): the
// page's ground; the side nav (the person's block: the picture, the name's
// line and the line under it; the list, an item, a link and the current
// one); the page's header (its title, its description); the section's grid
// (a column) and a card (the button it is, the card's ground and edge, its
// padding, its title, its line, the glyph at its right).
import { TYPE, LINK, PAGE } from '../../apple.com/pages/common.mjs';

const PAD = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'];
const EDGE = ['border-top-width', 'border-top-style', 'border-top-color', 'border-top-left-radius', 'background-color', 'box-shadow'];
export default {
  path: '/account/manage',
  probes: [
    { name: 'page', select: 'body', props: PAGE },
    { name: 'nav', select: '.layout__nav', props: ['width', ...PAD], box: true },
    { name: 'person.picture', select: '.person-info__image svg', props: ['width', 'height', 'margin-bottom'], box: true },
    { name: 'person.title', select: '.person-info__name', props: [...TYPE, 'margin-top'], box: true },
    { name: 'person.line', select: '.person-info__email', props: TYPE, box: true },
    { name: 'list', select: 'ul.side-nav__list', props: ['margin-top', ...PAD], box: true },
    { name: 'list.item', select: 'li.side-nav__item:not(.side-nav__item--current)', props: ['height', 'margin-bottom', ...PAD], box: true },
    { name: 'list.link', select: '.side-nav__item:not(.side-nav__item--current) .side-nav__link', props: [...TYPE, ...LINK, ...PAD, 'display'], box: true },
    { name: 'list.current', select: '.side-nav__item--current .side-nav__link', props: [...TYPE, ...LINK, ...PAD], box: true },
    { name: 'header', select: 'header.page-header', props: [...PAD, 'margin-bottom'], box: true },
    { name: 'title', select: 'h1.page-title', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'description', select: '.page-description', props: [...TYPE, 'margin-top'], box: true },
    { name: 'column', select: '.section .row > .column', props: ['width', ...PAD, 'margin-bottom'], box: true },
    { name: 'card.button', select: 'button.button-expand', props: [...PAD, ...EDGE, 'width', 'height', 'cursor'], box: true },
    { name: 'card', select: '.card.card-border', props: [...EDGE, 'border-left-color', 'width', 'height'], box: true },
    { name: 'card.top', select: '.card .card-top', props: [...PAD, 'margin-top'], box: true },
    { name: 'card.title', select: '.card .card-title', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'card.line', select: '.card .card-line', props: [...TYPE, 'margin-top'], box: true },
    { name: 'card.glyph', select: '.card .card-right-element svg', props: ['width', 'height', 'color', 'fill'], box: true },
  ],
};

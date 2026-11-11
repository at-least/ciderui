// @external: apple.com's search page with results (`airpods`): its field (a magnifier
// at its start, a reset at its end), the count over the results, and a result's
// tile, a picture over its name and price, in a grid.
import { TYPE, BOX, FLEX, LINK } from './common.mjs';

export default {
  url: 'https://www.apple.com/us/search/airpods?src=globalnav',
  wait: 3000,
  probes: [
    { name: 'search', select: '.rf-serp-search', props: [...BOX], box: true },
    { name: 'box', select: '.rf-serp-searchbox-wrapper', props: [...BOX, 'width'], box: true },
    { name: 'field', select: '.rf-serp-searchbox .form-textbox', props: [...BOX, 'width', 'position'], box: true },
    { name: 'input', select: '.rf-serp-searchbox input.form-textbox-input', props: [...TYPE, ...BOX, 'height', 'border-left-width', 'border-left-color', 'padding-left', 'padding-right', 'border-top-right-radius'], box: true, states: ['hover', 'focus'] },
    { name: 'label', select: '.rf-serp-searchbox label.form-textbox-label', props: [...TYPE, 'top', 'left', 'position'], box: true },
    { name: 'icon.left', select: '.rf-serp-searchbox .form-textbox-lefticon', props: ['position', 'left', 'top', 'width', 'height', 'color'], box: true },
    { name: 'icon.left.button', select: '.rf-serp-searchbox .form-textbox-lefticon button', props: ['width', 'height', 'color', 'font-size', 'line-height'], box: true },
    { name: 'icon.right', select: '.rf-serp-searchbox .form-textbox-righticon', props: ['position', 'right', 'top', 'width', 'height', 'visibility', 'opacity'], box: true },
    { name: 'count', select: '.rf-serp-resultcount', props: [...TYPE, ...BOX], box: true },
    { name: 'grid', select: '.rf-accessoriesgrid', props: [...FLEX, 'grid-template-columns', 'width'], box: true },
    { name: 'tile', select: '.rf-producttile', props: [...BOX, 'width', 'flex-basis'], box: true },
    { name: 'tile.gallery', select: '.rf-producttile-gallery', props: [...BOX, 'width', 'height', 'overflow'], box: true },
    { name: 'tile.img', select: '.rf-producttile-gallery img', props: ['width', 'height', 'object-fit', 'border-top-left-radius'], box: true },
    { name: 'tile.info', select: '.rf-producttile-info', props: [...BOX, 'text-align'], box: true },
    { name: 'tile.name', select: '.rf-producttile-name', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'tile.name.link', select: '.rf-producttile-name a', props: [...TYPE, ...LINK], states: ['hover', 'focus-visible'] },
    { name: 'tile.price', select: '.rf-producttile-priceinfo', props: [...TYPE, 'margin-top'], box: true },
  ],
};

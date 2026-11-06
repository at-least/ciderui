import { TYPE, BOX, LINK, PAGE } from './common.mjs';

export const MARK = ['width', 'height', 'margin-right', 'position', 'top', 'left', 'vertical-align', 'display', 'background-size', 'background-repeat'];
const WIDE = [1600, 1200];
const NARROW = [900, 390];
export default {
  url: 'https://developer.apple.com/system-status/',
  wait: 3000,
  probes: [
    { name: 'page', select: 'body', props: PAGE },
    { name: 'title', select: 'h1.section-headline', props: [...TYPE, 'margin-bottom'], box: true },
    { name: 'legend', select: 'section.section-legend', props: [...TYPE, 'margin-top', 'margin-bottom'], widths: WIDE },
    { name: 'legend.item', select: '.section-legend li', props: ['display', 'margin-right', 'padding-left', 'padding-right'], box: true, widths: WIDE },
    { name: 'legend.mark', select: '.section-legend li > div', props: MARK, image: ['background-image'], box: true, widths: WIDE },
    { name: 'legend.word', select: '.section-legend li > span:not(.a11y)', props: TYPE, widths: WIDE },
    { name: 'table', select: '.section-lights table', props: ['border-collapse', 'border-spacing', 'margin-top', 'margin-bottom'], box: true, widths: WIDE },
    { name: 'row', select: '.section-lights table tr', props: ['background-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true, widths: WIDE },
    { name: 'cell', select: '.section-lights td', props: [...TYPE, ...BOX, 'vertical-align', 'text-align'], box: true, widths: WIDE },
    { name: 'cell.mark', select: '.section-lights td .light-available', props: MARK, image: ['background-image'], box: true, widths: WIDE },
    { name: 'cell.name', select: '.section-lights td .light-name > span:not(.a11y)', props: TYPE, widths: WIDE },
    { name: 'cell.link', select: '.section-lights td a.light-url-icon', props: ['width', 'height', 'display', 'margin-left', 'background-size'], image: ['background-image'], box: true, widths: WIDE },
    { name: 'links', select: '.section-links p.link-copy', props: [...TYPE, 'margin-top'] },
    { name: 'links.link', select: '.section-links p.link-copy a', props: [...TYPE, ...LINK], states: ['hover'] },
    { name: 'date', select: '.section-date p.date-copy', props: [...TYPE, 'margin-top'] },
    { name: 'all', select: '.lights-small .light-container', props: [...BOX, 'display', 'width'], box: true, widths: NARROW },
    { name: 'all.mark', select: '.lights-small .light-available', props: MARK, image: ['background-image'], box: true, widths: NARROW },
    { name: 'all.words', select: '.lights-small h2', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true, widths: NARROW },
    { name: 'all.button', select: '.lights-small h2 button', props: [...TYPE, ...BOX, 'cursor', 'text-align'], box: true, widths: NARROW },
    { name: 'all.toggle', select: '.lights-small .light-toggle-icon', props: MARK, image: ['background-image'], box: true, widths: NARROW },
  ],
};

// The account's overview: the page's ground; the tab nav (a tab, its glyph,
// a paddle, disabled while every tab shows); a card (its ground, corners
// and padding); the program card's title, a section's picture, title, text
// and link; the membership's rows (a row, its label and the word in it, its
// value) and the info button; the link that adds, a plus glyph after it.
// Its emails card and the team's address are not read.
import { TYPE, LINK, PAGE } from '../../apple.com/pages/common.mjs';

const PAD = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left'];
const ROW = 'li[class*="TeamMembershipDetails-module__TeamMembershipDetailsRow"]';
export default {
  path: '/account',
  probes: [
    { name: 'page', select: 'body', props: PAGE },
    { name: 'ground', select: 'div.App.background', props: ['background-color'], box: true },
    { name: 'tab', select: 'a[class*="DeveloperTabNav-module__TabNavItemLink"]', props: [...TYPE, 'width', 'height'], box: true },
    { name: 'tab.glyph', select: 'figure[class*="DeveloperTabNav-module__TabNavItemIcon"]', props: ['width', 'height', 'margin-bottom', 'background-color'], box: true },
    { name: 'paddle.disabled', select: 'button[class*="DeveloperTabNav-module__TabNavPaddle--"]:disabled', props: ['width', 'height', 'background-color', 'color', 'border-top-left-radius', 'opacity'], box: true },
    { name: 'card', select: 'section[class*="Card-module__Card--"]', props: [...PAD, 'margin-bottom', 'background-color', 'border-top-left-radius', 'box-shadow', 'border-top-width'], box: true },
    { name: 'card.title', select: '.ProgramResourcesCard h4', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'section.picture', select: '.CardSection img', props: ['width', 'height', 'margin-bottom'], box: true },
    { name: 'section.title', select: '.CardSection h5', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'section.text', select: '.CardSection p', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'section.link', select: 'li[class*="Card-module__CardSectionLink"] a', props: [...TYPE, ...LINK], box: true },
    { name: 'row', select: ROW, props: [...PAD, 'margin-bottom', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color'], box: true },
    { name: 'row.label', select: `${ROW} > div[class*="TeamMembershipDetailsColumn"]`, props: TYPE, box: true },
    { name: 'row.label.word', select: `${ROW} > div[class*="TeamMembershipDetailsColumn"] span.lighter`, props: TYPE },
    { name: 'row.value', select: `${ROW} > div[class*="TeamMembershipDetailsColumn"].column`, props: TYPE, box: true },
    { name: 'info', select: 'button.form-tooltip-button', props: ['width', 'height', 'color', 'background-color'], box: true },
    { name: 'add', select: 'a.icon-pluscircle', props: [...TYPE, ...LINK], box: true },
  ],
};

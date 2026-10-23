// A feedback opened (the owner's one, already read: opening an unread item
// marks it read): its card in the list, now the current one; the detail's
// byline (the monogram, the name, the date); the title's row and the
// actions button; the status box (a label and its value); a section (its
// lead, a label, its text); an attachment (its glyph, its link, the line
// under it); and the reply area at the foot (its field, the attach glyph,
// the send button, disabled while the field is empty).
import { TYPE, BOX, LINK } from '../../apple.com/pages/common.mjs';

const SIDES = ['border-left-width', 'border-left-style', 'border-left-color', 'border-right-width', 'border-right-style', 'border-right-color'];
const BUTTON = [...TYPE, ...BOX, ...SIDES, 'height', 'cursor', 'opacity'];
const ATTACHMENT = 'a[class*="AttachmentsSection__AttachmentGlyphLink"]';
export default {
  path: '/feedback/<app>',
  probes: [
    { name: 'item.current', select: '[class*="ListItemWrapper"] a[href^="/feedback/"] [class^="Card-sc-"]', props: ['background-color', ...SIDES, 'border-bottom-width', 'border-bottom-color'], box: true },
    { name: 'byline', select: 'div:has(> [class^="Monogram-sc-"])', props: [...BOX, 'height'], box: true },
    { name: 'byline.monogram', select: '.sb-avatar__text', props: [...TYPE, 'width', 'height', 'background-color', 'border-top-left-radius'], box: true },
    { name: 'byline.name', select: '[class^="Name-sc-"]', props: TYPE, box: true },
    { name: 'byline.date', select: '[class^="Date-sc-"]', props: TYPE, box: true },
    { name: 'title.row', select: 'div:has(> p[class*="ItemDetail__FeedbackTitle"])', props: [...BOX, 'height'], box: true },
    { name: 'title', select: 'p[class*="ItemDetail__FeedbackTitle"]', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'actions', select: 'button[class*="ActionsMenuStyles__CustomPrimaryButton"]', props: BUTTON, box: true },
    { name: 'status', select: '[class^="Status-sc-"]', props: [...BOX, ...SIDES], box: true },
    { name: 'status.label', select: '[class*="Status__StatusLabel"]', props: [...TYPE, 'width', 'min-width'], box: true },
    { name: 'status.value', select: '[class*="Status__StatusValue"]', props: TYPE, box: true },
    { name: 'section', select: '[class^="Section-sc-"]', props: BOX, box: true },
    { name: 'section.lead', select: '[class^="Description-sc-"]', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'section.label', select: '[class^="Section-sc-"] [class^="Header-sc-"]', props: [...TYPE, 'margin-top', 'margin-bottom'], box: true },
    { name: 'section.text', select: '[class^="Section-sc-"] [class^="Body-sc-"]', props: [...TYPE, 'white-space', 'margin-bottom'], box: true },
    { name: 'attachment.glyph', select: `${ATTACHMENT} svg`, props: ['width', 'height', 'color', 'fill'], box: true },
    { name: 'attachment.link', select: ATTACHMENT, nth: 1, props: [...TYPE, ...LINK], box: true },
    { name: 'attachment.line', select: 'p.size-time', props: TYPE, box: true },
    { name: 'reply', select: '[class^="FeedbackResponse-sc-"]', props: [...BOX, ...SIDES, 'position', 'bottom'], box: true },
    { name: 'reply.field', select: 'textarea[class*="FormTextArea"]', props: [...TYPE, ...BOX, ...SIDES, 'height', 'resize'], box: true },
    { name: 'reply.attach', select: 'button[class*="feedbackResponseForm__AttachmentButton"] svg', props: ['width', 'height', 'color', 'fill'], box: true },
    { name: 'reply.send.disabled', select: '[class^="ActionRow-sc-"] button[class*="PrimaryButton__Button"]:disabled', props: BUTTON, box: true },
  ],
};

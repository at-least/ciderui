// AirPods Pro's page, a light one: its stat (a large one: its captions and
// value) and a footnote's mark in its text.
import { TYPE } from './common.mjs';
export default {
  url: 'https://www.apple.com/airpods-pro/',
  wait: 2500,
  probes: [
    { name: 'stat.content', select: 'main .stat .stat-content', props: ['padding-top', 'padding-bottom', 'margin-top', 'margin-bottom'] },
    { name: 'stat.caption', select: 'main .stat .stat-caption', props: [...TYPE, 'display'] },
    { name: 'stat.value', select: 'main .stat .stat-value', props: [...TYPE, 'margin-top', 'margin-bottom', 'display'] },
    { name: 'footnote', select: 'main sup.footnote', props: ['font-feature-settings', ...TYPE, 'vertical-align', 'position', 'top'] },
    { name: 'footnote.link', select: 'main sup.footnote a', props: ['font-feature-settings', 'color', 'text-decoration-line', 'font-size', 'vertical-align'] },
  ],
};

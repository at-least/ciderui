// The newer look's toolkit's control row (--tb-component-ControlRow-*: the
// room after a control in a row of them, a section's after a group), read
// off the root as the page computes it: tokens.json's pass named no
// ControlRow, and a fourth TOKENS pass would be a download.
const SIZES = ['sm', 'md', 'lg', 'xl'];
export default {
  path: '/apps/<app>/analytics/overview',
  probes: [
    {
      name: 'root',
      select: 'html',
      props: SIZES.flatMap((s) => [
        `--tb-component-ControlRow-anatomy-Button-variant-controlSize-${s}-marginRight`,
        `--tb-component-ControlRow-anatomy-Button-variant-controlSize-${s}-plainBackedMarginRight`,
        `--tb-component-ControlRow-anatomy-Section-variant-controlSize-${s}-marginRight`,
      ]),
    },
  ],
};

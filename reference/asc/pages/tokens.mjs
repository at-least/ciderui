// The newer look's toolkit's tokens, on an app's analytics overview (the
// only pages that load it): what the console components (cui-console-*)
// are built from and held to. The type, sizes, spaces, radii, borders,
// shadows and timing; every color role and both palettes (the dark side of
// a console color is its light role's step of the dark palette, so the
// light palette names the step and the dark one gives the value); its
// v2 colors (a metric's card's ground is v2-bgPlaceholder); and the
// components ciderui builds a console part from, without the variants it
// builds none of (a display list's grouped and plain lists; a tab bar's
// elevated, filled, tinted and gray-plain ones).
export default {
  path: '/apps/<app>/analytics/overview',
  tokens: {
    prefixes: [
      '--tb-meta-',
      '--tb-typography-', '--tb-fontSize-', '--tb-fontWeight-', '--tb-fontFamily-', '--tb-lineHeights-',
      '--tb-space-', '--tb-size-', '--tb-radius-', '--tb-border-', '--tb-boxShadow-', '--tb-timingFunctions-', '--tb-breakpoint-',
      '--tb-color-role-', '--tb-color-palette-', '--tb-color-system-', '--tb-color-effect-', '--tb-color-v2-', '--tb-color-black', '--tb-color-white', '--tb-color-transparent',
      ...['Sidebar', 'DisplayList', 'Content', 'Table', 'Badge', 'Banner', 'Tabs', 'EmptyState', 'Popover', 'Modal', 'Button', 'InfoSystemButton',
        'TextInput', 'Select', 'TextArea', 'FileUpload', 'Checkbox', 'Toggle', 'Label', 'HelpText', 'InputError', 'FormField',
        'Chip', 'DL', 'Progress', 'Text', 'Link'].map((c) => `--tb-component-${c}-`),
    ],
    exclude: [
      '^--tb-component-DisplayList-.*-(grouped|groupedGray|plain)-',
      '^--tb-component-Tabs-.*-(elevatedPlain|filledPlain|tintedPlain|grayPlain)-',
    ],
  },
};

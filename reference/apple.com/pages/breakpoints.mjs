// Where apple.com's breakpoints fall: a headline (the content's type) and the
// global nav (its own breakpoint) stepped across each boundary its stylesheets
// name (734/735, 1068/1069, 1440/1441; the nav's 833/834).
const TYPE = ['font-size', 'line-height', 'letter-spacing', 'font-weight', 'font-family'];
export default {
  url: 'https://www.apple.com/',
  widths: [733, 734, 735, 736, 832, 833, 834, 835, 1067, 1068, 1069, 1070, 1439, 1440, 1441, 1442],
  probes: [
    { name: 'type.headline', inject: '<p class="typography-headline">Headline</p>', props: TYPE },
    { name: 'globalnav', select: '#globalnav', props: ['height'], vars: ['--r-globalnav-height'], box: true },
  ],
};

// The property lists the page specs read.
export const TYPE = ['font-family', 'font-size', 'line-height', 'font-weight', 'letter-spacing', 'color'];
export const BOX = ['padding-top', 'padding-right', 'padding-bottom', 'padding-left', 'margin-top', 'margin-bottom',
  'border-top-width', 'border-top-style', 'border-top-color', 'border-bottom-width', 'border-bottom-style', 'border-bottom-color',
  'border-top-left-radius', 'background-color', 'box-shadow', 'min-width', 'min-height', 'max-width'];
export const MOTION = ['transition-property', 'transition-duration', 'transition-timing-function'];
export const BLUR = ['backdrop-filter', '-webkit-backdrop-filter'];
export const LINK = ['text-decoration-line', 'text-decoration-color', 'text-underline-offset'];
export const FOCUS = ['outline-style', 'outline-color', 'outline-width', 'outline-offset'];
// the page's own: its text, its ground, how its text is drawn
export const PAGE = [...TYPE, 'background-color', '-webkit-font-smoothing', 'font-synthesis', 'text-rendering'];

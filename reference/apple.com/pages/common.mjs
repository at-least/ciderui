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
// a control's text: how it sits, wraps and shows a link
export const TEXT = ['display', 'text-align', 'white-space', 'text-decoration-line', 'cursor', 'vertical-align', 'opacity'];
// a flex or grid box's own layout
export const FLEX = ['display', 'flex-direction', 'flex-wrap', 'align-items', 'justify-content', 'row-gap', 'column-gap'];
// an element's CSS animation, as it names it
export const ANIMATION = ['animation-name', 'animation-duration', 'animation-delay', 'animation-timing-function', 'animation-fill-mode', 'animation-iteration-count'];

// a spec read again on a touch screen (`touch`): its probes of these names,
// in its order (a probe that clicks is named with the ones after it), each
// with states forcing these alone (:hover)
export const touched = (spec, names, states = ['hover']) => {
  const probes = spec.probes.filter(p => names.includes(p.name));
  const missing = names.filter(n => !probes.some(p => p.name === n));
  if (missing.length) throw new Error(`no probe ${missing.join(', ')} in ${spec.url}'s spec`);
  return { ...spec, touch: true, probes: probes.map(p => p.states ? { ...p, states: p.states.filter(s => states.includes(s)) } : p) };
};

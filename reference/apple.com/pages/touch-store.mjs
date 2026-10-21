// The store on a touch screen: its buttons and its cards (a card's lift, its
// link), hovered; a card with the focus in it. Read only.
import store from './store.mjs';
import { touched } from './common.mjs';

export default touched(store, ['button', 'button.secondary', 'button.neutral', 'button.secondary-neutral', 'card', 'card.dark', 'card.link'], ['hover', 'focus-within']);

import store from './store.mjs';
import { touched } from './common.mjs';

export default touched(store, ['button', 'button.secondary', 'button.neutral', 'button.secondary-neutral', 'card', 'card.dark', 'card.link'], ['hover', 'focus-within']);

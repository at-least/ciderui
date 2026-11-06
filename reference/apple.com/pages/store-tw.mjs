// The Traditional Chinese store's type scale, `store`'s probes on its page, which
// sets them apart for its language (tokens.css's [lang|="zh"], type.css's :lang(zh)).
import store from './store.mjs';
export default { url: 'https://www.apple.com/tw/store', probes: store.probes.filter((p) => p.name === 'page.body' || p.name.startsWith('type.t-')) };

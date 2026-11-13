// @invariant: Apple Account's pages (account.apple.com), read and checked by App Store
// Connect's command line (../asc/read.mjs).
import { cli } from '../asc/read.mjs';

await cli(new URL('.', import.meta.url), 'https://account.apple.com');

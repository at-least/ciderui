// @invariant: The Apple Developer account's pages (developer.apple.com/account), read
// and checked by App Store Connect's command line (../asc/read.mjs).
import { cli } from '../asc/read.mjs';

await cli(new URL('.', import.meta.url), 'https://developer.apple.com');

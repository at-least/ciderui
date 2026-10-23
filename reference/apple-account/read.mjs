// Apple Account's pages (account.apple.com), read and checked as App Store
// Connect's are:
//   node read.mjs <spec>          the script to run in the page
//   node read.mjs --check <spec>  once results/<spec>.json is written
import { cli } from '../asc/read.mjs';

await cli(new URL('.', import.meta.url));

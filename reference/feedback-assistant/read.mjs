// @invariant: Feedback Assistant's pages, read and checked by App Store Connect's
// command line (../asc/read.mjs).
import { cli } from '../asc/read.mjs';

await cli(new URL('.', import.meta.url), 'https://feedbackassistant.apple.com');

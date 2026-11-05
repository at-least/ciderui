# The test image with Google Chrome beside Playwright's Chromium, for a probe
# whose page plays what only Chrome can (tv.apple.com's trailers: Widevine).
# Chrome is the version Playwright's installer fetches that day, not pinned:
# a result records the one it read in.
ARG BASE
FROM ${BASE}
RUN npx -y playwright@1.62.1 install --with-deps chrome

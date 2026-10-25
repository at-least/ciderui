# The image every browser step runs in (test/chromium.mjs refuses to start
# outside it): Playwright's own, at playwright-chromium's version, pinned by
# digest, its Chromium and fonts the same on every machine; zola at the
# version the docs are built with, its archive checked. The repository is
# mounted at run time (scripts/docker.mjs); node_modules is the image's.
FROM mcr.microsoft.com/playwright:v1.62.1-noble@sha256:dcc5531e97840b9b5e794f2814476b21571c5124a3fca2267d73041f56e7580e

ADD --checksum=sha256:8f5132b3522412d04e395e0b25f6d68613ad272a873e54a2b3ebf664873024a4 \
  https://github.com/getzola/zola/releases/download/v0.23.6/zola-v0.23.6-x86_64-unknown-linux-gnu.tar.gz /tmp/zola.tar.gz
RUN tar -xzf /tmp/zola.tar.gz -C /usr/local/bin zola && rm /tmp/zola.tar.gz

ENV CIDERUI_CONTAINER=1 PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1
WORKDIR /src/ciderui
COPY package.json package-lock.json ./
RUN npm ci --no-audit --no-fund && chmod -R a+rX node_modules

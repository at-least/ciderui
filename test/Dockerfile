# @external: Playwright's image at playwright-chromium's version, pinned by digest: its
# Chromium and fonts the same on every machine.
FROM mcr.microsoft.com/playwright:v1.62.1-noble@sha256:dcc5531e97840b9b5e794f2814476b21571c5124a3fca2267d73041f56e7580e

ENV CIDERUI_CONTAINER=1 PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1
WORKDIR /src/ciderui
COPY package.json package-lock.json ./
RUN npm ci --no-audit --no-fund && chmod -R a+rX node_modules

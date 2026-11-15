# @external: Playwright's image at playwright-go's driver version, pinned by digest: its
# Chromium and fonts the same on every machine.
FROM mcr.microsoft.com/playwright:v1.62.1-noble@sha256:dcc5531e97840b9b5e794f2814476b21571c5124a3fca2267d73041f56e7580e

ADD --checksum=sha256:7226736c2339719cda28fbf3a54d71aa63bcefe376e98cef00034268dab24840 https://go.dev/dl/go1.27.4.linux-amd64.tar.gz /tmp/go.tar.gz
RUN tar -C /usr/local -xzf /tmp/go.tar.gz && rm /tmp/go.tar.gz
ENV CIDERUI_CONTAINER=1 PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1 PATH=/usr/local/go/bin:$PATH GOTOOLCHAIN=local
WORKDIR /src/ciderui
COPY package.json package-lock.json ./
RUN npm ci --no-audit --no-fund && chmod -R a+rX node_modules
# @invariant: the module's requirements, the pinned tools and Playwright's driver in the
# image: a run fetches nothing, and its build cache (a volume, scripts/docker.mjs)
# compiles the repository alone
COPY go.mod go.sum ./
COPY internal/pinned internal/pinned
COPY scripts/fetch scripts/fetch
RUN go mod download && go run ./scripts/fetch && go build github.com/mxschmitt/playwright-go

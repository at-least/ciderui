# @external: Playwright's image at playwright-go's driver version, pinned by digest: its
# Chromium and fonts the same on every machine.
FROM mcr.microsoft.com/playwright:v1.62.1-noble@sha256:dcc5531e97840b9b5e794f2814476b21571c5124a3fca2267d73041f56e7580e

# @external: amd64, lab's Docker (user 2026-11-20: 「改在lab開發，可以用docker跑playright」;
# the Macs have no Docker). The image and its Go one architecture: under OrbStack an amd64 Go
# beside the arm64 Chromium picked an amd64 Tailwind that failed, "Dynamic loader not found".
ADD --checksum=sha256:7226736c2339719cda28fbf3a54d71aa63bcefe376e98cef00034268dab24840 https://go.dev/dl/go1.27.4.linux-amd64.tar.gz /tmp/go.tar.gz
RUN tar -C /usr/local -xzf /tmp/go.tar.gz && rm /tmp/go.tar.gz
ENV CIDERUI_CONTAINER=1 PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1 PATH=/usr/local/go/bin:$PATH GOTOOLCHAIN=local
WORKDIR /src/ciderui
# @invariant: the module's requirements, the pinned tools and Playwright's driver in the
# image: a run fetches nothing, and its build cache (a volume, scripts/docker)
# compiles the repository alone
COPY go.mod go.sum ./
COPY internal/pinned internal/pinned
COPY scripts/fetch scripts/fetch
RUN go mod download && go run ./scripts/fetch && go build github.com/mxschmitt/playwright-go

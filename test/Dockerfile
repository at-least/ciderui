# @external: Playwright's image at playwright-go's driver version, pinned by digest: its
# Chromium and fonts the same on every machine.
FROM mcr.microsoft.com/playwright:v1.62.1-noble@sha256:dcc5531e97840b9b5e794f2814476b21571c5124a3fca2267d73041f56e7580e

# @external: arm64, mini's OrbStack (user 2026-11-17: Docker runs on mini, not lab). Under
# OrbStack an amd64 Go runs through Rosetta beside the arm64 Chromium, and the amd64 Tailwind
# it then picks fails: "Dynamic loader not found: /lib64/ld-linux-x86-64.so.2".
ADD --checksum=sha256:3392582988a273a87ab5043ee1b2c9c17a2c3c00580bf9326c47a2e0aaaeac92 https://go.dev/dl/go1.27.4.linux-arm64.tar.gz /tmp/go.tar.gz
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

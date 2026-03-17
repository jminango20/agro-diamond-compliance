# ── Stage 1: Build ──
FROM node:20-alpine AS builder

RUN corepack enable && corepack prepare pnpm@10.26.1 --activate

WORKDIR /app

# Copy workspace root files for pnpm install
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml .npmrc* ./
COPY packages/app/package.json packages/app/package.json

# Install deps (only app workspace)
RUN pnpm install --frozen-lockfile --filter @diamond-erc3643/app...

# Copy app source
COPY packages/app/ packages/app/

# Build args → env vars for Vite (baked at build time)
ARG VITE_DIAMOND_ADDRESS
ARG VITE_INDEXER_URL
ARG VITE_RPC_URL
ARG VITE_COUNTRY_RESTRICT_MODULE
ARG VITE_MAX_BALANCE_MODULE
ARG VITE_MAX_HOLDERS_MODULE

WORKDIR /app/packages/app
RUN pnpm build

# ── Stage 2: Serve ──
FROM nginx:alpine

# Custom nginx config for SPA routing
COPY packages/app/nginx.conf /etc/nginx/conf.d/default.conf

# Copy built assets
COPY --from=builder /app/packages/app/dist /usr/share/nginx/html

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s \
  CMD wget -qO- http://127.0.0.1/ || exit 1

CMD ["nginx", "-g", "daemon off;"]

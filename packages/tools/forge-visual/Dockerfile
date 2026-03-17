# ── Stage 1: Build frontend + server ──
FROM node:20-alpine AS builder

RUN corepack enable && corepack prepare pnpm@10.26.1 --activate

WORKDIR /app

# Copy workspace root files for pnpm install
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml .npmrc* ./
COPY packages/tools/forge-visual/package.json packages/tools/forge-visual/package.json

# Install deps
RUN pnpm install --frozen-lockfile --filter @diamond-erc3643/forge-visual...

# Copy source
COPY packages/tools/forge-visual/ packages/tools/forge-visual/

WORKDIR /app/packages/tools/forge-visual

# Build frontend (Vite) + server (tsc)
RUN pnpm build

# ── Stage 2: Runtime with Foundry ──
FROM node:20-slim

# Install Foundry
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN curl -L https://foundry.paradigm.xyz | bash \
    && /root/.foundry/bin/foundryup

ENV PATH="/root/.foundry/bin:${PATH}"

# Install tsx for running TypeScript server
RUN npm i -g tsx

WORKDIR /app

# Copy built forge-visual (frontend dist + compiled server)
COPY --from=builder /app/packages/tools/forge-visual/dist ./dist
COPY --from=builder /app/packages/tools/forge-visual/server ./server
COPY --from=builder /app/packages/tools/forge-visual/package.json ./package.json

# Install production deps directly (avoids pnpm symlink issues)
RUN npm install --omit=dev

# Copy contracts source (forge needs it to run tests)
COPY packages/contracts/ /project/

# FORGE_PROJECT points to the Solidity project root
ENV FORGE_PROJECT=/project
ENV PORT=3999
ENV NODE_ENV=production

EXPOSE 3999

HEALTHCHECK --interval=30s --timeout=5s --start-period=15s \
  CMD curl -f http://localhost:3999/api/health || exit 1

CMD ["tsx", "server/index.ts"]

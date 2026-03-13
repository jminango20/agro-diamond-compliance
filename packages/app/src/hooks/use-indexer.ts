import { useQuery } from "@tanstack/react-query";

const INDEXER_URL = import.meta.env.VITE_INDEXER_URL ?? "";
const GRAPHQL_ENDPOINT = INDEXER_URL ? `${INDEXER_URL}/graphql` : "";

async function gqlFetch<T>(query: string, variables?: Record<string, unknown>): Promise<T> {
  const response = await fetch(GRAPHQL_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query, variables }),
  });

  if (!response.ok) {
    throw new Error(`Indexer request failed: ${response.status}`);
  }

  const json = await response.json();

  if (json.errors) {
    throw new Error(json.errors[0]?.message ?? "GraphQL error");
  }

  return json.data as T;
}

// ---------- Tokens ----------

export interface IndexerToken {
  id: string;
  totalSupply: string;
  holderCount: number;
}

interface TokensResponse {
  tokens: IndexerToken[];
}

const TOKENS_QUERY = `
  query Tokens {
    tokens {
      id
      totalSupply
      holderCount
    }
  }
`;

export function useIndexerTokens() {
  return useQuery<IndexerToken[]>({
    queryKey: ["indexer", "tokens"],
    queryFn: async () => {
      const data = await gqlFetch<TokensResponse>(TOKENS_QUERY);
      return data.tokens;
    },
    enabled: !!GRAPHQL_ENDPOINT,
  });
}

// ---------- Token Detail ----------

export interface IndexerTokenDetail {
  id: string;
  totalSupply: string;
  holderCount: number;
  holders: { address: string; balance: string }[];
  events: IndexerEvent[];
}

interface TokenDetailResponse {
  token: IndexerTokenDetail | null;
}

const TOKEN_DETAIL_QUERY = `
  query Token($id: String!) {
    token(id: $id) {
      id
      totalSupply
      holderCount
      holders(first: 50) {
        address
        balance
      }
      events(first: 20) {
        txHash
        block
        logIndex
        from
        to
        tokenId
        amount
        eventType
      }
    }
  }
`;

export function useIndexerToken(tokenId: string | undefined) {
  return useQuery<IndexerTokenDetail | null>({
    queryKey: ["indexer", "token", tokenId],
    queryFn: async () => {
      const data = await gqlFetch<TokenDetailResponse>(TOKEN_DETAIL_QUERY, { id: tokenId });
      return data.token;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!tokenId,
  });
}

// ---------- Events ----------

export interface IndexerEvent {
  txHash: string;
  block: number;
  logIndex: number;
  from: string;
  to: string;
  tokenId: string;
  amount: string;
  eventType: string;
}

interface EventsResponse {
  events: IndexerEvent[];
}

const EVENTS_QUERY = `
  query Events($first: Int!) {
    events(first: $first) {
      txHash
      block
      logIndex
      from
      to
      tokenId
      amount
      eventType
    }
  }
`;

export function useIndexerEvents(first: number = 20) {
  return useQuery<IndexerEvent[]>({
    queryKey: ["indexer", "events", first],
    queryFn: async () => {
      const data = await gqlFetch<EventsResponse>(EVENTS_QUERY, { first });
      return data.events;
    },
    enabled: !!GRAPHQL_ENDPOINT,
  });
}

// ---------- Holder ----------

export interface IndexerHolder {
  address: string;
  balance: string;
}

interface HolderResponse {
  holder: IndexerHolder | null;
}

const HOLDER_QUERY = `
  query Holder($tokenId: String!, $address: String!) {
    holder(tokenId: $tokenId, address: $address) {
      address
      balance
    }
  }
`;

export function useIndexerHolder(tokenId: string | undefined, address: string | undefined) {
  return useQuery<IndexerHolder | null>({
    queryKey: ["indexer", "holder", tokenId, address],
    queryFn: async () => {
      const data = await gqlFetch<HolderResponse>(HOLDER_QUERY, { tokenId, address });
      return data.holder;
    },
    enabled: !!GRAPHQL_ENDPOINT && !!tokenId && !!address,
  });
}

// ---------- Status ----------

export interface IndexerStatus {
  lastBlock: number;
  tokenCount: number;
}

interface StatusResponse {
  status: IndexerStatus;
}

const STATUS_QUERY = `
  query Status {
    status {
      lastBlock
      tokenCount
    }
  }
`;

export function useIndexerStatus() {
  return useQuery<IndexerStatus>({
    queryKey: ["indexer", "status"],
    queryFn: async () => {
      const data = await gqlFetch<StatusResponse>(STATUS_QUERY);
      return data.status;
    },
    enabled: !!GRAPHQL_ENDPOINT,
    refetchInterval: 10_000,
  });
}

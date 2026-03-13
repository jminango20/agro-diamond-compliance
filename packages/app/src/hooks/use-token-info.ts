
import { useReadContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

interface TokenInfo {
  name: string;
  symbol: string;
  uri: string;
  totalSupply: bigint;
  supplyCap: bigint;
  holderCount: bigint;
  issuer: `0x${string}`;
  paused: boolean;
}

export function useTokenInfo(tokenId: bigint | undefined) {
  const { data, isLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "tokenInfo",
    args: tokenId !== undefined ? [tokenId] : undefined,
    query: { enabled: tokenId !== undefined },
  });

  const info = data as TokenInfo | undefined;

  return {
    name: info?.name ?? "",
    symbol: info?.symbol ?? "",
    uri: info?.uri ?? "",
    totalSupply: info?.totalSupply ?? 0n,
    supplyCap: info?.supplyCap ?? 0n,
    holderCount: info?.holderCount ?? 0n,
    issuer: info?.issuer ?? ("0x0" as `0x${string}`),
    paused: info?.paused ?? false,
    isLoading,
  };
}


import { useMemo } from "react";
import type { Address } from "viem";
import { useReadContract, useReadContracts } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

export interface AssetConfig {
  tokenId: bigint;
  name: string;
  symbol: string;
  uri: string;
  supplyCap: bigint;
  totalSupply: bigint;
  identityProfileId: number;
  complianceModules: Address[];
  issuer: Address;
  paused: boolean;
  exists: boolean;
  allowedCountries: number[];
}

export function useAssets() {
  const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;

  const {
    data: tokenIds,
    isLoading: isLoadingIds,
  } = useReadContract({
    ...diamond,
    functionName: "getRegisteredTokenIds",
  });

  const configContracts = useMemo(() => {
    if (!tokenIds || !Array.isArray(tokenIds)) return [];

    return (tokenIds as bigint[]).map((tokenId) => ({
      ...diamond,
      functionName: "getAssetConfig" as const,
      args: [tokenId] as const,
    }));
  }, [tokenIds]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { data: configs, isLoading: isLoadingConfigs } = useReadContracts({
    contracts: configContracts as any,
    query: { enabled: !!tokenIds && (tokenIds as bigint[]).length > 0 },
  });

  const assets = useMemo(() => {
    if (!configs || !tokenIds) return [];

    return (tokenIds as bigint[]).reduce<AssetConfig[]>((acc, tokenId, i) => {
      const result = configs[i];
      if (result?.status === "success" && result.result) {
        acc.push({ tokenId, ...(result.result as Omit<AssetConfig, "tokenId">) });
      }
      return acc;
    }, []);
  }, [configs, tokenIds]);

  return {
    assets,
    tokenIds: (tokenIds as bigint[]) ?? [],
    isLoading: isLoadingIds || isLoadingConfigs,
  };
}

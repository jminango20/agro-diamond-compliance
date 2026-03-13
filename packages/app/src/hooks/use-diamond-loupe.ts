
import { useReadContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

export interface FacetInfo {
  facetAddress: `0x${string}`;
  functionSelectors: readonly `0x${string}`[];
}

export function useDiamondLoupe() {
  const { data, isLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "facets",
  });

  return {
    facets: (data as FacetInfo[]) ?? [],
    isLoading,
  };
}

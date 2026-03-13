
import { useMemo } from "react";
import type { Address } from "viem";
import { useReadContracts } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

export function useBalance(account: Address | undefined, tokenId: bigint | undefined) {
  const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;
  const enabled = !!account && tokenId !== undefined;

  const contracts = useMemo(() => {
    if (!enabled) return [];

    return [
      {
        ...diamond,
        functionName: "balanceOf" as const,
        args: [account!, tokenId!] as const,
      },
      {
        ...diamond,
        functionName: "partitionBalanceOf" as const,
        args: [account!, tokenId!] as const,
      },
    ] as const;
  }, [account, tokenId, enabled]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { data, isLoading } = useReadContracts({
    contracts: contracts as any,
    query: { enabled },
  });

  const balance = (data?.[0]?.result as bigint) ?? 0n;

  const partition = data?.[1]?.result as
    | readonly [bigint, bigint, bigint, bigint]
    | undefined;

  return {
    balance,
    free: partition?.[0] ?? 0n,
    locked: partition?.[1] ?? 0n,
    custody: partition?.[2] ?? 0n,
    pendingSettlement: partition?.[3] ?? 0n,
    isLoading,
  };
}

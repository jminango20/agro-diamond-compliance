
import type { Address } from "viem";
import { useReadContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

export function useCompliance(
  tokenId: bigint | undefined,
  from: Address | undefined,
  to: Address | undefined,
  amount: bigint | undefined,
) {
  const enabled =
    tokenId !== undefined &&
    !!from &&
    !!to &&
    amount !== undefined;

  const { data, isLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "canTransfer",
    args: enabled ? [tokenId!, from!, to!, amount!, "0x"] : undefined,
    query: { enabled },
  });

  return {
    canTransfer: (data as boolean) ?? false,
    isLoading,
  };
}

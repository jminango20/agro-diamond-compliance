
import type { Address } from "viem";
import { useWaitForTransactionReceipt, useWriteContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

export function useMint() {
  const { writeContract, data: hash, isPending } = useWriteContract();
  const { isSuccess } = useWaitForTransactionReceipt({ hash });

  function write(tokenId: bigint, to: Address, amount: bigint) {
    writeContract({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "mint",
      args: [tokenId, to, amount],
    });
  }

  return { write, isPending, isSuccess, hash };
}

export function useBurn() {
  const { writeContract, data: hash, isPending } = useWriteContract();
  const { isSuccess } = useWaitForTransactionReceipt({ hash });

  function write(tokenId: bigint, from: Address, amount: bigint) {
    writeContract({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "burn",
      args: [tokenId, from, amount],
    });
  }

  return { write, isPending, isSuccess, hash };
}

export function useForcedTransfer() {
  const { writeContract, data: hash, isPending } = useWriteContract();
  const { isSuccess } = useWaitForTransactionReceipt({ hash });

  function write(
    tokenId: bigint,
    from: Address,
    to: Address,
    amount: bigint,
    reason: string,
  ) {
    writeContract({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "forcedTransfer",
      args: [tokenId, from, to, amount, reason],
    });
  }

  return { write, isPending, isSuccess, hash };
}

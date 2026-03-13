
import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

export function useDiamond() {
  return { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;
}

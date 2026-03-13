
import { useMemo } from "react";
import { keccak256, toHex } from "viem";
import { useAccount, useReadContracts } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

const ISSUER_ROLE = keccak256(toHex("ISSUER_ROLE"));
const TRANSFER_AGENT = keccak256(toHex("TRANSFER_AGENT"));
const RECOVERY_AGENT = keccak256(toHex("RECOVERY_AGENT"));
const PAUSER_ROLE = keccak256(toHex("PAUSER_ROLE"));
const GOVERNANCE_ROLE = keccak256(toHex("GOVERNANCE_ROLE"));
const COMPLIANCE_ADMIN = keccak256(toHex("COMPLIANCE_ADMIN"));

export function useRole() {
  const { address } = useAccount();

  const contracts = useMemo(() => {
    if (!address) return [];

    const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;

    return [
      { ...diamond, functionName: "owner" },
      { ...diamond, functionName: "hasRole", args: [ISSUER_ROLE, address] },
      {
        ...diamond,
        functionName: "hasRole",
        args: [TRANSFER_AGENT, address],
      },
      {
        ...diamond,
        functionName: "hasRole",
        args: [RECOVERY_AGENT, address],
      },
      { ...diamond, functionName: "hasRole", args: [PAUSER_ROLE, address] },
      {
        ...diamond,
        functionName: "hasRole",
        args: [GOVERNANCE_ROLE, address],
      },
      {
        ...diamond,
        functionName: "hasRole",
        args: [COMPLIANCE_ADMIN, address],
      },
    ] as const;
  }, [address]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { data, isLoading } = useReadContracts({
    contracts: contracts as any,
    query: { enabled: !!address },
  });

  return {
    isOwner:
      !!data?.[0]?.result &&
      (data[0].result as string).toLowerCase() === address?.toLowerCase(),
    isIssuer: (data?.[1]?.result as boolean) ?? false,
    isAgent: (data?.[2]?.result as boolean) ?? false,
    isPauser: (data?.[4]?.result as boolean) ?? false,
    isRecoveryAgent: (data?.[3]?.result as boolean) ?? false,
    isGovernance: (data?.[5]?.result as boolean) ?? false,
    isComplianceAdmin: (data?.[6]?.result as boolean) ?? false,
    isLoading,
  };
}

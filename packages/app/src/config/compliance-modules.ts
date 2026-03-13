import type { Address } from "viem";

export interface ComplianceModuleInfo {
  name: string;
  description: string;
  address: Address | null;
}

/**
 * Available compliance modules.
 * Addresses are populated after deployment via DeployComplianceModules.s.sol.
 * Set to null if not yet deployed.
 */
export const COMPLIANCE_MODULES: ComplianceModuleInfo[] = [
  {
    name: "CountryRestrictModule",
    description: "Blocks transfers to/from wallets in restricted countries (deny-list per tokenId)",
    address: (import.meta.env.VITE_COUNTRY_RESTRICT_MODULE as Address) ?? null,
  },
  {
    name: "MaxBalanceModule",
    description: "Limits maximum token balance per wallet per tokenId (concentration risk)",
    address: (import.meta.env.VITE_MAX_BALANCE_MODULE as Address) ?? null,
  },
  {
    name: "MaxHoldersModule",
    description: "Limits maximum number of unique holders per tokenId (e.g., Reg D 506(b))",
    address: (import.meta.env.VITE_MAX_HOLDERS_MODULE as Address) ?? null,
  },
];

export function getDeployedModules(): ComplianceModuleInfo[] {
  return COMPLIANCE_MODULES.filter((m) => m.address != null);
}

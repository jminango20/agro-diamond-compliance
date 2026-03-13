import { keccak256, toBytes } from 'viem';

export const ROLES = {
  ISSUER_ROLE: keccak256(toBytes('ISSUER_ROLE')),
  TRANSFER_AGENT: keccak256(toBytes('TRANSFER_AGENT')),
  RECOVERY_AGENT: keccak256(toBytes('RECOVERY_AGENT')),
  PAUSER_ROLE: keccak256(toBytes('PAUSER_ROLE')),
  GOVERNANCE_ROLE: keccak256(toBytes('GOVERNANCE_ROLE')),
  COMPLIANCE_ADMIN: keccak256(toBytes('COMPLIANCE_ADMIN')),
} as const;

export const ROLE_NAMES: Record<string, string> = {
  [ROLES.ISSUER_ROLE]: 'Issuer',
  [ROLES.TRANSFER_AGENT]: 'Transfer Agent',
  [ROLES.RECOVERY_AGENT]: 'Recovery Agent',
  [ROLES.PAUSER_ROLE]: 'Pauser',
  [ROLES.GOVERNANCE_ROLE]: 'Governance',
  [ROLES.COMPLIANCE_ADMIN]: 'Compliance Admin',
};

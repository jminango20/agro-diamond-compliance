import type { Abi, Address } from "viem";

import AccessControlFacetAbi from "@/abi/AccessControlFacet.abi.json";
import AssetGroupFacetAbi from "@/abi/AssetGroupFacet.abi.json";
import AssetManagerFacetAbi from "@/abi/AssetManagerFacet.abi.json";
import ClaimTopicsFacetAbi from "@/abi/ClaimTopicsFacet.abi.json";
import ComplianceRouterFacetAbi from "@/abi/ComplianceRouterFacet.abi.json";
import DiamondCutFacetAbi from "@/abi/DiamondCutFacet.abi.json";
import DiamondLoupeFacetAbi from "@/abi/DiamondLoupeFacet.abi.json";
import DividendFacetAbi from "@/abi/DividendFacet.abi.json";
import EmergencyFacetAbi from "@/abi/EmergencyFacet.abi.json";
import ERC1155FacetAbi from "@/abi/ERC1155Facet.abi.json";
import FreezeFacetAbi from "@/abi/FreezeFacet.abi.json";
import IdentityRegistryFacetAbi from "@/abi/IdentityRegistryFacet.abi.json";
import MetadataFacetAbi from "@/abi/MetadataFacet.abi.json";
import OwnershipFacetAbi from "@/abi/OwnershipFacet.abi.json";
import PauseFacetAbi from "@/abi/PauseFacet.abi.json";
import RecoveryFacetAbi from "@/abi/RecoveryFacet.abi.json";
import SnapshotFacetAbi from "@/abi/SnapshotFacet.abi.json";
import SupplyFacetAbi from "@/abi/SupplyFacet.abi.json";
import TrustedIssuerFacetAbi from "@/abi/TrustedIssuerFacet.abi.json";

export const diamondAbi = [
  ...AccessControlFacetAbi,
  ...AssetGroupFacetAbi,
  ...AssetManagerFacetAbi,
  ...ClaimTopicsFacetAbi,
  ...ComplianceRouterFacetAbi,
  ...DiamondCutFacetAbi,
  ...DiamondLoupeFacetAbi,
  ...DividendFacetAbi,
  ...EmergencyFacetAbi,
  ...ERC1155FacetAbi,
  ...FreezeFacetAbi,
  ...IdentityRegistryFacetAbi,
  ...MetadataFacetAbi,
  ...OwnershipFacetAbi,
  ...PauseFacetAbi,
  ...RecoveryFacetAbi,
  ...SnapshotFacetAbi,
  ...SupplyFacetAbi,
  ...TrustedIssuerFacetAbi,
] as unknown as Abi;

export const DIAMOND_ADDRESS = (import.meta.env.VITE_DIAMOND_ADDRESS ??
  "0x") as Address;

export const diamondContract = {
  address: DIAMOND_ADDRESS,
  abi: diamondAbi,
} as const;

import { ConnectButton } from "@rainbow-me/rainbowkit";
import {
  Shield,
  Wallet,
  Building2,
  Users,
  FileCheck,
  Camera,
  Lock,
  Layers,
} from "lucide-react";
import { Link } from "react-router-dom";
import { useReadContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { LogoFull } from "@/components/ui/logo";
import { useAssets } from "@/hooks/use-assets";

const features = [
  {
    icon: Building2,
    title: "Asset Tokenization",
    description:
      "Register and manage RWA asset classes with per-token compliance rules and supply controls.",
  },
  {
    icon: Users,
    title: "Identity & KYC",
    description:
      "On-chain identity registry with ONCHAINID integration for investor verification.",
  },
  {
    icon: FileCheck,
    title: "Compliance Engine",
    description:
      "Pluggable compliance modules per asset — country restrict, max balance, max holders.",
  },
  {
    icon: Layers,
    title: "Asset Groups",
    description:
      "Hierarchical tokenization: buildings to apartments with lazy minting support.",
  },
  {
    icon: Camera,
    title: "Snapshots & Dividends",
    description:
      "Point-in-time balance snapshots with pro-rata dividend distribution to holders.",
  },
  {
    icon: Lock,
    title: "Security Controls",
    description:
      "Pause, freeze, recovery, and role-based access control for institutional operations.",
  },
];

export default function HomePage() {
  const { data: facetAddrs } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "facetAddresses",
  });

  const { tokenIds, isLoading: isLoadingAssets } = useAssets();

  const facetCount = Array.isArray(facetAddrs) ? facetAddrs.length : null;
  const assetCount = Array.isArray(tokenIds) ? tokenIds.length : null;

  const truncatedAddress = DIAMOND_ADDRESS
    ? `${DIAMOND_ADDRESS.slice(0, 6)}...${DIAMOND_ADDRESS.slice(-4)}`
    : "";

  return (
    <main className="flex min-h-screen flex-col items-center px-6 py-16">
      <section className="flex flex-col items-center gap-8 text-center max-w-3xl mt-8">
        <LogoFull className="h-14" />

        <div className="flex flex-col gap-4">
          <h1 className="text-5xl font-bold tracking-tight sm:text-6xl">
            <span className="bg-gradient-to-r from-indigo-400 via-indigo-300 to-purple-400 bg-clip-text text-transparent">
              Diamond ERC-3643
            </span>
            <br />
            <span className="text-white">Protocol</span>
          </h1>
          <p className="text-xl text-gray-300">
            Institutional-grade security token infrastructure for Real World Assets
          </p>
          <p className="text-sm text-gray-500 font-mono tracking-wide">
            ERC-3643 compliance + EIP-2535 Diamond Proxy + ERC-1155 multi-token
          </p>
        </div>

        <div className="mt-2">
          <ConnectButton />
        </div>
      </section>

      <section className="mt-20 w-full max-w-4xl">
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div className="glass-card flex flex-col items-center gap-2 p-6 text-center">
            <span className="text-3xl font-bold text-indigo-400">
              {facetCount !== null ? facetCount : "—"}
            </span>
            <span className="text-sm text-gray-400">Deployed Facets</span>
          </div>
          <div className="glass-card flex flex-col items-center gap-2 p-6 text-center">
            <span className="text-3xl font-bold text-indigo-400">
              {isLoadingAssets ? "—" : (assetCount ?? "—")}
            </span>
            <span className="text-sm text-gray-400">Registered Assets</span>
          </div>
          <div className="glass-card flex flex-col items-center gap-2 p-6 text-center">
            <span className="text-3xl font-bold text-indigo-400">
              Polygon Amoy
            </span>
            <span className="text-sm text-gray-400">Network</span>
          </div>
        </div>
      </section>

      <section className="mt-20 w-full max-w-5xl">
        <h2 className="mb-8 text-center text-2xl font-semibold text-white">
          Protocol Features
        </h2>
        <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-3">
          {features.map(({ icon: Icon, title, description }) => (
            <div
              key={title}
              className="glass-card flex flex-col gap-4 p-6 transition-all"
            >
              <div className="flex h-11 w-11 items-center justify-center rounded-lg bg-indigo-500/10 text-indigo-400">
                <Icon className="h-5 w-5" />
              </div>
              <div>
                <h3 className="text-base font-semibold text-white">{title}</h3>
                <p className="mt-1.5 text-sm leading-relaxed text-gray-400">
                  {description}
                </p>
              </div>
            </div>
          ))}
        </div>
      </section>

      <section className="mt-20 w-full max-w-3xl">
        <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
          <Link
            to="/admin"
            className="glass-card group flex flex-col items-center gap-4 p-8 transition-all accent-glow"
          >
            <div className="flex h-14 w-14 items-center justify-center rounded-xl bg-indigo-500/10 text-indigo-400 transition-colors group-hover:bg-indigo-500/20">
              <Shield className="h-7 w-7" />
            </div>
            <div className="text-center">
              <h2 className="text-xl font-semibold text-white">Admin Panel</h2>
              <p className="mt-1 text-sm text-gray-400">
                Manage assets, compliance, and identity registries
              </p>
            </div>
          </Link>

          <Link
            to="/portfolio"
            className="glass-card group flex flex-col items-center gap-4 p-8 transition-all accent-glow"
          >
            <div className="flex h-14 w-14 items-center justify-center rounded-xl bg-indigo-500/10 text-indigo-400 transition-colors group-hover:bg-indigo-500/20">
              <Wallet className="h-7 w-7" />
            </div>
            <div className="text-center">
              <h2 className="text-xl font-semibold text-white">Portfolio</h2>
              <p className="mt-1 text-sm text-gray-400">
                View holdings, transfer tokens, and claim dividends
              </p>
            </div>
          </Link>
        </div>
      </section>

      <footer className="mt-24 mb-8 flex flex-col items-center gap-2 text-center">
        <p className="text-sm text-gray-500">
          Deployed on{" "}
          <span className="text-gray-400">Polygon Amoy</span>
        </p>
        {DIAMOND_ADDRESS && DIAMOND_ADDRESS !== "0x" && (
          <a
            href={`https://amoy.polygonscan.com/address/${DIAMOND_ADDRESS}`}
            target="_blank"
            rel="noopener noreferrer"
            className="font-mono text-xs text-indigo-400/70 transition-colors hover:text-indigo-400"
          >
            {truncatedAddress}
          </a>
        )}
      </footer>
    </main>
  );
}

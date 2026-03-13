import { useState, useMemo } from "react";
import { Link } from "react-router-dom";
import { ExternalLink } from "lucide-react";
import type { Address } from "viem";
import {
  useWriteContract,
  useWaitForTransactionReceipt,
  useSimulateContract,
  useAccount,
  useReadContract,
} from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { COMPLIANCE_MODULES } from "@/config/compliance-modules";
import { useAssets } from "@/hooks/use-assets";

const EXPLORER = "https://amoy.polygonscan.com";

const COUNTRIES: Record<number, { name: string; flag: string }> = {
  76:  { name: "Brasil",       flag: "BR" },
  840: { name: "USA",          flag: "US" },
  826: { name: "UK",           flag: "GB" },
  276: { name: "Germany",      flag: "DE" },
  250: { name: "France",       flag: "FR" },
  392: { name: "Japan",        flag: "JP" },
  156: { name: "China",        flag: "CN" },
  380: { name: "Italy",        flag: "IT" },
  724: { name: "Spain",        flag: "ES" },
  756: { name: "Switzerland",  flag: "CH" },
};

function countryPill(code: number) {
  const c = COUNTRIES[code];
  if (!c) return { label: `ISO ${code}`, flag: "" };
  const flag = c.flag
    .toUpperCase()
    .split("")
    .map((ch) => String.fromCodePoint(0x1f1e6 + ch.charCodeAt(0) - 65))
    .join("");
  return { label: c.name, flag };
}

function moduleInfo(addr: string) {
  const lower = addr.toLowerCase();
  const found = COMPLIANCE_MODULES.find((m) => m.address?.toLowerCase() === lower);
  return {
    name: found?.name?.replace("Module", "") ?? "Unknown",
    description: found?.description ?? "",
    address: addr,
  };
}

function truncAddr(addr: string): string {
  return `${addr.slice(0, 6)}...${addr.slice(-4)}`;
}

const OWNER_ROLE_LABEL = "Owner";
const ADMIN_ROLE_LABEL = "COMPLIANCE_ADMIN";

function useIsAuthorized() {
  const { address } = useAccount();

  const { data: owner } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "owner",
  });

  const { data: hasAdminRole } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "hasRole",
    args: address
      ? [
          "0x7aad6637a6d629b595e3b29e3c1877d06518132030cb4f1fff2ef5593c6bad6d" as `0x${string}`,
          address,
        ]
      : undefined,
    query: { enabled: !!address },
  });

  const isOwner = !!address && !!owner && (owner as Address).toLowerCase() === address.toLowerCase();
  const isAdmin = !!hasAdminRole;

  return {
    isAuthorized: isOwner || isAdmin,
    role: isOwner ? OWNER_ROLE_LABEL : isAdmin ? ADMIN_ROLE_LABEL : null,
    connectedAddress: address,
    ownerAddress: owner as Address | undefined,
  };
}

function parseContractError(error: Error): string {
  const msg = error.message;

  if (msg.includes("AssetManagerFacet__Unauthorized")) {
    return "Unauthorized: only the Diamond owner or COMPLIANCE_ADMIN can register assets.";
  }
  if (msg.includes("AssetManagerFacet__ZeroAddress")) {
    return "Issuer address cannot be zero.";
  }
  if (msg.includes("AssetManagerFacet__EmptyString")) {
    return "Name and symbol cannot be empty.";
  }
  if (msg.includes("AssetManagerFacet__TooManyModules")) {
    return "Too many compliance modules selected.";
  }
  if (msg.includes("User rejected") || msg.includes("denied")) {
    return "Transaction rejected by user.";
  }

  const match = msg.match(/reason:\s*(.+?)(?:\n|$)/);
  if (match) return match[1];

  return msg.slice(0, 300);
}

export default function AssetsPage() {
  const { assets, isLoading } = useAssets();
  const { isAuthorized, role, connectedAddress, ownerAddress } = useIsAuthorized();

  const [name, setName] = useState("");
  const [symbol, setSymbol] = useState("");
  const [uri, setUri] = useState("");
  const [supplyCap, setSupplyCap] = useState("");
  const [identityProfileId, setIdentityProfileId] = useState("1");
  const [selectedModules, setSelectedModules] = useState<Set<string>>(new Set());
  const [issuer, setIssuer] = useState("");
  const [allowedCountries, setAllowedCountries] = useState("");

  const deployedModules = COMPLIANCE_MODULES.filter((m) => m.address != null);

  const toggleModule = (address: string) => {
    setSelectedModules((prev) => {
      const next = new Set(prev);
      if (next.has(address)) {
        next.delete(address);
      } else {
        next.add(address);
      }
      return next;
    });
  };

  const formReady = !!name && !!symbol && !!issuer;

  const registerArgs = useMemo(() => {
    if (!formReady) return undefined;

    const countries = allowedCountries
      .split(",")
      .map((c) => c.trim())
      .filter(Boolean)
      .map((c) => Number(c));

    const modules = Array.from(selectedModules) as Address[];

    return [
      {
        name,
        symbol,
        uri,
        supplyCap: BigInt(supplyCap || "0"),
        identityProfileId: Number(identityProfileId || "1"),
        complianceModules: modules,
        issuer: issuer as Address,
        allowedCountries: countries,
      },
    ] as const;
  }, [name, symbol, uri, supplyCap, identityProfileId, selectedModules, issuer, allowedCountries, formReady]);

  const {
    data: simulateData,
    error: simulateError,
  } = useSimulateContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "registerAsset",
    args: registerArgs,
    query: { enabled: formReady && isAuthorized },
  });

  const { writeContract, data: hash, isPending, error: writeError } = useWriteContract();
  const { isSuccess } = useWaitForTransactionReceipt({ hash });

  const handleRegister = () => {
    if (simulateData?.request) {
      writeContract(simulateData.request);
    }
  };

  const displayError = writeError
    ? parseContractError(writeError)
    : simulateError
      ? parseContractError(simulateError)
      : null;

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Asset Management</h1>

      {/* Register New Asset */}
      <div className="mb-8 rounded-xl bg-white/5 border border-white/10 p-6">
        <div className="mb-4 flex items-center justify-between">
          <h2 className="text-xl font-semibold text-indigo-400">Register New Asset</h2>
          {connectedAddress && (
            <div className="flex items-center gap-2">
              {isAuthorized ? (
                <span className="rounded-full bg-green-500/10 px-3 py-1 text-xs font-medium text-green-400">
                  {role}
                </span>
              ) : (
                <span className="rounded-full bg-red-500/10 px-3 py-1 text-xs font-medium text-red-400">
                  No permission
                </span>
              )}
              <span className="font-mono text-xs text-gray-500">
                {connectedAddress.slice(0, 6)}...{connectedAddress.slice(-4)}
              </span>
            </div>
          )}
        </div>

        {!connectedAddress && (
          <p className="mb-4 rounded-lg bg-yellow-500/10 border border-yellow-500/20 p-3 text-sm text-yellow-400">
            Connect your wallet to register assets.
          </p>
        )}

        {connectedAddress && !isAuthorized && (
          <p className="mb-4 rounded-lg bg-red-500/10 border border-red-500/20 p-3 text-sm text-red-400">
            Your wallet does not have permission to register assets. Required: Diamond owner
            {ownerAddress && (
              <span className="font-mono text-xs">
                {" "}({ownerAddress.slice(0, 6)}...{ownerAddress.slice(-4)})
              </span>
            )} or COMPLIANCE_ADMIN role.
          </p>
        )}

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Name</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="Real Estate Fund I"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Symbol</label>
            <input
              type="text"
              value={symbol}
              onChange={(e) => setSymbol(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="REF1"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">URI</label>
            <input
              type="text"
              value={uri}
              onChange={(e) => setUri(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="https://metadata.example.com/{id}"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Supply Cap (0 = unlimited)</label>
            <input
              type="text"
              value={supplyCap}
              onChange={(e) => setSupplyCap(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="1000000"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Identity Profile ID</label>
            <input
              type="text"
              value={identityProfileId}
              onChange={(e) => setIdentityProfileId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="1"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Issuer Address</label>
            <input
              type="text"
              value={issuer}
              onChange={(e) => setIssuer(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Allowed Countries (comma-separated ISO codes)</label>
            <input
              type="text"
              value={allowedCountries}
              onChange={(e) => setAllowedCountries(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="76, 840 (76=Brasil, 840=EUA)"
            />
          </div>
        </div>

        {/* Compliance Modules */}
        <div className="mt-4">
          <label className="mb-2 block text-sm text-gray-400">Compliance Modules (optional)</label>
          {deployedModules.length === 0 ? (
            <p className="text-sm text-gray-500 italic">
              No compliance modules deployed yet. Assets will be registered without compliance restrictions.
            </p>
          ) : (
            <div className="grid grid-cols-1 gap-2 sm:grid-cols-2 lg:grid-cols-3">
              {deployedModules.map((mod) => {
                const isSelected = selectedModules.has(mod.address!);
                return (
                  <button
                    key={mod.address}
                    type="button"
                    onClick={() => toggleModule(mod.address!)}
                    className={`flex flex-col gap-1 rounded-lg border p-3 text-left transition-colors ${
                      isSelected
                        ? "border-indigo-400 bg-indigo-500/10"
                        : "border-white/10 bg-white/5 hover:border-white/20"
                    }`}
                  >
                    <span className={`text-sm font-medium ${isSelected ? "text-indigo-300" : "text-white"}`}>
                      {isSelected ? "\u2713 " : ""}{mod.name}
                    </span>
                    <span className="text-xs text-gray-500">{mod.description}</span>
                    <span className="font-mono text-[10px] text-gray-600">
                      {mod.address!.slice(0, 6)}...{mod.address!.slice(-4)}
                    </span>
                  </button>
                );
              })}
            </div>
          )}
          {selectedModules.size > 0 && (
            <p className="mt-2 text-xs text-gray-500">
              {selectedModules.size} module{selectedModules.size > 1 ? "s" : ""} selected
            </p>
          )}
        </div>

        <button
          onClick={handleRegister}
          disabled={isPending || !formReady || !isAuthorized || !simulateData?.request}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {isPending ? "Registering..." : "Register Asset"}
        </button>
        {isSuccess && <p className="mt-2 text-sm text-green-400">Asset registered successfully!</p>}
        {displayError && (
          <p className="mt-2 text-sm text-red-400">
            {displayError}
          </p>
        )}
      </div>

      {/* Asset List */}
      <div>
        <div className="mb-5 flex items-center justify-between">
          <h2 className="text-xl font-semibold text-indigo-400">Registered Assets</h2>
          {!isLoading && assets.length > 0 && (
            <div className="flex items-center gap-3">
              <span className="text-sm text-gray-500">{assets.length} asset{assets.length > 1 ? "s" : ""}</span>
              <a
                href={`${EXPLORER}/address/${DIAMOND_ADDRESS}`}
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center gap-1 rounded-md bg-white/5 px-2.5 py-1 text-xs text-gray-400 transition-colors hover:bg-white/10 hover:text-white"
              >
                Diamond {truncAddr(DIAMOND_ADDRESS)} <ExternalLink className="h-3 w-3" />
              </a>
            </div>
          )}
        </div>

        {isLoading ? (
          <p className="text-gray-500">Loading assets...</p>
        ) : assets.length === 0 ? (
          <p className="text-gray-500">No assets registered yet.</p>
        ) : (
          <div className="grid grid-cols-1 gap-5 lg:grid-cols-2">
            {assets.map((asset) => {
              const capDisplay = asset.supplyCap === 0n ? "Unlimited" : Number(asset.supplyCap).toLocaleString();
              const supplyDisplay = Number(asset.totalSupply).toLocaleString();
              const modules = asset.complianceModules.map((m: string) => moduleInfo(m));

              return (
                <div
                  key={asset.tokenId.toString()}
                  className="rounded-xl border border-white/10 bg-white/[0.03] overflow-hidden"
                >
                  {/* Card Header */}
                  <div className="flex items-center justify-between border-b border-white/5 bg-white/[0.02] px-5 py-3">
                    <div className="flex items-center gap-3">
                      <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-indigo-500/15 font-mono text-base font-bold text-indigo-400">
                        {asset.tokenId.toString()}
                      </div>
                      <div>
                        <h3 className="text-lg font-semibold leading-tight text-white">{asset.name}</h3>
                        <span className="font-mono text-sm text-gray-500">{asset.symbol}</span>
                      </div>
                    </div>
                    <span className={`rounded-full px-3 py-1 text-xs font-semibold ${
                      asset.paused
                        ? "bg-red-500/15 text-red-400 border border-red-500/20"
                        : "bg-emerald-500/15 text-emerald-400 border border-emerald-500/20"
                    }`}>
                      {asset.paused ? "Paused" : "Active"}
                    </span>
                  </div>

                  {/* Card Body */}
                  <div className="px-5 py-4 space-y-4">
                    {/* Supply Stats */}
                    <div className="grid grid-cols-3 gap-3">
                      <div className="rounded-lg bg-white/[0.04] p-3 text-center">
                        <p className="text-xs text-gray-500 mb-0.5">Supply Cap</p>
                        <p className="text-sm font-semibold text-white">{capDisplay}</p>
                      </div>
                      <div className="rounded-lg bg-white/[0.04] p-3 text-center">
                        <p className="text-xs text-gray-500 mb-0.5">Minted</p>
                        <p className="text-sm font-semibold text-white">{supplyDisplay}</p>
                      </div>
                      <div className="rounded-lg bg-white/[0.04] p-3 text-center">
                        <p className="text-xs text-gray-500 mb-0.5">Profile</p>
                        <p className="text-sm font-semibold text-white">#{asset.identityProfileId}</p>
                      </div>
                    </div>

                    {/* Issuer */}
                    <div>
                      <p className="text-xs font-medium text-gray-500 mb-1.5">Issuer</p>
                      <a
                        href={`${EXPLORER}/address/${asset.issuer}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="inline-flex items-center gap-1.5 rounded-md bg-white/5 px-2.5 py-1.5 font-mono text-xs text-gray-300 transition-colors hover:bg-white/10 hover:text-white"
                      >
                        {asset.issuer}
                        <ExternalLink className="h-3 w-3 shrink-0 text-gray-500" />
                      </a>
                    </div>

                    {/* Allowed Countries */}
                    <div>
                      <p className="text-xs font-medium text-gray-500 mb-1.5">Allowed Countries</p>
                      {asset.allowedCountries.length === 0 ? (
                        <span className="text-xs text-gray-600 italic">No restrictions</span>
                      ) : (
                        <div className="flex flex-wrap gap-1.5">
                          {asset.allowedCountries.map((code: number) => {
                            const { label, flag } = countryPill(code);
                            return (
                              <span
                                key={code}
                                className="inline-flex items-center gap-1 rounded-full bg-white/5 border border-white/10 px-2.5 py-0.5 text-xs text-gray-300"
                              >
                                {flag && <span>{flag}</span>}
                                {label}
                              </span>
                            );
                          })}
                        </div>
                      )}
                    </div>

                    {/* Compliance Modules */}
                    <div>
                      <p className="text-xs font-medium text-gray-500 mb-1.5">Compliance</p>
                      {modules.length === 0 ? (
                        <span className="text-xs text-gray-600 italic">No compliance modules</span>
                      ) : (
                        <div className="flex flex-wrap gap-1.5">
                          {modules.map((mod) => (
                            <a
                              key={mod.address}
                              href={`${EXPLORER}/address/${mod.address}#code`}
                              target="_blank"
                              rel="noopener noreferrer"
                              title={`${mod.description}\n${mod.address}`}
                              className="inline-flex items-center gap-1 rounded-md bg-indigo-500/10 border border-indigo-500/20 px-2.5 py-1 text-xs font-medium text-indigo-300 transition-colors hover:bg-indigo-500/20 hover:text-indigo-200"
                            >
                              {mod.name}
                              <ExternalLink className="h-2.5 w-2.5 shrink-0 opacity-50" />
                            </a>
                          ))}
                        </div>
                      )}
                    </div>

                    {/* URI */}
                    {asset.uri && (
                      <div>
                        <p className="text-xs font-medium text-gray-500 mb-1">Metadata URI</p>
                        <p className="font-mono text-[11px] text-gray-500 break-all leading-relaxed">{asset.uri}</p>
                      </div>
                    )}
                  </div>

                  {/* Card Footer */}
                  <div className="flex items-center border-t border-white/5 px-5 py-3">
                    <Link
                      to={`/admin/assets/${asset.tokenId.toString()}`}
                      className="rounded-md bg-indigo-500/10 px-4 py-1.5 text-sm font-medium text-indigo-400 transition-colors hover:bg-indigo-500/20 hover:text-indigo-300"
                    >
                      Manage Asset
                    </Link>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>
    </div>
  );
}

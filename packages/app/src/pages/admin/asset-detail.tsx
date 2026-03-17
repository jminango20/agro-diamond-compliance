import { useState } from "react";
import { useParams, Link } from "react-router-dom";
import { ExternalLink, ArrowLeft } from "lucide-react";
import type { Address } from "viem";
import {
  useReadContract,
  useWriteContract,
  useWaitForTransactionReceipt,
  useSimulateContract,
  useAccount,
} from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { COMPLIANCE_MODULES } from "@/config/compliance-modules";
import { useTokenInfo } from "@/hooks/use-token-info";
import { useIndexerToken, useIndexerProtocolEvents } from "@/hooks/use-indexer";
import { ActivityFeed } from "@/components/ui/activity-feed";
import { truncateAddress } from "@/lib/format";

const EXPLORER = "https://amoy.polygonscan.com";

const COUNTRIES: Record<number, { name: string; flag: string }> = {
  76:  { name: "Brasil", flag: "BR" },
  840: { name: "USA",    flag: "US" },
  826: { name: "UK",     flag: "GB" },
  276: { name: "Germany", flag: "DE" },
  250: { name: "France",  flag: "FR" },
  392: { name: "Japan",   flag: "JP" },
};

function countryPill(code: number) {
  const c = COUNTRIES[code];
  if (!c) return { label: `ISO ${code}`, flag: "" };
  const flag = c.flag.toUpperCase().split("").map((ch) => String.fromCodePoint(0x1f1e6 + ch.charCodeAt(0) - 65)).join("");
  return { label: c.name, flag };
}

function moduleName(addr: string): string {
  const lower = addr.toLowerCase();
  const found = COMPLIANCE_MODULES.find((m) => m.address?.toLowerCase() === lower);
  return found?.name?.replace("Module", "") ?? `${addr.slice(0, 8)}...`;
}

function parseError(error: Error): string {
  const msg = error.message;
  if (msg.includes("Unauthorized")) return "Unauthorized: you don't have permission.";
  if (msg.includes("InsufficientBalance")) return "Insufficient balance.";
  if (msg.includes("SupplyCapExceeded")) return "Would exceed supply cap.";
  if (msg.includes("User rejected") || msg.includes("denied")) return "Transaction rejected.";
  const match = msg.match(/reason:\s*(.+?)(?:\n|$)/);
  if (match) return match[1];
  return msg.slice(0, 200);
}

export default function AssetDetailPage() {
  const { tokenId: tokenIdParam } = useParams();
  const tokenId = BigInt(tokenIdParam as string);
  const { address: connectedAddress } = useAccount();

  const { name, symbol, uri, totalSupply, supplyCap, holderCount, issuer, paused, isLoading } =
    useTokenInfo(tokenId);

  const { data: complianceModules, isLoading: modulesLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "getComplianceModules",
    args: [tokenId],
  });

  const { data: allowedCountries } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "allowedCountries",
    args: [tokenId],
  });

  // ── Mint ──
  const [mintTo, setMintTo] = useState("");
  const [mintAmount, setMintAmount] = useState("");
  const mintReady = !!mintTo && !!mintAmount && Number(mintAmount) > 0;

  const { data: mintSim, error: mintSimError } = useSimulateContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "mint",
    args: mintReady ? [tokenId, mintTo as Address, BigInt(mintAmount)] : undefined,
    query: { enabled: mintReady },
  });
  const { writeContract: writeMint, data: mintHash, isPending: mintPending, error: mintWriteError } = useWriteContract();
  const { isSuccess: mintSuccess } = useWaitForTransactionReceipt({ hash: mintHash });
  const handleMint = () => { if (mintSim?.request) writeMint(mintSim.request); };
  const mintError = mintWriteError ?? mintSimError;

  // ── Burn ──
  const [burnFrom, setBurnFrom] = useState("");
  const [burnAmount, setBurnAmount] = useState("");
  const burnReady = !!burnFrom && !!burnAmount && Number(burnAmount) > 0;

  const { data: burnSim, error: burnSimError } = useSimulateContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "burn",
    args: burnReady ? [tokenId, burnFrom as Address, BigInt(burnAmount)] : undefined,
    query: { enabled: burnReady },
  });
  const { writeContract: writeBurn, data: burnHash, isPending: burnPending, error: burnWriteError } = useWriteContract();
  const { isSuccess: burnSuccess } = useWaitForTransactionReceipt({ hash: burnHash });
  const handleBurn = () => { if (burnSim?.request) writeBurn(burnSim.request); };
  const burnError = burnWriteError ?? burnSimError;

  // ── Add compliance module ──
  const [moduleAddress, setModuleAddress] = useState("");
  const { writeContract: writeAddModule, data: addModuleHash, isPending: addModulePending } = useWriteContract();
  const { isSuccess: addModuleSuccess } = useWaitForTransactionReceipt({ hash: addModuleHash });
  const handleAddModule = () => {
    writeAddModule({ address: DIAMOND_ADDRESS, abi: diamondAbi, functionName: "addComplianceModule", args: [tokenId, moduleAddress as Address] });
  };

  // ── Remove compliance module ──
  const [removeAddress, setRemoveAddress] = useState("");
  const { writeContract: writeRemoveModule, data: removeModuleHash, isPending: removeModulePending } = useWriteContract();
  const { isSuccess: removeModuleSuccess } = useWaitForTransactionReceipt({ hash: removeModuleHash });
  const handleRemoveModule = () => {
    writeRemoveModule({ address: DIAMOND_ADDRESS, abi: diamondAbi, functionName: "removeComplianceModule", args: [tokenId, removeAddress as Address] });
  };

  // ── Set supply cap ──
  const [newSupplyCap, setNewSupplyCap] = useState("");
  const { writeContract: writeSupplyCap, data: supplyCapHash, isPending: supplyCapPending } = useWriteContract();
  const { isSuccess: supplyCapSuccess } = useWaitForTransactionReceipt({ hash: supplyCapHash });
  const handleSetSupplyCap = () => {
    writeSupplyCap({ address: DIAMOND_ADDRESS, abi: diamondAbi, functionName: "setSupplyCap", args: [tokenId, BigInt(newSupplyCap || "0")] });
  };

  // ── Set URI ──
  const [newUri, setNewUri] = useState("");
  const { writeContract: writeUri, data: uriHash, isPending: uriPending } = useWriteContract();
  const { isSuccess: uriSuccess } = useWaitForTransactionReceipt({ hash: uriHash });
  const handleSetUri = () => {
    writeUri({ address: DIAMOND_ADDRESS, abi: diamondAbi, functionName: "setAssetUri", args: [tokenId, newUri] });
  };

  // ── Set allowed countries ──
  const [countriesInput, setCountriesInput] = useState("");
  const { writeContract: writeCountries, data: countriesHash, isPending: countriesPending } = useWriteContract();
  const { isSuccess: countriesSuccess } = useWaitForTransactionReceipt({ hash: countriesHash });
  const handleSetCountries = () => {
    const list = countriesInput.split(",").map((c) => c.trim()).filter(Boolean).map((c) => Number(c));
    writeCountries({ address: DIAMOND_ADDRESS, abi: diamondAbi, functionName: "setAllowedCountries", args: [tokenId, list] });
  };

  // Indexer data: holders + asset events
  const { data: indexerToken } = useIndexerToken(tokenIdParam);
  const { data: assetEvents, isLoading: eventsLoading } = useIndexerProtocolEvents({
    first: 15,
    tokenId: tokenIdParam,
  });

  const inputClass = "w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none";

  if (isLoading) {
    return (
      <div className="min-h-screen bg-[#0a0a0f] p-8">
        <p className="text-gray-500">Loading asset details...</p>
      </div>
    );
  }

  const capDisplay = supplyCap === 0n ? "Unlimited" : Number(supplyCap).toLocaleString();
  const supplyDisplay = Number(totalSupply).toLocaleString();
  const pct = supplyCap > 0n ? Number((totalSupply * 100n) / supplyCap) : 0;
  const modules = (complianceModules as Address[]) ?? [];
  const countries = (allowedCountries as number[]) ?? [];

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      {/* Breadcrumb */}
      <Link to="/admin/assets" className="mb-6 inline-flex items-center gap-1.5 text-sm text-gray-500 transition-colors hover:text-gray-300">
        <ArrowLeft className="h-4 w-4" /> Back to Assets
      </Link>

      {/* Header */}
      <div className="mb-8 flex items-start justify-between">
        <div>
          <div className="flex items-center gap-3">
            <div className="flex h-12 w-12 items-center justify-center rounded-xl bg-indigo-500/15 font-mono text-lg font-bold text-indigo-400">
              {tokenId.toString()}
            </div>
            <div>
              <h1 className="text-3xl font-bold text-white">{name}</h1>
              <span className="font-mono text-lg text-gray-500">{symbol}</span>
            </div>
          </div>
        </div>
        <span className={`rounded-full px-3 py-1 text-sm font-semibold ${
          paused ? "bg-red-500/15 text-red-400 border border-red-500/20" : "bg-emerald-500/15 text-emerald-400 border border-emerald-500/20"
        }`}>
          {paused ? "Paused" : "Active"}
        </span>
      </div>

      {/* Stats */}
      <div className="mb-6 grid grid-cols-2 gap-4 lg:grid-cols-4">
        <div className="rounded-xl bg-white/5 border border-white/10 p-5">
          <p className="text-xs text-gray-500">Total Supply</p>
          <p className="mt-1 text-2xl font-bold text-white">{supplyDisplay}</p>
          {supplyCap > 0n && (
            <div className="mt-2">
              <div className="h-1.5 rounded-full bg-white/10 overflow-hidden">
                <div className="h-full rounded-full bg-indigo-500" style={{ width: `${Math.min(pct, 100)}%` }} />
              </div>
              <p className="mt-1 text-[10px] text-gray-500">{pct}% of cap</p>
            </div>
          )}
        </div>
        <div className="rounded-xl bg-white/5 border border-white/10 p-5">
          <p className="text-xs text-gray-500">Supply Cap</p>
          <p className="mt-1 text-2xl font-bold text-white">{capDisplay}</p>
        </div>
        <div className="rounded-xl bg-white/5 border border-white/10 p-5">
          <p className="text-xs text-gray-500">Holders</p>
          <p className="mt-1 text-2xl font-bold text-white">{holderCount.toString()}</p>
        </div>
        <div className="rounded-xl bg-white/5 border border-white/10 p-5">
          <p className="text-xs text-gray-500">Issuer</p>
          <a
            href={`${EXPLORER}/address/${issuer}`}
            target="_blank"
            rel="noopener noreferrer"
            className="mt-1 flex items-center gap-1 font-mono text-sm text-indigo-400 transition-colors hover:text-indigo-300"
          >
            {issuer.slice(0, 6)}...{issuer.slice(-4)} <ExternalLink className="h-3 w-3" />
          </a>
        </div>
      </div>

      {/* Info row: Countries + Compliance + URI */}
      <div className="mb-6 grid grid-cols-1 gap-4 lg:grid-cols-3">
        <div className="rounded-xl bg-white/5 border border-white/10 p-5">
          <p className="text-xs font-medium text-gray-500 mb-2">Allowed Countries</p>
          {countries.length === 0 ? (
            <span className="text-sm text-gray-600 italic">No restrictions</span>
          ) : (
            <div className="flex flex-wrap gap-1.5">
              {countries.map((code: number) => {
                const { label, flag } = countryPill(code);
                return (
                  <span key={code} className="inline-flex items-center gap-1 rounded-full bg-white/5 border border-white/10 px-2.5 py-0.5 text-xs text-gray-300">
                    {flag && <span>{flag}</span>} {label}
                  </span>
                );
              })}
            </div>
          )}
        </div>
        <div className="rounded-xl bg-white/5 border border-white/10 p-5">
          <p className="text-xs font-medium text-gray-500 mb-2">Compliance Modules</p>
          {modulesLoading ? (
            <span className="text-sm text-gray-600">Loading...</span>
          ) : modules.length === 0 ? (
            <span className="text-sm text-gray-600 italic">No modules</span>
          ) : (
            <div className="flex flex-wrap gap-1.5">
              {modules.map((addr) => (
                <a
                  key={addr}
                  href={`${EXPLORER}/address/${addr}#code`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="inline-flex items-center gap-1 rounded-md bg-indigo-500/10 border border-indigo-500/20 px-2.5 py-1 text-xs font-medium text-indigo-300 hover:bg-indigo-500/20"
                >
                  {moduleName(addr)} <ExternalLink className="h-2.5 w-2.5 opacity-50" />
                </a>
              ))}
            </div>
          )}
        </div>
        <div className="rounded-xl bg-white/5 border border-white/10 p-5">
          <p className="text-xs font-medium text-gray-500 mb-2">Metadata URI</p>
          <p className="font-mono text-xs text-gray-400 break-all">{uri || "Not set"}</p>
        </div>
      </div>

      {/* ═══════════════ HOLDERS & EVENTS ═══════════════ */}
      <div className="mb-6 grid grid-cols-1 gap-4 lg:grid-cols-2">
        {/* Holders Table */}
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <h2 className="mb-4 text-lg font-semibold text-indigo-400">
            Holders
            {indexerToken && <span className="ml-2 text-sm text-gray-500">({indexerToken.holders.length})</span>}
          </h2>
          {indexerToken && indexerToken.holders.length > 0 ? (
            <div className="overflow-x-auto max-h-64 overflow-y-auto">
              <table className="w-full text-left text-sm text-gray-300">
                <thead className="sticky top-0 bg-[#0a0a0f]">
                  <tr className="border-b border-white/10 text-gray-400">
                    <th className="pb-2 pr-4">Address</th>
                    <th className="pb-2 text-right">Balance</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-white/10">
                  {indexerToken.holders.map((h) => (
                    <tr key={h.address} className="hover:bg-white/[0.03]">
                      <td className="py-2 pr-4">
                        <a
                          href={`${EXPLORER}/address/${h.address}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="font-mono text-xs text-indigo-400 hover:text-indigo-300"
                        >
                          {truncateAddress(h.address)}
                        </a>
                      </td>
                      <td className="py-2 text-right font-mono text-xs">
                        {Number(h.balance).toLocaleString()}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <p className="text-sm text-gray-500">No holders yet.</p>
          )}
        </div>

        {/* Asset Events */}
        <ActivityFeed
          events={assetEvents ?? []}
          isLoading={eventsLoading}
          title="Asset Activity"
        />
      </div>

      {/* ═══════════════ MINT & BURN ═══════════════ */}
      <div className="mb-6 grid grid-cols-1 gap-4 lg:grid-cols-2">
        {/* Mint */}
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <h2 className="mb-4 text-lg font-semibold text-emerald-400">Mint Tokens</h2>
          <div className="space-y-3">
            <div>
              <label className="mb-1 block text-sm text-gray-400">Recipient</label>
              <input type="text" value={mintTo} onChange={(e) => setMintTo(e.target.value)} className={inputClass} placeholder={connectedAddress ?? "0x..."} />
            </div>
            <div>
              <label className="mb-1 block text-sm text-gray-400">Amount</label>
              <input type="text" value={mintAmount} onChange={(e) => setMintAmount(e.target.value)} className={inputClass} placeholder="1000" />
            </div>
          </div>
          <button
            onClick={handleMint}
            disabled={mintPending || !mintReady || !mintSim?.request}
            className="mt-4 w-full rounded-lg bg-emerald-500 px-6 py-2.5 font-medium text-white transition-colors hover:bg-emerald-600 disabled:opacity-50"
          >
            {mintPending ? "Minting..." : `Mint ${mintAmount || "0"} ${symbol}`}
          </button>
          {mintSuccess && <p className="mt-2 text-sm text-green-400">Tokens minted successfully!</p>}
          {mintError && <p className="mt-2 text-sm text-red-400">{parseError(mintError)}</p>}
        </div>

        {/* Burn */}
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <h2 className="mb-4 text-lg font-semibold text-red-400">Burn Tokens</h2>
          <div className="space-y-3">
            <div>
              <label className="mb-1 block text-sm text-gray-400">From Address</label>
              <input type="text" value={burnFrom} onChange={(e) => setBurnFrom(e.target.value)} className={inputClass} placeholder="0x..." />
            </div>
            <div>
              <label className="mb-1 block text-sm text-gray-400">Amount</label>
              <input type="text" value={burnAmount} onChange={(e) => setBurnAmount(e.target.value)} className={inputClass} placeholder="500" />
            </div>
          </div>
          <button
            onClick={handleBurn}
            disabled={burnPending || !burnReady || !burnSim?.request}
            className="mt-4 w-full rounded-lg bg-red-500 px-6 py-2.5 font-medium text-white transition-colors hover:bg-red-600 disabled:opacity-50"
          >
            {burnPending ? "Burning..." : `Burn ${burnAmount || "0"} ${symbol}`}
          </button>
          {burnSuccess && <p className="mt-2 text-sm text-green-400">Tokens burned successfully!</p>}
          {burnError && <p className="mt-2 text-sm text-red-400">{parseError(burnError)}</p>}
        </div>
      </div>

      {/* ═══════════════ SETTINGS ═══════════════ */}
      <h2 className="mb-4 mt-8 text-xl font-semibold text-white">Asset Settings</h2>

      <div className="grid grid-cols-1 gap-4 lg:grid-cols-2">
        {/* Compliance Modules */}
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <h3 className="mb-3 text-sm font-semibold text-indigo-400">Manage Compliance Modules</h3>
          <div className="space-y-3">
            <div>
              <label className="mb-1 block text-xs text-gray-500">Add Module</label>
              <div className="flex gap-2">
                <input type="text" value={moduleAddress} onChange={(e) => setModuleAddress(e.target.value)} className={`${inputClass} text-sm`} placeholder="0x..." />
                <button onClick={handleAddModule} disabled={addModulePending || !moduleAddress} className="shrink-0 rounded-lg bg-indigo-500 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-600 disabled:opacity-50">
                  {addModulePending ? "..." : "Add"}
                </button>
              </div>
              {addModuleSuccess && <p className="mt-1 text-xs text-green-400">Module added!</p>}
            </div>
            <div>
              <label className="mb-1 block text-xs text-gray-500">Remove Module</label>
              <div className="flex gap-2">
                <input type="text" value={removeAddress} onChange={(e) => setRemoveAddress(e.target.value)} className={`${inputClass} text-sm`} placeholder="0x..." />
                <button onClick={handleRemoveModule} disabled={removeModulePending || !removeAddress} className="shrink-0 rounded-lg bg-red-500/80 px-4 py-2 text-sm font-medium text-white hover:bg-red-500 disabled:opacity-50">
                  {removeModulePending ? "..." : "Remove"}
                </button>
              </div>
              {removeModuleSuccess && <p className="mt-1 text-xs text-green-400">Module removed!</p>}
            </div>
          </div>
        </div>

        {/* Supply Cap */}
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <h3 className="mb-3 text-sm font-semibold text-indigo-400">Supply Cap</h3>
          <p className="mb-2 text-xs text-gray-500">Current: {capDisplay}</p>
          <div className="flex gap-2">
            <input type="text" value={newSupplyCap} onChange={(e) => setNewSupplyCap(e.target.value)} className={inputClass} placeholder="New cap (0 = unlimited)" />
            <button onClick={handleSetSupplyCap} disabled={supplyCapPending} className="shrink-0 rounded-lg bg-indigo-500 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-600 disabled:opacity-50">
              {supplyCapPending ? "..." : "Set"}
            </button>
          </div>
          {supplyCapSuccess && <p className="mt-1 text-xs text-green-400">Cap updated!</p>}
        </div>

        {/* URI */}
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <h3 className="mb-3 text-sm font-semibold text-indigo-400">Metadata URI</h3>
          <p className="mb-2 text-xs text-gray-500 break-all">Current: {uri || "Not set"}</p>
          <div className="flex gap-2">
            <input type="text" value={newUri} onChange={(e) => setNewUri(e.target.value)} className={inputClass} placeholder="https://metadata.example.com/{id}" />
            <button onClick={handleSetUri} disabled={uriPending} className="shrink-0 rounded-lg bg-indigo-500 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-600 disabled:opacity-50">
              {uriPending ? "..." : "Set"}
            </button>
          </div>
          {uriSuccess && <p className="mt-1 text-xs text-green-400">URI updated!</p>}
        </div>

        {/* Countries */}
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <h3 className="mb-3 text-sm font-semibold text-indigo-400">Allowed Countries</h3>
          <p className="mb-2 text-xs text-gray-500">Current: {countries.length === 0 ? "None" : countries.join(", ")}</p>
          <div className="flex gap-2">
            <input type="text" value={countriesInput} onChange={(e) => setCountriesInput(e.target.value)} className={inputClass} placeholder="76, 840 (comma-separated ISO)" />
            <button onClick={handleSetCountries} disabled={countriesPending} className="shrink-0 rounded-lg bg-indigo-500 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-600 disabled:opacity-50">
              {countriesPending ? "..." : "Set"}
            </button>
          </div>
          {countriesSuccess && <p className="mt-1 text-xs text-green-400">Countries updated!</p>}
        </div>
      </div>
    </div>
  );
}

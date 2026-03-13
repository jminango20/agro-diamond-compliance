import { useState, useMemo } from "react";
import type { Address } from "viem";
import {
  useWriteContract,
  useWaitForTransactionReceipt,
  useSimulateContract,
  useAccount,
  useReadContract,
} from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useAssets } from "@/hooks/use-assets";

function useIsAuthorized() {
  const { address } = useAccount();
  const { data: owner } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "owner",
  });
  const { data: hasRole } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "hasRole",
    args: address
      ? ["0x7aad6637a6d629b595e3b29e3c1877d06518132030cb4f1fff2ef5593c6bad6d" as `0x${string}`, address]
      : undefined,
    query: { enabled: !!address },
  });
  const isOwner = !!address && !!owner && (owner as Address).toLowerCase() === address.toLowerCase();
  return { isAuthorized: isOwner || !!hasRole, address };
}

function parseError(error: Error): string {
  const msg = error.message;
  if (msg.includes("Unauthorized")) return "Unauthorized: you don't have permission for this operation.";
  if (msg.includes("InsufficientBalance")) return "Insufficient balance for this operation.";
  if (msg.includes("SupplyCapExceeded")) return "Minting would exceed the supply cap.";
  if (msg.includes("User rejected") || msg.includes("denied")) return "Transaction rejected by user.";
  const match = msg.match(/reason:\s*(.+?)(?:\n|$)/);
  if (match) return match[1];
  return msg.slice(0, 300);
}

export default function SupplyPage() {
  const { assets, isLoading: assetsLoading } = useAssets();
  const { isAuthorized, address: connectedAddress } = useIsAuthorized();

  // ── Mint ──
  const [mintTokenId, setMintTokenId] = useState("");
  const [mintTo, setMintTo] = useState("");
  const [mintAmount, setMintAmount] = useState("");

  const mintReady = !!mintTokenId && !!mintTo && !!mintAmount && Number(mintAmount) > 0;
  const mintArgs = useMemo(() => {
    if (!mintReady) return undefined;
    return [BigInt(mintTokenId), mintTo as Address, BigInt(mintAmount)] as const;
  }, [mintTokenId, mintTo, mintAmount, mintReady]);

  const { data: mintSim, error: mintSimError } = useSimulateContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "mint",
    args: mintArgs,
    query: { enabled: mintReady && isAuthorized },
  });
  const { writeContract: writeMint, data: mintHash, isPending: mintPending, error: mintWriteError } = useWriteContract();
  const { isSuccess: mintSuccess } = useWaitForTransactionReceipt({ hash: mintHash });

  const handleMint = () => { if (mintSim?.request) writeMint(mintSim.request); };
  const mintError = mintWriteError ?? mintSimError;

  // ── Burn ──
  const [burnTokenId, setBurnTokenId] = useState("");
  const [burnFrom, setBurnFrom] = useState("");
  const [burnAmount, setBurnAmount] = useState("");

  const burnReady = !!burnTokenId && !!burnFrom && !!burnAmount && Number(burnAmount) > 0;
  const burnArgs = useMemo(() => {
    if (!burnReady) return undefined;
    return [BigInt(burnTokenId), burnFrom as Address, BigInt(burnAmount)] as const;
  }, [burnTokenId, burnFrom, burnAmount, burnReady]);

  const { data: burnSim, error: burnSimError } = useSimulateContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "burn",
    args: burnArgs,
    query: { enabled: burnReady && isAuthorized },
  });
  const { writeContract: writeBurn, data: burnHash, isPending: burnPending, error: burnWriteError } = useWriteContract();
  const { isSuccess: burnSuccess } = useWaitForTransactionReceipt({ hash: burnHash });

  const handleBurn = () => { if (burnSim?.request) writeBurn(burnSim.request); };
  const burnError = burnWriteError ?? burnSimError;

  // ── Forced Transfer ──
  const [forcedTokenId, setForcedTokenId] = useState("");
  const [forcedFrom, setForcedFrom] = useState("");
  const [forcedTo, setForcedTo] = useState("");
  const [forcedAmount, setForcedAmount] = useState("");
  const [forcedReason, setForcedReason] = useState("");

  const { writeContract: writeForced, data: forcedHash, isPending: forcedPending, error: forcedError } = useWriteContract();
  const { isSuccess: forcedSuccess } = useWaitForTransactionReceipt({ hash: forcedHash });

  const handleForcedTransfer = () => {
    writeForced({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "forcedTransfer",
      args: [
        BigInt(forcedTokenId || "0"),
        forcedFrom as Address,
        forcedTo as Address,
        BigInt(forcedAmount || "0"),
        forcedReason,
      ],
    });
  };

  const selectedMintAsset = assets.find((a) => a.tokenId.toString() === mintTokenId);
  const selectedBurnAsset = assets.find((a) => a.tokenId.toString() === burnTokenId);

  const inputClass = "w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none";
  const selectClass = "w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white focus:border-indigo-400 focus:outline-none appearance-none";

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Supply Management</h1>

      {!connectedAddress && (
        <p className="mb-6 rounded-lg bg-yellow-500/10 border border-yellow-500/20 p-3 text-sm text-yellow-400">
          Connect your wallet to manage token supply.
        </p>
      )}

      {connectedAddress && !isAuthorized && (
        <p className="mb-6 rounded-lg bg-red-500/10 border border-red-500/20 p-3 text-sm text-red-400">
          Your wallet does not have permission for supply operations.
        </p>
      )}

      {/* Supply Overview */}
      <div className="mb-8 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Token Supply Overview</h2>
        {assetsLoading ? (
          <p className="text-gray-500">Loading...</p>
        ) : assets.length === 0 ? (
          <p className="text-gray-500">No assets registered.</p>
        ) : (
          <div className="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-3">
            {assets.map((asset) => {
              const cap = asset.supplyCap === 0n ? "Unlimited" : Number(asset.supplyCap).toLocaleString();
              const supply = Number(asset.totalSupply).toLocaleString();
              const pct = asset.supplyCap > 0n
                ? Number((asset.totalSupply * 100n) / asset.supplyCap)
                : 0;

              return (
                <div key={asset.tokenId.toString()} className="rounded-lg border border-white/10 bg-white/[0.03] p-4">
                  <div className="flex items-center justify-between mb-2">
                    <div>
                      <span className="text-sm font-semibold text-white">{asset.name}</span>
                      <span className="ml-2 font-mono text-xs text-gray-500">{asset.symbol}</span>
                    </div>
                    <span className="font-mono text-xs text-gray-500">#{asset.tokenId.toString()}</span>
                  </div>
                  <div className="flex items-end justify-between">
                    <div>
                      <p className="text-xs text-gray-500">Minted</p>
                      <p className="text-lg font-bold text-white">{supply}</p>
                    </div>
                    <div className="text-right">
                      <p className="text-xs text-gray-500">Cap</p>
                      <p className="text-sm text-gray-400">{cap}</p>
                    </div>
                  </div>
                  {asset.supplyCap > 0n && (
                    <div className="mt-2 h-1.5 rounded-full bg-white/10 overflow-hidden">
                      <div
                        className="h-full rounded-full bg-indigo-500 transition-all"
                        style={{ width: `${Math.min(pct, 100)}%` }}
                      />
                    </div>
                  )}
                </div>
              );
            })}
          </div>
        )}
      </div>

      {/* Mint */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-emerald-400">Mint Tokens</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Asset</label>
            <select
              value={mintTokenId}
              onChange={(e) => setMintTokenId(e.target.value)}
              className={selectClass}
            >
              <option value="">Select asset...</option>
              {assets.map((a) => (
                <option key={a.tokenId.toString()} value={a.tokenId.toString()}>
                  #{a.tokenId.toString()} — {a.name} ({a.symbol})
                </option>
              ))}
            </select>
            {selectedMintAsset && (
              <p className="mt-1 text-xs text-gray-500">
                Supply: {Number(selectedMintAsset.totalSupply).toLocaleString()} / {selectedMintAsset.supplyCap === 0n ? "Unlimited" : Number(selectedMintAsset.supplyCap).toLocaleString()}
              </p>
            )}
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Recipient</label>
            <input
              type="text"
              value={mintTo}
              onChange={(e) => setMintTo(e.target.value)}
              className={inputClass}
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Amount</label>
            <input
              type="text"
              value={mintAmount}
              onChange={(e) => setMintAmount(e.target.value)}
              className={inputClass}
              placeholder="1000"
            />
          </div>
        </div>
        <button
          onClick={handleMint}
          disabled={mintPending || !mintReady || !isAuthorized || !mintSim?.request}
          className="mt-4 rounded-lg bg-emerald-500 px-6 py-2 font-medium text-white transition-colors hover:bg-emerald-600 disabled:opacity-50"
        >
          {mintPending ? "Minting..." : "Mint"}
        </button>
        {mintSuccess && <p className="mt-2 text-sm text-green-400">Tokens minted successfully!</p>}
        {mintError && <p className="mt-2 text-sm text-red-400">{parseError(mintError)}</p>}
      </div>

      {/* Burn */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-red-400">Burn Tokens</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Asset</label>
            <select
              value={burnTokenId}
              onChange={(e) => setBurnTokenId(e.target.value)}
              className={selectClass}
            >
              <option value="">Select asset...</option>
              {assets.map((a) => (
                <option key={a.tokenId.toString()} value={a.tokenId.toString()}>
                  #{a.tokenId.toString()} — {a.name} ({a.symbol})
                </option>
              ))}
            </select>
            {selectedBurnAsset && (
              <p className="mt-1 text-xs text-gray-500">
                Current supply: {Number(selectedBurnAsset.totalSupply).toLocaleString()}
              </p>
            )}
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">From Address</label>
            <input
              type="text"
              value={burnFrom}
              onChange={(e) => setBurnFrom(e.target.value)}
              className={inputClass}
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Amount</label>
            <input
              type="text"
              value={burnAmount}
              onChange={(e) => setBurnAmount(e.target.value)}
              className={inputClass}
              placeholder="500"
            />
          </div>
        </div>
        <button
          onClick={handleBurn}
          disabled={burnPending || !burnReady || !isAuthorized || !burnSim?.request}
          className="mt-4 rounded-lg bg-red-500 px-6 py-2 font-medium text-white transition-colors hover:bg-red-600 disabled:opacity-50"
        >
          {burnPending ? "Burning..." : "Burn"}
        </button>
        {burnSuccess && <p className="mt-2 text-sm text-green-400">Tokens burned successfully!</p>}
        {burnError && <p className="mt-2 text-sm text-red-400">{parseError(burnError)}</p>}
      </div>

      {/* Forced Transfer */}
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-orange-400">Forced Transfer</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Asset</label>
            <select
              value={forcedTokenId}
              onChange={(e) => setForcedTokenId(e.target.value)}
              className={selectClass}
            >
              <option value="">Select asset...</option>
              {assets.map((a) => (
                <option key={a.tokenId.toString()} value={a.tokenId.toString()}>
                  #{a.tokenId.toString()} — {a.name} ({a.symbol})
                </option>
              ))}
            </select>
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">From Address</label>
            <input
              type="text"
              value={forcedFrom}
              onChange={(e) => setForcedFrom(e.target.value)}
              className={inputClass}
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">To Address</label>
            <input
              type="text"
              value={forcedTo}
              onChange={(e) => setForcedTo(e.target.value)}
              className={inputClass}
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Amount</label>
            <input
              type="text"
              value={forcedAmount}
              onChange={(e) => setForcedAmount(e.target.value)}
              className={inputClass}
              placeholder="100"
            />
          </div>
        </div>
        <div className="mt-4">
          <label className="mb-1 block text-sm text-gray-400">Reason</label>
          <input
            type="text"
            value={forcedReason}
            onChange={(e) => setForcedReason(e.target.value)}
            className={inputClass}
            placeholder="Court order #12345"
          />
        </div>
        <button
          onClick={handleForcedTransfer}
          disabled={forcedPending || !isAuthorized}
          className="mt-4 rounded-lg bg-orange-500 px-6 py-2 font-medium text-white transition-colors hover:bg-orange-600 disabled:opacity-50"
        >
          {forcedPending ? "Transferring..." : "Force Transfer"}
        </button>
        {forcedSuccess && <p className="mt-2 text-sm text-green-400">Forced transfer completed!</p>}
        {forcedError && <p className="mt-2 text-sm text-red-400">{parseError(forcedError)}</p>}
      </div>
    </div>
  );
}

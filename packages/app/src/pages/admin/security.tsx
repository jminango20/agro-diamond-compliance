import { useState } from "react";
import { keccak256, toHex } from "viem";
import type { Address } from "viem";
import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useAssets } from "@/hooks/use-assets";

const ROLES = [
  { label: "Issuer", value: keccak256(toHex("ISSUER_ROLE")) },
  { label: "Transfer Agent", value: keccak256(toHex("TRANSFER_AGENT")) },
  { label: "Recovery Agent", value: keccak256(toHex("RECOVERY_AGENT")) },
  { label: "Pauser", value: keccak256(toHex("PAUSER_ROLE")) },
] as const;

export default function SecurityPage() {
  const { assets, isLoading: assetsLoading } = useAssets();

  const { data: isProtocolPaused } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "isProtocolPaused",
  });

  // Protocol pause/unpause
  const { writeContract: writePause, data: pauseHash, isPending: pausePending } = useWriteContract();
  const { isSuccess: pauseSuccess } = useWaitForTransactionReceipt({ hash: pauseHash });

  const handleProtocolPause = () => {
    writePause({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "pauseProtocol",
    });
  };

  const handleProtocolUnpause = () => {
    writePause({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "unpauseProtocol",
    });
  };

  // Per-asset pause
  const { writeContract: writeAssetPause, data: assetPauseHash, isPending: assetPausePending } = useWriteContract();
  const { isSuccess: assetPauseSuccess } = useWaitForTransactionReceipt({ hash: assetPauseHash });
  const [pauseTokenId, setPauseTokenId] = useState("");

  const handlePauseAsset = () => {
    writeAssetPause({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "pauseAsset",
      args: [BigInt(pauseTokenId || "0")],
    });
  };

  const handleUnpauseAsset = () => {
    writeAssetPause({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "unpauseAsset",
      args: [BigInt(pauseTokenId || "0")],
    });
  };

  // Emergency pause
  const { writeContract: writeEmergency, data: emergencyHash, isPending: emergencyPending } = useWriteContract();
  const { isSuccess: emergencySuccess } = useWaitForTransactionReceipt({ hash: emergencyHash });

  const handleEmergencyPause = () => {
    writeEmergency({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "emergencyPause",
    });
  };

  // Freeze wallet
  const { writeContract: writeFreeze, data: freezeHash, isPending: freezePending } = useWriteContract();
  const { isSuccess: freezeSuccess } = useWaitForTransactionReceipt({ hash: freezeHash });
  const [freezeWallet, setFreezeWallet] = useState("");
  const [freezeValue, setFreezeValue] = useState(true);

  const handleSetWalletFrozen = () => {
    writeFreeze({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "setWalletFrozen",
      args: [freezeWallet as Address, freezeValue],
    });
  };

  // Per-asset wallet freeze
  const { writeContract: writeAssetFreeze, data: assetFreezeHash, isPending: assetFreezePending } = useWriteContract();
  const { isSuccess: assetFreezeSuccess } = useWaitForTransactionReceipt({ hash: assetFreezeHash });
  const [assetFreezeTokenId, setAssetFreezeTokenId] = useState("");
  const [assetFreezeWallet, setAssetFreezeWallet] = useState("");
  const [assetFreezeValue, setAssetFreezeValue] = useState(true);

  const handleSetAssetWalletFrozen = () => {
    writeAssetFreeze({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "setAssetWalletFrozen",
      args: [BigInt(assetFreezeTokenId || "0"), assetFreezeWallet as Address, assetFreezeValue],
    });
  };

  // Set frozen amount
  const { writeContract: writeFrozenAmount, data: frozenAmountHash, isPending: frozenAmountPending } = useWriteContract();
  const { isSuccess: frozenAmountSuccess } = useWaitForTransactionReceipt({ hash: frozenAmountHash });
  const [frozenTokenId, setFrozenTokenId] = useState("");
  const [frozenWallet, setFrozenWallet] = useState("");
  const [frozenAmount, setFrozenAmount] = useState("");

  const handleSetFrozenAmount = () => {
    writeFrozenAmount({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "setFrozenTokens",
      args: [BigInt(frozenTokenId || "0"), frozenWallet as Address, BigInt(frozenAmount || "0")],
    });
  };

  // Role management
  const { writeContract: writeRole, data: roleHash, isPending: rolePending } = useWriteContract();
  const { isSuccess: roleSuccess } = useWaitForTransactionReceipt({ hash: roleHash });
  const [roleAddress, setRoleAddress] = useState("");
  const [selectedRole, setSelectedRole] = useState(ROLES[0].value);

  const handleGrantRole = () => {
    writeRole({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "grantRole",
      args: [selectedRole, roleAddress as Address],
    });
  };

  const handleRevokeRole = () => {
    writeRole({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "revokeRole",
      args: [selectedRole, roleAddress as Address],
    });
  };

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Security Management</h1>

      {/* Protocol Pause */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Protocol Pause</h2>
        <div className="flex items-center gap-4">
          <p className="text-gray-300">
            Current status:{" "}
            <span className={isProtocolPaused ? "font-semibold text-red-400" : "font-semibold text-green-400"}>
              {isProtocolPaused ? "Paused" : "Active"}
            </span>
          </p>
          <button
            onClick={isProtocolPaused ? handleProtocolUnpause : handleProtocolPause}
            disabled={pausePending}
            className={`rounded-lg px-6 py-2 font-medium text-white transition-colors disabled:opacity-50 ${
              isProtocolPaused
                ? "bg-green-500 hover:bg-green-600"
                : "bg-red-500 hover:bg-red-600"
            }`}
          >
            {pausePending ? "Processing..." : isProtocolPaused ? "Unpause Protocol" : "Pause Protocol"}
          </button>
        </div>
        {pauseSuccess && <p className="mt-2 text-sm text-green-400">Protocol pause state updated!</p>}
      </div>

      {/* Emergency Pause */}
      <div className="mb-6 rounded-xl bg-white/5 border border-red-500/30 p-6">
        <h2 className="mb-4 text-xl font-semibold text-red-400">Emergency Pause</h2>
        <p className="mb-4 text-sm text-gray-400">
          This will immediately pause all protocol operations. Use only in emergencies.
        </p>
        <button
          onClick={handleEmergencyPause}
          disabled={emergencyPending}
          className="rounded-lg bg-red-600 px-6 py-2 font-bold text-white transition-colors hover:bg-red-700 disabled:opacity-50"
        >
          {emergencyPending ? "Pausing..." : "EMERGENCY PAUSE"}
        </button>
        {emergencySuccess && <p className="mt-2 text-sm text-green-400">Emergency pause activated!</p>}
      </div>

      {/* Per-Asset Pause */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Per-Asset Pause</h2>
        <div className="flex items-end gap-2">
          <div className="flex-1">
            <label className="mb-1 block text-sm text-gray-400">Token ID</label>
            <input
              type="text"
              value={pauseTokenId}
              onChange={(e) => setPauseTokenId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0"
            />
          </div>
          <button
            onClick={handlePauseAsset}
            disabled={assetPausePending}
            className="rounded-lg bg-red-500 px-4 py-2 font-medium text-white transition-colors hover:bg-red-600 disabled:opacity-50"
          >
            Pause
          </button>
          <button
            onClick={handleUnpauseAsset}
            disabled={assetPausePending}
            className="rounded-lg bg-green-500 px-4 py-2 font-medium text-white transition-colors hover:bg-green-600 disabled:opacity-50"
          >
            Unpause
          </button>
        </div>
        {assetPauseSuccess && <p className="mt-2 text-sm text-green-400">Asset pause state updated!</p>}

        {!assetsLoading && assets.length > 0 && (
          <div className="mt-4 overflow-x-auto">
            <table className="w-full text-left text-sm text-gray-300">
              <thead>
                <tr className="border-b border-white/10 text-gray-400">
                  <th className="pb-3 pr-4">Token ID</th>
                  <th className="pb-3 pr-4">Name</th>
                  <th className="pb-3">Status</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-white/10">
                {assets.map((a) => (
                  <tr key={a.tokenId.toString()}>
                    <td className="py-3 pr-4 font-mono">{a.tokenId.toString()}</td>
                    <td className="py-3 pr-4">{a.name}</td>
                    <td className="py-3">
                      <span className={a.paused ? "text-red-400" : "text-green-400"}>
                        {a.paused ? "Paused" : "Active"}
                      </span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>

      {/* Freeze Wallet (Global) */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Freeze Wallet (Global)</h2>
        <div className="flex items-end gap-2">
          <div className="flex-1">
            <label className="mb-1 block text-sm text-gray-400">Wallet Address</label>
            <input
              type="text"
              value={freezeWallet}
              onChange={(e) => setFreezeWallet(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <select
            value={freezeValue ? "freeze" : "unfreeze"}
            onChange={(e) => setFreezeValue(e.target.value === "freeze")}
            className="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white focus:border-indigo-400 focus:outline-none"
          >
            <option value="freeze">Freeze</option>
            <option value="unfreeze">Unfreeze</option>
          </select>
          <button
            onClick={handleSetWalletFrozen}
            disabled={freezePending}
            className="rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
          >
            {freezePending ? "Processing..." : "Submit"}
          </button>
        </div>
        {freezeSuccess && <p className="mt-2 text-sm text-green-400">Wallet freeze state updated!</p>}
      </div>

      {/* Freeze Wallet Per-Asset */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Freeze Wallet (Per-Asset)</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Token ID</label>
            <input
              type="text"
              value={assetFreezeTokenId}
              onChange={(e) => setAssetFreezeTokenId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Wallet Address</label>
            <input
              type="text"
              value={assetFreezeWallet}
              onChange={(e) => setAssetFreezeWallet(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Action</label>
            <select
              value={assetFreezeValue ? "freeze" : "unfreeze"}
              onChange={(e) => setAssetFreezeValue(e.target.value === "freeze")}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white focus:border-indigo-400 focus:outline-none"
            >
              <option value="freeze">Freeze</option>
              <option value="unfreeze">Unfreeze</option>
            </select>
          </div>
        </div>
        <button
          onClick={handleSetAssetWalletFrozen}
          disabled={assetFreezePending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {assetFreezePending ? "Processing..." : "Submit"}
        </button>
        {assetFreezeSuccess && <p className="mt-2 text-sm text-green-400">Asset wallet freeze state updated!</p>}
      </div>

      {/* Set Frozen Amount */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Set Frozen Token Amount</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Token ID</label>
            <input
              type="text"
              value={frozenTokenId}
              onChange={(e) => setFrozenTokenId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Wallet Address</label>
            <input
              type="text"
              value={frozenWallet}
              onChange={(e) => setFrozenWallet(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Frozen Amount</label>
            <input
              type="text"
              value={frozenAmount}
              onChange={(e) => setFrozenAmount(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="500"
            />
          </div>
        </div>
        <button
          onClick={handleSetFrozenAmount}
          disabled={frozenAmountPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {frozenAmountPending ? "Setting..." : "Set Frozen Amount"}
        </button>
        {frozenAmountSuccess && <p className="mt-2 text-sm text-green-400">Frozen amount updated!</p>}
      </div>

      {/* Role Management */}
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Role Management</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Wallet Address</label>
            <input
              type="text"
              value={roleAddress}
              onChange={(e) => setRoleAddress(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Role</label>
            <select
              value={selectedRole}
              onChange={(e) => setSelectedRole(e.target.value as `0x${string}`)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white focus:border-indigo-400 focus:outline-none"
            >
              {ROLES.map((r) => (
                <option key={r.value} value={r.value}>
                  {r.label}
                </option>
              ))}
            </select>
          </div>
        </div>
        <div className="mt-4 flex gap-2">
          <button
            onClick={handleGrantRole}
            disabled={rolePending}
            className="rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
          >
            {rolePending ? "Processing..." : "Grant Role"}
          </button>
          <button
            onClick={handleRevokeRole}
            disabled={rolePending}
            className="rounded-lg bg-red-500 px-6 py-2 font-medium text-white transition-colors hover:bg-red-600 disabled:opacity-50"
          >
            {rolePending ? "Processing..." : "Revoke Role"}
          </button>
        </div>
        {roleSuccess && <p className="mt-2 text-sm text-green-400">Role updated!</p>}

        <div className="mt-6">
          <h3 className="mb-2 text-sm font-medium text-gray-400">Role Reference</h3>
          <div className="overflow-x-auto">
            <table className="w-full text-left text-sm text-gray-300">
              <thead>
                <tr className="border-b border-white/10 text-gray-400">
                  <th className="pb-2 pr-4">Role</th>
                  <th className="pb-2">Bytes32 Hash</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-white/10">
                {ROLES.map((r) => (
                  <tr key={r.value}>
                    <td className="py-2 pr-4">{r.label}</td>
                    <td className="py-2 font-mono text-xs text-gray-500">{r.value}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}

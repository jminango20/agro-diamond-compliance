import { useState } from "react";
import type { Address } from "viem";
import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useAssets } from "@/hooks/use-assets";

export default function CompliancePage() {
  const { assets, isLoading: assetsLoading } = useAssets();

  // Create claim topic profile
  const { writeContract: writeCreateProfile, data: createProfileHash, isPending: createProfilePending } = useWriteContract();
  const { isSuccess: createProfileSuccess } = useWaitForTransactionReceipt({ hash: createProfileHash });
  const [profileName, setProfileName] = useState("");

  const handleCreateProfile = () => {
    writeCreateProfile({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "createProfile",
      args: [profileName],
    });
  };

  // Set profile claim topics
  const { writeContract: writeSetTopics, data: setTopicsHash, isPending: setTopicsPending } = useWriteContract();
  const { isSuccess: setTopicsSuccess } = useWaitForTransactionReceipt({ hash: setTopicsHash });
  const [profileId, setProfileId] = useState("");
  const [claimTopics, setClaimTopics] = useState("");

  const handleSetClaimTopics = () => {
    const topics = claimTopics
      .split(",")
      .map((t) => t.trim())
      .filter(Boolean)
      .map((t) => BigInt(t));

    writeSetTopics({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "setProfileClaimTopics",
      args: [BigInt(profileId || "0"), topics],
    });
  };

  // Add trusted issuer
  const { writeContract: writeAddIssuer, data: addIssuerHash, isPending: addIssuerPending } = useWriteContract();
  const { isSuccess: addIssuerSuccess } = useWaitForTransactionReceipt({ hash: addIssuerHash });
  const [issuerAddress, setIssuerAddress] = useState("");
  const [issuerClaimTopics, setIssuerClaimTopics] = useState("");

  const handleAddIssuer = () => {
    const topics = issuerClaimTopics
      .split(",")
      .map((t) => t.trim())
      .filter(Boolean)
      .map((t) => BigInt(t));

    writeAddIssuer({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "addTrustedIssuer",
      args: [issuerAddress as Address, topics],
    });
  };

  // Remove trusted issuer
  const { writeContract: writeRemoveIssuer, data: removeIssuerHash, isPending: removeIssuerPending } = useWriteContract();
  const { isSuccess: removeIssuerSuccess } = useWaitForTransactionReceipt({ hash: removeIssuerHash });
  const [removeIssuerAddr, setRemoveIssuerAddr] = useState("");

  const handleRemoveIssuer = () => {
    writeRemoveIssuer({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "removeTrustedIssuer",
      args: [removeIssuerAddr as Address],
    });
  };

  // View modules per asset
  const [viewTokenId, setViewTokenId] = useState("");
  const { data: modules, isLoading: modulesLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "getComplianceModules",
    args: viewTokenId ? [BigInt(viewTokenId)] : undefined,
    query: { enabled: !!viewTokenId },
  });

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Compliance Management</h1>

      {/* Create Claim Topic Profile */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Create Claim Topic Profile</h2>
        <div className="flex gap-2">
          <input
            type="text"
            value={profileName}
            onChange={(e) => setProfileName(e.target.value)}
            className="flex-1 rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
            placeholder="Profile name (e.g., KYC_AML)"
          />
          <button
            onClick={handleCreateProfile}
            disabled={createProfilePending}
            className="rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
          >
            {createProfilePending ? "Creating..." : "Create Profile"}
          </button>
        </div>
        {createProfileSuccess && <p className="mt-2 text-sm text-green-400">Profile created!</p>}
      </div>

      {/* Set Profile Claim Topics */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Set Profile Claim Topics</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Profile ID</label>
            <input
              type="text"
              value={profileId}
              onChange={(e) => setProfileId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="1"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Claim Topics (comma-separated)</label>
            <input
              type="text"
              value={claimTopics}
              onChange={(e) => setClaimTopics(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="1, 2, 3"
            />
          </div>
        </div>
        <button
          onClick={handleSetClaimTopics}
          disabled={setTopicsPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {setTopicsPending ? "Setting..." : "Set Topics"}
        </button>
        {setTopicsSuccess && <p className="mt-2 text-sm text-green-400">Claim topics set!</p>}
      </div>

      {/* Add Trusted Issuer */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Add Trusted Issuer</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Issuer Address</label>
            <input
              type="text"
              value={issuerAddress}
              onChange={(e) => setIssuerAddress(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Claim Topics (comma-separated)</label>
            <input
              type="text"
              value={issuerClaimTopics}
              onChange={(e) => setIssuerClaimTopics(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="1, 2"
            />
          </div>
        </div>
        <button
          onClick={handleAddIssuer}
          disabled={addIssuerPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {addIssuerPending ? "Adding..." : "Add Issuer"}
        </button>
        {addIssuerSuccess && <p className="mt-2 text-sm text-green-400">Trusted issuer added!</p>}
      </div>

      {/* Remove Trusted Issuer */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Remove Trusted Issuer</h2>
        <div className="flex gap-2">
          <input
            type="text"
            value={removeIssuerAddr}
            onChange={(e) => setRemoveIssuerAddr(e.target.value)}
            className="flex-1 rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
            placeholder="0x..."
          />
          <button
            onClick={handleRemoveIssuer}
            disabled={removeIssuerPending}
            className="rounded-lg bg-red-500 px-6 py-2 font-medium text-white transition-colors hover:bg-red-600 disabled:opacity-50"
          >
            {removeIssuerPending ? "Removing..." : "Remove Issuer"}
          </button>
        </div>
        {removeIssuerSuccess && <p className="mt-2 text-sm text-green-400">Issuer removed!</p>}
      </div>

      {/* View Active Modules Per Asset */}
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Active Modules Per Asset</h2>
        <div className="mb-4 flex gap-2">
          <select
            value={viewTokenId}
            onChange={(e) => setViewTokenId(e.target.value)}
            className="flex-1 rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white focus:border-indigo-400 focus:outline-none"
          >
            <option value="">Select an asset...</option>
            {!assetsLoading &&
              assets.map((a) => (
                <option key={a.tokenId.toString()} value={a.tokenId.toString()}>
                  {a.name} ({a.symbol}) - ID: {a.tokenId.toString()}
                </option>
              ))}
          </select>
        </div>

        {viewTokenId && (
          <div>
            {modulesLoading ? (
              <p className="text-gray-500">Loading modules...</p>
            ) : !modules || (modules as Address[]).length === 0 ? (
              <p className="text-gray-500">No compliance modules for this asset.</p>
            ) : (
              <ul className="space-y-2">
                {(modules as Address[]).map((addr) => (
                  <li key={addr} className="font-mono text-sm text-gray-300">
                    {addr}
                  </li>
                ))}
              </ul>
            )}
          </div>
        )}
      </div>
    </div>
  );
}

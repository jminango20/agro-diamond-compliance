import { useState } from "react";
import type { Address } from "viem";
import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";

export default function IdentityPage() {
  // Register identity
  const { writeContract: writeRegister, data: registerHash, isPending: registerPending } = useWriteContract();
  const { isSuccess: registerSuccess } = useWaitForTransactionReceipt({ hash: registerHash });

  const [wallet, setWallet] = useState("");
  const [identity, setIdentity] = useState("");
  const [country, setCountry] = useState("");

  const handleRegister = () => {
    writeRegister({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "registerIdentity",
      args: [wallet as Address, identity as Address, Number(country)],
    });
  };

  // Batch register
  const { writeContract: writeBatch, data: batchHash, isPending: batchPending } = useWriteContract();
  const { isSuccess: batchSuccess } = useWaitForTransactionReceipt({ hash: batchHash });

  const [batchWallets, setBatchWallets] = useState("");
  const [batchIdentities, setBatchIdentities] = useState("");
  const [batchCountries, setBatchCountries] = useState("");

  const handleBatchRegister = () => {
    const wallets = batchWallets.split(",").map((w) => w.trim()) as Address[];
    const identities = batchIdentities.split(",").map((i) => i.trim()) as Address[];
    const countries = batchCountries.split(",").map((c) => Number(c.trim()));

    writeBatch({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "batchRegisterIdentity",
      args: [wallets, identities, countries],
    });
  };

  // Search
  const [searchAddress, setSearchAddress] = useState("");
  const [searchTriggered, setSearchTriggered] = useState(false);
  const searchAddr = searchTriggered ? (searchAddress as Address) : undefined;

  const { data: containsResult } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "contains",
    args: searchAddr ? [searchAddr] : undefined,
    query: { enabled: !!searchAddr },
  });

  const { data: identityResult } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "getIdentity",
    args: searchAddr ? [searchAddr] : undefined,
    query: { enabled: !!searchAddr && containsResult === true },
  });

  const { data: countryResult } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "getCountry",
    args: searchAddr ? [searchAddr] : undefined,
    query: { enabled: !!searchAddr && containsResult === true },
  });

  const handleSearch = () => {
    setSearchTriggered(false);
    setTimeout(() => setSearchTriggered(true), 0);
  };

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Identity Registry</h1>

      {/* Register Single Identity */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Register Identity</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Wallet Address</label>
            <input
              type="text"
              value={wallet}
              onChange={(e) => setWallet(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Identity Address (ONCHAINID)</label>
            <input
              type="text"
              value={identity}
              onChange={(e) => setIdentity(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Country Code (uint16)</label>
            <input
              type="number"
              value={country}
              onChange={(e) => setCountry(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="55"
            />
          </div>
        </div>
        <button
          onClick={handleRegister}
          disabled={registerPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {registerPending ? "Registering..." : "Register Identity"}
        </button>
        {registerSuccess && <p className="mt-2 text-sm text-green-400">Identity registered successfully!</p>}
      </div>

      {/* Search */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Search Identity</h2>
        <div className="flex gap-2">
          <input
            type="text"
            value={searchAddress}
            onChange={(e) => {
              setSearchAddress(e.target.value);
              setSearchTriggered(false);
            }}
            className="flex-1 rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
            placeholder="Enter wallet address to search..."
          />
          <button
            onClick={handleSearch}
            className="rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600"
          >
            Search
          </button>
        </div>

        {searchTriggered && searchAddr && (
          <div className="mt-4 rounded-lg border border-white/10 bg-white/[0.03] p-4">
            {containsResult === false ? (
              <p className="text-gray-400">Address not found in identity registry.</p>
            ) : containsResult === true ? (
              <div className="space-y-2">
                <div>
                  <span className="text-sm text-gray-400">Registered: </span>
                  <span className="text-green-400">Yes</span>
                </div>
                <div>
                  <span className="text-sm text-gray-400">Identity (ONCHAINID): </span>
                  <span className="font-mono text-sm text-indigo-400">{identityResult as string}</span>
                </div>
                <div>
                  <span className="text-sm text-gray-400">Country Code: </span>
                  <span className="text-indigo-400">{countryResult?.toString()}</span>
                </div>
              </div>
            ) : (
              <p className="text-gray-500">Querying...</p>
            )}
          </div>
        )}
      </div>

      {/* Batch Register */}
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Batch Register Identities</h2>
        <div className="space-y-4">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Wallet Addresses (comma-separated)</label>
            <textarea
              value={batchWallets}
              onChange={(e) => setBatchWallets(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              rows={2}
              placeholder="0xabc..., 0xdef..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Identity Addresses (comma-separated)</label>
            <textarea
              value={batchIdentities}
              onChange={(e) => setBatchIdentities(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              rows={2}
              placeholder="0x111..., 0x222..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Country Codes (comma-separated)</label>
            <input
              type="text"
              value={batchCountries}
              onChange={(e) => setBatchCountries(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="55, 1, 44"
            />
          </div>
        </div>
        <button
          onClick={handleBatchRegister}
          disabled={batchPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {batchPending ? "Registering..." : "Batch Register"}
        </button>
        {batchSuccess && <p className="mt-2 text-sm text-green-400">Batch registration successful!</p>}
      </div>
    </div>
  );
}

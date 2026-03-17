import { useState } from "react";
import type { Address } from "viem";
import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useAssets } from "@/hooks/use-assets";
import { useIndexerProtocolEvents } from "@/hooks/use-indexer";
import { ActivityFeed } from "@/components/ui/activity-feed";

export default function SnapshotsPage() {
  const { assets, isLoading: assetsLoading } = useAssets();

  // Create snapshot
  const { writeContract: writeSnapshot, data: snapshotHash, isPending: snapshotPending } = useWriteContract();
  const { isSuccess: snapshotSuccess } = useWaitForTransactionReceipt({ hash: snapshotHash });
  const [snapshotTokenId, setSnapshotTokenId] = useState("");

  const handleCreateSnapshot = () => {
    writeSnapshot({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "createSnapshot",
      args: [BigInt(snapshotTokenId || "0")],
    });
  };

  // View snapshots
  const [viewSnapshotTokenId, setViewSnapshotTokenId] = useState("");
  const { data: snapshotCount, isLoading: snapshotCountLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "getSnapshotCount",
    args: viewSnapshotTokenId ? [BigInt(viewSnapshotTokenId)] : undefined,
    query: { enabled: !!viewSnapshotTokenId },
  });

  // Create dividend
  const { writeContract: writeDividend, data: dividendHash, isPending: dividendPending } = useWriteContract();
  const { isSuccess: dividendSuccess } = useWaitForTransactionReceipt({ hash: dividendHash });
  const [divTokenId, setDivTokenId] = useState("");
  const [divSnapshotId, setDivSnapshotId] = useState("");
  const [divTotalAmount, setDivTotalAmount] = useState("");
  const [divPaymentToken, setDivPaymentToken] = useState("");

  const handleCreateDividend = () => {
    writeDividend({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "createDividend",
      args: [
        BigInt(divTokenId || "0"),
        BigInt(divSnapshotId || "0"),
        BigInt(divTotalAmount || "0"),
        divPaymentToken as Address,
      ],
    });
  };

  // View dividends
  const [viewDivTokenId, setViewDivTokenId] = useState("");
  const { data: dividendCount, isLoading: dividendCountLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "getDividendCount",
    args: viewDivTokenId ? [BigInt(viewDivTokenId)] : undefined,
    query: { enabled: !!viewDivTokenId },
  });

  // Indexer: snapshot + dividend events
  const { data: sdEvents, isLoading: sdEventsLoading } = useIndexerProtocolEvents({ first: 20 });
  const filteredSDEvents = sdEvents?.filter((e) =>
    ["snapshot_created", "dividend_created", "dividend_claimed"].includes(e.eventType)
  ) ?? [];

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Snapshots & Dividends</h1>

      {/* Create Snapshot */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Create Snapshot</h2>
        <div className="flex gap-2">
          <div className="flex-1">
            <label className="mb-1 block text-sm text-gray-400">Token ID</label>
            <input
              type="text"
              value={snapshotTokenId}
              onChange={(e) => setSnapshotTokenId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0"
            />
          </div>
          <div className="flex items-end">
            <button
              onClick={handleCreateSnapshot}
              disabled={snapshotPending}
              className="rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
            >
              {snapshotPending ? "Creating..." : "Create Snapshot"}
            </button>
          </div>
        </div>
        {snapshotSuccess && <p className="mt-2 text-sm text-green-400">Snapshot created!</p>}
      </div>

      {/* View Snapshots */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">View Snapshots</h2>
        <div className="flex gap-2">
          <select
            value={viewSnapshotTokenId}
            onChange={(e) => setViewSnapshotTokenId(e.target.value)}
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
        {viewSnapshotTokenId && (
          <div className="mt-4 rounded-lg border border-white/10 bg-white/[0.03] p-4">
            {snapshotCountLoading ? (
              <p className="text-gray-500">Loading...</p>
            ) : (
              <div>
                <p className="text-gray-300">
                  Total snapshots:{" "}
                  <span className="font-semibold text-indigo-400">
                    {snapshotCount?.toString() ?? "0"}
                  </span>
                </p>
                {snapshotCount != null && Number(snapshotCount) > 0 && (
                  <p className="mt-2 text-sm text-gray-400">
                    Snapshot IDs: 1 through {snapshotCount.toString()}
                  </p>
                )}
              </div>
            )}
          </div>
        )}
      </div>

      {/* Create Dividend */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Create Dividend</h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">Token ID</label>
            <input
              type="text"
              value={divTokenId}
              onChange={(e) => setDivTokenId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Snapshot ID</label>
            <input
              type="text"
              value={divSnapshotId}
              onChange={(e) => setDivSnapshotId(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="1"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Total Amount</label>
            <input
              type="text"
              value={divTotalAmount}
              onChange={(e) => setDivTotalAmount(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="1000000"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">Payment Token Address</label>
            <input
              type="text"
              value={divPaymentToken}
              onChange={(e) => setDivPaymentToken(e.target.value)}
              className="w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none"
              placeholder="0x... (ERC-20 token address)"
            />
          </div>
        </div>
        <button
          onClick={handleCreateDividend}
          disabled={dividendPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {dividendPending ? "Creating..." : "Create Dividend"}
        </button>
        {dividendSuccess && <p className="mt-2 text-sm text-green-400">Dividend created!</p>}
      </div>

      {/* View Dividends */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">View Dividends</h2>
        <div className="flex gap-2">
          <select
            value={viewDivTokenId}
            onChange={(e) => setViewDivTokenId(e.target.value)}
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
        {viewDivTokenId && (
          <div className="mt-4 rounded-lg border border-white/10 bg-white/[0.03] p-4">
            {dividendCountLoading ? (
              <p className="text-gray-500">Loading...</p>
            ) : (
              <p className="text-gray-300">
                Total dividends:{" "}
                <span className="font-semibold text-indigo-400">
                  {dividendCount?.toString() ?? "0"}
                </span>
              </p>
            )}
          </div>
        )}
      </div>

      {/* Snapshot & Dividend Event History */}
      <ActivityFeed
        events={filteredSDEvents}
        isLoading={sdEventsLoading}
        title="Snapshot & Dividend History"
      />
    </div>
  );
}

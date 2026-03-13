import { useState, useMemo } from "react";
import type { Address } from "viem";
import {
  useReadContract,
  useReadContracts,
  useWriteContract,
  useWaitForTransactionReceipt,
} from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useAssets } from "@/hooks/use-assets";

const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;

interface AssetGroup {
  parentTokenId: bigint;
  name: string;
  maxUnits: bigint;
  unitCount: bigint;
  nextUnitIndex: bigint;
  exists: boolean;
}

export default function GroupsPage() {
  const { assets, tokenIds, isLoading: assetsLoading } = useAssets();

  // ─── Create Group ───
  const [createParentTokenId, setCreateParentTokenId] = useState("");
  const [createName, setCreateName] = useState("");
  const [createMaxUnits, setCreateMaxUnits] = useState("");

  const {
    writeContract: writeCreateGroup,
    data: createHash,
    isPending: createPending,
  } = useWriteContract();
  const { isSuccess: createSuccess } = useWaitForTransactionReceipt({
    hash: createHash,
  });

  const handleCreateGroup = () => {
    writeCreateGroup({
      ...diamond,
      functionName: "createGroup",
      args: [
        {
          name: createName,
          parentTokenId: BigInt(createParentTokenId || "0"),
          maxUnits: BigInt(createMaxUnits || "0"),
        },
      ],
    });
  };

  // ─── Mint Unit ───
  const [mintGroupId, setMintGroupId] = useState("");
  const [mintName, setMintName] = useState("");
  const [mintSymbol, setMintSymbol] = useState("");
  const [mintUri, setMintUri] = useState("");
  const [mintSupplyCap, setMintSupplyCap] = useState("");
  const [mintRecipient, setMintRecipient] = useState("");
  const [mintAmount, setMintAmount] = useState("");

  const {
    writeContract: writeMintUnit,
    data: mintHash,
    isPending: mintPending,
  } = useWriteContract();
  const { isSuccess: mintSuccess } = useWaitForTransactionReceipt({
    hash: mintHash,
  });

  const handleMintUnit = () => {
    writeMintUnit({
      ...diamond,
      functionName: "mintUnit",
      args: [
        {
          groupId: BigInt(mintGroupId || "0"),
          name: mintName,
          symbol: mintSymbol,
          uri: mintUri,
          supplyCap: BigInt(mintSupplyCap || "0"),
          investor: mintRecipient as Address,
          amount: BigInt(mintAmount || "0"),
        },
      ],
    });
  };

  // ─── Batch Mint Units ───
  const [batchGroupId, setBatchGroupId] = useState("");
  const [batchEntries, setBatchEntries] = useState("");

  const {
    writeContract: writeBatchMint,
    data: batchHash,
    isPending: batchPending,
  } = useWriteContract();
  const { isSuccess: batchSuccess } = useWaitForTransactionReceipt({
    hash: batchHash,
  });

  const handleBatchMint = () => {
    const lines = batchEntries
      .split("\n")
      .map((l) => l.trim())
      .filter(Boolean);

    const params = lines.map((line) => {
      const [investor, amount, name, symbol, uri, supplyCap] = line
        .split(",")
        .map((s) => s.trim());
      return {
        groupId: BigInt(batchGroupId || "0"),
        name: name || "",
        symbol: symbol || "",
        uri: uri || "",
        supplyCap: BigInt(supplyCap || "0"),
        investor: investor as Address,
        amount: BigInt(amount || "0"),
      };
    });

    writeBatchMint({
      ...diamond,
      functionName: "mintUnitBatch",
      args: [params],
    });
  };

  // ─── View Groups ───
  const { data: groupIds, isLoading: groupsLoading } = useReadContract({
    ...diamond,
    functionName: "getRegisteredGroupIds",
  });

  const groupContracts = useMemo(() => {
    if (!groupIds || !Array.isArray(groupIds)) return [];
    return (groupIds as bigint[]).map((gid) => ({
      ...diamond,
      functionName: "getGroup" as const,
      args: [gid] as const,
    }));
  }, [groupIds]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { data: groupResults, isLoading: groupDataLoading } = useReadContracts({
    contracts: groupContracts as any,
    query: { enabled: !!groupIds && (groupIds as bigint[]).length > 0 },
  });

  const groups = useMemo(() => {
    if (!groupResults || !groupIds) return [];
    return (groupIds as bigint[]).reduce<
      { groupId: bigint; group: AssetGroup }[]
    >((acc, gid, i) => {
      const result = groupResults[i];
      if (result?.status === "success" && result.result) {
        acc.push({ groupId: gid, group: result.result as AssetGroup });
      }
      return acc;
    }, []);
  }, [groupResults, groupIds]);

  // ─── Expanded Group Children ───
  const [expandedGroupId, setExpandedGroupId] = useState<string | null>(null);

  const { data: childrenData } = useReadContract({
    ...diamond,
    functionName: "getGroupChildren",
    args: [BigInt(expandedGroupId || "0")],
    query: { enabled: !!expandedGroupId },
  });

  const childTokenIds = (childrenData as bigint[]) ?? [];

  const childGroupContracts = useMemo(() => {
    if (!childTokenIds.length) return [];
    return childTokenIds.map((cid) => ({
      ...diamond,
      functionName: "getChildGroup" as const,
      args: [cid] as const,
    }));
  }, [childTokenIds]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { data: childGroupResults } = useReadContracts({
    contracts: childGroupContracts as any,
    query: { enabled: childTokenIds.length > 0 },
  });

  const inputClass =
    "w-full rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white placeholder-gray-500 focus:border-indigo-400 focus:outline-none";

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">
        Asset Groups
      </h1>

      {/* Info Box */}
      <div className="mb-8 rounded-xl border border-indigo-500/30 bg-indigo-500/10 p-4">
        <p className="text-sm text-indigo-200">
          Asset Groups enable hierarchical tokenization. Create a group from a
          parent asset (e.g., a building registered as tokenId 1) and mint child
          units (e.g., individual apartments). Each child inherits the
          parent&apos;s compliance rules, identity profile, and allowed
          countries.
        </p>
      </div>

      {/* Section 1: Create Group */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">
          Create Group
        </h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Parent Token ID
            </label>
            <select
              value={createParentTokenId}
              onChange={(e) => setCreateParentTokenId(e.target.value)}
              className={inputClass}
            >
              <option value="" className="bg-[#0a0a0f]">
                Select parent asset...
              </option>
              {assetsLoading ? (
                <option disabled className="bg-[#0a0a0f]">
                  Loading...
                </option>
              ) : (
                assets.map((asset) => (
                  <option
                    key={asset.tokenId.toString()}
                    value={asset.tokenId.toString()}
                    className="bg-[#0a0a0f]"
                  >
                    #{asset.tokenId.toString()} - {asset.name} ({asset.symbol})
                  </option>
                ))
              )}
            </select>
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Group Name
            </label>
            <input
              type="text"
              value={createName}
              onChange={(e) => setCreateName(e.target.value)}
              className={inputClass}
              placeholder="Tower A Apartments"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Max Units
            </label>
            <input
              type="text"
              value={createMaxUnits}
              onChange={(e) => setCreateMaxUnits(e.target.value)}
              className={inputClass}
              placeholder="100"
            />
          </div>
        </div>
        <button
          onClick={handleCreateGroup}
          disabled={createPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {createPending ? "Creating..." : "Create Group"}
        </button>
        {createSuccess && (
          <p className="mt-2 text-sm text-green-400">
            Group created successfully!
          </p>
        )}
      </div>

      {/* Section 2: Mint Unit */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">
          Mint Unit
        </h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Group ID
            </label>
            <input
              type="text"
              value={mintGroupId}
              onChange={(e) => setMintGroupId(e.target.value)}
              className={inputClass}
              placeholder="1"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Unit Name
            </label>
            <input
              type="text"
              value={mintName}
              onChange={(e) => setMintName(e.target.value)}
              className={inputClass}
              placeholder="Apartment 101"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Unit Symbol
            </label>
            <input
              type="text"
              value={mintSymbol}
              onChange={(e) => setMintSymbol(e.target.value)}
              className={inputClass}
              placeholder="APT101"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Metadata URI
            </label>
            <input
              type="text"
              value={mintUri}
              onChange={(e) => setMintUri(e.target.value)}
              className={inputClass}
              placeholder="https://metadata.example.com/{id}"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Supply Cap
            </label>
            <input
              type="text"
              value={mintSupplyCap}
              onChange={(e) => setMintSupplyCap(e.target.value)}
              className={inputClass}
              placeholder="1000"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Recipient Address
            </label>
            <input
              type="text"
              value={mintRecipient}
              onChange={(e) => setMintRecipient(e.target.value)}
              className={inputClass}
              placeholder="0x..."
            />
          </div>
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Amount (fractions)
            </label>
            <input
              type="text"
              value={mintAmount}
              onChange={(e) => setMintAmount(e.target.value)}
              className={inputClass}
              placeholder="100"
            />
          </div>
        </div>
        <button
          onClick={handleMintUnit}
          disabled={mintPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {mintPending ? "Minting..." : "Mint Unit"}
        </button>
        {mintSuccess && (
          <p className="mt-2 text-sm text-green-400">
            Unit minted successfully!
          </p>
        )}
      </div>

      {/* Section 3: Batch Mint Units */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">
          Batch Mint Units
        </h2>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-1 block text-sm text-gray-400">
              Group ID
            </label>
            <input
              type="text"
              value={batchGroupId}
              onChange={(e) => setBatchGroupId(e.target.value)}
              className={inputClass}
              placeholder="1"
            />
          </div>
        </div>
        <div className="mt-4">
          <label className="mb-1 block text-sm text-gray-400">
            Entries (one per line: address, amount, name, symbol, uri,
            supplyCap)
          </label>
          <textarea
            value={batchEntries}
            onChange={(e) => setBatchEntries(e.target.value)}
            rows={5}
            className={inputClass}
            placeholder={
              "0xABC..., 100, Apt 201, APT201, https://meta.io/201, 1000\n0xDEF..., 200, Apt 202, APT202, https://meta.io/202, 1000"
            }
          />
        </div>
        <button
          onClick={handleBatchMint}
          disabled={batchPending}
          className="mt-4 rounded-lg bg-indigo-500 px-6 py-2 font-medium text-white transition-colors hover:bg-indigo-600 disabled:opacity-50"
        >
          {batchPending ? "Minting..." : "Batch Mint"}
        </button>
        {batchSuccess && (
          <p className="mt-2 text-sm text-green-400">
            Batch mint completed successfully!
          </p>
        )}
      </div>

      {/* Section 4: View Groups */}
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">
          Registered Groups
        </h2>
        {groupsLoading || groupDataLoading ? (
          <p className="text-gray-500">Loading groups...</p>
        ) : groups.length === 0 ? (
          <p className="text-gray-500">No groups registered yet.</p>
        ) : (
          <div className="space-y-4">
            {groups.map(({ groupId, group }) => (
              <div
                key={groupId.toString()}
                className="rounded-lg border border-white/10 bg-white/5 p-4"
              >
                <div className="flex items-center justify-between">
                  <div>
                    <h3 className="text-lg font-medium text-white">
                      {group.name}
                      <span className="ml-2 text-sm text-gray-400">
                        (Group #{groupId.toString()})
                      </span>
                    </h3>
                    <div className="mt-1 flex gap-6 text-sm text-gray-400">
                      <span>
                        Parent Token:{" "}
                        <span className="font-mono text-gray-300">
                          {group.parentTokenId.toString()}
                        </span>
                      </span>
                      <span>
                        Units:{" "}
                        <span className="text-gray-300">
                          {group.unitCount.toString()} /{" "}
                          {group.maxUnits.toString()}
                        </span>
                      </span>
                    </div>
                  </div>
                  <button
                    onClick={() =>
                      setExpandedGroupId(
                        expandedGroupId === groupId.toString()
                          ? null
                          : groupId.toString()
                      )
                    }
                    className="rounded-lg border border-white/10 px-4 py-1.5 text-sm text-gray-300 transition-colors hover:bg-white/5"
                  >
                    {expandedGroupId === groupId.toString()
                      ? "Collapse"
                      : "View Children"}
                  </button>
                </div>

                {expandedGroupId === groupId.toString() && (
                  <div className="mt-4 border-t border-white/10 pt-4">
                    {childTokenIds.length === 0 ? (
                      <p className="text-sm text-gray-500">
                        No children minted yet.
                      </p>
                    ) : (
                      <div className="overflow-x-auto">
                        <table className="w-full text-left text-sm text-gray-300">
                          <thead>
                            <tr className="border-b border-white/10 text-gray-400">
                              <th className="pb-2 pr-4">Child Token ID</th>
                              <th className="pb-2">Parent Group ID</th>
                            </tr>
                          </thead>
                          <tbody className="divide-y divide-white/10">
                            {childTokenIds.map((cid, idx) => {
                              const childResult = childGroupResults?.[idx];
                              const parentGroupId =
                                childResult?.status === "success"
                                  ? (childResult.result as bigint)
                                  : null;
                              return (
                                <tr key={cid.toString()}>
                                  <td className="py-2 pr-4 font-mono">
                                    {cid.toString()}
                                  </td>
                                  <td className="py-2 font-mono">
                                    {parentGroupId !== null
                                      ? parentGroupId.toString()
                                      : "-"}
                                  </td>
                                </tr>
                              );
                            })}
                          </tbody>
                        </table>
                      </div>
                    )}
                  </div>
                )}
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

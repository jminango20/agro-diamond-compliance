import { useCallback, useMemo } from 'react';
import { Link } from 'react-router-dom';
import {
  useAccount,
  useReadContract,
  useReadContracts,
  useWriteContract,
  useWaitForTransactionReceipt,
} from 'wagmi';

import { DIAMOND_ADDRESS, diamondAbi } from '@/config/contracts';
import { formatTokenAmount, truncateAddress } from '@/lib/format';

const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;

type AssetConfig = {
  name: string;
  symbol: string;
};

type DividendRow = {
  dividendId: bigint;
  snapshotId: bigint;
  tokenId: bigint;
  totalAmount: bigint;
  paymentToken: string;
  claimedAmount: bigint;
  createdAt: bigint;
  claimable: bigint;
  claimed: boolean;
  tokenName: string;
};

function ClaimButton({ dividendId }: { dividendId: bigint }) {
  const {
    writeContract,
    data: txHash,
    isPending,
    error,
    reset,
  } = useWriteContract();

  const { isLoading: isTxLoading, isSuccess } = useWaitForTransactionReceipt({
    hash: txHash,
  });

  const handleClaim = useCallback(() => {
    reset();
    writeContract({
      ...diamond,
      functionName: 'claimDividend',
      args: [dividendId],
    });
  }, [dividendId, writeContract, reset]);

  if (isSuccess) {
    return (
      <span className="text-green-400 text-sm font-medium">Claimed!</span>
    );
  }

  return (
    <div>
      <button
        onClick={handleClaim}
        disabled={isPending || isTxLoading}
        className="px-3 py-1.5 rounded-lg bg-indigo-600 text-white text-sm font-medium hover:bg-indigo-500 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {isPending ? 'Confirm...' : isTxLoading ? 'Waiting...' : 'Claim'}
      </button>
      {error && (
        <p className="text-red-400 text-xs mt-1">{error.message.slice(0, 80)}</p>
      )}
    </div>
  );
}

export default function DividendsPage() {
  const { address, isConnected } = useAccount();

  const { data: tokenIds, isLoading: idsLoading } = useReadContract({
    ...diamond,
    functionName: 'getRegisteredTokenIds',
  });

  const registeredIds = (tokenIds as bigint[]) ?? [];

  const configContracts = useMemo(() => {
    if (registeredIds.length === 0) return [];
    return registeredIds.map((id) => ({
      ...diamond,
      functionName: 'getAssetConfig' as const,
      args: [id] as const,
    }));
  }, [tokenIds]);

  const { data: configsData } = useReadContracts({
    contracts: configContracts,
    query: { enabled: configContracts.length > 0 },
  });

  const dividendIdsContracts = useMemo(() => {
    if (registeredIds.length === 0) return [];
    return registeredIds.map((id) => ({
      ...diamond,
      functionName: 'getTokenDividends' as const,
      args: [id] as const,
    }));
  }, [tokenIds]);

  const { data: dividendIdsData, isLoading: dividendsLoading } = useReadContracts({
    contracts: dividendIdsContracts,
    query: { enabled: dividendIdsContracts.length > 0 },
  });

  // Flatten all dividend IDs across all tokens
  const allDividendIds = useMemo(() => {
    const result: { dividendId: bigint; tokenIndex: number }[] = [];
    dividendIdsData?.forEach((r, tokenIndex) => {
      const ids = r.result as bigint[] | undefined;
      if (ids) {
        ids.forEach((dividendId) => {
          result.push({ dividendId, tokenIndex });
        });
      }
    });
    return result;
  }, [dividendIdsData]);

  // Fetch details for all dividends
  const infoContracts = useMemo(() => {
    if (allDividendIds.length === 0) return [];
    return allDividendIds.map(({ dividendId }) => ({
      ...diamond,
      functionName: 'getDividend' as const,
      args: [dividendId] as const,
    }));
  }, [allDividendIds]);

  const { data: dividendInfos } = useReadContracts({
    contracts: infoContracts,
    query: { enabled: infoContracts.length > 0 },
  });

  const claimableContracts = useMemo(() => {
    if (!isConnected || !address || allDividendIds.length === 0) return [];
    return allDividendIds.map(({ dividendId }) => ({
      ...diamond,
      functionName: 'claimableAmount' as const,
      args: [dividendId, address] as const,
    }));
  }, [allDividendIds, isConnected, address]);

  const { data: claimableAmounts } = useReadContracts({
    contracts: claimableContracts,
    query: { enabled: claimableContracts.length > 0 },
  });

  const claimedContracts = useMemo(() => {
    if (!isConnected || !address || allDividendIds.length === 0) return [];
    return allDividendIds.map(({ dividendId }) => ({
      ...diamond,
      functionName: 'hasClaimed' as const,
      args: [dividendId, address] as const,
    }));
  }, [allDividendIds, isConnected, address]);

  const { data: claimedStatuses } = useReadContracts({
    contracts: claimedContracts,
    query: { enabled: claimedContracts.length > 0 },
  });

  const isLoading = idsLoading || dividendsLoading;

  const rows: DividendRow[] = allDividendIds.map(({ dividendId, tokenIndex }, i) => {
    const rawInfo = dividendInfos?.[i]?.result as
      | [bigint, bigint, bigint, string, bigint, bigint]
      | undefined;
    const config = configsData?.[tokenIndex]?.result as AssetConfig | undefined;
    return {
      dividendId,
      snapshotId: rawInfo?.[0] ?? 0n,
      tokenId: rawInfo?.[1] ?? 0n,
      totalAmount: rawInfo?.[2] ?? 0n,
      paymentToken: rawInfo?.[3] ?? '',
      claimedAmount: rawInfo?.[4] ?? 0n,
      createdAt: rawInfo?.[5] ?? 0n,
      claimable: (claimableAmounts?.[i]?.result as bigint) ?? 0n,
      claimed: (claimedStatuses?.[i]?.result as boolean) ?? false,
      tokenName: config?.name ?? `Token #${registeredIds[tokenIndex]?.toString()}`,
    };
  });

  if (!isConnected) {
    return (
      <div className="min-h-screen bg-[#0a0a0f] p-8">
        <h1 className="mb-8 text-3xl font-bold text-white">Dividends</h1>
        <div className="rounded-xl bg-white/5 border border-white/10 p-6 text-center">
          <p className="text-gray-400 text-lg">Connect your wallet to view dividends.</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <Link to="/portfolio" className="text-indigo-400 hover:text-indigo-300 text-sm mb-6 inline-block">
        &larr; Back to Portfolio
      </Link>
      <h1 className="mb-8 text-3xl font-bold text-white">Dividends</h1>

      {isLoading ? (
        <div className="space-y-4">
          {[1, 2].map((i) => (
            <div key={i} className="rounded-xl bg-white/5 border border-white/10 p-6 animate-pulse">
              <div className="h-5 w-32 rounded bg-white/10 mb-3" />
              <div className="h-4 w-48 rounded bg-white/10" />
            </div>
          ))}
        </div>
      ) : rows.length === 0 ? (
        <div className="rounded-xl bg-white/5 border border-white/10 p-6 text-center">
          <p className="text-gray-400 text-lg">No dividends available.</p>
          <p className="text-gray-500 text-sm mt-2">
            Dividends will appear here when they are distributed for tokens you hold.
          </p>
        </div>
      ) : (
        <div className="space-y-4">
          {rows.map((row) => (
            <div
              key={row.dividendId.toString()}
              className="rounded-xl bg-white/5 border border-white/10 p-6"
            >
              <div className="flex items-start justify-between">
                <div>
                  <div className="flex items-center gap-3 mb-2">
                    <h3 className="text-lg font-semibold text-white">
                      Dividend #{row.dividendId.toString()}
                    </h3>
                    <span className="text-xs font-mono text-gray-500 bg-white/5 px-2 py-0.5 rounded">
                      {row.tokenName}
                    </span>
                  </div>
                  <div className="grid grid-cols-2 gap-x-8 gap-y-2 text-sm">
                    <div>
                      <span className="text-gray-500">Total Amount</span>
                      <p className="text-white font-mono">
                        {formatTokenAmount(row.totalAmount)}
                      </p>
                    </div>
                    <div>
                      <span className="text-gray-500">Payment Token</span>
                      <p className="text-indigo-400 font-mono">
                        {row.paymentToken === '0x0000000000000000000000000000000000000000'
                          ? 'ETH'
                          : truncateAddress(row.paymentToken)}
                      </p>
                    </div>
                    <div>
                      <span className="text-gray-500">Claimable</span>
                      <p className="text-white font-mono">
                        {formatTokenAmount(row.claimable)}
                      </p>
                    </div>
                    <div>
                      <span className="text-gray-500">Status</span>
                      <p className={row.claimed ? 'text-gray-500' : 'text-green-400'}>
                        {row.claimed ? 'Claimed' : 'Available'}
                      </p>
                    </div>
                  </div>
                </div>

                <div className="flex-shrink-0 ml-4">
                  {row.claimed ? (
                    <span className="px-3 py-1.5 rounded-lg bg-white/5 text-gray-500 text-sm">
                      Claimed
                    </span>
                  ) : row.claimable > 0n ? (
                    <ClaimButton dividendId={row.dividendId} />
                  ) : (
                    <span className="text-gray-600 text-sm">No claimable amount</span>
                  )}
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

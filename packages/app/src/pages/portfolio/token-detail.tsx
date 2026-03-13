import { Link, useParams } from 'react-router-dom';
import { useAccount, useReadContract } from 'wagmi';

import { DIAMOND_ADDRESS, diamondAbi } from '@/config/contracts';
import { formatTokenAmount, truncateAddress } from '@/lib/format';

const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;

type TokenInfo = [string, string, string, bigint, bigint, bigint, string, boolean];

type PartitionBalance = [bigint, bigint, bigint, bigint];

export default function TokenDetailPage() {
  const { tokenId } = useParams();
  const tokenIdBigInt = BigInt(tokenId!);
  const { address, isConnected } = useAccount();

  const { data: tokenInfo, isLoading: infoLoading } = useReadContract({
    ...diamond,
    functionName: 'tokenInfo',
    args: [tokenIdBigInt],
  });

  const { data: partitionData, isLoading: partitionLoading } = useReadContract({
    ...diamond,
    functionName: 'partitionBalanceOf',
    args: [address!, tokenIdBigInt],
    query: { enabled: isConnected },
  });

  const { data: isVerified, isLoading: verifiedLoading } = useReadContract({
    ...diamond,
    functionName: 'contains',
    args: [address!],
    query: { enabled: isConnected },
  });

  const info = tokenInfo as TokenInfo | undefined;
  const partition = partitionData as PartitionBalance | undefined;

  const name = info?.[0] ?? '';
  const symbol = info?.[1] ?? '';
  const uri = info?.[2] ?? '';
  const totalSupply = info?.[3] ?? 0n;
  const supplyCap = info?.[4] ?? 0n;
  const holderCount = info?.[5] ?? 0n;
  const issuer = info?.[6] ?? '';
  const paused = info?.[7] ?? false;

  const free = partition?.[0] ?? 0n;
  const locked = partition?.[1] ?? 0n;
  const custody = partition?.[2] ?? 0n;
  const pendingSettlement = partition?.[3] ?? 0n;
  const totalBalance = free + locked + custody + pendingSettlement;

  const isLoading = infoLoading || partitionLoading || verifiedLoading;

  function BalanceBar({ label, value, total, color }: { label: string; value: bigint; total: bigint; color: string }) {
    const pct = total > 0n ? Number((value * 10000n) / total) / 100 : 0;
    return (
      <div className="mb-3">
        <div className="flex items-center justify-between mb-1">
          <span className="text-sm text-gray-400">{label}</span>
          <span className="text-sm font-mono text-white">
            {formatTokenAmount(value)} <span className="text-gray-500">({pct.toFixed(1)}%)</span>
          </span>
        </div>
        <div className="h-2 w-full rounded-full bg-white/10">
          <div className={`h-2 rounded-full ${color}`} style={{ width: `${Math.min(pct, 100)}%` }} />
        </div>
      </div>
    );
  }

  if (!isConnected) {
    return (
      <div className="min-h-screen bg-[#0a0a0f] p-8">
        <h1 className="mb-8 text-3xl font-bold text-white">Token Detail</h1>
        <div className="rounded-xl bg-white/5 border border-white/10 p-6 text-center">
          <p className="text-gray-400 text-lg">Connect your wallet to view token details.</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <Link to="/portfolio" className="text-indigo-400 hover:text-indigo-300 text-sm mb-6 inline-block">
        &larr; Back to Portfolio
      </Link>

      {isLoading ? (
        <div className="space-y-4">
          {[1, 2, 3].map((i) => (
            <div key={i} className="rounded-xl bg-white/5 border border-white/10 p-6 animate-pulse">
              <div className="h-6 w-48 rounded bg-white/10 mb-3" />
              <div className="h-4 w-32 rounded bg-white/10" />
            </div>
          ))}
        </div>
      ) : (
        <div className="space-y-6">
          {/* Token Metadata */}
          <div className="rounded-xl bg-white/5 border border-white/10 p-6">
            <div className="flex items-center justify-between mb-4">
              <div>
                <h1 className="text-2xl font-bold text-white">{name || `Token #${tokenId}`}</h1>
                <p className="text-gray-400 mt-1">
                  <span className="font-mono text-indigo-400">{symbol}</span>
                  <span className="mx-2 text-gray-600">|</span>
                  Token ID: {tokenId}
                </p>
              </div>
              {paused && (
                <span className="px-3 py-1 rounded-full text-xs font-semibold bg-red-500/20 text-red-400 border border-red-500/30">
                  Paused
                </span>
              )}
            </div>
            {uri && (
              <p className="text-sm text-gray-500 truncate">
                URI: <span className="font-mono text-gray-400">{uri}</span>
              </p>
            )}
          </div>

          {/* Balance Breakdown */}
          <div className="rounded-xl bg-white/5 border border-white/10 p-6">
            <h2 className="text-lg font-semibold text-white mb-1">Balance Breakdown</h2>
            <p className="text-3xl font-bold text-indigo-400 mb-6">{formatTokenAmount(totalBalance)} <span className="text-base text-gray-400">{symbol}</span></p>

            <BalanceBar label="Free (Available)" value={free} total={totalBalance} color="bg-green-500" />
            <BalanceBar label="Locked" value={locked} total={totalBalance} color="bg-yellow-500" />
            <BalanceBar label="Custody" value={custody} total={totalBalance} color="bg-blue-500" />
            <BalanceBar label="Pending Settlement" value={pendingSettlement} total={totalBalance} color="bg-purple-500" />

            <div className="mt-6">
              <Link
                to={`/portfolio/transfer?tokenId=${tokenId}`}
                className="inline-flex items-center px-4 py-2 rounded-lg bg-indigo-600 text-white text-sm font-medium hover:bg-indigo-500 transition-colors"
              >
                Transfer Tokens
              </Link>
            </div>
          </div>

          {/* Token Stats */}
          <div className="rounded-xl bg-white/5 border border-white/10 p-6">
            <h2 className="text-lg font-semibold text-white mb-4">Token Statistics</h2>
            <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
              <div>
                <p className="text-sm text-gray-400">Total Supply</p>
                <p className="text-xl font-bold text-white">{formatTokenAmount(totalSupply)}</p>
              </div>
              <div>
                <p className="text-sm text-gray-400">Supply Cap</p>
                <p className="text-xl font-bold text-white">
                  {supplyCap === 0n ? 'Unlimited' : formatTokenAmount(supplyCap)}
                </p>
              </div>
              <div>
                <p className="text-sm text-gray-400">Holders</p>
                <p className="text-xl font-bold text-white">{holderCount.toString()}</p>
              </div>
            </div>
            {issuer && (
              <div className="mt-4 pt-4 border-t border-white/10">
                <p className="text-sm text-gray-400">Issuer</p>
                <p className="text-sm font-mono text-indigo-400">{truncateAddress(issuer)}</p>
              </div>
            )}
          </div>

          {/* Compliance Status */}
          <div className="rounded-xl bg-white/5 border border-white/10 p-6">
            <h2 className="text-lg font-semibold text-white mb-4">Compliance Status</h2>
            <div className="flex items-center gap-3">
              {isVerified ? (
                <>
                  <div className="h-3 w-3 rounded-full bg-green-500" />
                  <p className="text-green-400 font-medium">Identity Verified</p>
                </>
              ) : (
                <>
                  <div className="h-3 w-3 rounded-full bg-red-500" />
                  <p className="text-red-400 font-medium">Identity Not Verified</p>
                </>
              )}
            </div>
            <p className="text-sm text-gray-500 mt-2">
              Your wallet must be registered in the Identity Registry to send and receive tokens.
            </p>
          </div>
        </div>
      )}
    </div>
  );
}

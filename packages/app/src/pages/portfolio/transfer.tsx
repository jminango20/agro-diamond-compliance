import { useCallback, useEffect, useMemo, useState } from 'react';
import { Link, useSearchParams } from 'react-router-dom';
import { type Address, isAddress } from 'viem';
import {
  useAccount,
  useReadContract,
  useReadContracts,
  useWriteContract,
  useWaitForTransactionReceipt,
} from 'wagmi';

import { DIAMOND_ADDRESS, diamondAbi } from '@/config/contracts';
import { formatTokenAmount } from '@/lib/format';

const diamond = { address: DIAMOND_ADDRESS, abi: diamondAbi } as const;

type AssetConfig = {
  name: string;
  symbol: string;
  uri: string;
  supplyCap: bigint;
  totalSupply: bigint;
  identityProfileId: number;
  complianceModule: string;
  issuer: string;
  paused: boolean;
  exists: boolean;
  allowedCountries: number[];
};

export default function TransferPage() {
  const [searchParams] = useSearchParams();
  const initialTokenId = searchParams.get('tokenId') ?? '';

  const { address, isConnected } = useAccount();
  const [selectedTokenId, setSelectedTokenId] = useState(initialTokenId);
  const [recipient, setRecipient] = useState('');
  const [amount, setAmount] = useState('');
  const [complianceResult, setComplianceResult] = useState<{
    checked: boolean;
    ok: boolean;
    reason: string;
  }>({ checked: false, ok: false, reason: '' });

  const { data: tokenIds } = useReadContract({
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

  const selectedBigInt = selectedTokenId ? BigInt(selectedTokenId) : undefined;

  const { data: balance } = useReadContract({
    ...diamond,
    functionName: 'balanceOf',
    args: [address!, selectedBigInt!],
    query: { enabled: isConnected && selectedBigInt !== undefined },
  });

  const { refetch: checkCompliance } = useReadContract({
    ...diamond,
    functionName: 'canTransfer',
    args: [
      selectedBigInt!,
      address!,
      recipient as Address,
      amount ? BigInt(amount) : 0n,
      '0x' as `0x${string}`,
    ],
    query: { enabled: false },
  });

  const {
    writeContract,
    data: txHash,
    isPending: isWritePending,
    error: writeError,
    reset: resetWrite,
  } = useWriteContract();

  const { isLoading: isTxLoading, isSuccess: isTxSuccess } =
    useWaitForTransactionReceipt({ hash: txHash });

  const handleCheckCompliance = useCallback(async () => {
    if (!selectedTokenId || !recipient || !amount || !isAddress(recipient)) return;
    const result = await checkCompliance();
    if (result.data) {
      const [ok, reason] = result.data as [boolean, `0x${string}`];
      setComplianceResult({
        checked: true,
        ok,
        reason: ok ? '' : reason,
      });
    }
  }, [selectedTokenId, recipient, amount, checkCompliance]);

  const handleTransfer = useCallback(() => {
    if (!selectedBigInt || !address || !isAddress(recipient)) return;
    resetWrite();
    writeContract({
      ...diamond,
      functionName: 'safeTransferFrom',
      args: [address, recipient as Address, selectedBigInt, BigInt(amount), '0x'],
    });
  }, [selectedBigInt, address, recipient, amount, writeContract, resetWrite]);

  useEffect(() => {
    setComplianceResult({ checked: false, ok: false, reason: '' });
  }, [selectedTokenId, recipient, amount]);

  const formValid =
    selectedTokenId && recipient && amount && isAddress(recipient) && BigInt(amount || '0') > 0n;

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <Link to="/portfolio" className="text-indigo-400 hover:text-indigo-300 text-sm mb-6 inline-block">
        &larr; Back to Portfolio
      </Link>
      <h1 className="mb-8 text-3xl font-bold text-white">Transfer Tokens</h1>
      <div className="max-w-2xl">
        {!isConnected ? (
          <div className="rounded-xl bg-white/5 border border-white/10 p-6 text-center">
            <p className="text-gray-400 text-lg">Connect your wallet to transfer tokens.</p>
          </div>
        ) : (
          <div className="space-y-6">
            {/* Transfer Form */}
            <div className="rounded-xl bg-white/5 border border-white/10 p-6">
              <h2 className="text-lg font-semibold text-white mb-6">Transfer Details</h2>

              <div className="space-y-4">
                {/* Token Selection */}
                <div>
                  <label className="block text-sm text-gray-400 mb-1">Token</label>
                  <select
                    value={selectedTokenId}
                    onChange={(e) => setSelectedTokenId(e.target.value)}
                    className="w-full rounded-lg bg-white/5 border border-white/10 px-4 py-2.5 text-white focus:outline-none focus:border-indigo-400"
                  >
                    <option value="" className="bg-[#0a0a0f]">
                      Select a token...
                    </option>
                    {registeredIds.map((id, i) => {
                      const config = configsData?.[i]?.result as AssetConfig | undefined;
                      return (
                        <option key={id.toString()} value={id.toString()} className="bg-[#0a0a0f]">
                          {config?.name ?? `Token #${id.toString()}`} ({config?.symbol ?? '???'})
                        </option>
                      );
                    })}
                  </select>
                </div>

                {/* Balance Display */}
                {selectedTokenId && balance !== undefined && (
                  <div className="text-sm text-gray-400">
                    Available balance:{' '}
                    <span className="font-mono text-indigo-400">
                      {formatTokenAmount(balance as bigint)}
                    </span>
                  </div>
                )}

                {/* Recipient */}
                <div>
                  <label className="block text-sm text-gray-400 mb-1">Recipient Address</label>
                  <input
                    type="text"
                    value={recipient}
                    onChange={(e) => setRecipient(e.target.value)}
                    placeholder="0x..."
                    className="w-full rounded-lg bg-white/5 border border-white/10 px-4 py-2.5 text-white font-mono text-sm focus:outline-none focus:border-indigo-400 placeholder:text-gray-600"
                  />
                  {recipient && !isAddress(recipient) && (
                    <p className="text-red-400 text-xs mt-1">Invalid address</p>
                  )}
                </div>

                {/* Amount */}
                <div>
                  <label className="block text-sm text-gray-400 mb-1">Amount</label>
                  <input
                    type="text"
                    value={amount}
                    onChange={(e) => {
                      const v = e.target.value.replace(/[^0-9]/g, '');
                      setAmount(v);
                    }}
                    placeholder="0"
                    className="w-full rounded-lg bg-white/5 border border-white/10 px-4 py-2.5 text-white font-mono focus:outline-none focus:border-indigo-400 placeholder:text-gray-600"
                  />
                </div>
              </div>
            </div>

            {/* Compliance Check */}
            <div className="rounded-xl bg-white/5 border border-white/10 p-6">
              <h2 className="text-lg font-semibold text-white mb-4">Compliance Check</h2>

              {!complianceResult.checked ? (
                <div>
                  <p className="text-sm text-gray-400 mb-4">
                    Verify that this transfer meets all compliance requirements before sending.
                  </p>
                  <button
                    onClick={handleCheckCompliance}
                    disabled={!formValid}
                    className="px-4 py-2 rounded-lg bg-white/10 border border-white/10 text-white text-sm font-medium hover:bg-white/15 transition-colors disabled:opacity-40 disabled:cursor-not-allowed"
                  >
                    Check Compliance
                  </button>
                </div>
              ) : complianceResult.ok ? (
                <div className="flex items-center gap-3">
                  <div className="h-8 w-8 rounded-full bg-green-500/20 flex items-center justify-center">
                    <span className="text-green-400 text-lg">&#10003;</span>
                  </div>
                  <div>
                    <p className="text-green-400 font-medium">Transfer Allowed</p>
                    <p className="text-sm text-gray-500">All compliance checks passed.</p>
                  </div>
                </div>
              ) : (
                <div className="flex items-center gap-3">
                  <div className="h-8 w-8 rounded-full bg-red-500/20 flex items-center justify-center">
                    <span className="text-red-400 text-lg">&#10007;</span>
                  </div>
                  <div>
                    <p className="text-red-400 font-medium">Transfer Rejected</p>
                    <p className="text-sm text-gray-500 font-mono">Reason: {complianceResult.reason}</p>
                  </div>
                </div>
              )}
            </div>

            {/* Transfer Action */}
            {complianceResult.checked && complianceResult.ok && (
              <div className="rounded-xl bg-white/5 border border-white/10 p-6">
                {isTxSuccess ? (
                  <div className="text-center">
                    <div className="h-12 w-12 rounded-full bg-green-500/20 flex items-center justify-center mx-auto mb-3">
                      <span className="text-green-400 text-2xl">&#10003;</span>
                    </div>
                    <p className="text-green-400 font-semibold text-lg">Transfer Successful</p>
                    <p className="text-sm text-gray-500 font-mono mt-1 break-all">{txHash}</p>
                    <Link
                      to="/portfolio"
                      className="inline-block mt-4 text-indigo-400 hover:text-indigo-300 text-sm"
                    >
                      &larr; Back to Portfolio
                    </Link>
                  </div>
                ) : (
                  <div>
                    <button
                      onClick={handleTransfer}
                      disabled={isWritePending || isTxLoading}
                      className="w-full px-4 py-3 rounded-lg bg-indigo-600 text-white font-medium hover:bg-indigo-500 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                      {isWritePending
                        ? 'Confirm in Wallet...'
                        : isTxLoading
                          ? 'Waiting for Confirmation...'
                          : 'Transfer'}
                    </button>
                    {writeError && (
                      <p className="text-red-400 text-sm mt-2">
                        Error: {writeError.message.slice(0, 120)}
                      </p>
                    )}
                  </div>
                )}
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}

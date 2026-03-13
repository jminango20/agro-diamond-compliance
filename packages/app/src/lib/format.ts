export function truncateAddress(addr: string): string {
  return addr.slice(0, 6) + '...' + addr.slice(-4);
}

export function formatTokenAmount(amount: bigint, decimals = 0): string {
  if (decimals === 0) return amount.toString();
  const str = amount.toString().padStart(decimals + 1, '0');
  return str.slice(0, -decimals) + '.' + str.slice(-decimals);
}

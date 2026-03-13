export function formatGas(n: number | null): string {
  if (n === null) return "—";
  return n.toLocaleString("en-US");
}

export function formatDuration(ms: number): string {
  if (ms < 1) return "<1ms";
  if (ms < 1000) return `${Math.round(ms)}ms`;
  return `${(ms / 1000).toFixed(1)}s`;
}

export function formatPct(n: number): string {
  return `${n.toFixed(1)}%`;
}

export function gasLevel(gas: number, p25: number, p75: number): "low" | "mid" | "high" {
  if (gas <= p25) return "low";
  if (gas <= p75) return "mid";
  return "high";
}

export function coverageColor(pct: number): string {
  if (pct >= 80) return "text-emerald-400";
  if (pct >= 50) return "text-yellow-400";
  return "text-red-400";
}

export function coverageBg(pct: number): string {
  if (pct >= 80) return "bg-emerald-500/20";
  if (pct >= 50) return "bg-yellow-500/20";
  return "bg-red-500/20";
}

export function statusColor(status: string): string {
  switch (status) {
    case "success":
    case "pass":
      return "text-emerald-400";
    case "failure":
    case "fail":
      return "text-red-400";
    default:
      return "text-gray-500";
  }
}

export function statusBg(status: string): string {
  switch (status) {
    case "success":
    case "pass":
      return "bg-emerald-500/15";
    case "failure":
    case "fail":
      return "bg-red-500/15";
    default:
      return "bg-gray-500/15";
  }
}

export function kindBadge(kind: string): string {
  switch (kind) {
    case "CALL": return "bg-blue-500/20 text-blue-400";
    case "DELEGATECALL": return "bg-purple-500/20 text-purple-400";
    case "STATICCALL": return "bg-cyan-500/20 text-cyan-400";
    case "EMIT": return "bg-yellow-500/20 text-yellow-400";
    case "REVERT": return "bg-red-500/20 text-red-400";
    case "CREATE": return "bg-green-500/20 text-green-400";
    default: return "bg-gray-500/20 text-gray-400";
  }
}

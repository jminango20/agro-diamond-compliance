import { useState } from "react";
import { Search } from "lucide-react";
import { formatGas } from "../lib/format";
import type { GasReportData } from "../types";

interface Props {
  report: GasReportData;
}

export function GasTab({ report }: Props) {
  const [search, setSearch] = useState("");

  const filtered = report.contracts
    .map((c) => ({
      ...c,
      functions: c.functions.filter(
        (f) =>
          f.name.toLowerCase().includes(search.toLowerCase()) ||
          c.contract.toLowerCase().includes(search.toLowerCase()),
      ),
    }))
    .filter((c) => c.functions.length > 0);

  const alertDiffs = report.diffs.filter((d) => d.alert);

  return (
    <div className="p-6 space-y-4">
      {alertDiffs.length > 0 && (
        <div className="bg-red-500/10 border border-red-500/30 rounded-lg p-4">
          <p className="text-sm font-semibold text-red-400 mb-2">
            {alertDiffs.length} function(s) exceeded {report.threshold}% threshold
          </p>
          <div className="space-y-1">
            {alertDiffs.map((d) => (
              <p key={`${d.contract}.${d.function}`} className="text-xs text-red-300">
                <span className="text-white">{d.contract}.{d.function}</span>{" "}
                {formatGas(d.previousAvg)} → {formatGas(d.currentAvg)}{" "}
                <span className="text-red-400">(+{d.deltaPct.toFixed(1)}%)</span>
              </p>
            ))}
          </div>
        </div>
      )}

      <div className="relative">
        <Search size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
        <input
          type="text"
          placeholder="Search functions..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-full pl-9 pr-3 py-2 text-xs bg-surface-card border border-surface-border rounded-lg text-gray-200 placeholder:text-gray-600 focus:outline-none focus:border-accent/50"
        />
      </div>

      {filtered.map((contract) => (
        <div key={contract.contract} className="bg-surface-card border border-surface-border rounded-lg overflow-hidden">
          <div className="px-4 py-3 border-b border-surface-border">
            <span className="text-sm font-medium text-white">{contract.contract}</span>
            <span className="text-xs text-gray-500 ml-2">({contract.functions.length} functions)</span>
          </div>
          <table className="w-full text-xs">
            <thead>
              <tr className="text-gray-500 border-b border-surface-border/50">
                <th className="text-left p-2.5 pl-4">Function</th>
                <th className="text-right p-2.5">Min</th>
                <th className="text-right p-2.5">Avg</th>
                <th className="text-right p-2.5">Median</th>
                <th className="text-right p-2.5">Max</th>
                <th className="text-right p-2.5 pr-4">Calls</th>
              </tr>
            </thead>
            <tbody>
              {contract.functions.map((f) => {
                const diff = report.diffs.find(
                  (d) => d.contract === contract.contract && d.function === f.name,
                );
                return (
                  <tr key={f.name} className="border-b border-surface-border/30 hover:bg-surface-hover/50">
                    <td className="p-2.5 pl-4 text-gray-200">
                      {f.name}
                      {diff && (
                        <span
                          className={`ml-2 text-[10px] ${
                            diff.alert ? "text-red-400" : diff.delta > 0 ? "text-yellow-400" : "text-emerald-400"
                          }`}
                        >
                          {diff.delta > 0 ? "+" : ""}
                          {diff.deltaPct.toFixed(1)}%
                        </span>
                      )}
                    </td>
                    <td className="p-2.5 text-right text-emerald-400/70">{formatGas(f.min)}</td>
                    <td className="p-2.5 text-right text-orange-400">{formatGas(f.avg)}</td>
                    <td className="p-2.5 text-right text-blue-400">{formatGas(f.median)}</td>
                    <td className="p-2.5 text-right text-red-400/70">{formatGas(f.max)}</td>
                    <td className="p-2.5 text-right text-gray-500 pr-4">{f.calls}</td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      ))}
    </div>
  );
}

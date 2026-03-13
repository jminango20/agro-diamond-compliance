import { useState, useMemo } from "react";
import { formatPct, coverageColor, coverageBg } from "../lib/format";
import type { CoverageData } from "../types";

interface Props {
  data: CoverageData;
}

export function CoverageTab({ data }: Props) {
  const [expandedFile, setExpandedFile] = useState<string | null>(null);

  const categories = useMemo(() => {
    const map = new Map<string, typeof data.files>();
    for (const f of data.files) {
      const list = map.get(f.category) ?? [];
      list.push(f);
      map.set(f.category, list);
    }
    return [...map.entries()].sort(([a], [b]) => a.localeCompare(b));
  }, [data]);

  return (
    <div className="p-6 space-y-6">
      {/* Summary */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div className={`rounded-lg p-4 ${coverageBg(data.linePct)} border border-surface-border`}>
          <p className="text-xs text-gray-500 uppercase">Line Coverage</p>
          <p className={`text-2xl font-bold ${coverageColor(data.linePct)}`}>
            {formatPct(data.linePct)}
          </p>
          <p className="text-xs text-gray-500 mt-1">
            {data.totalLinesHit}/{data.totalLines} lines
          </p>
        </div>
        <div className={`rounded-lg p-4 ${coverageBg(data.branchPct)} border border-surface-border`}>
          <p className="text-xs text-gray-500 uppercase">Branch Coverage</p>
          <p className={`text-2xl font-bold ${coverageColor(data.branchPct)}`}>
            {formatPct(data.branchPct)}
          </p>
          <p className="text-xs text-gray-500 mt-1">
            {data.totalBranchesHit}/{data.totalBranches} branches
          </p>
        </div>
      </div>

      {/* Facet Heatmap */}
      <div>
        <h3 className="text-sm font-semibold text-gray-400 mb-3">Facet Coverage Heatmap</h3>
        <div className="grid grid-cols-4 md:grid-cols-6 lg:grid-cols-8 gap-2">
          {data.files
            .filter((f) => f.category !== "Libraries" && f.category !== "Other")
            .sort((a, b) => b.linePct - a.linePct)
            .map((f) => (
              <button
                key={f.path}
                onClick={() => setExpandedFile(expandedFile === f.path ? null : f.path)}
                className={`p-2 rounded text-center text-[10px] border border-surface-border hover:border-accent/40 transition-colors ${
                  f.linePct >= 80
                    ? "bg-emerald-500/10"
                    : f.linePct >= 50
                      ? "bg-yellow-500/10"
                      : "bg-red-500/10"
                }`}
              >
                <p className="text-gray-300 truncate">{f.facet}</p>
                <p className={`font-bold ${coverageColor(f.linePct)}`}>
                  {formatPct(f.linePct)}
                </p>
              </button>
            ))}
        </div>
      </div>

      {/* Category Tables */}
      {categories.map(([category, files]) => (
        <div key={category} className="bg-surface-card border border-surface-border rounded-lg overflow-hidden">
          <div className="px-4 py-3 border-b border-surface-border">
            <span className="text-sm font-medium text-white">{category}</span>
            <span className="text-xs text-gray-500 ml-2">({files.length} files)</span>
          </div>
          <table className="w-full text-xs">
            <thead>
              <tr className="text-gray-500 border-b border-surface-border/50">
                <th className="text-left p-2.5 pl-4">File</th>
                <th className="text-right p-2.5">Lines</th>
                <th className="text-right p-2.5">Line %</th>
                <th className="text-right p-2.5">Branches</th>
                <th className="text-right p-2.5 pr-4">Branch %</th>
              </tr>
            </thead>
            <tbody>
              {files.sort((a, b) => a.linePct - b.linePct).map((f) => (
                <tr
                  key={f.path}
                  className="border-b border-surface-border/30 hover:bg-surface-hover/50 cursor-pointer"
                  onClick={() => setExpandedFile(expandedFile === f.path ? null : f.path)}
                >
                  <td className="p-2.5 pl-4 text-gray-200">{f.facet}</td>
                  <td className="p-2.5 text-right text-gray-400">
                    {f.linesHit}/{f.linesTotal}
                  </td>
                  <td className="p-2.5 text-right">
                    <span className={coverageColor(f.linePct)}>{formatPct(f.linePct)}</span>
                  </td>
                  <td className="p-2.5 text-right text-gray-400">
                    {f.branchesHit}/{f.branchesTotal}
                  </td>
                  <td className="p-2.5 text-right pr-4">
                    <span className={coverageColor(f.branchPct)}>{formatPct(f.branchPct)}</span>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          {files.some((f) => expandedFile === f.path) && (
            <div className="px-4 py-3 border-t border-surface-border bg-surface/50">
              {files
                .filter((f) => expandedFile === f.path)
                .map((f) => (
                  <div key={f.path}>
                    <p className="text-xs text-gray-400 mb-1">{f.path}</p>
                    {f.uncoveredLines.length > 0 && (
                      <p className="text-[10px] text-red-400/70">
                        Uncovered: {f.uncoveredLines.join(", ")}
                      </p>
                    )}
                  </div>
                ))}
            </div>
          )}
        </div>
      ))}
    </div>
  );
}

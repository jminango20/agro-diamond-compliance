import { StatCard } from "./StatCard";
import { formatPct, formatDuration, formatGas, coverageColor } from "../lib/format";
import type { TestReport, GasReportData, CoverageData, SelectorMapData, ProjectInfo } from "../types";

interface Props {
  testReport: TestReport | null;
  gasReport: GasReportData | null;
  coverage: CoverageData | null;
  selectorMap: SelectorMapData | null;
  project: ProjectInfo | null;
}

export function DashboardTab({ testReport, gasReport, coverage, selectorMap, project }: Props) {
  const hasData = testReport || gasReport || coverage;

  if (!hasData) {
    return (
      <div className="flex flex-col items-center justify-center h-full text-gray-500">
        <p className="text-lg font-semibold mb-2">No data yet</p>
        <p className="text-sm">Click "Run All" to execute tests and generate reports</p>
      </div>
    );
  }

  const topGas = gasReport?.contracts
    .flatMap((c) => c.functions.map((f) => ({ ...f, contract: c.contract })))
    .sort((a, b) => b.avg - a.avg)
    .slice(0, 10);

  return (
    <div className="space-y-6 p-6">
      {project && (
        <div className="flex items-center gap-4 bg-surface-card border border-surface-border rounded-lg px-4 py-3">
          <div className="flex-1 min-w-0">
            <h2 className="text-sm font-semibold text-white truncate">{project.name}</h2>
            <div className="flex items-center gap-3 mt-0.5 text-[11px] text-gray-500">
              {project.solcVersion && <span>Solidity {project.solcVersion}</span>}
              <span>{project.contractCount} contracts</span>
              <span>{project.isDiamond ? "Diamond Proxy" : "Standard"}</span>
              <span className="text-gray-700">{project.srcDir}/</span>
            </div>
          </div>
        </div>
      )}

      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        {testReport && (
          <>
            <StatCard
              label="Tests"
              value={testReport.totalTests}
              sub={formatDuration(testReport.totalDurationMs)}
            />
            <StatCard
              label="Pass Rate"
              value={formatPct(testReport.passRate)}
              color={testReport.passRate === 100 ? "text-emerald-400" : "text-yellow-400"}
            />
            <StatCard
              label="Passed"
              value={testReport.totalPassed}
              color="text-emerald-400"
            />
            <StatCard
              label="Failed"
              value={testReport.totalFailed}
              color={testReport.totalFailed > 0 ? "text-red-400" : "text-gray-500"}
            />
          </>
        )}
      </div>

      {coverage && (
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          <StatCard
            label="Line Coverage"
            value={formatPct(coverage.linePct)}
            color={coverageColor(coverage.linePct)}
            sub={`${coverage.totalLinesHit}/${coverage.totalLines} lines`}
          />
          <StatCard
            label="Branch Coverage"
            value={formatPct(coverage.branchPct)}
            color={coverageColor(coverage.branchPct)}
            sub={`${coverage.totalBranchesHit}/${coverage.totalBranches} branches`}
          />
          <StatCard label="Files" value={coverage.files.length} />
          {selectorMap && (
            <StatCard
              label="Selectors"
              value={selectorMap.totalSelectors}
              sub={`${selectorMap.facetCount} facets${selectorMap.collisions.length > 0 ? ` / ${selectorMap.collisions.length} collisions` : ""}`}
              color={selectorMap.collisions.length > 0 ? "text-red-400" : "text-white"}
            />
          )}
        </div>
      )}

      {gasReport && gasReport.alertsCount > 0 && (
        <div className="bg-red-500/10 border border-red-500/30 rounded-lg p-4">
          <p className="text-sm text-red-400 font-semibold">
            {gasReport.alertsCount} function(s) exceeded {gasReport.threshold}% gas threshold
          </p>
          <ul className="mt-2 text-xs text-red-300 space-y-1">
            {gasReport.diffs
              .filter((d) => d.alert)
              .map((d) => (
                <li key={`${d.contract}.${d.function}`}>
                  {d.contract}.{d.function}: {formatGas(d.previousAvg)} → {formatGas(d.currentAvg)}{" "}
                  (+{d.deltaPct.toFixed(1)}%)
                </li>
              ))}
          </ul>
        </div>
      )}

      {topGas && topGas.length > 0 && (
        <div>
          <h3 className="text-sm font-semibold text-gray-400 mb-3">Top 10 Gas Consumers</h3>
          <div className="bg-surface-card border border-surface-border rounded-lg overflow-hidden">
            <table className="w-full text-xs">
              <thead>
                <tr className="border-b border-surface-border text-gray-500">
                  <th className="text-left p-3">Function</th>
                  <th className="text-left p-3">Contract</th>
                  <th className="text-right p-3">Avg Gas</th>
                  <th className="text-right p-3">Max Gas</th>
                  <th className="text-right p-3">Calls</th>
                </tr>
              </thead>
              <tbody>
                {topGas.map((f, i) => (
                  <tr key={i} className="border-b border-surface-border/50 hover:bg-surface-hover">
                    <td className="p-3 text-white">{f.name}</td>
                    <td className="p-3 text-gray-400">{f.contract}</td>
                    <td className="p-3 text-right text-orange-400">{formatGas(f.avg)}</td>
                    <td className="p-3 text-right text-red-400">{formatGas(f.max)}</td>
                    <td className="p-3 text-right text-gray-500">{f.calls}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      )}
    </div>
  );
}

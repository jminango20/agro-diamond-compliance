import { useState, useMemo } from "react";
import { Search, ChevronDown, ChevronRight } from "lucide-react";
import { formatGas, formatDuration, statusColor, statusBg } from "../lib/format";
import { TestDetailPanel } from "./TestDetailPanel";
import type { TestReport, TestCase, TestSuite, TraceStepData, RunCommand } from "../types";

interface Props {
  report: TestReport;
  traceSteps: TraceStepData | null;
  running: boolean;
  connected: boolean;
  onRun: (cmd: RunCommand, args?: string[]) => void;
}

interface SelectedTest {
  test: TestCase;
  contractName: string;
}

export function TestsTab({ report, traceSteps, running, connected, onRun }: Props) {
  const [search, setSearch] = useState("");
  const [statusFilter, setStatusFilter] = useState<"all" | "success" | "failure" | "skipped">("all");
  const [expandedSuites, setExpandedSuites] = useState<Set<string>>(() => {
    // Auto-expand first suite or any suite with failures
    const initial = new Set<string>();
    for (const suite of report.suites) {
      if (suite.failed > 0) initial.add(suite.contract);
    }
    if (initial.size === 0 && report.suites.length > 0) {
      initial.add(report.suites[0].contract);
    }
    return initial;
  });
  const [selectedTest, setSelectedTest] = useState<SelectedTest | null>(null);

  const filtered = useMemo(() => {
    return report.suites
      .map((suite) => {
        const tests = suite.tests.filter((t) => {
          if (statusFilter !== "all" && t.status !== statusFilter) return false;
          if (search && !t.name.toLowerCase().includes(search.toLowerCase()) &&
              !suite.contract.toLowerCase().includes(search.toLowerCase())) return false;
          return true;
        });
        return { ...suite, tests };
      })
      .filter((s) => s.tests.length > 0);
  }, [report, search, statusFilter]);

  const toggleSuite = (contract: string) => {
    setExpandedSuites((prev) => {
      const next = new Set(prev);
      if (next.has(contract)) next.delete(contract);
      else next.add(contract);
      return next;
    });
  };

  const totalFiltered = filtered.reduce((s, suite) => s + suite.tests.length, 0);

  return (
    <>
      <div className="p-6 space-y-4">
        {/* Summary bar */}
        <div className="flex items-center gap-4 text-xs text-gray-500">
          <span className="text-white font-medium">{report.totalTests} tests</span>
          <span className="text-emerald-400">{report.totalPassed} passed</span>
          {report.totalFailed > 0 && (
            <span className="text-red-400">{report.totalFailed} failed</span>
          )}
          {report.totalSkipped > 0 && (
            <span>{report.totalSkipped} skipped</span>
          )}
          <span>{formatDuration(report.totalDurationMs)}</span>
          <span className="ml-auto text-gray-600">
            Click any test to see execution details
          </span>
        </div>

        {/* Filters */}
        <div className="flex items-center gap-3">
          <div className="relative flex-1">
            <Search size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
            <input
              type="text"
              placeholder="Search tests or contracts..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="w-full pl-9 pr-3 py-2 text-xs bg-surface-card border border-surface-border rounded-lg text-gray-200 placeholder:text-gray-600 focus:outline-none focus:border-accent/50"
            />
          </div>
          <div className="flex gap-1">
            {(["all", "success", "failure", "skipped"] as const).map((s) => (
              <button
                key={s}
                onClick={() => setStatusFilter(s)}
                className={`px-2.5 py-1.5 text-xs rounded transition-colors ${
                  statusFilter === s
                    ? "bg-accent/20 text-accent"
                    : "text-gray-500 hover:text-gray-300 hover:bg-surface-hover"
                }`}
              >
                {s === "all" ? `All (${report.totalTests})` : `${s.charAt(0).toUpperCase() + s.slice(1)} (${
                  s === "success" ? report.totalPassed : s === "failure" ? report.totalFailed : report.totalSkipped
                })`}
              </button>
            ))}
          </div>
        </div>

        {/* Suites */}
        <div className="space-y-2">
          {filtered.map((suite) => (
            <SuiteCard
              key={suite.contract}
              suite={suite}
              expanded={expandedSuites.has(suite.contract)}
              onToggle={() => toggleSuite(suite.contract)}
              onSelectTest={(test) =>
                setSelectedTest({ test, contractName: suite.contract })
              }
            />
          ))}
        </div>

        {filtered.length === 0 && search && (
          <div className="text-center py-12 text-gray-500">
            <p className="text-sm">No tests match "{search}"</p>
          </div>
        )}
      </div>

      {/* Detail panel overlay */}
      {selectedTest && (
        <TestDetailPanel
          test={selectedTest.test}
          contractName={selectedTest.contractName}
          traceSteps={traceSteps}
          running={running}
          connected={connected}
          onRun={onRun}
          onClose={() => setSelectedTest(null)}
        />
      )}
    </>
  );
}

function SuiteCard({
  suite,
  expanded,
  onToggle,
  onSelectTest,
}: {
  suite: TestSuite;
  expanded: boolean;
  onToggle: () => void;
  onSelectTest: (test: TestCase) => void;
}) {
  const passRate = suite.tests.length > 0
    ? (suite.passed / suite.tests.length) * 100
    : 0;

  return (
    <div className="bg-surface-card border border-surface-border rounded-lg overflow-hidden">
      <button
        onClick={onToggle}
        className="w-full flex items-center gap-3 p-3 hover:bg-surface-hover transition-colors"
      >
        <span className="text-gray-500">
          {expanded ? <ChevronDown size={14} /> : <ChevronRight size={14} />}
        </span>

        <span className="text-sm font-medium text-white">{suite.contract}</span>
        <span className="text-xs text-gray-600">{suite.file}</span>

        <div className="ml-auto flex items-center gap-3 text-xs">
          {/* Mini progress bar */}
          <div className="w-16 h-1.5 bg-surface-border rounded-full overflow-hidden">
            <div
              className={`h-full rounded-full ${
                passRate === 100 ? "bg-emerald-400" : passRate > 0 ? "bg-yellow-400" : "bg-red-400"
              }`}
              style={{ width: `${passRate}%` }}
            />
          </div>

          <span className="text-emerald-400">{suite.passed}</span>
          {suite.failed > 0 && <span className="text-red-400">{suite.failed}</span>}
          {suite.skipped > 0 && <span className="text-gray-500">{suite.skipped}</span>}
          <span className="text-gray-600">{formatDuration(suite.durationMs)}</span>
        </div>
      </button>

      {expanded && (
        <div className="border-t border-surface-border">
          {suite.tests.map((t) => (
            <TestRow
              key={t.name}
              test={t}
              onClick={() => onSelectTest(t)}
            />
          ))}
        </div>
      )}
    </div>
  );
}

function TestRow({ test, onClick }: { test: TestCase; onClick: () => void }) {
  return (
    <div
      onClick={onClick}
      className="flex items-center gap-3 px-4 py-2.5 border-b border-surface-border/30 hover:bg-accent/5 cursor-pointer transition-colors group"
    >
      {/* Status dot */}
      <div
        className={`w-2 h-2 rounded-full flex-none ${
          test.status === "success"
            ? "bg-emerald-400"
            : test.status === "failure"
              ? "bg-red-400"
              : "bg-gray-500"
        }`}
      />

      {/* Test name */}
      <div className="flex-1 min-w-0">
        <span className="text-xs text-gray-200 font-mono group-hover:text-accent transition-colors">
          {test.name}
        </span>
        {test.reason && (
          <p className="text-red-400/80 mt-0.5 text-[10px] truncate">{test.reason}</p>
        )}
      </div>

      {/* Status badge */}
      <span
        className={`flex-none px-1.5 py-0.5 rounded text-[10px] font-medium ${statusBg(test.status)} ${statusColor(test.status)}`}
      >
        {test.status}
      </span>

      {/* Kind */}
      <span className="flex-none text-[10px] text-gray-500 w-20 text-right">
        {test.kind.type}
        {test.kind.type === "Fuzz" && ` (${test.kind.runs})`}
        {test.kind.type === "Invariant" && ` (${test.kind.calls}c)`}
      </span>

      {/* Gas */}
      <span className="flex-none text-[10px] text-gray-400 w-20 text-right font-mono">
        {formatGas(test.gas)}
      </span>

      {/* Duration */}
      <span className="flex-none text-[10px] text-gray-600 w-14 text-right">
        {formatDuration(test.durationMs)}
      </span>

      {/* Hover arrow */}
      <ChevronRight
        size={12}
        className="flex-none text-gray-600 group-hover:text-accent transition-colors"
      />
    </div>
  );
}

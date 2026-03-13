import { useState, useMemo, useEffect, useCallback } from "react";
import {
  Search,
  Play,
  Loader2,
  CheckSquare,
  Square,
  ChevronDown,
  ChevronRight,
  FileCode,
  FlaskConical,
  X,
  RefreshCw,
  Filter,
  Minus,
} from "lucide-react";
import { formatGas, formatDuration, statusColor, statusBg } from "../lib/format";
import type { TestListData, TestReport, RunCommand } from "../types";

type RunMode = "contract" | "test";

interface Props {
  testList: TestListData | null;
  testReport: TestReport | null;
  running: boolean;
  connected: boolean;
  onRun: (cmd: RunCommand, args?: string[]) => void;
}

export function TestRunnerTab({ testList, testReport, running, connected, onRun }: Props) {
  const [mode, setMode] = useState<RunMode>("contract");
  const [search, setSearch] = useState("");
  const [selectedContracts, setSelectedContracts] = useState<Set<string>>(new Set());
  const [selectedTests, setSelectedTests] = useState<Set<string>>(new Set());
  const [expandedContracts, setExpandedContracts] = useState<Set<string>>(new Set());

  // Auto-discover tests on mount
  useEffect(() => {
    if (!testList && connected && !running) {
      onRun("list_tests");
    }
  }, [testList, connected, running, onRun]);

  const filteredContracts = useMemo(() => {
    if (!testList) return [];
    const q = search.toLowerCase();
    return testList.contracts
      .map((c) => ({
        ...c,
        tests: c.tests.filter((t) =>
          mode === "contract"
            ? c.name.toLowerCase().includes(q) || t.toLowerCase().includes(q)
            : t.toLowerCase().includes(q) || c.name.toLowerCase().includes(q)
        ),
      }))
      .filter((c) => c.tests.length > 0);
  }, [testList, search, mode]);

  const totalFilteredTests = useMemo(
    () => filteredContracts.reduce((s, c) => s + c.tests.length, 0),
    [filteredContracts]
  );

  const toggleContract = useCallback(
    (name: string) => {
      setSelectedContracts((prev) => {
        const next = new Set(prev);
        if (next.has(name)) {
          next.delete(name);
          // Also deselect all tests of this contract
          const contract = testList?.contracts.find((c) => c.name === name);
          if (contract) {
            setSelectedTests((pt) => {
              const nt = new Set(pt);
              contract.tests.forEach((t) => nt.delete(`${name}::${t}`));
              return nt;
            });
          }
        } else {
          next.add(name);
          // Also select all tests of this contract
          const contract = testList?.contracts.find((c) => c.name === name);
          if (contract) {
            setSelectedTests((pt) => {
              const nt = new Set(pt);
              contract.tests.forEach((t) => nt.add(`${name}::${t}`));
              return nt;
            });
          }
        }
        return next;
      });
    },
    [testList]
  );

  const toggleTest = useCallback(
    (contractName: string, testName: string) => {
      const key = `${contractName}::${testName}`;
      setSelectedTests((prev) => {
        const next = new Set(prev);
        if (next.has(key)) {
          next.delete(key);
        } else {
          next.add(key);
        }
        return next;
      });

      // Update contract selection based on children
      const contract = testList?.contracts.find((c) => c.name === contractName);
      if (contract) {
        setSelectedContracts((prev) => {
          const next = new Set(prev);
          // Check if any test of this contract is selected
          setSelectedTests((tests) => {
            const hasAny = contract.tests.some((t) => tests.has(`${contractName}::${t}`));
            if (hasAny) next.add(contractName);
            else next.delete(contractName);
            return tests;
          });
          return next;
        });
      }
    },
    [testList]
  );

  const toggleExpand = useCallback((name: string) => {
    setExpandedContracts((prev) => {
      const next = new Set(prev);
      if (next.has(name)) next.delete(name);
      else next.add(name);
      return next;
    });
  }, []);

  const selectAll = useCallback(() => {
    if (!testList) return;
    const allContracts = new Set(filteredContracts.map((c) => c.name));
    const allTests = new Set(
      filteredContracts.flatMap((c) => c.tests.map((t) => `${c.name}::${t}`))
    );
    setSelectedContracts(allContracts);
    setSelectedTests(allTests);
  }, [testList, filteredContracts]);

  const clearSelection = useCallback(() => {
    setSelectedContracts(new Set());
    setSelectedTests(new Set());
  }, []);

  const runSelected = useCallback(() => {
    if (mode === "contract" && selectedContracts.size > 0) {
      const pattern = [...selectedContracts].join("|");
      onRun("test_filtered", ["--match-contract", pattern]);
    } else if (mode === "test" && selectedTests.size > 0) {
      // Extract unique test names (without contract prefix)
      const testNames = [...selectedTests].map((k) => k.split("::")[1]);
      const uniqueTests = [...new Set(testNames)];
      const pattern = uniqueTests.join("|");
      // Also filter by contract if from same contracts
      const contracts = [...new Set([...selectedTests].map((k) => k.split("::")[0]))];
      const args = ["--match-test", pattern];
      if (contracts.length < (testList?.totalContracts ?? 999)) {
        args.push("--match-contract", contracts.join("|"));
      }
      onRun("test_filtered", args);
    }
  }, [mode, selectedContracts, selectedTests, onRun, testList]);

  const selectionCount =
    mode === "contract" ? selectedContracts.size : selectedTests.size;

  // Get results for a specific contract
  const getContractResult = useCallback(
    (contractName: string) => {
      if (!testReport) return null;
      return testReport.suites.find((s) => s.contract === contractName) ?? null;
    },
    [testReport]
  );

  // Get result for a specific test
  const getTestResult = useCallback(
    (contractName: string, testName: string) => {
      if (!testReport) return null;
      const suite = testReport.suites.find((s) => s.contract === contractName);
      return suite?.tests.find((t) => t.name === testName) ?? null;
    },
    [testReport]
  );

  // Check partial selection for contract
  const isContractPartial = useCallback(
    (contractName: string) => {
      const contract = testList?.contracts.find((c) => c.name === contractName);
      if (!contract) return false;
      const selected = contract.tests.filter((t) =>
        selectedTests.has(`${contractName}::${t}`)
      ).length;
      return selected > 0 && selected < contract.tests.length;
    },
    [testList, selectedTests]
  );

  if (!testList) {
    return (
      <div className="flex flex-col items-center justify-center h-full text-gray-500">
        {running ? (
          <>
            <Loader2 size={32} className="text-accent animate-spin mb-4" />
            <p className="text-sm text-accent">Discovering tests...</p>
          </>
        ) : (
          <>
            <FlaskConical size={48} className="text-gray-600 mb-4" />
            <p className="text-lg font-semibold mb-2">Test Runner</p>
            <p className="text-sm mb-4">Discover and run tests by contract or individually</p>
            <button
              onClick={() => onRun("list_tests")}
              disabled={!connected}
              className="flex items-center gap-2 px-4 py-2 text-sm rounded bg-accent hover:bg-accent-dim text-white disabled:opacity-40 transition-colors"
            >
              <RefreshCw size={14} />
              Discover Tests
            </button>
          </>
        )}
      </div>
    );
  }

  return (
    <div className="flex flex-col h-full">
      {/* Toolbar */}
      <div className="flex-none p-4 border-b border-surface-border space-y-3">
        {/* Mode toggle + stats */}
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <div className="flex bg-surface-card border border-surface-border rounded-lg overflow-hidden">
              <button
                onClick={() => setMode("contract")}
                className={`flex items-center gap-1.5 px-3 py-1.5 text-xs font-medium transition-colors ${
                  mode === "contract"
                    ? "bg-accent/20 text-accent"
                    : "text-gray-400 hover:text-gray-200"
                }`}
              >
                <FileCode size={13} />
                By Contract
              </button>
              <button
                onClick={() => setMode("test")}
                className={`flex items-center gap-1.5 px-3 py-1.5 text-xs font-medium transition-colors ${
                  mode === "test"
                    ? "bg-accent/20 text-accent"
                    : "text-gray-400 hover:text-gray-200"
                }`}
              >
                <FlaskConical size={13} />
                By Test
              </button>
            </div>

            <span className="text-xs text-gray-500 ml-2">
              {testList.totalContracts} contracts, {testList.totalTests} tests
            </span>
          </div>

          <div className="flex items-center gap-2">
            <button
              onClick={() => onRun("list_tests")}
              disabled={running}
              className="p-1.5 rounded text-gray-500 hover:text-gray-300 hover:bg-surface-hover disabled:opacity-40 transition-colors"
              title="Refresh test list"
            >
              <RefreshCw size={14} />
            </button>
          </div>
        </div>

        {/* Search + actions */}
        <div className="flex items-center gap-2">
          <div className="relative flex-1">
            <Search size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
            <input
              type="text"
              placeholder={
                mode === "contract" ? "Search contracts..." : "Search tests..."
              }
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="w-full pl-9 pr-3 py-2 text-xs bg-surface-card border border-surface-border rounded-lg text-gray-200 placeholder:text-gray-600 focus:outline-none focus:border-accent/50"
            />
          </div>

          <button
            onClick={selectAll}
            className="flex items-center gap-1 px-2.5 py-2 text-xs text-gray-400 hover:text-gray-200 rounded border border-surface-border hover:border-accent/40 transition-colors"
          >
            <CheckSquare size={12} />
            All
          </button>
          <button
            onClick={clearSelection}
            className="flex items-center gap-1 px-2.5 py-2 text-xs text-gray-400 hover:text-gray-200 rounded border border-surface-border hover:border-accent/40 transition-colors"
          >
            <X size={12} />
            Clear
          </button>
        </div>

        {/* Run button */}
        <div className="flex items-center justify-between">
          <div className="text-xs text-gray-500">
            {selectionCount > 0 ? (
              <span className="text-accent">
                {selectionCount} {mode === "contract" ? "contract(s)" : "test(s)"} selected
              </span>
            ) : (
              <span>Select {mode === "contract" ? "contracts" : "tests"} to run</span>
            )}
          </div>

          <button
            onClick={runSelected}
            disabled={running || !connected || selectionCount === 0}
            className="flex items-center gap-1.5 px-4 py-2 text-xs font-medium rounded bg-accent hover:bg-accent-dim disabled:opacity-40 disabled:cursor-not-allowed text-white transition-colors"
          >
            {running ? (
              <Loader2 size={14} className="animate-spin" />
            ) : (
              <Play size={14} />
            )}
            Run Selected
          </button>
        </div>
      </div>

      {/* Contract/Test list */}
      <div className="flex-1 overflow-y-auto p-4 space-y-1.5">
        {filteredContracts.map((contract) => {
          const isExpanded = expandedContracts.has(contract.name);
          const isSelected = selectedContracts.has(contract.name);
          const partial = isContractPartial(contract.name);
          const result = getContractResult(contract.name);

          return (
            <div
              key={contract.name}
              className="bg-surface-card border border-surface-border rounded-lg overflow-hidden"
            >
              {/* Contract row */}
              <div className="flex items-center">
                {/* Checkbox area */}
                <button
                  onClick={() => toggleContract(contract.name)}
                  className="flex items-center justify-center w-10 h-full py-3 text-gray-500 hover:text-accent transition-colors"
                >
                  {isSelected && !partial ? (
                    <CheckSquare size={16} className="text-accent" />
                  ) : partial ? (
                    <div className="relative">
                      <Square size={16} className="text-accent" />
                      <Minus size={10} className="absolute top-[3px] left-[3px] text-accent" />
                    </div>
                  ) : (
                    <Square size={16} />
                  )}
                </button>

                {/* Expand + contract info */}
                <button
                  onClick={() => toggleExpand(contract.name)}
                  className="flex-1 flex items-center gap-2 py-3 pr-3 hover:bg-surface-hover/50 transition-colors"
                >
                  {isExpanded ? (
                    <ChevronDown size={14} className="text-gray-500" />
                  ) : (
                    <ChevronRight size={14} className="text-gray-500" />
                  )}

                  <FileCode size={14} className="text-accent/60" />
                  <span className="text-sm font-medium text-white">{contract.name}</span>
                  <span className="text-xs text-gray-600">{contract.file}</span>

                  <span className="ml-auto text-xs text-gray-500">
                    {contract.tests.length} tests
                  </span>

                  {/* Inline result badges */}
                  {result && (
                    <div className="flex items-center gap-1.5 ml-2">
                      {result.passed > 0 && (
                        <span className="px-1.5 py-0.5 rounded text-[10px] font-medium bg-emerald-500/15 text-emerald-400">
                          {result.passed} pass
                        </span>
                      )}
                      {result.failed > 0 && (
                        <span className="px-1.5 py-0.5 rounded text-[10px] font-medium bg-red-500/15 text-red-400">
                          {result.failed} fail
                        </span>
                      )}
                      {result.skipped > 0 && (
                        <span className="px-1.5 py-0.5 rounded text-[10px] font-medium bg-gray-500/15 text-gray-500">
                          {result.skipped} skip
                        </span>
                      )}
                      <span className="text-[10px] text-gray-600">
                        {formatDuration(result.durationMs)}
                      </span>
                    </div>
                  )}
                </button>
              </div>

              {/* Expanded test list */}
              {isExpanded && (
                <div className="border-t border-surface-border/50">
                  {contract.tests.map((testName) => {
                    const testKey = `${contract.name}::${testName}`;
                    const isTestSelected = selectedTests.has(testKey);
                    const testResult = getTestResult(contract.name, testName);

                    return (
                      <div
                        key={testKey}
                        className="flex items-center hover:bg-surface-hover/30 transition-colors"
                      >
                        {/* Checkbox */}
                        <button
                          onClick={() => toggleTest(contract.name, testName)}
                          className="flex items-center justify-center w-10 py-2 ml-4 text-gray-500 hover:text-accent transition-colors"
                        >
                          {isTestSelected ? (
                            <CheckSquare size={14} className="text-accent" />
                          ) : (
                            <Square size={14} />
                          )}
                        </button>

                        {/* Test info */}
                        <div className="flex-1 flex items-center gap-2 py-2 pr-3">
                          <FlaskConical size={12} className="text-gray-600" />
                          <span className="text-xs text-gray-300 font-mono">
                            {testName}
                          </span>

                          {/* Test result inline */}
                          {testResult && (
                            <div className="ml-auto flex items-center gap-2">
                              <span
                                className={`inline-block px-1.5 py-0.5 rounded text-[10px] font-medium ${statusBg(testResult.status)} ${statusColor(testResult.status)}`}
                              >
                                {testResult.status}
                              </span>
                              {testResult.gas && (
                                <span className="text-[10px] text-gray-500">
                                  {formatGas(testResult.gas)} gas
                                </span>
                              )}
                              <span className="text-[10px] text-gray-600">
                                {formatDuration(testResult.durationMs)}
                              </span>
                              {testResult.reason && (
                                <span className="text-[10px] text-red-400 max-w-[300px] truncate">
                                  {testResult.reason}
                                </span>
                              )}
                            </div>
                          )}
                        </div>
                      </div>
                    );
                  })}
                </div>
              )}
            </div>
          );
        })}

        {filteredContracts.length === 0 && search && (
          <div className="text-center py-12 text-gray-500">
            <Filter size={24} className="mx-auto mb-2 opacity-50" />
            <p className="text-sm">No matches for "{search}"</p>
          </div>
        )}
      </div>

      {/* Results summary footer */}
      {testReport && (
        <div className="flex-none border-t border-surface-border bg-surface-card px-4 py-3">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-4 text-xs">
              <span className="text-gray-500">Last run:</span>
              <span className="text-white font-medium">{testReport.totalTests} tests</span>
              <span className="text-emerald-400">{testReport.totalPassed} passed</span>
              {testReport.totalFailed > 0 && (
                <span className="text-red-400">{testReport.totalFailed} failed</span>
              )}
              {testReport.totalSkipped > 0 && (
                <span className="text-gray-500">{testReport.totalSkipped} skipped</span>
              )}
              <span className="text-gray-500">
                {formatDuration(testReport.totalDurationMs)}
              </span>
            </div>
            <div className="flex items-center gap-1.5">
              <div
                className={`w-2 h-2 rounded-full ${
                  testReport.totalFailed > 0
                    ? "bg-red-400"
                    : testReport.passRate === 100
                      ? "bg-emerald-400"
                      : "bg-yellow-400"
                }`}
              />
              <span
                className={`text-xs font-medium ${
                  testReport.totalFailed > 0
                    ? "text-red-400"
                    : testReport.passRate === 100
                      ? "text-emerald-400"
                      : "text-yellow-400"
                }`}
              >
                {testReport.passRate.toFixed(1)}%
              </span>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

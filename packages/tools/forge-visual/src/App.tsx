import { useState, useEffect, useRef } from "react";
import { useForgeSocket } from "./hooks/useForgeSocket";
import { Sidebar, type Tab } from "./components/Sidebar";
import { Header } from "./components/Header";
import { DashboardTab } from "./components/DashboardTab";
import { TestRunnerTab } from "./components/TestRunnerTab";
import { TestsTab } from "./components/TestsTab";
import { GasTab } from "./components/GasTab";
import { CoverageTab } from "./components/CoverageTab";
import { TracesTab } from "./components/TracesTab";
import { SelectorsTab } from "./components/SelectorsTab";
import { LogsTab } from "./components/LogsTab";
import { RunningOverlay } from "./components/RunningOverlay";
import type { RunCommand } from "./types";

export function App() {
  const { state, run } = useForgeSocket();
  const [tab, setTab] = useState<Tab>("dashboard");
  const [sidebarCollapsed, setSidebarCollapsed] = useState(false);
  const prevRunning = useRef(false);

  // Auto-switch to logs when a command starts (except metadata/inline commands)
  useEffect(() => {
    if (state.running && !prevRunning.current) {
      const step = state.currentStep;
      const skipAutoSwitch =
        step === "list_tests" ||
        step === "trace_test" ||
        step === "Discovering tests..." ||
        (typeof step === "string" && step.startsWith("Tracing "));
      if (!skipAutoSwitch) {
        setTab("logs");
      }
    }
    prevRunning.current = state.running;
  }, [state.running, state.currentStep]);

  const hasResults = {
    tests: !!state.testReport,
    gas: !!state.gasReport,
    coverage: !!state.coverage,
    traces: !!state.traces,
    selectors: !!state.selectorMap,
  };

  const handleRun = (cmd: RunCommand, args?: string[]) => {
    run(cmd, args);
  };

  return (
    <div className="flex h-screen overflow-hidden">
      <Sidebar
        active={tab}
        onNavigate={setTab}
        collapsed={sidebarCollapsed}
        onToggle={() => setSidebarCollapsed(!sidebarCollapsed)}
        hasResults={hasResults}
        running={state.running}
        projectName={state.project?.name}
        isDiamond={state.project?.isDiamond}
      />

      <div className="flex-1 flex flex-col overflow-hidden">
        <Header
          connected={state.connected}
          running={state.running}
          currentStep={state.currentStep}
          gitRef={state.gitRef}
          gitBranch={state.gitBranch}
          onRun={handleRun}
          logCount={state.logs.length}
        />

        <main className="flex-1 overflow-y-auto relative">
          {state.error && (
            <div className="mx-6 mt-4 bg-red-500/10 border border-red-500/30 rounded-lg p-4">
              <p className="text-sm text-red-400">{state.error}</p>
            </div>
          )}

          {tab === "dashboard" && (
            <DashboardTab
              testReport={state.testReport}
              gasReport={state.gasReport}
              coverage={state.coverage}
              selectorMap={state.selectorMap}
              project={state.project}
            />
          )}

          {tab === "runner" && (
            <TestRunnerTab
              testList={state.testList}
              testReport={state.testReport}
              running={state.running}
              connected={state.connected}
              onRun={handleRun}
            />
          )}

          {tab === "tests" &&
            (state.testReport ? (
              <TestsTab
                report={state.testReport}
                traceSteps={state.traceSteps}
                running={state.running}
                connected={state.connected}
                onRun={handleRun}
              />
            ) : (
              <EmptyState
                label="Tests"
                onRun={() => handleRun("test")}
                running={state.running}
                currentStep={state.currentStep}
              />
            ))}

          {tab === "gas" &&
            (state.gasReport ? (
              <GasTab report={state.gasReport} />
            ) : (
              <EmptyState
                label="Gas Report"
                onRun={() => handleRun("gas")}
                running={state.running}
                currentStep={state.currentStep}
              />
            ))}

          {tab === "coverage" &&
            (state.coverage ? (
              <CoverageTab data={state.coverage} />
            ) : (
              <EmptyState
                label="Coverage"
                onRun={() => handleRun("coverage")}
                running={state.running}
                currentStep={state.currentStep}
              />
            ))}

          {tab === "traces" &&
            (state.traces && state.traces.length > 0 ? (
              <TracesTab traces={state.traces} />
            ) : (
              <EmptyState
                label="Traces"
                onRun={() => handleRun("trace")}
                running={state.running}
                currentStep={state.currentStep}
              />
            ))}

          {tab === "selectors" &&
            (state.selectorMap ? (
              <SelectorsTab data={state.selectorMap} />
            ) : (
              <EmptyState
                label="Selector Map"
                onRun={() => handleRun("selectors")}
                running={state.running}
                currentStep={state.currentStep}
              />
            ))}

          {tab === "logs" && (
            <LogsTab
              logs={state.logs}
              running={state.running}
              currentStep={state.currentStep}
            />
          )}

          {/* Floating running indicator when not on logs tab */}
          {state.running && tab !== "logs" && (
            <RunningOverlay
              currentStep={state.currentStep}
              logCount={state.logs.length}
              lastLog={state.logs[state.logs.length - 1] ?? null}
              onViewLogs={() => setTab("logs")}
            />
          )}
        </main>
      </div>
    </div>
  );
}

function EmptyState({
  label,
  onRun,
  running,
  currentStep,
}: {
  label: string;
  onRun: () => void;
  running: boolean;
  currentStep: string | null;
}) {
  return (
    <div className="flex flex-col items-center justify-center h-full text-gray-500">
      {running ? (
        <>
          <div className="w-8 h-8 border-2 border-accent/30 border-t-accent rounded-full animate-spin mb-4" />
          <p className="text-sm text-accent">{currentStep ?? "Running..."}</p>
        </>
      ) : (
        <>
          <p className="text-lg font-semibold mb-2">No {label} data</p>
          <button
            onClick={onRun}
            className="mt-2 px-4 py-2 text-sm rounded bg-accent hover:bg-accent-dim text-white transition-colors"
          >
            Run {label}
          </button>
        </>
      )}
    </div>
  );
}

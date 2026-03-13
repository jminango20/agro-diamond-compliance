import { Play, Loader2, Wifi, WifiOff, Terminal } from "lucide-react";
import type { RunCommand } from "../types";

interface HeaderProps {
  connected: boolean;
  running: boolean;
  currentStep: string | null;
  gitRef: string | null;
  gitBranch: string | null;
  onRun: (cmd: RunCommand, args?: string[]) => void;
  logCount: number;
}

export function Header({ connected, running, currentStep, gitRef, gitBranch, onRun, logCount }: HeaderProps) {
  return (
    <header className="flex items-center justify-between h-14 px-4 border-b border-surface-border bg-surface-card">
      <div className="flex items-center gap-3">
        {connected ? (
          <div className="flex items-center gap-1.5">
            <Wifi size={14} className="text-emerald-400" />
            <span className="text-[10px] text-emerald-400/70">connected</span>
          </div>
        ) : (
          <div className="flex items-center gap-1.5">
            <WifiOff size={14} className="text-red-400" />
            <span className="text-[10px] text-red-400/70">disconnected</span>
          </div>
        )}

        {gitBranch && (
          <span className="text-xs text-gray-500">
            {gitBranch}
            {gitRef && <span className="text-gray-600 ml-1">({gitRef})</span>}
          </span>
        )}

        {running && (
          <div className="flex items-center gap-2 px-3 py-1 rounded-full bg-accent/10 border border-accent/20">
            <Loader2 size={12} className="text-accent animate-spin" />
            <span className="text-xs text-accent font-medium">
              {currentStep ?? "Running..."}
            </span>
            {logCount > 0 && (
              <span className="flex items-center gap-1 text-[10px] text-accent/60">
                <Terminal size={10} />
                {logCount}
              </span>
            )}
          </div>
        )}
      </div>

      <div className="flex items-center gap-2">
        <button
          onClick={() => onRun("report")}
          disabled={running || !connected}
          className="flex items-center gap-1.5 px-3 py-1.5 text-xs font-medium rounded bg-accent hover:bg-accent-dim disabled:opacity-40 disabled:cursor-not-allowed text-white transition-colors"
        >
          {running ? <Loader2 size={12} className="animate-spin" /> : <Play size={12} />}
          Run All
        </button>
        <button
          onClick={() => onRun("test")}
          disabled={running || !connected}
          className="px-2.5 py-1.5 text-xs text-gray-400 hover:text-white rounded border border-surface-border hover:border-accent/40 disabled:opacity-40 transition-colors"
        >
          Tests
        </button>
        <button
          onClick={() => onRun("gas")}
          disabled={running || !connected}
          className="px-2.5 py-1.5 text-xs text-gray-400 hover:text-white rounded border border-surface-border hover:border-accent/40 disabled:opacity-40 transition-colors"
        >
          Gas
        </button>
        <button
          onClick={() => onRun("coverage")}
          disabled={running || !connected}
          className="px-2.5 py-1.5 text-xs text-gray-400 hover:text-white rounded border border-surface-border hover:border-accent/40 disabled:opacity-40 transition-colors"
        >
          Coverage
        </button>
      </div>
    </header>
  );
}

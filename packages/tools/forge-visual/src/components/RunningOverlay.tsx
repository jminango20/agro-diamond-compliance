import { Terminal } from "lucide-react";

interface Props {
  currentStep: string | null;
  logCount: number;
  lastLog: string | null;
  onViewLogs: () => void;
}

export function RunningOverlay({ currentStep, logCount, lastLog, onViewLogs }: Props) {
  return (
    <div className="fixed bottom-4 right-4 z-50">
      <button
        onClick={onViewLogs}
        className="flex items-start gap-3 bg-surface-card border border-accent/30 rounded-lg p-4 shadow-xl shadow-black/30 max-w-sm hover:border-accent/50 transition-colors"
      >
        <div className="mt-0.5">
          <div className="w-5 h-5 border-2 border-accent/30 border-t-accent rounded-full animate-spin" />
        </div>
        <div className="text-left min-w-0">
          <p className="text-xs font-medium text-accent">
            {currentStep ?? "Running..."}
          </p>
          {lastLog && (
            <p className="text-[10px] text-gray-500 mt-1 truncate max-w-[280px]">
              {lastLog}
            </p>
          )}
          <p className="text-[10px] text-gray-600 mt-1 flex items-center gap-1">
            <Terminal size={10} />
            {logCount} log lines — click to view
          </p>
        </div>
      </button>
    </div>
  );
}

import { useEffect, useRef } from "react";

interface Props {
  logs: string[];
  running: boolean;
  currentStep: string | null;
}

export function LogsTab({ logs, running, currentStep }: Props) {
  const endRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    endRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [logs.length]);

  return (
    <div className="p-6 h-full flex flex-col">
      {running && currentStep && (
        <div className="flex items-center gap-2 mb-3 text-xs text-accent">
          <span className="inline-block w-2 h-2 rounded-full bg-accent animate-pulse" />
          Running: {currentStep}
        </div>
      )}

      <div className="flex-1 bg-surface-card border border-surface-border rounded-lg p-4 overflow-y-auto font-mono text-[11px] leading-5">
        {logs.length === 0 && (
          <p className="text-gray-600">No logs yet. Run a command to see output.</p>
        )}
        {logs.map((line, i) => (
          <div
            key={i}
            className={`whitespace-pre-wrap ${
              line.includes("PASS") || line.includes("✓")
                ? "text-emerald-400"
                : line.includes("FAIL") || line.includes("✗") || line.includes("Error")
                  ? "text-red-400"
                  : line.includes("WARN")
                    ? "text-yellow-400"
                    : "text-gray-400"
            }`}
          >
            {line}
          </div>
        ))}
        <div ref={endRef} />
      </div>
    </div>
  );
}

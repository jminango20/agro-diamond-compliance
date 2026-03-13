import { useState, useEffect } from "react";
import {
  X,
  Loader2,
  Zap,
  Send,
  Eye,
  Bell,
  AlertTriangle,
  CheckCircle2,
  Settings,
  PlusCircle,
  ArrowLeftRight,
  Flame,
  Clock,
  Hash,
  RefreshCw,
  Copy,
  Check,
} from "lucide-react";
import { formatGas, formatDuration, statusColor, statusBg } from "../lib/format";
import type { TestCase, TraceStepData, HumanStep, RunCommand } from "../types";

interface Props {
  test: TestCase;
  contractName: string;
  traceSteps: TraceStepData | null;
  running: boolean;
  connected: boolean;
  onRun: (cmd: RunCommand, args?: string[]) => void;
  onClose: () => void;
}

const ICON_MAP: Record<string, typeof Zap> = {
  setup: Settings,
  call: Send,
  delegatecall: ArrowLeftRight,
  read: Eye,
  event: Bell,
  revert: AlertTriangle,
  check: CheckCircle2,
  create: PlusCircle,
  return: CheckCircle2,
};

const ICON_BG: Record<string, string> = {
  setup: "bg-slate-500/15 border-slate-500/25 text-slate-400",
  call: "bg-blue-500/15 border-blue-500/25 text-blue-400",
  delegatecall: "bg-purple-500/15 border-purple-500/25 text-purple-400",
  read: "bg-cyan-500/15 border-cyan-500/25 text-cyan-400",
  event: "bg-amber-500/15 border-amber-500/25 text-amber-400",
  revert: "bg-red-500/15 border-red-500/25 text-red-400",
  check: "bg-emerald-500/15 border-emerald-500/25 text-emerald-400",
  create: "bg-green-500/15 border-green-500/25 text-green-400",
  return: "bg-slate-500/15 border-slate-500/25 text-slate-400",
};

const STEP_LABEL: Record<string, string> = {
  setup: "Setup",
  call: "Function Call",
  delegatecall: "Delegate Call",
  read: "Read Data",
  event: "Event Emitted",
  revert: "Error",
  check: "Verification",
  create: "Deploy",
  return: "Return",
};

const LINE_COLOR: Record<string, string> = {
  setup: "bg-slate-500/30",
  call: "bg-blue-500/30",
  read: "bg-cyan-500/30",
  event: "bg-amber-500/30",
  revert: "bg-red-500/30",
  check: "bg-emerald-500/30",
  create: "bg-green-500/30",
};

export function TestDetailPanel({
  test,
  contractName,
  traceSteps,
  running,
  connected,
  onRun,
  onClose,
}: Props) {
  const isTracing = running && (traceSteps === null || traceSteps.testName !== test.name);
  const hasTrace = traceSteps !== null && traceSteps.testName === test.name;

  // Auto-trigger trace on open
  const [autoTriggered, setAutoTriggered] = useState(false);
  useEffect(() => {
    if (!hasTrace && !isTracing && !autoTriggered && connected && !running) {
      setAutoTriggered(true);
      onRun("trace_test", [test.name]);
    }
  }, [hasTrace, isTracing, autoTriggered, connected, running, onRun, test.name]);

  // Count step types
  const stepCounts = hasTrace && traceSteps
    ? traceSteps.steps.reduce((acc, s) => {
        acc[s.icon] = (acc[s.icon] || 0) + 1;
        return acc;
      }, {} as Record<string, number>)
    : {};

  // Close on Escape
  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      if (e.key === "Escape") onClose();
    };
    window.addEventListener("keydown", handler);
    return () => window.removeEventListener("keydown", handler);
  }, [onClose]);

  return (
    <div className="fixed inset-0 z-50 flex">
      <div className="absolute inset-0 bg-black/50 backdrop-blur-sm" onClick={onClose} />

      <div className="relative ml-auto w-full max-w-4xl bg-[#0c0f16] border-l border-surface-border flex flex-col overflow-hidden shadow-2xl">
        {/* Status bar */}
        <div
          className={`h-[3px] ${
            test.status === "success"
              ? "bg-gradient-to-r from-emerald-500 to-emerald-400"
              : test.status === "failure"
                ? "bg-gradient-to-r from-red-500 to-red-400"
                : "bg-gray-500"
          }`}
        />

        {/* Header */}
        <div className="flex-none px-6 py-4 border-b border-surface-border bg-surface-card/50">
          <div className="flex items-start justify-between">
            <div className="flex-1 min-w-0">
              <div className="flex items-center gap-2.5 mb-2">
                <span
                  className={`inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-[11px] font-semibold ${statusBg(test.status)} ${statusColor(test.status)}`}
                >
                  {test.status === "success" ? (
                    <CheckCircle2 size={12} />
                  ) : test.status === "failure" ? (
                    <AlertTriangle size={12} />
                  ) : null}
                  {test.status === "success" ? "PASSED" : test.status === "failure" ? "FAILED" : "SKIPPED"}
                </span>
                <span className="text-xs text-gray-500">{contractName}</span>
              </div>

              <h2 className="text-[15px] font-semibold text-white font-mono leading-tight mb-3">
                {test.name}
              </h2>

              <div className="flex items-center gap-5 text-xs">
                <div className="flex items-center gap-1.5 text-orange-400">
                  <Flame size={12} />
                  <span className="font-mono font-medium">{formatGas(test.gas)}</span>
                  <span className="text-orange-400/50">gas</span>
                </div>
                <div className="flex items-center gap-1.5 text-gray-400">
                  <Clock size={12} />
                  <span className="font-mono">{formatDuration(test.durationMs)}</span>
                </div>
                <div className="flex items-center gap-1.5 text-gray-400">
                  <Hash size={12} />
                  <span>{test.kind.type}</span>
                  {test.kind.type === "Fuzz" && <span className="text-gray-600">({test.kind.runs} runs)</span>}
                  {test.kind.type === "Invariant" && <span className="text-gray-600">({test.kind.calls} calls)</span>}
                </div>
              </div>

              {test.reason && (
                <div className="mt-3 flex items-start gap-2 p-3 bg-red-500/8 border border-red-500/20 rounded-lg">
                  <AlertTriangle size={14} className="text-red-400 flex-none mt-0.5" />
                  <p className="text-xs text-red-300 leading-relaxed">{test.reason}</p>
                </div>
              )}
            </div>

            <div className="flex items-center gap-2 ml-4 flex-none">
              {hasTrace && (
                <button
                  onClick={() => { setAutoTriggered(true); onRun("trace_test", [test.name]); }}
                  disabled={running || !connected}
                  className="p-1.5 rounded text-gray-500 hover:text-accent hover:bg-surface-hover disabled:opacity-40 transition-colors"
                  title="Re-run trace"
                >
                  <RefreshCw size={15} />
                </button>
              )}
              <button
                onClick={onClose}
                className="p-1.5 rounded text-gray-500 hover:text-gray-300 hover:bg-surface-hover transition-colors"
              >
                <X size={18} />
              </button>
            </div>
          </div>
        </div>

        {/* Content */}
        <div className="flex-1 overflow-y-auto">
          {isTracing && (
            <div className="flex flex-col items-center justify-center py-24">
              <div className="w-16 h-16 rounded-full border-2 border-accent/20 flex items-center justify-center mb-6">
                <Loader2 size={28} className="text-accent animate-spin" />
              </div>
              <p className="text-sm font-medium text-accent mb-1">Analyzing execution flow</p>
              <p className="text-xs text-gray-600">Running test with full trace...</p>
            </div>
          )}

          {hasTrace && traceSteps && (
            <div className="px-6 py-5">
              {/* Summary row */}
              <div className="grid grid-cols-4 gap-3 mb-5">
                <MiniCard label="Steps" value={String(traceSteps.steps.length)} color="text-white" />
                <MiniCard label="Gas Used" value={formatGas(traceSteps.gasTotal)} color="text-orange-400" mono />
                <MiniCard label="Events" value={String(stepCounts.event || 0)} color="text-amber-400" />
                <MiniCard
                  label="Result"
                  value={traceSteps.status === "pass" ? "Passed" : "Failed"}
                  color={traceSteps.status === "pass" ? "text-emerald-400" : "text-red-400"}
                />
              </div>

              {/* Type legend */}
              <div className="flex items-center gap-3 mb-5 flex-wrap">
                {Object.entries(stepCounts)
                  .sort(([, a], [, b]) => b - a)
                  .map(([type, count]) => {
                    const bg = ICON_BG[type] || "";
                    const textColor = bg.split(" ").find((c) => c.startsWith("text-")) || "text-gray-400";
                    return (
                      <span
                        key={type}
                        className={`inline-flex items-center gap-1.5 px-2 py-0.5 rounded-full text-[10px] font-medium ${bg}`}
                      >
                        {STEP_LABEL[type] || type} ({count})
                      </span>
                    );
                  })}
              </div>

              {/* Flat timeline */}
              <div className="relative">
                {traceSteps.steps.map((step, i) => (
                  <TimelineStep
                    key={step.id}
                    step={step}
                    index={i}
                    isLast={i === traceSteps.steps.length - 1}
                    totalGas={traceSteps.gasTotal}
                    totalSteps={traceSteps.steps.length}
                  />
                ))}
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

function MiniCard({ label, value, color, mono }: { label: string; value: string; color: string; mono?: boolean }) {
  return (
    <div className="bg-surface-card border border-surface-border rounded-lg px-3 py-2.5">
      <p className="text-[10px] text-gray-600 uppercase tracking-wider mb-0.5">{label}</p>
      <p className={`text-lg font-bold ${color} ${mono ? "font-mono text-base" : ""}`}>{value}</p>
    </div>
  );
}

function TimelineStep({
  step,
  index,
  isLast,
  totalGas,
  totalSteps,
}: {
  step: HumanStep;
  index: number;
  isLast: boolean;
  totalGas: number;
  totalSteps: number;
}) {
  const [showDetails, setShowDetails] = useState(false);
  const Icon = ICON_MAP[step.icon] || Zap;
  const iconBg = ICON_BG[step.icon] || ICON_BG.call;
  const lineColor = LINE_COLOR[step.icon] || "bg-surface-border";
  const gasPercent = totalGas > 0 ? (step.gasUsed / totalGas) * 100 : 0;

  return (
    <div className="relative flex gap-4">
      {/* Timeline column */}
      <div className="flex flex-col items-center flex-none w-10">
        <div className={`w-10 h-10 rounded-xl flex items-center justify-center border z-10 ${iconBg}`}>
          <Icon size={16} />
        </div>
        {!isLast && <div className={`flex-1 w-px min-h-[12px] ${lineColor}`} />}
      </div>

      {/* Content */}
      <div className="flex-1 min-w-0 pb-4">
        {/* Step label row */}
        <div className="flex items-center gap-2 mb-1.5">
          <span className="text-[10px] text-gray-600 font-mono w-5">
            {index + 1}
          </span>
          <span className={`text-[10px] font-semibold uppercase tracking-widest ${iconBg.split(" ").find((c) => c.startsWith("text-")) || "text-gray-400"}`}>
            {STEP_LABEL[step.icon] || step.icon}
          </span>

          {gasPercent > 0 && (
            <div className="flex items-center gap-1.5 ml-auto">
              <div className="w-16 h-1 bg-surface-border rounded-full overflow-hidden">
                <div
                  className={`h-full rounded-full ${
                    gasPercent > 40 ? "bg-orange-500" : gasPercent > 15 ? "bg-yellow-500/60" : "bg-accent/40"
                  }`}
                  style={{ width: `${Math.min(gasPercent, 100)}%` }}
                />
              </div>
              <span className="text-[10px] text-gray-600 font-mono">
                {formatGas(step.gasUsed)}
              </span>
            </div>
          )}
        </div>

        {/* Main content card */}
        <div
          className={`rounded-lg overflow-hidden transition-all ${
            step.isError
              ? "bg-red-500/5 border border-red-500/25"
              : step.isEvent
                ? "bg-amber-500/5 border border-amber-500/15"
                : "bg-surface-card border border-surface-border"
          } ${step.details ? "cursor-pointer" : ""}`}
          onClick={() => step.details && setShowDetails(!showDetails)}
        >
          <div className="px-4 py-3">
            {/* Title */}
            <h4 className={`text-[13px] font-medium leading-tight ${
              step.isError ? "text-red-300" : step.isEvent ? "text-amber-200" : "text-gray-100"
            }`}>
              {step.title}
            </h4>

            {/* Description */}
            {step.description && (
              <p className="text-xs text-gray-500 mt-1.5 leading-relaxed">
                {step.description}
              </p>
            )}
          </div>

          {/* Expandable raw details */}
          {showDetails && step.details && (
            <div className="border-t border-surface-border/30 px-4 py-2.5 bg-[#080a0f]">
              <div className="flex items-center justify-between mb-1">
                <span className="text-[10px] text-gray-600 uppercase tracking-wider">Raw Arguments</span>
                <CopyButton text={step.details} />
              </div>
              <pre className="text-[11px] text-gray-500 font-mono whitespace-pre-wrap break-all leading-relaxed">
                {step.details}
              </pre>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false);

  const handleCopy = (e: React.MouseEvent) => {
    e.stopPropagation();
    navigator.clipboard.writeText(text);
    setCopied(true);
    setTimeout(() => setCopied(false), 1500);
  };

  return (
    <button
      onClick={handleCopy}
      className="p-1 rounded text-gray-600 hover:text-gray-400 transition-colors"
    >
      {copied ? <Check size={11} className="text-emerald-400" /> : <Copy size={11} />}
    </button>
  );
}

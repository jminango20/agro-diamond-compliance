import { useState, useMemo } from "react";
import { formatGas, kindBadge, statusColor } from "../lib/format";
import type { TraceTest, TraceNode } from "../types";

interface Props {
  traces: TraceTest[];
}

export function TracesTab({ traces }: Props) {
  const [selectedTest, setSelectedTest] = useState<string>(traces[0]?.name ?? "");
  const [hideReturns, setHideReturns] = useState(true);
  const [hideVm, setHideVm] = useState(true);

  const currentTrace = useMemo(
    () => traces.find((t) => t.name === selectedTest),
    [traces, selectedTest],
  );

  return (
    <div className="p-6 space-y-4">
      <div className="flex items-center gap-3 flex-wrap">
        <select
          value={selectedTest}
          onChange={(e) => setSelectedTest(e.target.value)}
          className="px-3 py-1.5 text-xs bg-surface-card border border-surface-border rounded-lg text-gray-200 focus:outline-none focus:border-accent/50"
        >
          {traces.map((t) => (
            <option key={t.name} value={t.name}>
              {t.name} ({t.status}) — {formatGas(t.gas)} gas
            </option>
          ))}
        </select>
        <label className="flex items-center gap-1.5 text-xs text-gray-500">
          <input
            type="checkbox"
            checked={hideReturns}
            onChange={(e) => setHideReturns(e.target.checked)}
            className="rounded border-gray-600"
          />
          Hide Returns
        </label>
        <label className="flex items-center gap-1.5 text-xs text-gray-500">
          <input
            type="checkbox"
            checked={hideVm}
            onChange={(e) => setHideVm(e.target.checked)}
            className="rounded border-gray-600"
          />
          Hide VM
        </label>
      </div>

      {currentTrace && (
        <div className="space-y-3">
          <div className="flex items-center gap-3">
            <span className={`text-sm font-medium ${statusColor(currentTrace.status)}`}>
              [{currentTrace.status.toUpperCase()}]
            </span>
            <span className="text-sm text-white">{currentTrace.name}</span>
            <span className="text-xs text-gray-500">{formatGas(currentTrace.gas)} gas</span>
          </div>

          <div className="bg-surface-card border border-surface-border rounded-lg p-4 overflow-x-auto">
            <div className="space-y-0.5 font-mono text-[11px]">
              {currentTrace.nodes.map((node, i) => (
                <TraceNodeView
                  key={i}
                  node={node}
                  hideReturns={hideReturns}
                  hideVm={hideVm}
                  maxGas={currentTrace.gas}
                />
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

function TraceNodeView({
  node,
  hideReturns,
  hideVm,
  maxGas,
}: {
  node: TraceNode;
  hideReturns: boolean;
  hideVm: boolean;
  maxGas: number;
}) {
  const [collapsed, setCollapsed] = useState(false);

  if (hideReturns && (node.kind === "RETURN" || node.kind === "STOP")) return null;
  if (hideVm && node.kind === "VM") return null;

  const indent = node.depth * 20;
  const gasBarWidth = maxGas > 0 ? (node.gas / maxGas) * 100 : 0;

  return (
    <>
      <div
        className="flex items-center gap-2 py-0.5 hover:bg-surface-hover/50 rounded px-1 group cursor-pointer"
        style={{ paddingLeft: indent }}
        onClick={() => node.children.length > 0 && setCollapsed(!collapsed)}
      >
        {node.children.length > 0 && (
          <span className="text-gray-600 text-[10px] w-3">{collapsed ? "▸" : "▾"}</span>
        )}
        {node.children.length === 0 && <span className="w-3" />}

        <span className={`px-1 py-0.5 rounded text-[9px] font-medium ${kindBadge(node.kind)}`}>
          {node.kind}
        </span>

        {node.gas > 0 && (
          <span className="text-gray-600 text-[10px] min-w-[50px]">[{formatGas(node.gas)}]</span>
        )}

        {node.contract && (
          <span className="text-purple-400">{node.contract}::</span>
        )}
        <span className={node.isRevert ? "text-red-400" : node.isEmit ? "text-yellow-400" : "text-gray-200"}>
          {node.func}
        </span>
        {node.args && (
          <span className="text-gray-500 truncate max-w-[400px]">({node.args})</span>
        )}
        {node.returnData && (
          <span className="text-gray-600 truncate max-w-[200px]">→ {node.returnData}</span>
        )}

        {gasBarWidth > 1 && (
          <div className="ml-auto w-20 h-1 bg-surface-border rounded overflow-hidden opacity-0 group-hover:opacity-100 transition-opacity">
            <div
              className="h-full bg-orange-500/50 rounded"
              style={{ width: `${Math.min(gasBarWidth, 100)}%` }}
            />
          </div>
        )}
      </div>

      {!collapsed &&
        node.children.map((child, i) => (
          <TraceNodeView
            key={i}
            node={child}
            hideReturns={hideReturns}
            hideVm={hideVm}
            maxGas={maxGas}
          />
        ))}
    </>
  );
}

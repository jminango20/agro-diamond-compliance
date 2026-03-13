import { useCallback, useEffect, useRef, useState } from "react";
import type {
  WsMessage,
  RunCommand,
  TestReport,
  GasReportData,
  CoverageData,
  TraceTest,
  SelectorMapData,
  TestListData,
  TraceStepData,
  ProjectInfo,
} from "../types";

export interface ForgeState {
  connected: boolean;
  running: boolean;
  currentStep: string | null;
  logs: string[];
  testReport: TestReport | null;
  gasReport: GasReportData | null;
  coverage: CoverageData | null;
  traces: TraceTest[] | null;
  selectorMap: SelectorMapData | null;
  testList: TestListData | null;
  traceSteps: TraceStepData | null;
  error: string | null;
  gitRef: string | null;
  gitBranch: string | null;
  project: ProjectInfo | null;
}

const initialState: ForgeState = {
  connected: false,
  running: false,
  currentStep: null,
  logs: [],
  testReport: null,
  gasReport: null,
  coverage: null,
  traces: null,
  selectorMap: null,
  testList: null,
  traceSteps: null,
  error: null,
  gitRef: null,
  gitBranch: null,
  project: null,
};

export function useForgeSocket() {
  const [state, setState] = useState<ForgeState>(initialState);
  const wsRef = useRef<WebSocket | null>(null);
  const reconnectRef = useRef<ReturnType<typeof setTimeout> | undefined>(undefined);

  const connect = useCallback(() => {
    if (wsRef.current?.readyState === WebSocket.OPEN) return;

    // Connect directly to the backend server (port 3999), not through Vite proxy
    const wsHost = window.location.hostname;
    const wsPort = import.meta.env.DEV ? "3999" : window.location.port;
    const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
    const ws = new WebSocket(`${protocol}//${wsHost}:${wsPort}/ws`);
    wsRef.current = ws;

    ws.onopen = () => {
      setState((s) => ({ ...s, connected: true, error: null }));
    };

    ws.onclose = () => {
      setState((s) => ({ ...s, connected: false, running: false }));
      reconnectRef.current = setTimeout(connect, 2000);
    };

    ws.onerror = () => {
      ws.close();
    };

    ws.onmessage = (event) => {
      const msg: WsMessage = JSON.parse(event.data);
      setState((prev) => handleMessage(prev, msg));
    };
  }, []);

  useEffect(() => {
    connect();
    return () => {
      clearTimeout(reconnectRef.current);
      wsRef.current?.close();
    };
  }, [connect]);

  const run = useCallback((command: RunCommand, args?: string[]) => {
    if (!wsRef.current || wsRef.current.readyState !== WebSocket.OPEN) return;
    setState((s) => ({
      ...s,
      running: true,
      currentStep: command,
      logs: [],
      error: null,
    }));
    wsRef.current.send(JSON.stringify({ command, args }));
  }, []);

  const clearResults = useCallback(() => {
    setState((s) => ({
      ...s,
      testReport: null,
      gasReport: null,
      coverage: null,
      traces: null,
      selectorMap: null,
      logs: [],
      error: null,
      currentStep: null,
    }));
  }, []);

  return { state, run, clearResults };
}

function handleMessage(prev: ForgeState, msg: WsMessage): ForgeState {
  switch (msg.type) {
    case "status":
      return {
        ...prev,
        currentStep: (msg.data as { step: string }).step,
        running: true,
      };
    case "log":
      return {
        ...prev,
        logs: [...prev.logs.slice(-500), msg.data as string],
      };
    case "test_result":
      return { ...prev, testReport: msg.data as TestReport };
    case "gas_result":
      return { ...prev, gasReport: msg.data as GasReportData };
    case "coverage_result":
      return { ...prev, coverage: msg.data as CoverageData };
    case "trace_result":
      return { ...prev, traces: msg.data as TraceTest[] };
    case "selector_result":
      return { ...prev, selectorMap: msg.data as SelectorMapData };
    case "test_list":
      return { ...prev, testList: msg.data as TestListData };
    case "trace_step_result":
      return { ...prev, traceSteps: msg.data as TraceStepData };
    case "dashboard": {
      const d = msg.data as {
        gitRef: string | null;
        gitBranch: string | null;
        project?: ProjectInfo;
      };
      return {
        ...prev,
        gitRef: d.gitRef,
        gitBranch: d.gitBranch,
        ...(d.project ? { project: d.project } : {}),
      };
    }
    case "error":
      return { ...prev, error: msg.data as string, running: false };
    case "complete":
      return { ...prev, running: false, currentStep: null };
    default:
      return prev;
  }
}

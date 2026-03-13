// ── Test Report ──

export interface TestSuite {
  file: string;
  contract: string;
  facet: string | null;
  durationMs: number;
  tests: TestCase[];
  passed: number;
  failed: number;
  skipped: number;
}

export interface TestCase {
  name: string;
  status: "success" | "failure" | "skipped";
  reason: string | null;
  durationMs: number;
  gas: number | null;
  meanGas: number | null;
  medianGas: number | null;
  kind: TestKind;
}

export type TestKind =
  | { type: "Standard" }
  | { type: "Fuzz"; runs: number }
  | { type: "Invariant"; runs: number; calls: number; reverts: number };

export interface TestReport {
  suites: TestSuite[];
  totalTests: number;
  totalPassed: number;
  totalFailed: number;
  totalSkipped: number;
  totalDurationMs: number;
  passRate: number;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

// ── Gas Report ──

export interface FunctionGas {
  name: string;
  min: number;
  avg: number;
  median: number;
  max: number;
  calls: number;
}

export interface ContractGasReport {
  contract: string;
  functions: FunctionGas[];
}

export interface GasHistoryEntry {
  timestamp: string;
  gitRef: string | null;
  gitBranch: string | null;
  reports: ContractGasReport[];
}

export interface GasHistory {
  entries: GasHistoryEntry[];
}

export interface GasDiff {
  function: string;
  contract: string;
  previousAvg: number;
  currentAvg: number;
  delta: number;
  deltaPct: number;
  alert: boolean;
}

export interface GasReportData {
  contracts: ContractGasReport[];
  diffs: GasDiff[];
  hasPrevious: boolean;
  threshold: number;
  alertsCount: number;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

// ── Coverage ──

export interface FileCoverage {
  path: string;
  facet: string | null;
  category: string;
  linesHit: number;
  linesTotal: number;
  branchesHit: number;
  branchesTotal: number;
  linePct: number;
  branchPct: number;
  uncoveredLines: number[];
  level: "high" | "medium" | "low";
}

export interface CoverageData {
  files: FileCoverage[];
  totalLinesHit: number;
  totalLines: number;
  totalBranchesHit: number;
  totalBranches: number;
  linePct: number;
  branchPct: number;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

// ── Traces ──

export interface TraceTest {
  name: string;
  status: "pass" | "fail";
  gas: number;
  nodes: TraceNode[];
}

export interface TraceNode {
  depth: number;
  gas: number;
  kind: "CALL" | "DELEGATECALL" | "STATICCALL" | "CREATE" | "EMIT" | "RETURN" | "REVERT" | "STOP" | "VM" | "OTHER";
  contract: string;
  func: string;
  args: string;
  returnData: string | null;
  isRevert: boolean;
  isEmit: boolean;
  children: TraceNode[];
}

// ── Selector Map ──

export interface SelectorEntry {
  selector: string;
  signature: string;
  facet: string;
}

export interface SelectorMapData {
  entries: SelectorEntry[];
  collisions: { selector: string; facets: string[] }[];
  facetCount: number;
  totalSelectors: number;
}

// ── Dashboard ──

export interface DashboardData {
  testReport: TestReport | null;
  gasReport: GasReportData | null;
  coverage: CoverageData | null;
  selectorMap: SelectorMapData | null;
  generatedAt: string;
  gitRef: string | null;
  gitBranch: string | null;
}

export interface WsMessage {
  type: WsMessageType;
  data: unknown;
  timestamp: string;
}

export type RunCommand = "test" | "gas" | "coverage" | "trace" | "selectors" | "report" | "list_tests" | "test_filtered" | "trace_test";

// ── Humanized Trace ──

export interface TraceStepData {
  testName: string;
  status: "pass" | "fail";
  gasTotal: number;
  steps: HumanStep[];
}

export interface HumanStep {
  id: number;
  icon: "setup" | "call" | "delegatecall" | "read" | "event" | "revert" | "check" | "create" | "return";
  title: string;
  description: string;
  gasUsed: number;
  gasPercent: number;
  depth: number;
  isError: boolean;
  isEvent: boolean;
  details?: string;
  children: HumanStep[];
}

export interface RunRequest {
  command: RunCommand;
  args?: string[];
}

// ── Test Discovery ──

export interface DiscoveredTest {
  contract: string;
  file: string;
  name: string;
}

export interface TestListData {
  contracts: {
    name: string;
    file: string;
    tests: string[];
  }[];
  totalContracts: number;
  totalTests: number;
}

// ── Project Info ──

export interface ProjectInfo {
  name: string;
  solcVersion: string | null;
  srcDir: string;
  testDir: string;
  isDiamond: boolean;
  contractCount: number;
  remappings: string[];
}

// ── WebSocket Messages ──

export type WsMessageType =
  | "status"
  | "log"
  | "test_result"
  | "gas_result"
  | "coverage_result"
  | "trace_result"
  | "selector_result"
  | "dashboard"
  | "test_list"
  | "trace_step_result"
  | "error"
  | "complete";

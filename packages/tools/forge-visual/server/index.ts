import express from "express";
import { createServer } from "node:http";
import { WebSocketServer, WebSocket } from "ws";
import { spawn, execSync } from "node:child_process";
import { existsSync, readFileSync, mkdirSync, writeFileSync } from "node:fs";
import { join, resolve, dirname, basename } from "node:path";

// ── Config ──

const PORT = parseInt(process.env.PORT ?? "3999", 10);
const PROJECT_ROOT = findProjectRoot();
const REPORTS_DIR = join(PROJECT_ROOT, "reports");
const MAX_HISTORY = 50;
const PROJECT_INFO = detectProjectInfo();

console.log(`[forge-visual] Project: ${PROJECT_INFO.name} (${PROJECT_ROOT})`);

// ── Express + WebSocket ──

const app = express();
const server = createServer(app);
const wss = new WebSocketServer({ server, path: "/ws" });

app.get("/api/health", (_req, res) => {
  res.json({ ok: true, project: PROJECT_ROOT });
});

// ── WebSocket Handler ──

wss.on("connection", (ws) => {
  console.log("[ws] client connected");

  send(ws, "dashboard", {
    gitRef: gitRef(),
    gitBranch: gitBranch(),
    project: PROJECT_INFO,
  });

  ws.on("message", (raw) => {
    try {
      const msg = JSON.parse(raw.toString());
      handleCommand(ws, msg.command, msg.args ?? []);
    } catch (e) {
      send(ws, "error", `Invalid message: ${e}`);
    }
  });

  ws.on("close", () => console.log("[ws] client disconnected"));
});

// ── Command Router ──

async function handleCommand(ws: WebSocket, command: string, args: string[]) {
  try {
    switch (command) {
      case "test":
        await runTests(ws);
        break;
      case "gas":
        await runGas(ws);
        break;
      case "coverage":
        await runCoverage(ws);
        break;
      case "trace":
        await runTrace(ws);
        break;
      case "selectors":
        await runSelectors(ws);
        break;
      case "report":
        await runFullReport(ws);
        break;
      case "list_tests":
        await listTests(ws);
        break;
      case "test_filtered":
        await runTestsFiltered(ws, args);
        break;
      case "trace_test":
        await runTraceTest(ws, args);
        break;
      default:
        send(ws, "error", `Unknown command: ${command}`);
        return;
    }
    send(ws, "complete", null);
  } catch (e) {
    send(ws, "error", `${e}`);
  }
}

// ── Test Runner ──

async function runTests(ws: WebSocket) {
  send(ws, "status", { step: "Running tests..." });
  send(ws, "log", "[forge-visual] Running forge test (streaming output) ...");

  // Phase 1: Run verbose tests with live streaming so the user sees progress
  await forgeExecMerged(ws, ["test", "-v"]);

  // Phase 2: Run JSON mode (fast — already compiled/cached) to get structured data
  send(ws, "status", { step: "Parsing test results..." });
  send(ws, "log", "[forge-visual] Parsing structured results ...");
  const jsonOutput = await forgeExecQuiet(["test", "--json"]);
  const suites = parseForgeTestJson(jsonOutput);

  const totalTests = suites.reduce((s, suite) => s + suite.tests.length, 0);
  const totalPassed = suites.reduce((s, suite) => s + suite.passed, 0);
  const totalFailed = suites.reduce((s, suite) => s + suite.failed, 0);
  const totalSkipped = suites.reduce((s, suite) => s + suite.skipped, 0);
  const totalDurationMs = suites.reduce((s, suite) => s + suite.durationMs, 0);

  const report = {
    suites,
    totalTests,
    totalPassed,
    totalFailed,
    totalSkipped,
    totalDurationMs,
    passRate: totalTests > 0 ? (totalPassed / totalTests) * 100 : 0,
    generatedAt: now(),
    gitRef: gitRef(),
    gitBranch: gitBranch(),
  };

  send(ws, "log", `[forge-visual] Tests complete: ${report.totalPassed} passed, ${report.totalFailed} failed, ${report.totalSkipped} skipped`);
  send(ws, "test_result", report);
}

// ── Test Discovery ──

async function listTests(ws: WebSocket) {
  send(ws, "status", { step: "Discovering tests..." });
  send(ws, "log", "[forge-visual] Running forge test --list --json ...");

  const output = await forgeExecQuiet(["test", "--list", "--json"]);
  const raw = JSON.parse(output);

  // Format: { "file.sol": { "ContractName": ["test1", "test2"] } }
  const contractMap = new Map<string, { file: string; tests: string[] }>();

  for (const [file, fileData] of Object.entries(raw) as [string, any][]) {
    for (const [contractName, tests] of Object.entries(fileData) as [string, any][]) {
      const testNames = Array.isArray(tests) ? tests : Object.keys(tests);
      if (testNames.length === 0) continue;
      contractMap.set(contractName, { file, tests: testNames.sort() });
    }
  }

  const contracts = [...contractMap.entries()]
    .map(([name, data]) => ({ name, file: data.file, tests: data.tests }))
    .sort((a, b) => a.name.localeCompare(b.name));

  const totalTests = contracts.reduce((s, c) => s + c.tests.length, 0);

  send(ws, "test_list", {
    contracts,
    totalContracts: contracts.length,
    totalTests,
  });

  send(ws, "log", `[forge-visual] Found ${contracts.length} contracts, ${totalTests} tests`);
}

// ── Filtered Test Runner ──

async function runTestsFiltered(ws: WebSocket, args: string[]) {
  // args format: ["--match-contract", "Foo|Bar", "--match-test", "test_something"]
  // or just forge test flags
  const forgeArgs = ["test", "-v", ...args];

  send(ws, "status", { step: "Running filtered tests..." });
  send(ws, "log", `[forge-visual] forge ${forgeArgs.join(" ")}`);

  // Phase 1: stream verbose output
  await forgeExecMerged(ws, forgeArgs);

  // Phase 2: get structured JSON (same filters)
  send(ws, "status", { step: "Parsing results..." });
  const jsonArgs = ["test", "--json", ...args];
  const jsonOutput = await forgeExecQuiet(jsonArgs);
  const suites = parseForgeTestJson(jsonOutput);

  const totalTests = suites.reduce((s, suite) => s + suite.tests.length, 0);
  const totalPassed = suites.reduce((s, suite) => s + suite.passed, 0);
  const totalFailed = suites.reduce((s, suite) => s + suite.failed, 0);
  const totalSkipped = suites.reduce((s, suite) => s + suite.skipped, 0);
  const totalDurationMs = suites.reduce((s, suite) => s + suite.durationMs, 0);

  const report = {
    suites,
    totalTests,
    totalPassed,
    totalFailed,
    totalSkipped,
    totalDurationMs,
    passRate: totalTests > 0 ? (totalPassed / totalTests) * 100 : 0,
    generatedAt: now(),
    gitRef: gitRef(),
    gitBranch: gitBranch(),
  };

  send(ws, "log", `[forge-visual] Filtered tests: ${totalPassed} passed, ${totalFailed} failed, ${totalSkipped} skipped`);
  send(ws, "test_result", report);
}

// ── Single Test Trace (Humanized) ──

async function runTraceTest(ws: WebSocket, args: string[]) {
  const testName = args[0];
  if (!testName) {
    send(ws, "error", "Test name required");
    return;
  }

  send(ws, "status", { step: `Tracing ${testName}...` });
  send(ws, "log", `[forge-visual] forge test --match-test ${testName} -vvvv`);

  const output = await forgeExecMerged(ws, ["test", "--match-test", testName, "-vvvv"]);
  const traces = parseTraceOutput(output);

  // Find the exact test
  const trace = traces.find((t) => t.name === testName) ?? traces[0];
  if (!trace) {
    send(ws, "error", `No trace found for ${testName}`);
    return;
  }

  const steps = humanizeTrace(trace.nodes, trace.gas);
  const result = {
    testName: trace.name,
    status: trace.status,
    gasTotal: trace.gas,
    steps,
  };

  send(ws, "trace_step_result", result);
  send(ws, "log", `[forge-visual] Trace complete: ${steps.length} steps`);
}

interface HumanStep {
  id: number;
  icon: string;
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

let stepCounter = 0;

function humanizeTrace(nodes: TraceNode[], totalGas: number): HumanStep[] {
  stepCounter = 0;
  // Flatten: walk the full tree and produce a linear timeline
  // promoting events, reads, reverts to top-level for visibility
  const flat: HumanStep[] = [];
  for (const node of nodes) {
    flattenNode(node, totalGas, flat);
  }
  return flat;
}

/**
 * Recursively walks the trace tree and emits steps into a flat list.
 * Diamond::fallback → delegatecall patterns are collapsed into one "call" step.
 * Events, reads, and reverts inside calls are promoted to top-level.
 */
function flattenNode(node: TraceNode, totalGas: number, out: HumanStep[]): void {
  // Skip returns/stops — noise
  if (node.kind === "RETURN" || node.kind === "STOP") return;

  // VM ops
  if (node.kind === "VM" || node.contract === "VM") {
    stepCounter++;
    const vmStep = humanizeVmOp(node, stepCounter, 0);
    if (vmStep) out.push(vmStep);
    return;
  }

  // Events → always top-level
  if (node.kind === "EMIT" || node.isEmit) {
    stepCounter++;
    out.push({
      id: stepCounter,
      icon: "event",
      title: humanizeEventName(node.func),
      description: humanizeEventArgs(node.func, node.args),
      gasUsed: node.gas,
      gasPercent: 0,
      depth: 0,
      isError: false,
      isEvent: true,
      details: node.args || undefined,
      children: [],
    });
    return;
  }

  // Reverts → always top-level
  if (node.kind === "REVERT" || node.isRevert) {
    stepCounter++;
    out.push({
      id: stepCounter,
      icon: "revert",
      title: "Transaction Reverted",
      description: humanizeRevert(node.func, node.args, node.returnData),
      gasUsed: node.gas,
      gasPercent: 0,
      depth: 0,
      isError: true,
      isEvent: false,
      details: node.returnData || undefined,
      children: [],
    });
    return;
  }

  // Static calls (reads) → top-level with return value extracted
  if (node.kind === "STATICCALL") {
    stepCounter++;
    const returnChild = findDeepReturn(node.children);
    const returnVal = returnChild?.returnData;
    // Resolve the actual function: if it's Diamond::fallback with a delegatecall child, use that
    let contract = node.contract;
    let func = node.func;
    const delegateChild = node.children.find((c) => c.kind === "DELEGATECALL");
    if ((node.contract === "Diamond" && node.func === "fallback") && delegateChild) {
      contract = delegateChild.contract;
      func = delegateChild.func;
    }
    out.push({
      id: stepCounter,
      icon: "read",
      title: humanizeReadCall(contract, func),
      description: returnVal
        ? `Result: ${humanizeReturnValue(func, returnVal)}`
        : humanizeCallDescription(contract, func, node.args),
      gasUsed: node.gas,
      gasPercent: totalGas > 0 ? (node.gas / totalGas) * 100 : 0,
      depth: 0,
      isError: false,
      isEvent: false,
      details: node.args || undefined,
      children: [],
    });
    return;
  }

  // Diamond::fallback → delegatecall → collapse into one step, then recurse children
  const isDiamondFallback = node.contract === "Diamond" && node.func === "fallback";
  if (isDiamondFallback || node.kind === "DELEGATECALL") {
    // Find the actual function being called
    let realContract = node.contract;
    let realFunc = node.func;
    let realArgs = node.args;
    let realChildren = node.children;

    if (isDiamondFallback) {
      const dc = node.children.find((c) => c.kind === "DELEGATECALL");
      if (dc) {
        realContract = dc.contract;
        realFunc = dc.func;
        realArgs = dc.args;
        realChildren = dc.children;
      }
    }

    // Emit the call step
    stepCounter++;
    const gasPercent = totalGas > 0 ? (node.gas / totalGas) * 100 : 0;
    out.push({
      id: stepCounter,
      icon: "call",
      title: humanizeFunctionCall(realContract, realFunc),
      description: humanizeCallDescription(realContract, realFunc, realArgs),
      gasUsed: node.gas,
      gasPercent,
      depth: 0,
      isError: false,
      isEvent: false,
      details: realArgs || undefined,
      children: [], // children are flattened to top-level
    });

    // Recurse children → events, sub-calls, reads will be promoted
    for (const child of realChildren) {
      flattenNode(child, totalGas, out);
    }
    return;
  }

  // CREATE
  if (node.kind === "CREATE") {
    stepCounter++;
    out.push({
      id: stepCounter,
      icon: "create",
      title: `Deploy ${node.contract || "Contract"}`,
      description: "New contract deployed",
      gasUsed: node.gas,
      gasPercent: totalGas > 0 ? (node.gas / totalGas) * 100 : 0,
      depth: 0,
      isError: false,
      isEvent: false,
      children: [],
    });
    for (const child of node.children) {
      flattenNode(child, totalGas, out);
    }
    return;
  }

  // Generic CALL — emit step, then recurse children
  stepCounter++;
  const gasPercent = totalGas > 0 ? (node.gas / totalGas) * 100 : 0;
  out.push({
    id: stepCounter,
    icon: "call",
    title: humanizeFunctionCall(node.contract, node.func),
    description: humanizeCallDescription(node.contract, node.func, node.args),
    gasUsed: node.gas,
    gasPercent,
    depth: 0,
    isError: false,
    isEvent: false,
    details: node.args || undefined,
    children: [],
  });

  for (const child of node.children) {
    flattenNode(child, totalGas, out);
  }
}

/** Walk children recursively to find first Return node with data */
function findDeepReturn(nodes: TraceNode[]): TraceNode | null {
  for (const n of nodes) {
    if ((n.kind === "RETURN" || n.kind === "STOP") && n.returnData) return n;
    if (n.kind === "DELEGATECALL" || n.kind === "CALL") {
      const found = findDeepReturn(n.children);
      if (found) return found;
    }
  }
  return null;
}

// ── Humanization Helpers ──

function humanizeVmOp(node: TraceNode, id: number, _gasPercent: number): HumanStep | null {
  const func = node.func || "";
  const args = node.args || "";

  if (func === "prank" || func.includes("prank")) {
    const who = extractLabel(args);
    return {
      id,
      icon: "setup",
      title: "Switch Caller",
      description: `Next call will be made as ${who}`,
      gasUsed: 0,
      gasPercent: 0,
      depth: node.depth,
      isError: false,
      isEvent: false,
      children: [],
    };
  }

  if (func === "expectRevert" || func.includes("expectRevert")) {
    return {
      id,
      icon: "check",
      title: "Expect Error",
      description: "Next call should fail with an error (this is expected behavior)",
      gasUsed: 0,
      gasPercent: 0,
      depth: node.depth,
      isError: false,
      isEvent: false,
      children: [],
    };
  }

  if (func === "expectEmit" || func.includes("expectEmit")) {
    return {
      id,
      icon: "check",
      title: "Expect Event",
      description: "Verifying that the next operation emits the correct event",
      gasUsed: 0,
      gasPercent: 0,
      depth: node.depth,
      isError: false,
      isEvent: false,
      children: [],
    };
  }

  if (func === "startPrank") {
    const who = extractLabel(args);
    return {
      id,
      icon: "setup",
      title: "Impersonate User",
      description: `All following calls will be made as ${who}`,
      gasUsed: 0,
      gasPercent: 0,
      depth: node.depth,
      isError: false,
      isEvent: false,
      children: [],
    };
  }

  if (func === "stopPrank") {
    return {
      id,
      icon: "setup",
      title: "Stop Impersonation",
      description: "Returning to default caller",
      gasUsed: 0,
      gasPercent: 0,
      depth: node.depth,
      isError: false,
      isEvent: false,
      children: [],
    };
  }

  if (func === "deal") {
    return {
      id,
      icon: "setup",
      title: "Set Balance",
      description: `Setting ETH/token balance for an address`,
      gasUsed: 0,
      gasPercent: 0,
      depth: node.depth,
      isError: false,
      isEvent: false,
      children: [],
    };
  }

  if (func === "label") {
    return null; // Just labeling, not interesting
  }

  if (func === "assertEq" || func === "assertTrue" || func === "assertFalse" || func.startsWith("assert")) {
    return {
      id,
      icon: "check",
      title: "Verify Result",
      description: `Asserting that the result matches expected value`,
      gasUsed: 0,
      gasPercent: 0,
      depth: node.depth,
      isError: false,
      isEvent: false,
      children: [],
    };
  }

  // Generic VM op — skip if not interesting
  return null;
}

function extractLabel(args: string): string {
  // "owner: [0x7c89...]" → "owner"
  // "[0x7c89...]" → short address
  const m = args.match(/^(\w+):\s*\[/);
  if (m) return m[1];
  const addr = args.match(/\[?(0x[0-9a-fA-F]{4,})/);
  if (addr) return addr[1].slice(0, 6) + "..." + addr[1].slice(-4);
  return args || "unknown";
}

function humanizeEventName(name: string): string {
  // TransferSingle → "Token Transfer"
  // Minted → "Token Minted"
  // ComplianceModuleAdded → "Compliance Module Added"
  const map: Record<string, string> = {
    TransferSingle: "Token Transfer",
    TransferBatch: "Batch Transfer",
    Minted: "Tokens Minted",
    Burned: "Tokens Burned",
    ForcedTransfer: "Forced Transfer",
    ComplianceModuleAdded: "Compliance Module Added",
    ComplianceModuleRemoved: "Compliance Module Removed",
    AssetRegistered: "Asset Registered",
    AssetConfigUpdated: "Asset Config Updated",
    WalletFrozen: "Wallet Frozen",
    WalletUnfrozen: "Wallet Unfrozen",
    AssetFrozen: "Asset Frozen",
    AssetUnfrozen: "Asset Unfrozen",
    PartialFreeze: "Partial Token Freeze",
    IdentityRegistered: "Identity Registered",
    IdentityDeleted: "Identity Deleted",
    OwnershipTransferred: "Ownership Transferred",
    RoleGranted: "Role Granted",
    RoleRevoked: "Role Revoked",
    EmergencyPause: "Emergency Pause",
    Paused: "Protocol Paused",
    Unpaused: "Protocol Unpaused",
    ApprovalForAll: "Operator Approved",
    URI: "Metadata URI Updated",
    DividendCreated: "Dividend Created",
    DividendClaimed: "Dividend Claimed",
    SnapshotCreated: "Snapshot Created",
  };
  return map[name] || splitCamelCase(name);
}

function humanizeEventArgs(name: string, args: string): string {
  if (!args) return "";

  // Try to build human-readable event description
  if (name === "TransferSingle") {
    const parts = parseEventArgs(args);
    const from = parts.from || parts[1];
    const to = parts.to || parts[2];
    const amount = parts.value || parts.amount || parts[4];
    if (from?.includes("0x000000")) {
      return `Minted ${amount || ""} tokens to ${extractLabel(to || "")}`;
    }
    if (to?.includes("0x000000")) {
      return `Burned ${amount || ""} tokens from ${extractLabel(from || "")}`;
    }
    return `${extractLabel(from || "")} sent ${amount || ""} tokens to ${extractLabel(to || "")}`;
  }

  if (name === "Minted") {
    const parts = parseEventArgs(args);
    const to = parts.to || parts[1];
    const amount = parts.amount || parts[2];
    return `${amount || ""} tokens minted to ${extractLabel(to || "")}`;
  }

  if (name === "Burned") {
    const parts = parseEventArgs(args);
    const from = parts.from || parts[1];
    const amount = parts.amount || parts[2];
    return `${amount || ""} tokens burned from ${extractLabel(from || "")}`;
  }

  // Generic: just list named params
  return args.length > 120 ? args.slice(0, 120) + "..." : args;
}

function parseEventArgs(args: string): Record<string, string> {
  const result: Record<string, string> = {};
  let idx = 0;
  for (const part of args.split(",")) {
    const trimmed = part.trim();
    const m = trimmed.match(/^(\w+):\s*(.+)$/);
    if (m) {
      result[m[1]] = m[2].trim();
    }
    result[String(idx)] = trimmed;
    idx++;
  }
  return result;
}

function humanizeFunctionCall(contract: string, func: string): string {
  if (!contract || !func) return func || contract || "Unknown Call";

  // Common patterns
  if (func === "mint") return `Mint Tokens (${contract})`;
  if (func === "burn") return `Burn Tokens (${contract})`;
  if (func === "forcedTransfer") return `Forced Transfer (${contract})`;
  if (func === "safeTransferFrom") return `Transfer Tokens (${contract})`;
  if (func === "safeBatchTransferFrom") return `Batch Transfer (${contract})`;
  if (func === "balanceOf") return `Check Balance (${contract})`;
  if (func === "totalSupply") return `Check Total Supply (${contract})`;
  if (func === "holderCount") return `Check Holder Count (${contract})`;
  if (func === "isHolder") return `Check If Holder (${contract})`;
  if (func === "canTransfer") return `Check Transfer Permission (${contract})`;
  if (func === "isVerified") return `Check Identity Verified (${contract})`;
  if (func === "setApprovalForAll") return `Set Operator Approval (${contract})`;
  if (func === "addComplianceModule") return `Add Compliance Module (${contract})`;
  if (func === "registerAsset") return `Register Asset (${contract})`;
  if (func === "grantRole") return `Grant Role (${contract})`;
  if (func === "revokeRole") return `Revoke Role (${contract})`;
  if (func === "minted") return `Notify Compliance: Minted (${contract})`;
  if (func === "burned") return `Notify Compliance: Burned (${contract})`;
  if (func === "transferred") return `Notify Compliance: Transferred (${contract})`;
  if (func === "mintedCount") return `Read Minted Count (${contract})`;
  if (func === "freezeWallet") return `Freeze Wallet (${contract})`;
  if (func === "unfreezeWallet") return `Unfreeze Wallet (${contract})`;
  if (func === "recoverWallet") return `Recover Wallet (${contract})`;
  if (func === "fallback") return `Call Diamond Proxy`;

  return `${splitCamelCase(func)} (${contract})`;
}

function humanizeCallDescription(_contract: string, func: string, args: string): string {
  if (!args) return "";

  if (func === "mint") {
    const parts = args.split(",").map((s) => s.trim());
    return `Minting ${parts[2] || ""} tokens of asset #${parts[0] || ""} to ${extractLabel(parts[1] || "")}`;
  }
  if (func === "burn") {
    const parts = args.split(",").map((s) => s.trim());
    return `Burning ${parts[2] || ""} tokens of asset #${parts[0] || ""} from ${extractLabel(parts[1] || "")}`;
  }
  if (func === "balanceOf") {
    const parts = args.split(",").map((s) => s.trim());
    return `Querying balance of ${extractLabel(parts[0] || "")} for asset #${parts[1] || ""}`;
  }
  if (func === "totalSupply") {
    return `Querying total supply of asset #${args.trim()}`;
  }
  if (func === "holderCount") {
    return `Querying number of holders for asset #${args.trim()}`;
  }
  if (func === "isHolder") {
    const parts = args.split(",").map((s) => s.trim());
    return `Checking if ${extractLabel(parts[1] || "")} holds asset #${parts[0] || ""}`;
  }

  return args.length > 150 ? args.slice(0, 150) + "..." : args;
}

function humanizeReadCall(contract: string, func: string): string {
  if (func === "balanceOf") return `Read Balance (${contract})`;
  if (func === "totalSupply") return `Read Total Supply (${contract})`;
  if (func === "holderCount") return `Read Holder Count (${contract})`;
  if (func === "isHolder") return `Read Is Holder (${contract})`;
  if (func === "mintedCount") return `Read Minted Count (${contract})`;
  if (func === "isVerified") return `Read Identity Status (${contract})`;
  if (func === "getComplianceModules") return `Read Compliance Modules (${contract})`;
  if (func === "supportsInterface") return `Check Interface Support (${contract})`;
  return `Read ${splitCamelCase(func)} (${contract})`;
}

function humanizeReturnValue(_func: string, val: string): string {
  if (val === "true") return "Yes";
  if (val === "false") return "No";
  // Just a number
  if (/^\d+$/.test(val.trim())) return val.trim();
  return val.length > 80 ? val.slice(0, 80) + "..." : val;
}

function humanizeRevert(func: string, args: string, returnData: string | null): string {
  const error = returnData || args || func;
  // Extract custom error name
  const m = error.match(/(\w+__\w+)\(\)/);
  if (m) {
    const parts = m[1].split("__");
    return `${parts[0]} rejected: ${splitCamelCase(parts[1])}`;
  }
  const m2 = error.match(/(\w+)\(\)/);
  if (m2) return `Error: ${splitCamelCase(m2[1])}`;
  return `Reverted: ${error.length > 120 ? error.slice(0, 120) + "..." : error}`;
}

function splitCamelCase(s: string): string {
  return s
    .replace(/([a-z])([A-Z])/g, "$1 $2")
    .replace(/([A-Z]+)([A-Z][a-z])/g, "$1 $2")
    .replace(/^./, (c) => c.toUpperCase());
}

async function runGas(ws: WebSocket) {
  send(ws, "status", { step: "Running gas report..." });
  send(ws, "log", "[forge-visual] Running forge test --gas-report ...");

  const output = await forgeExecMerged(ws, ["test", "--gas-report"]);
  const contracts = parseGasReport(output);

  const diffs = computeGasDiffs(contracts, 10);
  appendGasHistory(contracts);

  const data = {
    contracts,
    diffs,
    hasPrevious: diffs.length > 0,
    threshold: 10,
    alertsCount: diffs.filter((d: GasDiff) => d.alert).length,
    generatedAt: now(),
    gitRef: gitRef(),
    gitBranch: gitBranch(),
  };

  const totalFuncs = contracts.reduce((s: number, c: ContractGasReport) => s + c.functions.length, 0);
  send(ws, "log", `[forge-visual] Gas complete: ${contracts.length} contracts, ${totalFuncs} functions`);
  send(ws, "gas_result", data);
}

async function runCoverage(ws: WebSocket) {
  send(ws, "status", { step: "Running coverage (slow)..." });
  send(ws, "log", "[forge-visual] Running forge coverage --report lcov --ir-minimum ...");

  await forgeExecMerged(ws, ["coverage", "--report", "lcov", "--ir-minimum"]);

  const lcovPath = join(PROJECT_ROOT, "lcov.info");
  if (!existsSync(lcovPath)) {
    send(ws, "error", "lcov.info not found after coverage run");
    return;
  }

  const content = readFileSync(lcovPath, "utf-8");
  const files = parseLcov(content);

  const totalLinesHit = files.reduce((s: number, f: FileCov) => s + f.linesHit, 0);
  const totalLines = files.reduce((s: number, f: FileCov) => s + f.linesTotal, 0);
  const totalBranchesHit = files.reduce((s: number, f: FileCov) => s + f.branchesHit, 0);
  const totalBranches = files.reduce((s: number, f: FileCov) => s + f.branchesTotal, 0);

  const data = {
    files,
    totalLinesHit,
    totalLines,
    totalBranchesHit,
    totalBranches,
    linePct: totalLines > 0 ? (totalLinesHit / totalLines) * 100 : 100,
    branchPct: totalBranches > 0 ? (totalBranchesHit / totalBranches) * 100 : 100,
    generatedAt: now(),
    gitRef: gitRef(),
    gitBranch: gitBranch(),
  };

  send(ws, "log", `[forge-visual] Coverage complete: ${data.linePct.toFixed(1)}% lines, ${data.branchPct.toFixed(1)}% branches`);
  send(ws, "coverage_result", data);
}

async function runTrace(ws: WebSocket) {
  send(ws, "status", { step: "Running traces..." });
  send(ws, "log", "[forge-visual] Running forge test -vvvv ...");

  const output = await forgeExecMerged(ws, ["test", "-vvvv"]);
  const traces = parseTraceOutput(output);

  send(ws, "log", `[forge-visual] Traces complete: ${traces.length} tests traced`);
  send(ws, "trace_result", traces);
}

async function runSelectors(ws: WebSocket) {
  send(ws, "status", { step: "Inspecting selectors..." });
  send(ws, "log", "[forge-visual] Inspecting facet selectors ...");

  const facets = detectFacets();
  if (facets.length === 0) {
    send(ws, "error", "No facets found");
    return;
  }

  const entries: SelectorEntry[] = [];
  for (const facet of facets) {
    try {
      const output = execSync(`forge inspect ${facet} methods --json`, {
        cwd: PROJECT_ROOT,
        encoding: "utf-8",
        maxBuffer: 10 * 1024 * 1024,
        stdio: ["pipe", "pipe", "pipe"],
      });
      const methods = JSON.parse(output) as Record<string, string>;
      for (const [signature, selector] of Object.entries(methods)) {
        entries.push({
          selector: selector.startsWith("0x") ? selector : `0x${selector}`,
          signature,
          facet,
        });
      }
    } catch { /* skip */ }
  }

  entries.sort((a, b) => a.facet.localeCompare(b.facet) || a.selector.localeCompare(b.selector));

  const selectorMap = new Map<string, string[]>();
  for (const e of entries) {
    const list = selectorMap.get(e.selector) ?? [];
    if (!list.includes(e.facet)) list.push(e.facet);
    selectorMap.set(e.selector, list);
  }

  const collisions = [...selectorMap.entries()]
    .filter(([, fs]) => fs.length > 1)
    .map(([selector, fs]) => ({ selector, facets: fs }));

  send(ws, "selector_result", {
    entries,
    collisions,
    facetCount: new Set(entries.map((e) => e.facet)).size,
    totalSelectors: entries.length,
  });
}

async function runFullReport(ws: WebSocket) {
  await runTests(ws);
  await runGas(ws);
  await runSelectors(ws);

  // Coverage is slow, run last
  try {
    await runCoverage(ws);
  } catch (e) {
    send(ws, "log", `Coverage skipped: ${e}`);
  }

  // Trace is slow too
  try {
    await runTrace(ws);
  } catch (e) {
    send(ws, "log", `Traces skipped: ${e}`);
  }
}

// ── Forge Execution ──

/** Runs forge silently — no WebSocket streaming. Returns stdout. */
function forgeExecQuiet(args: string[]): Promise<string> {
  return new Promise((resolve, reject) => {
    const child = spawn("forge", args, {
      cwd: PROJECT_ROOT,
      stdio: ["pipe", "pipe", "pipe"],
    });

    let stdout = "";
    let stderr = "";

    child.stdout.on("data", (chunk: Buffer) => {
      stdout += chunk.toString();
    });

    child.stderr.on("data", (chunk: Buffer) => {
      stderr += chunk.toString();
    });

    child.on("close", (code) => {
      if (code === 0 || stdout.length > 0) {
        resolve(stdout);
      } else {
        reject(new Error(`forge ${args.join(" ")} failed (exit ${code}): ${stderr.slice(0, 500)}`));
      }
    });

    child.on("error", reject);
  });
}

function forgeExecMerged(ws: WebSocket, args: string[]): Promise<string> {
  return new Promise((resolve, reject) => {
    const child = spawn("forge", args, {
      cwd: PROJECT_ROOT,
      stdio: ["pipe", "pipe", "pipe"],
    });

    let combined = "";

    child.stdout.on("data", (chunk: Buffer) => {
      const text = chunk.toString();
      combined += text;
      for (const line of text.split("\n")) {
        if (line.trim()) send(ws, "log", line);
      }
    });

    child.stderr.on("data", (chunk: Buffer) => {
      const text = chunk.toString();
      combined += text;
      for (const line of text.split("\n")) {
        if (line.trim()) send(ws, "log", line);
      }
    });

    child.on("close", (code) => {
      if (code === 0 || combined.length > 100) {
        resolve(combined);
      } else {
        reject(new Error(`forge ${args.join(" ")} failed (exit ${code})`));
      }
    });

    child.on("error", reject);
  });
}

// ── Parsers (server-side, same logic as client) ──

interface TestSuite {
  file: string;
  contract: string;
  facet: string | null;
  durationMs: number;
  tests: TestCase[];
  passed: number;
  failed: number;
  skipped: number;
}

interface TestCase {
  name: string;
  status: "success" | "failure" | "skipped";
  reason: string | null;
  durationMs: number;
  gas: number | null;
  meanGas: number | null;
  medianGas: number | null;
  kind: unknown;
}

function parseForgeTestJson(jsonStr: string): TestSuite[] {
  const raw = JSON.parse(jsonStr);
  const suites: TestSuite[] = [];

  for (const [key, suiteResult] of Object.entries(raw) as [string, any][]) {
    const idx = key.lastIndexOf(":");
    const file = idx === -1 ? key : key.slice(0, idx);
    const contract = idx === -1 ? key : key.slice(idx + 1);
    const facet = contract.replace(/Tests?$/, "") || null;
    const durationMs = parseDuration(suiteResult.duration);
    const tests: TestCase[] = [];

    for (const [testName, result] of Object.entries(suiteResult.test_results) as [string, any][]) {
      const status = result.status === "Success" ? "success" : result.status === "Failure" ? "failure" : "skipped";
      const kindObj = result.kind ?? {};
      let kind: unknown = { type: "Standard" };
      let gas: number | null = null;
      let meanGas: number | null = null;
      let medianGas: number | null = null;

      if ("Standard" in kindObj) {
        gas = typeof kindObj.Standard === "number" ? kindObj.Standard : null;
      } else if ("Fuzz" in kindObj) {
        const fuzz = kindObj.Fuzz;
        kind = { type: "Fuzz", runs: fuzz.runs ?? 0 };
        meanGas = fuzz.mean_gas ?? null;
        medianGas = fuzz.median_gas ?? null;
        gas = fuzz.first_case?.gas ?? meanGas;
      } else if ("Invariant" in kindObj) {
        const inv = kindObj.Invariant;
        kind = { type: "Invariant", runs: inv.runs ?? 0, calls: inv.calls ?? 0, reverts: inv.reverts ?? 0 };
      }

      tests.push({ name: testName, status: status as TestCase["status"], reason: result.reason, durationMs: parseDuration(result.duration), gas, meanGas, medianGas, kind });
    }

    tests.sort((a, b) => a.name.localeCompare(b.name));
    suites.push({ file, contract, facet, durationMs, tests, passed: tests.filter(t => t.status === "success").length, failed: tests.filter(t => t.status === "failure").length, skipped: tests.filter(t => t.status === "skipped").length });
  }

  suites.sort((a, b) => a.contract.localeCompare(b.contract));
  return suites;
}

function parseDuration(d: any): number {
  if (typeof d === "object" && d !== null) return (d.secs ?? 0) * 1000 + (d.nanos ?? 0) / 1_000_000;
  if (typeof d === "string") {
    let ms = 0;
    for (const p of d.split(/\s+/)) {
      const m = p.match(/^([\d.]+)(ns|µs|us|ms|s)$/);
      if (!m) continue;
      const v = parseFloat(m[1]);
      if (m[2] === "s") ms += v * 1000;
      else if (m[2] === "ms") ms += v;
      else if (m[2] === "µs" || m[2] === "us") ms += v / 1000;
      else if (m[2] === "ns") ms += v / 1_000_000;
    }
    return ms;
  }
  return 0;
}

interface FunctionGas { name: string; min: number; avg: number; median: number; max: number; calls: number; }
interface ContractGasReport { contract: string; functions: FunctionGas[]; }

function parseGasReport(text: string): ContractGasReport[] {
  const reports: ContractGasReport[] = [];
  let currentContract: string | null = null;
  let fns: FunctionGas[] = [];
  let inFn = false;

  for (const line of text.split("\n")) {
    const t = line.trim();
    if (!t || t.startsWith("+-") || t.startsWith("|--")) continue;
    if (!t.startsWith("|")) { inFn = false; continue; }
    const cells = t.split("|").map(s => s.trim()).filter(s => s.length > 0);
    if (cells.length === 0) continue;
    if (cells.length >= 1 && /[Cc]ontract\s*$/.test(cells[0])) {
      if (currentContract && fns.length > 0) reports.push({ contract: currentContract, functions: fns });
      const h = cells[0].replace(/\s*[Cc]ontract\s*$/, "").trim();
      const ci = h.lastIndexOf(":");
      currentContract = ci >= 0 ? h.slice(ci + 1) : h;
      fns = [];
      inFn = false;
      continue;
    }
    if (cells[0] === "Function Name") { inFn = true; continue; }
    if (cells[0] === "Deployment Cost" || cells[0] === "Deployment Size") continue;
    if (inFn && cells.length >= 6 && currentContract) {
      const [min, avg, median, max, callCount] = [1,2,3,4,5].map(i => parseInt(cells[i].replace(/,/g, ""), 10));
      if (![min, avg, median, max, callCount].some(isNaN)) {
        fns.push({ name: cells[0], min, avg, median, max, calls: callCount });
      }
    }
  }
  if (currentContract && fns.length > 0) reports.push({ contract: currentContract, functions: fns });
  return reports;
}

interface FileCov { path: string; facet: string | null; category: string; linesHit: number; linesTotal: number; branchesHit: number; branchesTotal: number; linePct: number; branchPct: number; uncoveredLines: number[]; level: string; }

function categorizePath(filePath: string): string {
  const p = filePath.toLowerCase();
  // Common Solidity project directory patterns
  if (p.includes("/facets/")) return "Facets";
  if (p.includes("/libraries/") || p.includes("/libs/") || p.includes("/lib/")) return "Libraries";
  if (p.includes("/interfaces/") || p.includes("/interface/")) return "Interfaces";
  if (p.includes("/compliance/") || p.includes("/modules/")) return "Compliance";
  if (p.includes("/identity/") || p.includes("/registry/")) return "Identity";
  if (p.includes("/initializers/") || p.includes("/init/")) return "Initializers";
  if (p.includes("/extensions/") || p.includes("/mixins/")) return "Extensions";
  if (p.includes("/utils/") || p.includes("/helpers/")) return "Utilities";
  if (p.includes("/governance/")) return "Governance";
  if (p.includes("/tokens/") || p.includes("/token/")) return "Token";
  if (p.includes("/access/") || p.includes("/auth/")) return "Access Control";
  if (p.includes("/proxy/") || p.includes("/upgradeable/")) return "Proxy";
  if (p.includes("/oracle/") || p.includes("/oracles/")) return "Oracles";
  if (p.includes("/mocks/") || p.includes("/test/")) return "Test";
  // Derive from first directory after src/
  const srcMatch = filePath.match(/src\/([^/]+)\//);
  if (srcMatch) {
    const dir = srcMatch[1];
    return dir.charAt(0).toUpperCase() + dir.slice(1);
  }
  return "Core";
}

function parseLcov(content: string): FileCov[] {
  const files: FileCov[] = [];
  let cp: string | null = null;
  let lh = 0, lt = 0, bh = 0, bt = 0;
  let unc: number[] = [];

  for (const line of content.split("\n")) {
    const t = line.trim();
    if (t.startsWith("SF:")) { cp = t.slice(3); lh=lt=bh=bt=0; unc=[]; }
    else if (t.startsWith("DA:")) { const p=t.slice(3).split(","); if(parseInt(p[1],10)===0) unc.push(parseInt(p[0],10)); }
    else if (t.startsWith("LF:")) lt=parseInt(t.slice(3),10);
    else if (t.startsWith("LH:")) lh=parseInt(t.slice(3),10);
    else if (t.startsWith("BRF:")) bt=parseInt(t.slice(4),10);
    else if (t.startsWith("BRH:")) bh=parseInt(t.slice(4),10);
    else if (t==="end_of_record"&&cp) {
      const lp=lt>0?(lh/lt)*100:100;
      const bp=bt>0?(bh/bt)*100:100;
      const parts=cp.split("/");
      const facet=parts[parts.length-1]?.replace(".sol","")??null;
      const cat = categorizePath(cp);
      files.push({path:cp,facet,category:cat,linesHit:lh,linesTotal:lt,branchesHit:bh,branchesTotal:bt,linePct:lp,branchPct:bp,uncoveredLines:unc,level:lp>=80?"high":lp>=50?"medium":"low"});
      cp=null;
    }
  }
  return files;
}

interface TraceTest { name: string; status: string; gas: number; nodes: TraceNode[]; }
interface TraceNode { depth: number; gas: number; kind: string; contract: string; func: string; args: string; returnData: string | null; isRevert: boolean; isEmit: boolean; children: TraceNode[]; }

function parseTraceOutput(text: string): TraceTest[] {
  const tests: TraceTest[] = [];
  const lines = text.split("\n");
  let i = 0;
  while (i < lines.length) {
    const m = lines[i].match(/\[(PASS|FAIL)\]\s+(\S+)\s+\(gas:\s*(\d+)\)/);
    if (m) {
      const status = m[1].toLowerCase();
      const name = m[2];
      const gas = parseInt(m[3], 10);
      i++;
      if (i<lines.length&&lines[i].trim()==="Traces:") i++;
      const tl: string[] = [];
      while (i<lines.length) {
        if(lines[i].match(/^\[(PASS|FAIL)\]/)||lines[i].match(/^Suite result:/)||lines[i].match(/^Ran \d+ test/)) break;
        tl.push(lines[i]); i++;
      }
      tests.push({ name, status, gas, nodes: parseTraceLines(tl) });
    } else { i++; }
  }
  return tests;
}

function parseTraceLines(lines: string[]): TraceNode[] {
  const roots: TraceNode[] = [];
  const stack: { depth: number; node: TraceNode }[] = [];
  for (const line of lines) {
    if (!line.trim()) continue;
    let d=0;
    for(let i=0;i<line.length;i++){const c=line[i];if(c===" ")continue;if(c==="│"||c==="├"||c==="└"||c==="─")d++;else break;}
    d=Math.floor(d/2);
    const content=line.replace(/^[\s│├└─]+/,"").trim();
    if(!content) continue;
    const node=parseTraceLine(content,d);
    if(!node) continue;
    while(stack.length>0&&stack[stack.length-1].depth>=d)stack.pop();
    if(stack.length>0)stack[stack.length-1].node.children.push(node);
    else roots.push(node);
    stack.push({depth:d,node});
  }
  return roots;
}

function parseTraceLine(c: string, d: number): TraceNode|null {
  let m=c.match(/^←\s+\[(Return|Revert|Stop)\]\s*(.*)?$/);
  if(m) return {depth:d,gas:0,kind:m[1]==="Revert"?"REVERT":m[1]==="Stop"?"STOP":"RETURN",contract:"",func:m[1],args:"",returnData:m[2]?.trim()||null,isRevert:m[1]==="Revert",isEmit:false,children:[]};
  m=c.match(/^emit\s+(\w+)\((.*)?\)$/);
  if(m) return {depth:d,gas:0,kind:"EMIT",contract:"",func:m[1],args:m[2]||"",returnData:null,isRevert:false,isEmit:true,children:[]};
  m=c.match(/^\[(\d+)\]\s+(\w+)::(\w+)\((.*)?\)(?:\s+\[(\w+)\])?$/);
  if(m) return {depth:d,gas:parseInt(m[1],10),kind:(m[5]?.toUpperCase()||"CALL"),contract:m[2],func:m[3],args:m[4]||"",returnData:null,isRevert:false,isEmit:false,children:[]};
  m=c.match(/^\[(\d+)\]\s+(\w+)::(\w+)\(\)$/);
  if(m) return {depth:d,gas:parseInt(m[1],10),kind:"CALL",contract:m[2],func:m[3],args:"",returnData:null,isRevert:false,isEmit:false,children:[]};
  return null;
}

// ── Gas History ──

interface GasDiff { function: string; contract: string; previousAvg: number; currentAvg: number; delta: number; deltaPct: number; alert: boolean; }

interface SelectorEntry { selector: string; signature: string; facet: string; }

function computeGasDiffs(current: ContractGasReport[], threshold: number): GasDiff[] {
  const histPath = join(REPORTS_DIR, ".forge-visual", "gas-history.json");
  if (!existsSync(histPath)) return [];
  const history = JSON.parse(readFileSync(histPath, "utf-8"));
  const previous = history.entries?.[history.entries.length - 1]?.reports;
  if (!previous) return [];

  const diffs: GasDiff[] = [];
  for (const curr of current) {
    for (const func of curr.functions) {
      const prevC = previous.find((c: any) => c.contract === curr.contract);
      const prevF = prevC?.functions.find((f: any) => f.name === func.name);
      if (prevF) {
        const delta = func.avg - prevF.avg;
        const deltaPct = prevF.avg > 0 ? (delta / prevF.avg) * 100 : 0;
        diffs.push({ function: func.name, contract: curr.contract, previousAvg: prevF.avg, currentAvg: func.avg, delta, deltaPct, alert: deltaPct > threshold });
      }
    }
  }
  diffs.sort((a, b) => b.deltaPct - a.deltaPct);
  return diffs;
}

function appendGasHistory(reports: ContractGasReport[]) {
  const dir = join(REPORTS_DIR, ".forge-visual");
  mkdirSync(dir, { recursive: true });
  const histPath = join(dir, "gas-history.json");
  const history = existsSync(histPath) ? JSON.parse(readFileSync(histPath, "utf-8")) : { entries: [] };
  history.entries.push({ timestamp: new Date().toISOString(), gitRef: gitRef(), gitBranch: gitBranch(), reports });
  if (history.entries.length > MAX_HISTORY) history.entries.splice(0, history.entries.length - MAX_HISTORY);
  writeFileSync(histPath, JSON.stringify(history, null, 2));
}

// ── Contract Detection ──

function detectFacets(): string[] {
  const srcDir = findSrcDir();
  if (!srcDir) return [];
  const solFiles = walkSol(srcDir);

  // Try Diamond facets first
  const facetFiles = solFiles.filter(f => f.includes("/facets/"));
  if (facetFiles.length > 0) {
    return facetFiles.map(f => basename(f, ".sol")).filter(n => !n.startsWith("I"));
  }

  // Fallback: all non-interface, non-library, non-test .sol files
  return solFiles
    .map(f => basename(f, ".sol"))
    .filter(n => !n.startsWith("I") && !n.endsWith("Test") && !n.startsWith("Mock") && !n.startsWith("Script"));
}

interface ProjectInfo {
  name: string;
  solcVersion: string | null;
  srcDir: string;
  testDir: string;
  isDiamond: boolean;
  contractCount: number;
  remappings: string[];
}

function detectProjectInfo(): ProjectInfo {
  const tomlPath = join(PROJECT_ROOT, "foundry.toml");
  let name = basename(PROJECT_ROOT);
  let solcVersion: string | null = null;
  let srcDir = "src";
  let testDir = "test";
  const remappings: string[] = [];

  if (existsSync(tomlPath)) {
    try {
      const content = readFileSync(tomlPath, "utf-8");
      const srcMatch = content.match(/^\s*src\s*=\s*["']([^"']+)["']/m);
      if (srcMatch) srcDir = srcMatch[1];
      const testMatch = content.match(/^\s*test\s*=\s*["']([^"']+)["']/m);
      if (testMatch) testDir = testMatch[1];
      const solcMatch = content.match(/^\s*solc_version\s*=\s*["']([^"']+)["']/m);
      if (solcMatch) solcVersion = solcMatch[1];
      // Collect remappings
      const remapSection = content.match(/^\s*remappings\s*=\s*\[([\s\S]*?)\]/m);
      if (remapSection) {
        for (const m of remapSection[1].matchAll(/["']([^"']+)["']/g)) {
          remappings.push(m[1]);
        }
      }
    } catch { /* */ }
  }

  // Try package.json for name
  const pkgPath = join(PROJECT_ROOT, "package.json");
  if (existsSync(pkgPath)) {
    try {
      const pkg = JSON.parse(readFileSync(pkgPath, "utf-8"));
      if (pkg.name) name = pkg.name;
    } catch { /* */ }
  }

  // Detect Diamond pattern
  const srcFullDir = join(PROJECT_ROOT, srcDir);
  const solFiles = existsSync(srcFullDir) ? walkSol(srcFullDir) : [];
  const isDiamond = solFiles.some(f =>
    f.includes("Diamond") || f.includes("/facets/") || f.includes("DiamondCut") || f.includes("LibDiamond")
  );
  const contractCount = solFiles.filter(f => !basename(f).startsWith("I")).length;

  return { name, solcVersion, srcDir, testDir, isDiamond, contractCount, remappings };
}

function findSrcDir(): string | null {
  const tomlPath = join(PROJECT_ROOT, "foundry.toml");
  if (existsSync(tomlPath)) {
    try {
      const content = readFileSync(tomlPath, "utf-8");
      const m = content.match(/^\s*src\s*=\s*["']([^"']+)["']/m);
      if (m) {
        const dir = join(PROJECT_ROOT, m[1]);
        if (existsSync(dir)) return dir;
      }
    } catch { /* */ }
  }
  const def = join(PROJECT_ROOT, "src");
  return existsSync(def) ? def : null;
}

function walkSol(dir: string): string[] {
  const results: string[] = [];
  try {
    const { readdirSync } = require("node:fs") as typeof import("node:fs");
    for (const entry of readdirSync(dir, { withFileTypes: true })) {
      const full = join(dir, entry.name);
      if (entry.isDirectory()) results.push(...walkSol(full));
      else if (entry.name.endsWith(".sol")) results.push(full);
    }
  } catch { /* */ }
  return results;
}

// ── Helpers ──

function send(ws: WebSocket, type: string, data: unknown) {
  if (ws.readyState === WebSocket.OPEN) {
    ws.send(JSON.stringify({ type, data, timestamp: new Date().toISOString() }));
  }
}

function now(): string {
  return new Date().toISOString().replace("T", " ").slice(0, 19) + " UTC";
}

function gitRef(): string | null {
  try { return execSync("git rev-parse --short HEAD", { cwd: PROJECT_ROOT, encoding: "utf-8" }).trim(); } catch { return null; }
}

function gitBranch(): string | null {
  try { return execSync("git rev-parse --abbrev-ref HEAD", { cwd: PROJECT_ROOT, encoding: "utf-8" }).trim(); } catch { return null; }
}

function findProjectRoot(): string {
  // Check for FORGE_PROJECT env, then walk up from cwd
  if (process.env.FORGE_PROJECT) return resolve(process.env.FORGE_PROJECT);

  let dir = process.cwd();
  while (true) {
    if (existsSync(join(dir, "foundry.toml"))) return dir;
    const parent = resolve(dir, "..");
    if (parent === dir) {
      // Default to contracts package
      const contractsDir = resolve(__dirname, "../../..");
      if (existsSync(join(contractsDir, "foundry.toml"))) return contractsDir;
      throw new Error("Could not find foundry.toml. Set FORGE_PROJECT env var.");
    }
    dir = parent;
  }
}

// ── Start ──

server.listen(PORT, () => {
  console.log(`[forge-visual] Server running at http://localhost:${PORT}`);
  console.log(`[forge-visual] WebSocket at ws://localhost:${PORT}/ws`);
});

import type {
  TestCase,
  TestKind,
  TestSuite,
  ContractGasReport,
  FunctionGas,
  FileCoverage,
  TraceTest,
  TraceNode,
} from "../types";

// ── Test JSON Parser ──

interface RawSuiteResult {
  duration: string | { secs: number; nanos: number };
  test_results: Record<string, RawTestResult>;
}

interface RawTestResult {
  status: string;
  reason: string | null;
  duration: string | { secs: number; nanos: number };
  kind: Record<string, unknown>;
}

export function parseForgeTestJson(jsonStr: string): TestSuite[] {
  const raw: Record<string, RawSuiteResult> = JSON.parse(jsonStr);
  const suites: TestSuite[] = [];

  for (const [key, suiteResult] of Object.entries(raw)) {
    const idx = key.lastIndexOf(":");
    const file = idx === -1 ? key : key.slice(0, idx);
    const contract = idx === -1 ? key : key.slice(idx + 1);
    const facet = contract.replace(/Tests?$/, "") || null;
    const durationMs = parseDuration(suiteResult.duration);

    const tests: TestCase[] = [];

    for (const [testName, result] of Object.entries(suiteResult.test_results)) {
      const status =
        result.status === "Success"
          ? "success"
          : result.status === "Failure"
            ? "failure"
            : "skipped";

      const { kind, gas, meanGas, medianGas } = parseKindAndGas(result.kind);

      tests.push({
        name: testName,
        status: status as TestCase["status"],
        reason: result.reason,
        durationMs: parseDuration(result.duration),
        gas,
        meanGas,
        medianGas,
        kind,
      });
    }

    tests.sort((a, b) => a.name.localeCompare(b.name));

    const passed = tests.filter((t) => t.status === "success").length;
    const failed = tests.filter((t) => t.status === "failure").length;
    const skipped = tests.filter((t) => t.status === "skipped").length;

    suites.push({ file, contract, facet, durationMs, tests, passed, failed, skipped });
  }

  suites.sort((a, b) => a.contract.localeCompare(b.contract));
  return suites;
}

function parseDuration(d: string | { secs: number; nanos: number }): number {
  if (typeof d === "object" && d !== null) {
    return (d.secs ?? 0) * 1000 + (d.nanos ?? 0) / 1_000_000;
  }
  if (typeof d === "string") {
    let totalMs = 0;
    for (const part of d.split(/\s+/)) {
      const match = part.match(/^([\d.]+)(ns|µs|us|ms|s)$/);
      if (!match) continue;
      const val = parseFloat(match[1]);
      switch (match[2]) {
        case "s": totalMs += val * 1000; break;
        case "ms": totalMs += val; break;
        case "µs":
        case "us": totalMs += val / 1000; break;
        case "ns": totalMs += val / 1_000_000; break;
      }
    }
    return totalMs;
  }
  return 0;
}

function parseKindAndGas(kindObj: Record<string, unknown>): {
  kind: TestKind;
  gas: number | null;
  meanGas: number | null;
  medianGas: number | null;
} {
  if ("Standard" in kindObj) {
    return {
      kind: { type: "Standard" },
      gas: typeof kindObj.Standard === "number" ? kindObj.Standard : null,
      meanGas: null,
      medianGas: null,
    };
  }

  if ("Fuzz" in kindObj) {
    const fuzz = kindObj.Fuzz as Record<string, unknown>;
    const runs = (fuzz.runs as number) ?? 0;
    const meanGas = (fuzz.mean_gas as number) ?? null;
    const medianGas = (fuzz.median_gas as number) ?? null;
    const firstCaseGas = (fuzz.first_case as Record<string, unknown>)?.gas as number | undefined;
    return { kind: { type: "Fuzz", runs }, gas: firstCaseGas ?? meanGas, meanGas, medianGas };
  }

  if ("Invariant" in kindObj) {
    const inv = kindObj.Invariant as Record<string, unknown>;
    return {
      kind: {
        type: "Invariant",
        runs: (inv.runs as number) ?? 0,
        calls: (inv.calls as number) ?? 0,
        reverts: (inv.reverts as number) ?? 0,
      },
      gas: null,
      meanGas: null,
      medianGas: null,
    };
  }

  return { kind: { type: "Standard" }, gas: null, meanGas: null, medianGas: null };
}

// ── Gas Report Parser ──

export function parseGasReport(text: string): ContractGasReport[] {
  const reports: ContractGasReport[] = [];
  let currentContract: string | null = null;
  let currentFunctions: FunctionGas[] = [];
  let inFunctionSection = false;

  for (const line of text.split("\n")) {
    const trimmed = line.trim();
    if (!trimmed || trimmed.startsWith("+-") || trimmed.startsWith("|--")) continue;
    if (!trimmed.startsWith("|")) { inFunctionSection = false; continue; }

    const cells = trimmed.split("|").map((s) => s.trim()).filter((s) => s.length > 0);
    if (cells.length === 0) continue;

    if (cells.length >= 1 && /[Cc]ontract\s*$/.test(cells[0])) {
      if (currentContract && currentFunctions.length > 0) {
        reports.push({ contract: currentContract, functions: currentFunctions });
      }
      const header = cells[0].replace(/\s*[Cc]ontract\s*$/, "").trim();
      const colonIdx = header.lastIndexOf(":");
      currentContract = colonIdx >= 0 ? header.slice(colonIdx + 1) : header;
      currentFunctions = [];
      inFunctionSection = false;
      continue;
    }

    if (cells[0] === "Function Name") { inFunctionSection = true; continue; }
    if (cells[0] === "Deployment Cost" || cells[0] === "Deployment Size") continue;

    if (inFunctionSection && cells.length >= 6 && currentContract) {
      const min = parseInt(cells[1].replace(/,/g, ""), 10);
      const avg = parseInt(cells[2].replace(/,/g, ""), 10);
      const median = parseInt(cells[3].replace(/,/g, ""), 10);
      const max = parseInt(cells[4].replace(/,/g, ""), 10);
      const callCount = parseInt(cells[5].replace(/,/g, ""), 10);
      if (![min, avg, median, max, callCount].some(isNaN)) {
        currentFunctions.push({ name: cells[0], min, avg, median, max, calls: callCount });
      }
    }
  }

  if (currentContract && currentFunctions.length > 0) {
    reports.push({ contract: currentContract, functions: currentFunctions });
  }

  return reports;
}

// ── LCOV Parser ──

export function parseLcov(content: string): FileCoverage[] {
  const files: FileCoverage[] = [];
  let currentPath: string | null = null;
  let linesHit = 0;
  let linesTotal = 0;
  let branchesHit = 0;
  let branchesTotal = 0;
  let uncovered: number[] = [];

  for (const line of content.split("\n")) {
    const trimmed = line.trim();
    if (trimmed.startsWith("SF:")) {
      currentPath = trimmed.slice(3);
      linesHit = 0; linesTotal = 0; branchesHit = 0; branchesTotal = 0; uncovered = [];
    } else if (trimmed.startsWith("DA:")) {
      const parts = trimmed.slice(3).split(",");
      if (parseInt(parts[1], 10) === 0) uncovered.push(parseInt(parts[0], 10));
    } else if (trimmed.startsWith("LF:")) {
      linesTotal = parseInt(trimmed.slice(3), 10);
    } else if (trimmed.startsWith("LH:")) {
      linesHit = parseInt(trimmed.slice(3), 10);
    } else if (trimmed.startsWith("BRF:")) {
      branchesTotal = parseInt(trimmed.slice(4), 10);
    } else if (trimmed.startsWith("BRH:")) {
      branchesHit = parseInt(trimmed.slice(4), 10);
    } else if (trimmed === "end_of_record" && currentPath) {
      const linePct = linesTotal > 0 ? (linesHit / linesTotal) * 100 : 100;
      const branchPct = branchesTotal > 0 ? (branchesHit / branchesTotal) * 100 : 100;
      const path = currentPath;
      const parts = path.split("/");
      const facet = parts[parts.length - 1]?.replace(".sol", "") ?? null;
      let category = "Other";
      if (path.includes("/facets/")) {
        if (path.includes("/token/")) category = "Token";
        else if (path.includes("/compliance/")) category = "Compliance";
        else if (path.includes("/identity/")) category = "Identity";
        else if (path.includes("/rwa/")) category = "RWA";
        else if (path.includes("/core/")) category = "Core";
        else if (path.includes("/security/")) category = "Security";
        else category = "Facets";
      } else if (path.includes("/compliance/modules/")) category = "Compliance Modules";
      else if (path.includes("/storage/") || path.includes("/libraries/")) category = "Libraries";

      files.push({
        path, facet, category, linesHit, linesTotal, branchesHit, branchesTotal,
        linePct, branchPct, uncoveredLines: uncovered,
        level: linePct >= 80 ? "high" : linePct >= 50 ? "medium" : "low",
      });
      currentPath = null;
    }
  }
  return files;
}

// ── Trace Parser ──

export function parseTraceOutput(text: string): TraceTest[] {
  const tests: TraceTest[] = [];
  const lines = text.split("\n");
  let i = 0;

  while (i < lines.length) {
    const line = lines[i];
    const statusMatch = line.match(/\[(PASS|FAIL)\]\s+(\S+)\s+\(gas:\s*(\d+)\)/);
    if (statusMatch) {
      const status = statusMatch[1].toLowerCase() as "pass" | "fail";
      const name = statusMatch[2];
      const gas = parseInt(statusMatch[3], 10);
      i++;
      if (i < lines.length && lines[i].trim() === "Traces:") i++;

      const traceLines: string[] = [];
      while (i < lines.length) {
        const tl = lines[i];
        if (tl.match(/^\[(PASS|FAIL)\]/) || tl.match(/^Suite result:/) || tl.match(/^Ran \d+ test/)) break;
        traceLines.push(tl);
        i++;
      }

      const nodes = parseTraceLines(traceLines);
      tests.push({ name, status, gas, nodes });
    } else {
      i++;
    }
  }
  return tests;
}

function parseTraceLines(lines: string[]): TraceNode[] {
  const rootNodes: TraceNode[] = [];
  const stack: { depth: number; node: TraceNode }[] = [];

  for (const line of lines) {
    if (!line.trim()) continue;
    const depth = countDepth(line);
    const content = line.replace(/^[\s│├└─]+/, "").trim();
    if (!content) continue;

    const node = parseTraceLine(content, depth);
    if (!node) continue;

    while (stack.length > 0 && stack[stack.length - 1].depth >= depth) stack.pop();

    if (stack.length > 0) {
      stack[stack.length - 1].node.children.push(node);
    } else {
      rootNodes.push(node);
    }
    stack.push({ depth, node });
  }
  return rootNodes;
}

function countDepth(line: string): number {
  let depth = 0;
  for (let i = 0; i < line.length; i++) {
    const ch = line[i];
    if (ch === " ") continue;
    if (ch === "│" || ch === "├" || ch === "└" || ch === "─") depth++;
    else break;
  }
  return Math.floor(depth / 2);
}

function parseTraceLine(content: string, depth: number): TraceNode | null {
  const returnMatch = content.match(/^←\s+\[(Return|Revert|Stop)\]\s*(.*)?$/);
  if (returnMatch) {
    const isRevert = returnMatch[1] === "Revert";
    return {
      depth, gas: 0,
      kind: isRevert ? "REVERT" : returnMatch[1] === "Stop" ? "STOP" : "RETURN",
      contract: "", func: returnMatch[1], args: "",
      returnData: returnMatch[2]?.trim() || null,
      isRevert, isEmit: false, children: [],
    };
  }

  const emitMatch = content.match(/^emit\s+(\w+)\((.*)?\)$/);
  if (emitMatch) {
    return {
      depth, gas: 0, kind: "EMIT", contract: "", func: emitMatch[1],
      args: emitMatch[2] || "", returnData: null, isRevert: false, isEmit: true, children: [],
    };
  }

  const callMatch = content.match(/^\[(\d+)\]\s+(\w+)::(\w+)\((.*)?\)(?:\s+\[(\w+)\])?$/);
  if (callMatch) {
    return {
      depth, gas: parseInt(callMatch[1], 10),
      kind: (callMatch[5]?.toUpperCase() || "CALL") as TraceNode["kind"],
      contract: callMatch[2], func: callMatch[3], args: callMatch[4] || "",
      returnData: null, isRevert: false, isEmit: false, children: [],
    };
  }

  const topMatch = content.match(/^\[(\d+)\]\s+(\w+)::(\w+)\(\)$/);
  if (topMatch) {
    return {
      depth, gas: parseInt(topMatch[1], 10), kind: "CALL",
      contract: topMatch[2], func: topMatch[3], args: "",
      returnData: null, isRevert: false, isEmit: false, children: [],
    };
  }

  return null;
}

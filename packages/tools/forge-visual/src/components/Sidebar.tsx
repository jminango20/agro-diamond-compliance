import {
  LayoutDashboard,
  FlaskConical,
  Fuel,
  Shield,
  Activity,
  Hash,
  Terminal,
  ChevronLeft,
  ChevronRight,
  PlayCircle,
} from "lucide-react";

export type Tab = "dashboard" | "runner" | "tests" | "gas" | "coverage" | "traces" | "selectors" | "logs";

interface SidebarProps {
  active: Tab;
  onNavigate: (tab: Tab) => void;
  collapsed: boolean;
  onToggle: () => void;
  hasResults: {
    tests: boolean;
    gas: boolean;
    coverage: boolean;
    traces: boolean;
    selectors: boolean;
  };
  running: boolean;
  projectName?: string;
  isDiamond?: boolean;
}

const tabs: { id: Tab; label: string; icon: typeof LayoutDashboard }[] = [
  { id: "dashboard", label: "Dashboard", icon: LayoutDashboard },
  { id: "runner", label: "Test Runner", icon: PlayCircle },
  { id: "tests", label: "Results", icon: FlaskConical },
  { id: "gas", label: "Gas", icon: Fuel },
  { id: "coverage", label: "Coverage", icon: Shield },
  { id: "traces", label: "Traces", icon: Activity },
  { id: "selectors", label: "Selectors", icon: Hash },
  { id: "logs", label: "Logs", icon: Terminal },
];

function DiamondLogo({ size = 28 }: { size?: number }) {
  return (
    <svg
      width={size}
      height={size}
      viewBox="0 0 32 32"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      {/* Outer diamond shape */}
      <path
        d="M16 2L28 12L16 30L4 12L16 2Z"
        fill="url(#diamondGrad)"
        stroke="url(#diamondStroke)"
        strokeWidth="1.5"
        strokeLinejoin="round"
      />
      {/* Top facet line */}
      <path
        d="M4 12H28"
        stroke="url(#diamondStroke)"
        strokeWidth="1"
        opacity="0.6"
      />
      {/* Left facet line */}
      <path
        d="M10 2L4 12L16 30"
        stroke="url(#diamondStroke)"
        strokeWidth="0.75"
        opacity="0.3"
        fill="none"
      />
      {/* Right facet line */}
      <path
        d="M22 2L28 12L16 30"
        stroke="url(#diamondStroke)"
        strokeWidth="0.75"
        opacity="0.3"
        fill="none"
      />
      {/* Inner highlight — top facet */}
      <path
        d="M16 2L10 12H22L16 2Z"
        fill="url(#topFacet)"
        opacity="0.5"
      />
      {/* Anvil / hammer symbol in center */}
      <path
        d="M13 14H19V16H18V20H14V16H13V14Z"
        fill="url(#anvilGrad)"
        opacity="0.9"
      />
      {/* Spark */}
      <circle cx="21" cy="10" r="1" fill="#fbbf24" opacity="0.8" />
      <circle cx="22.5" cy="8.5" r="0.5" fill="#fbbf24" opacity="0.5" />

      <defs>
        <linearGradient id="diamondGrad" x1="16" y1="2" x2="16" y2="30" gradientUnits="userSpaceOnUse">
          <stop stopColor="#6366f1" stopOpacity="0.15" />
          <stop offset="0.5" stopColor="#8b5cf6" stopOpacity="0.08" />
          <stop offset="1" stopColor="#6366f1" stopOpacity="0.02" />
        </linearGradient>
        <linearGradient id="diamondStroke" x1="16" y1="2" x2="16" y2="30" gradientUnits="userSpaceOnUse">
          <stop stopColor="#818cf8" />
          <stop offset="1" stopColor="#6366f1" stopOpacity="0.4" />
        </linearGradient>
        <linearGradient id="topFacet" x1="16" y1="2" x2="16" y2="12" gradientUnits="userSpaceOnUse">
          <stop stopColor="#c4b5fd" stopOpacity="0.3" />
          <stop offset="1" stopColor="#6366f1" stopOpacity="0" />
        </linearGradient>
        <linearGradient id="anvilGrad" x1="16" y1="14" x2="16" y2="20" gradientUnits="userSpaceOnUse">
          <stop stopColor="#e2e8f0" />
          <stop offset="1" stopColor="#94a3b8" />
        </linearGradient>
      </defs>
    </svg>
  );
}

export function Sidebar({ active, onNavigate, collapsed, onToggle, hasResults, running, projectName, isDiamond }: SidebarProps) {
  return (
    <aside
      className={`flex flex-col bg-surface-card border-r border-surface-border transition-all duration-200 ${
        collapsed ? "w-14" : "w-56"
      }`}
    >
      {/* Logo area */}
      <div className="flex items-center gap-2.5 px-3 h-14 border-b border-surface-border">
        <div className="flex-none">
          <DiamondLogo size={collapsed ? 24 : 28} />
        </div>
        {!collapsed && (
          <div className="flex-1 min-w-0">
            <div className="text-sm font-bold tracking-tight">
              <span className="text-accent">forge</span>
              <span className="text-gray-500">-</span>
              <span className="text-gray-300">visual</span>
            </div>
            <p className="text-[9px] text-gray-600 tracking-wider uppercase truncate">
              {projectName || (isDiamond ? "Diamond Project" : "Solidity Project")}
            </p>
          </div>
        )}
        <button
          onClick={onToggle}
          className="flex-none p-1 rounded hover:bg-surface-hover text-gray-600 hover:text-gray-400 transition-colors"
        >
          {collapsed ? <ChevronRight size={14} /> : <ChevronLeft size={14} />}
        </button>
      </div>

      <nav className="flex-1 py-2 space-y-0.5">
        {tabs.map(({ id, label, icon: Icon }) => {
          const isActive = active === id;
          const hasDot =
            id !== "dashboard" && id !== "logs" && id !== "runner" && hasResults[id as keyof typeof hasResults];
          const isLogRunning = id === "logs" && running;
          return (
            <button
              key={id}
              onClick={() => onNavigate(id)}
              className={`w-full flex items-center gap-3 px-3 py-2 text-sm transition-colors ${
                isActive
                  ? "bg-accent/10 text-accent border-r-2 border-accent"
                  : "text-gray-400 hover:text-gray-200 hover:bg-surface-hover"
              }`}
            >
              <div className="relative flex-none">
                <Icon size={18} />
                {isLogRunning && (
                  <span className="absolute -top-0.5 -right-0.5 w-2 h-2 rounded-full bg-accent animate-pulse" />
                )}
              </div>
              {!collapsed && <span className="flex-1 text-left">{label}</span>}
              {!collapsed && hasDot && (
                <span className="w-1.5 h-1.5 rounded-full bg-emerald-400" />
              )}
              {!collapsed && isLogRunning && (
                <span className="text-[10px] text-accent animate-pulse">live</span>
              )}
            </button>
          );
        })}
      </nav>

      {/* Footer */}
      {!collapsed && (
        <div className="px-3 py-3 border-t border-surface-border">
          <p className="text-[10px] text-gray-700 text-center">
            Foundry Test Visualizer
          </p>
        </div>
      )}
    </aside>
  );
}

import { ConnectButton } from "@rainbow-me/rainbowkit";
import { useLocation, Link } from "react-router-dom";
import { useAccount, useChainId } from "wagmi";
import { ChevronRight } from "lucide-react";
import { LogoFull } from "@/components/ui/logo";

function NetworkIndicator() {
  const { isConnected } = useAccount();
  const chainId = useChainId();

  const chainName = (() => {
    switch (chainId) {
      case 80002:
        return "Polygon Amoy";
      case 137:
        return "Polygon";
      case 1:
        return "Ethereum";
      case 11155111:
        return "Sepolia";
      default:
        return `Chain ${chainId}`;
    }
  })();

  return (
    <div className="flex items-center gap-2 rounded-full border border-white/[0.06] bg-white/[0.03] px-3 py-1.5">
      <span className="relative flex h-2 w-2">
        {isConnected && (
          <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75" />
        )}
        <span
          className={`relative inline-flex h-2 w-2 rounded-full ${
            isConnected ? "bg-emerald-400" : "bg-red-400"
          }`}
        />
      </span>
      <span className="text-xs font-medium text-gray-400">
        {isConnected ? chainName : "Disconnected"}
      </span>
    </div>
  );
}

function Breadcrumb() {
  const { pathname } = useLocation();
  const segments = pathname.split("/").filter(Boolean);

  if (segments.length === 0) return null;

  return (
    <nav className="hidden items-center gap-1 text-sm md:flex">
      {segments.map((segment, index) => {
        const href = "/" + segments.slice(0, index + 1).join("/");
        const isLast = index === segments.length - 1;
        const label = segment
          .split("-")
          .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
          .join(" ");

        return (
          <span key={href} className="flex items-center gap-1">
            {index > 0 && (
              <ChevronRight className="h-3.5 w-3.5 text-gray-600" />
            )}
            {isLast ? (
              <span className="rounded-md bg-white/[0.06] px-2 py-0.5 text-xs font-medium text-gray-200">
                {label}
              </span>
            ) : (
              <Link
                to={href}
                className="rounded-md px-2 py-0.5 text-xs text-gray-400 transition-colors hover:bg-white/[0.04] hover:text-gray-200"
              >
                {label}
              </Link>
            )}
          </span>
        );
      })}
    </nav>
  );
}

export function Header() {
  return (
    <header className="fixed left-0 right-0 top-0 z-50 bg-[#0a0a12]/80 backdrop-blur-xl">
      <div className="flex h-16 items-center justify-between px-6">
        <Link to="/" className="flex-shrink-0 transition-opacity hover:opacity-80">
          <LogoFull className="h-9" />
        </Link>

        <div className="absolute left-1/2 -translate-x-1/2">
          <Breadcrumb />
        </div>

        <div className="flex items-center gap-3">
          <NetworkIndicator />
          <ConnectButton
            accountStatus="avatar"
            chainStatus="none"
            showBalance={false}
          />
        </div>
      </div>

      <div className="h-[1px] w-full bg-gradient-to-r from-transparent via-indigo-500/40 to-transparent" />
    </header>
  );
}

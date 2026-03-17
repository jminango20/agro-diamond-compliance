import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { useReadContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useAssets } from "@/hooks/use-assets";
import { useRole } from "@/hooks/use-role";
import {
  useIndexerStatus,
  useIndexerProtocolEvents,
  useIndexerTokens,
} from "@/hooks/use-indexer";
import { ActivityFeed } from "@/components/ui/activity-feed";

const quickActions = [
  { title: "Asset Management", description: "Register and manage tokenized assets", href: "/admin/assets" },
  { title: "Identity Registry", description: "Register investor identities and KYC", href: "/admin/identity" },
  { title: "Compliance", description: "Claim topics, trusted issuers, modules", href: "/admin/compliance" },
  { title: "Supply Management", description: "Mint, burn, and forced transfers", href: "/admin/supply" },
  { title: "Security", description: "Pause, freeze, and role management", href: "/admin/security" },
  { title: "Asset Groups", description: "Hierarchical asset fractionalization", href: "/admin/groups" },
  { title: "Diamond Info", description: "Facets, selectors, and interfaces", href: "/admin/diamond" },
  { title: "Snapshots & Dividends", description: "Create snapshots and distribute dividends", href: "/admin/snapshots" },
];

export default function AdminDashboardPage() {
  const { assets, isLoading: assetsLoading } = useAssets();
  const { isOwner } = useRole();

  const { data: ownerAddress, isLoading: ownerLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "owner",
  });

  const { data: isProtocolPaused, isLoading: pauseLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "isProtocolPaused",
  });

  // Indexer data
  const { data: indexerStatus } = useIndexerStatus();
  const { data: indexerTokens } = useIndexerTokens();
  const { data: recentEvents, isLoading: eventsLoading } = useIndexerProtocolEvents({ first: 15 });

  const [indexerHealth, setIndexerHealth] = useState<"healthy" | "down" | "loading">("loading");

  useEffect(() => {
    const url = import.meta.env.VITE_INDEXER_URL;
    if (!url) {
      setIndexerHealth("down");
      return;
    }
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), 3000);
    fetch(`${url}/health`, { signal: controller.signal })
      .then((res) => {
        clearTimeout(timeout);
        setIndexerHealth(res.ok ? "healthy" : "down");
      })
      .catch(() => {
        clearTimeout(timeout);
        setIndexerHealth("down");
      });
    return () => { clearTimeout(timeout); controller.abort(); };
  }, []);

  // Compute total holders across all tokens
  const totalHolders = indexerTokens?.reduce((sum, t) => sum + t.holderCount, 0) ?? 0;

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Admin Dashboard</h1>

      {/* Protocol Status Cards */}
      <div className="mb-8 grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4 xl:grid-cols-6">
        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <p className="text-sm text-gray-400">Protocol Status</p>
          {pauseLoading ? (
            <p className="mt-1 text-lg text-gray-500">Loading...</p>
          ) : (
            <p className={`mt-1 text-lg font-semibold ${isProtocolPaused ? "text-red-400" : "text-green-400"}`}>
              {isProtocolPaused ? "Paused" : "Active"}
            </p>
          )}
        </div>

        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <p className="text-sm text-gray-400">Registered Assets</p>
          {assetsLoading ? (
            <p className="mt-1 text-lg text-gray-500">Loading...</p>
          ) : (
            <p className="mt-1 text-lg font-semibold text-indigo-400">{assets.length}</p>
          )}
        </div>

        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <p className="text-sm text-gray-400">Total Holders</p>
          <p className="mt-1 text-lg font-semibold text-indigo-400">{totalHolders}</p>
        </div>

        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <p className="text-sm text-gray-400">Owner</p>
          {ownerLoading ? (
            <p className="mt-1 text-lg text-gray-500">Loading...</p>
          ) : (
            <p className="mt-1 truncate text-sm font-mono text-indigo-400">
              {ownerAddress as string}
            </p>
          )}
        </div>

        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <p className="text-sm text-gray-400">Indexer</p>
          <p
            className={`mt-1 text-lg font-semibold ${
              indexerHealth === "healthy"
                ? "text-green-400"
                : indexerHealth === "down"
                  ? "text-red-400"
                  : "text-gray-500"
            }`}
          >
            {indexerHealth === "loading" ? "Checking..." : indexerHealth === "healthy" ? "Healthy" : "Unreachable"}
          </p>
        </div>

        <div className="rounded-xl bg-white/5 border border-white/10 p-6">
          <p className="text-sm text-gray-400">Last Indexed Block</p>
          <p className="mt-1 text-lg font-semibold text-indigo-400 font-mono">
            {indexerStatus ? indexerStatus.lastBlock.toLocaleString() : "-"}
          </p>
        </div>
      </div>

      {/* Main content: Quick Actions + Activity Feed */}
      <div className="grid grid-cols-1 gap-8 lg:grid-cols-3">
        {/* Quick Actions */}
        <div className="lg:col-span-2">
          <h2 className="mb-4 text-xl font-semibold text-white">Quick Actions</h2>
          <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
            {quickActions.map((action) => (
              <Link key={action.href} to={action.href}>
                <div className="rounded-xl bg-white/5 border border-white/10 p-6 transition-colors hover:border-indigo-400/50 hover:bg-white/[0.08]">
                  <h3 className="text-lg font-semibold text-indigo-400">{action.title}</h3>
                  <p className="mt-1 text-sm text-gray-400">{action.description}</p>
                </div>
              </Link>
            ))}
          </div>
        </div>

        {/* Activity Feed */}
        <div>
          <ActivityFeed
            events={recentEvents ?? []}
            isLoading={eventsLoading}
            title="Recent Protocol Activity"
            compact
          />
        </div>
      </div>
    </div>
  );
}

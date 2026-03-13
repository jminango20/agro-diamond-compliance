import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { useReadContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useAssets } from "@/hooks/use-assets";
import { useRole } from "@/hooks/use-role";

const quickActions = [
  { title: "Asset Management", description: "Register and manage tokenized assets", href: "/admin/assets" },
  { title: "Identity Registry", description: "Register investor identities and KYC", href: "/admin/identity" },
  { title: "Compliance", description: "Claim topics, trusted issuers, modules", href: "/admin/compliance" },
  { title: "Supply Management", description: "Mint, burn, and forced transfers", href: "/admin/supply" },
  { title: "Security", description: "Pause, freeze, and role management", href: "/admin/security" },
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

  const [indexerStatus, setIndexerStatus] = useState<"healthy" | "down" | "loading">("loading");

  useEffect(() => {
    const url = import.meta.env.VITE_INDEXER_URL;
    if (!url) {
      setIndexerStatus("down");
      return;
    }
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), 3000);
    fetch(`${url}/health`, { signal: controller.signal })
      .then((res) => {
        clearTimeout(timeout);
        setIndexerStatus(res.ok ? "healthy" : "down");
      })
      .catch(() => {
        clearTimeout(timeout);
        setIndexerStatus("down");
      });
    return () => { clearTimeout(timeout); controller.abort(); };
  }, []);

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Admin Dashboard</h1>

      {/* Protocol Status Cards */}
      <div className="mb-8 grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
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
              indexerStatus === "healthy"
                ? "text-green-400"
                : indexerStatus === "down"
                  ? "text-red-400"
                  : "text-gray-500"
            }`}
          >
            {indexerStatus === "loading" ? "Checking..." : indexerStatus === "healthy" ? "Healthy" : "Unreachable"}
          </p>
        </div>
      </div>

      {/* Quick Actions */}
      <h2 className="mb-4 text-xl font-semibold text-white">Quick Actions</h2>
      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
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
  );
}

import { Link, useLocation } from "react-router-dom";
import {
  LayoutDashboard,
  Coins,
  UserCheck,
  Shield,
  ArrowUpDown,
  Lock,
  Building2,
  Diamond,
  Camera,
  Wallet,
} from "lucide-react";

const navItems = [
  { href: "/admin", label: "Dashboard", icon: LayoutDashboard },
  { href: "/admin/assets", label: "Assets", icon: Coins },
  { href: "/admin/identity", label: "Identity", icon: UserCheck },
  { href: "/admin/compliance", label: "Compliance", icon: Shield },
  { href: "/admin/supply", label: "Supply", icon: ArrowUpDown },
  { href: "/admin/security", label: "Security", icon: Lock },
  { href: "/admin/groups", label: "Asset Groups", icon: Building2 },
  { href: "/admin/diamond", label: "Diamond", icon: Diamond },
  { href: "/admin/snapshots", label: "Snapshots", icon: Camera },
];

export function AdminSidebar() {
  const { pathname } = useLocation();

  return (
    <aside className="flex h-full w-64 flex-col border-r border-white/10 bg-white/5 backdrop-blur">
      <nav className="flex flex-1 flex-col gap-1 p-4">
        {navItems.map(({ href, label, icon: Icon }) => {
          const isActive =
            href === "/admin"
              ? pathname === "/admin"
              : pathname.startsWith(href);

          return (
            <Link
              key={href}
              to={href}
              className={`flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium transition-colors ${
                isActive
                  ? "bg-indigo-500/15 text-indigo-400"
                  : "text-gray-400 hover:bg-white/5 hover:text-gray-200"
              }`}
            >
              <Icon className="h-4.5 w-4.5" />
              {label}
            </Link>
          );
        })}
      </nav>

      <div className="border-t border-white/10 p-4">
        <Link
          to="/portfolio"
          className="flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-gray-400 transition-colors hover:bg-white/5 hover:text-gray-200"
        >
          <Wallet className="h-4.5 w-4.5" />
          Switch to Portfolio
        </Link>
      </div>
    </aside>
  );
}

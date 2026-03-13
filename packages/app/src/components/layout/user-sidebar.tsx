import { Link, useLocation } from "react-router-dom";
import { Wallet, Send, DollarSign, Shield } from "lucide-react";

const navItems = [
  { href: "/portfolio", label: "Portfolio", icon: Wallet },
  { href: "/portfolio/transfer", label: "Transfer", icon: Send },
  { href: "/portfolio/dividends", label: "Dividends", icon: DollarSign },
];

export function UserSidebar() {
  const { pathname } = useLocation();

  return (
    <aside className="flex h-full w-64 flex-col border-r border-white/10 bg-white/5 backdrop-blur">
      <nav className="flex flex-1 flex-col gap-1 p-4">
        {navItems.map(({ href, label, icon: Icon }) => {
          const isActive =
            href === "/portfolio"
              ? pathname === "/portfolio"
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
          to="/admin"
          className="flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-gray-400 transition-colors hover:bg-white/5 hover:text-gray-200"
        >
          <Shield className="h-4.5 w-4.5" />
          Switch to Admin
        </Link>
      </div>
    </aside>
  );
}

import { useState, useMemo } from "react";
import { Search, Copy, AlertTriangle } from "lucide-react";
import type { SelectorMapData } from "../types";

interface Props {
  data: SelectorMapData;
}

export function SelectorsTab({ data }: Props) {
  const [search, setSearch] = useState("");
  const [facetFilter, setFacetFilter] = useState<string>("all");

  const facets = useMemo(() => {
    return [...new Set(data.entries.map((e) => e.facet))].sort();
  }, [data]);

  const filtered = useMemo(() => {
    return data.entries.filter((e) => {
      if (facetFilter !== "all" && e.facet !== facetFilter) return false;
      if (search) {
        const q = search.toLowerCase();
        return (
          e.selector.toLowerCase().includes(q) ||
          e.signature.toLowerCase().includes(q) ||
          e.facet.toLowerCase().includes(q)
        );
      }
      return true;
    });
  }, [data, search, facetFilter]);

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
  };

  return (
    <div className="p-6 space-y-4">
      {data.collisions.length > 0 && (
        <div className="bg-red-500/10 border border-red-500/30 rounded-lg p-4">
          <div className="flex items-center gap-2 mb-2">
            <AlertTriangle size={14} className="text-red-400" />
            <span className="text-sm font-semibold text-red-400">
              {data.collisions.length} Selector Collision(s)
            </span>
          </div>
          {data.collisions.map((c) => (
            <p key={c.selector} className="text-xs text-red-300">
              <code className="bg-red-500/20 px-1 rounded">{c.selector}</code> in{" "}
              {c.facets.join(", ")}
            </p>
          ))}
        </div>
      )}

      <div className="flex items-center gap-3">
        <div className="relative flex-1">
          <Search size={14} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
          <input
            type="text"
            placeholder="Search selector, signature, or facet..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-full pl-9 pr-3 py-2 text-xs bg-surface-card border border-surface-border rounded-lg text-gray-200 placeholder:text-gray-600 focus:outline-none focus:border-accent/50"
          />
        </div>
        <select
          value={facetFilter}
          onChange={(e) => setFacetFilter(e.target.value)}
          className="px-3 py-2 text-xs bg-surface-card border border-surface-border rounded-lg text-gray-200 focus:outline-none focus:border-accent/50"
        >
          <option value="all">All Facets ({data.facetCount})</option>
          {facets.map((f) => (
            <option key={f} value={f}>{f}</option>
          ))}
        </select>
      </div>

      <div className="text-xs text-gray-500">
        {filtered.length} of {data.totalSelectors} selectors
      </div>

      <div className="bg-surface-card border border-surface-border rounded-lg overflow-hidden">
        <table className="w-full text-xs">
          <thead>
            <tr className="text-gray-500 border-b border-surface-border/50">
              <th className="text-left p-2.5 pl-4 w-28">Selector</th>
              <th className="text-left p-2.5">Signature</th>
              <th className="text-left p-2.5 pr-4">Facet</th>
            </tr>
          </thead>
          <tbody>
            {filtered.map((e, i) => {
              const isCollision = data.collisions.some((c) => c.selector === e.selector);
              return (
                <tr
                  key={i}
                  className={`border-b border-surface-border/30 hover:bg-surface-hover/50 ${
                    isCollision ? "bg-red-500/5" : ""
                  }`}
                >
                  <td className="p-2.5 pl-4">
                    <button
                      onClick={() => copyToClipboard(e.selector)}
                      className="flex items-center gap-1.5 text-accent hover:text-accent-dim group"
                    >
                      <code>{e.selector}</code>
                      <Copy size={10} className="opacity-0 group-hover:opacity-100 transition-opacity" />
                    </button>
                  </td>
                  <td className="p-2.5 text-gray-200">
                    <code>{e.signature}</code>
                  </td>
                  <td className="p-2.5 pr-4 text-purple-400">{e.facet}</td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
}

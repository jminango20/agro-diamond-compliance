import { useReadContract } from "wagmi";

import { DIAMOND_ADDRESS, diamondAbi } from "@/config/contracts";
import { useDiamondLoupe } from "@/hooks/use-diamond-loupe";

const INTERFACE_IDS: { name: string; id: `0x${string}` }[] = [
  { name: "ERC165", id: "0x01ffc9a7" },
  { name: "ERC1155", id: "0xd9b67a26" },
  { name: "ERC1155MetadataURI", id: "0x0e89341c" },
  { name: "DiamondCut", id: "0x1f931c1c" },
  { name: "DiamondLoupe", id: "0x48e2b093" },
];

export default function DiamondPage() {
  const { facets, isLoading: facetsLoading } = useDiamondLoupe();

  const { data: ownerAddress, isLoading: ownerLoading } = useReadContract({
    address: DIAMOND_ADDRESS,
    abi: diamondAbi,
    functionName: "owner",
  });

  // Check interface support
  const interfaceChecks = INTERFACE_IDS.map((iface) => {
    const { data } = useReadContract({
      address: DIAMOND_ADDRESS,
      abi: diamondAbi,
      functionName: "supportsInterface",
      args: [iface.id],
    });
    return { ...iface, supported: data as boolean | undefined };
  });

  return (
    <div className="min-h-screen bg-[#0a0a0f] p-8">
      <h1 className="mb-8 text-3xl font-bold text-white">Diamond Info</h1>

      {/* Owner */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-2 text-xl font-semibold text-indigo-400">Owner</h2>
        {ownerLoading ? (
          <p className="text-gray-500">Loading...</p>
        ) : (
          <p className="font-mono text-sm text-gray-300">{ownerAddress as string}</p>
        )}
        <p className="mt-2 text-sm text-gray-400">
          Diamond Address: <span className="font-mono text-indigo-400">{DIAMOND_ADDRESS}</span>
        </p>
      </div>

      {/* Interface Support */}
      <div className="mb-6 rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">Interface Support (ERC-165)</h2>
        <div className="overflow-x-auto">
          <table className="w-full text-left text-sm text-gray-300">
            <thead>
              <tr className="border-b border-white/10 text-gray-400">
                <th className="pb-3 pr-4">Interface</th>
                <th className="pb-3 pr-4">ID</th>
                <th className="pb-3">Supported</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-white/10">
              {interfaceChecks.map((iface) => (
                <tr key={iface.id}>
                  <td className="py-3 pr-4">{iface.name}</td>
                  <td className="py-3 pr-4 font-mono text-xs">{iface.id}</td>
                  <td className="py-3">
                    {iface.supported === undefined ? (
                      <span className="text-gray-500">Checking...</span>
                    ) : iface.supported ? (
                      <span className="text-green-400">Yes</span>
                    ) : (
                      <span className="text-red-400">No</span>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      {/* Facets */}
      <div className="rounded-xl bg-white/5 border border-white/10 p-6">
        <h2 className="mb-4 text-xl font-semibold text-indigo-400">
          Facets {!facetsLoading && `(${facets.length})`}
        </h2>
        {facetsLoading ? (
          <p className="text-gray-500">Loading facets...</p>
        ) : facets.length === 0 ? (
          <p className="text-gray-500">No facets found.</p>
        ) : (
          <div className="overflow-x-auto">
            <table className="w-full text-left text-sm text-gray-300">
              <thead>
                <tr className="border-b border-white/10 text-gray-400">
                  <th className="pb-3 pr-4">Facet Address</th>
                  <th className="pb-3 pr-4">Selectors Count</th>
                  <th className="pb-3">Selectors</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-white/10">
                {facets.map((facet) => (
                  <tr key={facet.facetAddress}>
                    <td className="py-3 pr-4 font-mono text-xs">{facet.facetAddress}</td>
                    <td className="py-3 pr-4 text-indigo-400">
                      {facet.functionSelectors.length}
                    </td>
                    <td className="py-3">
                      <div className="flex flex-wrap gap-1">
                        {facet.functionSelectors.map((sel) => (
                          <span
                            key={sel}
                            className="inline-block rounded bg-white/5 px-1.5 py-0.5 font-mono text-xs text-gray-400"
                          >
                            {sel}
                          </span>
                        ))}
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

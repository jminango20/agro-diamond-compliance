import { getDefaultConfig } from "@rainbow-me/rainbowkit";
import { http } from "wagmi";
import { polygonAmoy } from "wagmi/chains";

const rpcUrl =
  import.meta.env.VITE_RPC_URL ??
  "https://rpc-amoy.polygon.technology";

export const config = getDefaultConfig({
  appName: "Diamond ERC-3643",
  projectId: "demo",
  chains: [polygonAmoy],
  transports: {
    [polygonAmoy.id]: http(rpcUrl),
  },
});

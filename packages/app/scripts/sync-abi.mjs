import { readdirSync, readFileSync, writeFileSync, mkdirSync, existsSync } from "node:fs";
import { resolve, basename, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const appRoot = resolve(__dirname, "..");
const destDir = resolve(appRoot, "src", "abi");

const indexerAbiDir = resolve(appRoot, "..", "indexer", "abi");
const contractsOutDir = resolve(appRoot, "..", "contracts", "out");

const forgeArtifacts = [
  "DiamondLoupeFacet.sol/DiamondLoupeFacet.json",
  "DiamondCutFacet.sol/DiamondCutFacet.json",
  "AssetGroupFacet.sol/AssetGroupFacet.json",
];

mkdirSync(destDir, { recursive: true });

let copied = 0;

// 1. Copy all ABIs from indexer/abi/
if (existsSync(indexerAbiDir)) {
  const files = readdirSync(indexerAbiDir).filter((f) => f.endsWith(".json"));
  for (const file of files) {
    const src = resolve(indexerAbiDir, file);
    const dest = resolve(destDir, file);
    const content = readFileSync(src, "utf-8");
    writeFileSync(dest, content);
    copied++;
    console.log(`[indexer] ${file}`);
  }
} else {
  console.warn(`WARN: indexer ABI dir not found at ${indexerAbiDir}`);
}

// 2. Extract ABI from Foundry build artifacts
for (const artifact of forgeArtifacts) {
  const src = resolve(contractsOutDir, artifact);
  if (!existsSync(src)) {
    console.warn(`WARN: artifact not found: ${artifact}`);
    continue;
  }

  const json = JSON.parse(readFileSync(src, "utf-8"));
  const abi = json.abi;

  if (!abi) {
    console.warn(`WARN: no ABI field in ${artifact}`);
    continue;
  }

  const name = basename(artifact, ".json");
  const destFile = resolve(destDir, `${name}.abi.json`);
  writeFileSync(destFile, JSON.stringify(abi, null, 2) + "\n");
  copied++;
  console.log(`[forge]   ${name}.abi.json`);
}

console.log(`\nSynced ${copied} ABI files to src/abi/`);

import { Routes, Route } from "react-router-dom";

import { AdminLayout } from "@/components/layout/admin-layout";
import { PortfolioLayout } from "@/components/layout/portfolio-layout";

import HomePage from "@/pages/home";
import AdminDashboardPage from "@/pages/admin/dashboard";
import AssetsPage from "@/pages/admin/assets";
import AssetDetailPage from "@/pages/admin/asset-detail";
import CompliancePage from "@/pages/admin/compliance";
import DiamondPage from "@/pages/admin/diamond";
import GroupsPage from "@/pages/admin/groups";
import IdentityPage from "@/pages/admin/identity";
import SecurityPage from "@/pages/admin/security";
import SnapshotsPage from "@/pages/admin/snapshots";
import SupplyPage from "@/pages/admin/supply";
import PortfolioPage from "@/pages/portfolio/portfolio";
import TokenDetailPage from "@/pages/portfolio/token-detail";
import DividendsPage from "@/pages/portfolio/dividends";
import TransferPage from "@/pages/portfolio/transfer";

export function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />

      <Route element={<AdminLayout />}>
        <Route path="/admin" element={<AdminDashboardPage />} />
        <Route path="/admin/assets" element={<AssetsPage />} />
        <Route path="/admin/assets/:tokenId" element={<AssetDetailPage />} />
        <Route path="/admin/compliance" element={<CompliancePage />} />
        <Route path="/admin/diamond" element={<DiamondPage />} />
        <Route path="/admin/groups" element={<GroupsPage />} />
        <Route path="/admin/identity" element={<IdentityPage />} />
        <Route path="/admin/security" element={<SecurityPage />} />
        <Route path="/admin/snapshots" element={<SnapshotsPage />} />
        <Route path="/admin/supply" element={<SupplyPage />} />
      </Route>

      <Route element={<PortfolioLayout />}>
        <Route path="/portfolio" element={<PortfolioPage />} />
        <Route path="/portfolio/:tokenId" element={<TokenDetailPage />} />
        <Route path="/portfolio/dividends" element={<DividendsPage />} />
        <Route path="/portfolio/transfer" element={<TransferPage />} />
      </Route>
    </Routes>
  );
}

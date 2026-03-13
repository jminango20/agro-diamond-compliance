import { Outlet } from "react-router-dom";
import { Header } from "@/components/layout/header";
import { AdminSidebar } from "@/components/layout/admin-sidebar";

export function AdminLayout() {
  return (
    <div className="flex min-h-screen flex-col">
      <Header />
      <div className="flex flex-1 pt-16">
        <AdminSidebar />
        <main className="flex-1 overflow-y-auto p-6">
          <Outlet />
        </main>
      </div>
    </div>
  );
}

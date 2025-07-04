import { Button } from "@/components/ui/button"
import { HomeIcon, BarChart3Icon, UsersIcon, SettingsIcon } from "lucide-react"
import Link from "next/link"

const menuItems = [
  { name: "Ana Sayfa", icon: HomeIcon, path: "/" },
  { name: "Jobs", icon: UsersIcon, path: "/users" },
  /*{ name: "", icon: BarChart3Icon, path: "/statistics" },
  { name: "Ayarlar", icon: SettingsIcon, path: "/settings" },*/
]

export function Sidebar() {
  return (
    <div className="pb-12 min-h-screen">
      <div className="space-y-4 py-4">
        <div className="px-3 py-2">
          <h2 className="mb-2 px-4 text-lg font-semibold">Dashboard</h2>
          <div className="space-y-1">
            {menuItems.map((item) => (
              <Link key={item.path} href={item.path}>
                <Button variant="ghost" className="w-full justify-start">
                  <item.icon className="mr-2 h-4 w-4" />
                  {item.name}
                </Button>
              </Link>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}

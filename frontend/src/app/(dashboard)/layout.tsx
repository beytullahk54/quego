'use client' // Bu satırla component'i Client Component olarak işaretle

import { Sidebar } from "@/components/sidebar"
import { Topbar } from "@/components/topbar"
import { redirect } from 'next/navigation'

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {

  
  return (
    <div className="flex h-screen">
      <div className="w-64 border-r">
        <Sidebar />
      </div>
      
      <div className="flex-1 flex flex-col">
        <Topbar />
        <div className="flex-1 p-8">
          {children}
        </div>
      </div>
    </div>
  )
}

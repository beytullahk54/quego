"use client"

import {  User, LogOut, UserCircle } from "lucide-react"
import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { AuthService } from "@/lib/auth/AuthService";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export function Topbar() {
  const router = useRouter();
  //const { data: session } = useSession();
  const auth = AuthService.getInstance();
  const [userName, setUserName] = useState<string | null>(null);

  useEffect(() => {
    // Sadece client tarafında çalışacak
    const user = auth.getUser();
    setUserName(user?.name || null);
  }, []);

  const handleSignOut = async () => {
    await auth.logout();
    router.push('/login');
  };

  return (
    <div className="border-b">
      <div className="flex h-16 items-center px-4">
        <div className="flex-1">
          {/* <div className="relative w-96">
            <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Ara..."
              className="pl-8"
            />
          </div>*/}
        </div>
        <div className="flex items-center space-x-4">
          {/* <Button variant="ghost" size="icon">
            <Bell className="h-5 w-5" />
          </Button> */}
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" className="flex items-center gap-2">
                <User className="h-5 w-5" />
                <span className="font-medium">
                  {userName}
                </span>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuItem>
                <UserCircle className="mr-2 h-4 w-4" />
                <span>Profilim</span>
              </DropdownMenuItem>
              <DropdownMenuItem onClick={handleSignOut}>
                <LogOut className="mr-2 h-4 w-4"  />
                <span>Çıkış Yap</span>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>
    </div>
  )
} 
"use client"

import { Table, TableCaption, TableHeader, TableRow, TableHead, TableBody, TableCell } from "@/components/ui/table"
import { useUserStore } from "@/store/useUserStore";
import { useEffect, useCallback } from "react";
import { Button } from "@/components/ui/button";
import { Pencil, Trash2 } from "lucide-react";
import { AuthService } from "@/lib/auth/AuthService";

export default function Users() {
  const { users, isLoading, error, fetchUsers } = useUserStore();
  const auth = AuthService.getInstance();

  const fetchUsersCallback = useCallback(() => {
    const token = auth.getToken();
    fetchUsers(token || '');
    
  }, [ fetchUsers]);

  useEffect(() => {
    fetchUsersCallback();
  }, [fetchUsersCallback]);

  if (isLoading) {
    return (
      <>
        <h1 className="text-3xl font-bold mb-8">Kullanıcılar</h1>
        <div>Yükleniyor...</div>
      </>
    );
  }

  if (error) {
    return (
      <>
        <h1 className="text-3xl font-bold mb-8">Kullanıcılar</h1>
        <div className="text-red-500">{error}</div>
      </>
    );
  }
  
  return (
    <>
      <h1 className="text-3xl font-bold mb-8">Kullanıcılar</h1>
      
      <div className="rounded-md border">
        <Table>
          <TableCaption>Toplam {users.length} kullanıcı</TableCaption>
          
          <TableHeader>
            <TableRow>
              <TableHead>Adı</TableHead>
              <TableHead>Email</TableHead>
              <TableHead>Kayıt Tarihi</TableHead>
              <TableHead className="text-right">İşlemler</TableHead>
            </TableRow>
          </TableHeader>
          
          <TableBody>
            {users.map((user) => (
              <TableRow key={user.id}>
                <TableCell className="font-medium">{user.name}</TableCell>
                <TableCell>{user.email}</TableCell>
                <TableCell>
                  {new Date(user.created_at).toLocaleDateString('tr-TR')}
                </TableCell>
                <TableCell className="text-right">
                  <div className="flex justify-end gap-2">
                    <Button variant="ghost" size="icon">
                      <Pencil className="h-4 w-4" />
                    </Button>
                    <Button variant="ghost" size="icon" className="text-red-500 hover:text-red-700">
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </>
  );
}

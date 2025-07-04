'use client'
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { AuthService } from "@/lib/auth/AuthService";
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export default function Home() {
  const auth = AuthService.getInstance();
  const router = useRouter();

  useEffect(() => {
    const token = auth.getToken();
    if (!token) {
      router.push('/login');
    }
  }, [router]);

  return (
      <div className="flex-1 p-4">
        <h1 className="text-3xl font-bold mb-8">Hoş Geldiniz</h1>
        
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                Toplam Gelir
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">45.231 ₺</div>
              <p className="text-xs text-muted-foreground">
                +20.1% geçen aydan
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                Aktif Kullanıcılar
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">+2350</div>
              <p className="text-xs text-muted-foreground">
                +180 son 24 saatte
              </p>
            </CardContent>
          </Card>

          {/* Diğer kartlar buraya eklenebilir */}
        </div>
      </div>
  )
}

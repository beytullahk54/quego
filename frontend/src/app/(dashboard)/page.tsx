'use client'
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { AuthService } from "@/lib/auth/AuthService";
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';
import { Alert, AlertTitle, AlertDescription } from "@/components/ui/alert";
import { CheckCircle2Icon } from "lucide-react";

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

        <Alert>
          <CheckCircle2Icon />
          <AlertTitle>İşler hakkında</AlertTitle>
          <AlertDescription>
            İşlerinizi "işler" menüsü altından ekleyebilir,, güncelleyebilirsiniz.
          </AlertDescription>
        </Alert>
      </div>
  )
}

import '@/app/globals.css'
import { Inter } from 'next/font/google'
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const inter = Inter({ subsets: ['latin'] })

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {

  return (
    <html lang="tr" suppressHydrationWarning>
      <body className={inter.className} suppressHydrationWarning>
        <div className="min-h-screen bg-background">
          {children}
        </div>
        <ToastContainer />
      </body>
    </html>
  )
}

// Metadata ekleyelim
export const metadata = {
  title: 'Dashboard',
  description: 'Dashboard uygulaması',
}

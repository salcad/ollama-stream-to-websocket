import './globals.css';
import { Inter } from "next/font/google";
import ClientComponent from './clientComponent';

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: 'Ollama Chat',
  description: 'Ollama chat interface',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ClientComponent>{children}</ClientComponent>
      </body>
    </html>
  );
}

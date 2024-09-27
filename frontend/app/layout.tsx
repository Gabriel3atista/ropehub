import "./globals.css";
import type { Metadata } from "next";
import { Open_Sans, Roboto_Mono } from "next/font/google";

const openSans = Open_Sans({
  subsets: ['latin'],
  variable: '--font-sans',
});

const robotoMono = Roboto_Mono({
  subsets: ['latin'],
  variable: '--font-mono',
});

export const metadata: Metadata = {
  title: "Rope Hub",
  description: "Transforme o mundo ao seu redor!",
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <head />
      <body
        className={`${openSans.variable} ${robotoMono.variable} antialiased text-balance`}
      >
        {children}
      </body>
    </html>
  );
}

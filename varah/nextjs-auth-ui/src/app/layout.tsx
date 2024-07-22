// src/app/layout.tsx

import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ReactNode } from 'react';
import Head from 'next/head';

interface LayoutProps {
    children: ReactNode;
}

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Perfbit App",
    description: "Perfbit is an advanced Developer Performance Analytics Tool that offers in-depth insights into developer performance, code contributions, and overall development tracking. Designed for project managers, product owners, company owners, and team leads, Perfbit helps enhance team efficiency and productivity.",
};

const RootLayout = ({ children }: LayoutProps) => {
    return (
        <html lang="en">
        <Head>
            <title>Perfbit App</title>
            <link rel="icon" href="/siteicon.svg" type="image/svg+xml" />
        </Head>
        <body className={inter.className}>
        {children}
        </body>
        </html>
    );
};

export default RootLayout;
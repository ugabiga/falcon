import type {Metadata} from 'next'
import {Inter} from 'next/font/google'
import './globals.css'
import React from "react";
import Providers from "@/app/providers";
import {NavigationBar} from "@/components/navigation-bar";
import {Toaster} from "@/components/ui/sonner"
import {getServerSession} from "next-auth";
import SessionProvider from "@/lib/session"
import { Analytics } from '@vercel/analytics/react';

const inter = Inter({subsets: ['latin']})

export const metadata: Metadata = {
    title: 'Falcon',
    description: 'Falcon is a crypto trading platform',
    manifest: '/manifest.json',
    icons: [
        {
            rel: 'icon',
            type: 'image/png',
            sizes: '32x32',
            url: process.env.PUBLIC_URL + '/favicon-32x32.png',
        },
        {
            rel: 'icon',
            type: 'image/png',
            sizes: '16x16',
            url: process.env.PUBLIC_URL + '/favicon-16x16.png',
        },
        {
            rel: 'apple-touch-icon',
            sizes: '180x180',
            url: process.env.PUBLIC_URL + '/apple-touch-icon.png',
        },
        {
            rel: 'mask-icon',
            url: process.env.PUBLIC_URL + '/safari-pinned-tab.svg',
            color: '#5bbad5',
        },
    ],
}


export default async function RootLayout({children,}: { children: React.ReactNode }) {
    const session = await getServerSession()

    return (
        <html lang="en" suppressHydrationWarning>
        <body className={inter.className}>
        <Providers>
            <SessionProvider session={session}>
                <NavigationBar/>
                {children}
                <Analytics />
                <Toaster
                    richColors={true}
                />
            </SessionProvider>
        </Providers>
        </body>
        </html>
    )
}

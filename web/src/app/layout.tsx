import type {Metadata, Viewport} from 'next'
import {Inter} from 'next/font/google'
import './globals.css'
import React, {Suspense} from "react";
import Providers from "@/app/providers";
import {Toaster} from "@/components/ui/sonner"
import {getServerSession} from "next-auth";
import SessionProvider from "@/lib/session"
import {Analytics} from '@vercel/analytics/react';
import {SpeedInsights} from "@vercel/speed-insights/next"
import NavigationBar from "@/components/navigation-bar";
import GtmAnalytics from "@/lib/gtm-head";

const inter = Inter({subsets: ['latin']})

export const viewport: Viewport = {
    width: 'device-width',
    initialScale: 1,
    minimumScale: 1,
    maximumScale: 1,
}


export const metadata: Metadata = {
    title: 'Falcon',
    description: 'Falcon is a crypto trading platform',
    manifest: '/manifest.json',
    icons: [
        {
            rel: 'icon',
            type: 'image/png',
            sizes: 'any',
            url: '/favicon-32x32.png',
        },
        {
            rel: 'icon',
            type: 'image/png',
            sizes: '32x32',
            url: '/favicon-32x32.png',
        },
        {
            rel: 'icon',
            type: 'image/png',
            sizes: '16x16',
            url: '/favicon-16x16.png',
        },
        {
            rel: 'apple-touch-icon',
            sizes: '180x180',
            url: '/apple-touch-icon.png',
        },
        // {
        //     rel: 'mask-icon',
        //     url: '/safari-pinned-tab.svg',
        //     color: '#5bbad5',
        // },
    ],
}


export default async function RootLayout(
    {
        children,
    }: {
        children: React.ReactNode
    }) {
    const session = await getServerSession()

    return (
        <html lang="en" suppressHydrationWarning>

        <body className={inter.className}>
        <Providers>
            <SessionProvider session={session}>
                <NavigationBar/>
                {children}
                <Analytics/>
                <SpeedInsights/>
                <Toaster
                    richColors={true}
                />
                <Suspense>
                    <GtmAnalytics/>
                </Suspense>
            </SessionProvider>
        </Providers>
        </body>
        </html>
    )
}

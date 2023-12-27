"use client";

import React from "react";
import {Card, CardContent, CardHeader} from "@/components/ui/card";

export default function Home() {
    return (
        <main className="flex-1 flex flex-col min-h-screen">
            <section className="w-full py-12 md:py-24 lg:py-32">
                <div className="container space-y-12 px-4 md:px-6 justify-center">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        Welcome to CryptoDCA
                    </h2>
                    <p className="text-center max-w-[900px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 mx-auto">
                        Take control of your cryptocurrency investments with our Dollar Cost Averaging (DCA)
                        platform.
                    </p>
                </div>
            </section>
            <section className="w-full py-12 md:py-24 lg:py-32 border-y">
                <div className="container space-y-12 px-4 md:px-6">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        What is Dollar Cost Averaging (DCA)?
                    </h2>
                    <p className="max-w-[900px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 mx-auto">
                        Dollar Cost Averaging (DCA) is an investment strategy that involves purchasing a fixed
                        amount of an asset
                        on a regular schedule, regardless of the asset's price. This technique can help mitigate the
                        impact of
                        volatility on large purchases of financial assets such as equities or cryptocurrencies.
                    </p>
                </div>
            </section>
            <section className="w-full py-12 md:py-24 lg:py-32 bg-gray-100 dark:bg-gray-800">
                <div className="container space-y-12 px-4 md:px-6">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        Available Cryptocurrency Exchanges
                    </h2>
                    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-8">
                        <Card>
                            <CardHeader>
                                <h3 className="text-xl font-semibold">Upbit</h3>
                            </CardHeader>
                            <CardContent>
                                <p className="text-gray-500">
                                    A leading cryptocurrency exchange that offers a wide variety of digital assets.
                                </p>
                            </CardContent>
                        </Card>
                        <Card>
                            <CardHeader>
                                <h3 className="text-xl font-semibold">Binance</h3>
                            </CardHeader>
                            <CardContent>
                                <p className="text-gray-500">
                                    The world's largest cryptocurrency exchange by trading volume and users.
                                </p>
                            </CardContent>
                        </Card>
                    </div>
                </div>
            </section>
        </main>
    )
}
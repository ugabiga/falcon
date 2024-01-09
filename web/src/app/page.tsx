"use client";

import React from "react";
import {Card, CardContent, CardHeader} from "@/components/ui/card";
import {useTranslation} from "react-i18next";
import {Button} from "@/components/ui/button";
import Link from "next/link";
import {ManualKRMain} from "@/lib/ref-url";

export default function Home() {
    const {t} = useTranslation();
    return (
        <main className="flex-1 flex flex-col min-h-screen">
            <section className="w-full py-12 md:py-24 lg:py-24">
                <div className="container space-y-12 px-4 md:px-6 justify-center flex flex-col items-center">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        {t("home.title")}
                    </h2>
                    <h3 className="text-xl font-semibold text-center sm:text-2xl">
                        {t("home.subTitle")}
                    </h3>
                    <p className="text-center max-w-[900px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 mx-auto">
                        {t("home.description")}
                    </p>
                    <Button onClick={() => {
                        window.open(ManualKRMain, '_blank')
                    }}>
                        {t("home.find_out_more.btn")}
                    </Button>
                </div>
            </section>
            <section className="w-full py-12 md:py-24 lg:py-24 border-y">
                <div className="container space-y-12 px-4 md:px-6">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        {t("home.dca.title")}
                    </h2>
                    <p className="max-w-[900px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 mx-auto">
                        {t("home.dca.description")}
                    </p>
                </div>
            </section>
            <section className="w-full py-12 md:py-24 lg:py-32">
                <div className="container space-y-12 px-4 md:px-6">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        {t("home.exchanges.title")}
                    </h2>
                    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-8">
                        <Card>
                            <CardHeader>
                                <h3 className="text-xl font-semibold">
                                    {t("home.exchanges.upbit.title")}
                                </h3>
                            </CardHeader>
                            <CardContent>
                                <p className="text-gray-500">
                                    {t("home.exchanges.upbit.description")}
                                </p>
                            </CardContent>
                        </Card>
                        <Card>
                            <CardHeader>
                                <h3 className="text-xl font-semibold">
                                    {t("home.exchanges.binance.title")}
                                </h3>
                            </CardHeader>
                            <CardContent>
                                <p className="text-gray-500">
                                    {t("home.exchanges.binance.description")}
                                </p>
                            </CardContent>
                        </Card>
                    </div>
                </div>
            </section>
        </main>
    )
}
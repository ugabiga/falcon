import React from "react";
import {Card, CardContent, CardHeader} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import {ManualKRMain} from "@/lib/ref-url";
import {useTranslation} from "@/lib/i18n-server";
import {capture} from "@/lib/posthog";

export default async function Home() {
    const {t} = await useTranslation()

    capture('home_page_viewed')

    return (
        <main className="flex-1 flex flex-col min-h-screen">
            {/* Hero */}
            <section className="w-full py-12 md:py-24 lg:py-24">
                <div className="container space-y-12 px-4 md:px-6 justify-center flex flex-col items-center">
                    <h1 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        {t("home.title")}
                    </h1>
                    <h3 className="text-xl font-semibold text-center sm:text-2xl">
                        {t("home.subTitle")}
                    </h3>
                    <p className="text-center max-w-[900px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 mx-auto">
                        <span className="mb-2">
                            {t("home.description.1")}
                        </span>
                        <span>
                            {t("home.description.2")}
                        </span>
                    </p>
                </div>
            </section>

            {/* Why */}
            <section className="w-full py-12 md:py-24 lg:py-24 border-y">
                <div className="container space-y-12 px-4 md:px-6">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        {t("home.why.title")}
                    </h2>
                    <p className="max-w-[1000px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 mx-auto">
                        <span className="mb-2">
                            {t("home.why.1.problem")}
                        </span>
                        <span>
                            {t("home.why.1.solution")}
                        </span>
                    </p>
                    <p className="max-w-[1000px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400 mx-auto">
                        <span className="mb-2">
                            {t("home.why.2.problem")}
                        </span>
                        <span>
                            {t("home.why.2.solution")}
                        </span>
                    </p>
                </div>
            </section>

            {/* Strategy */}
            <section className="w-full py-12 md:py-24 lg:py-24 border-y ">
                <div className="container space-y-12 px-4 md:px-6">
                    <h2 className="text-3xl font-bold tracking-tighter text-center sm:text-5xl">
                        {t("home.strategy.title")}
                    </h2>
                    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-8">
                        <Card>
                            <CardHeader>
                                <h3 className="text-xl font-semibold">
                                    {t("home.strategy.dca.title")}
                                </h3>
                            </CardHeader>
                            <CardContent>
                                <p className="text-gray-500">
                                    {t("home.strategy.dca.description")}
                                </p>
                            </CardContent>
                        </Card>
                        <Card>
                            <CardHeader>
                                <h3 className="text-xl font-semibold">
                                    {t("home.strategy.buying_grid.title")}
                                </h3>
                            </CardHeader>
                            <CardContent>
                                <p className="text-gray-500">
                                    {t("home.strategy.buying_grid.description")}
                                </p>
                            </CardContent>
                        </Card>
                    </div>
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
                                    {t("home.exchanges.binance_future.title")}
                                </h3>
                            </CardHeader>
                            <CardContent>
                                <p className="text-gray-500">
                                    {t("home.exchanges.binance_future.description")}
                                </p>
                            </CardContent>
                        </Card>
                    </div>
                </div>
            </section>

            {/* Footer */}
            <div className="container md:h-12 items-center">
                <p className="text-balance text-center text-sm leading-loose text-muted-foreground">
                    {t("home.contact.title")}
                    <a
                        href="mailto:vultor.xyz@gmail.com"
                        target="_blank" rel="noreferrer"
                        className="font-medium underline underline-offset-4">
                        vultor.xyz@gmail.com
                    </a>
                </p>
            </div>

        </main>
    )
}
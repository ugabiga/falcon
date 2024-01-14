"use client";

import {TaskTable} from "@/app/tasks/table";
import {useQuery} from "@apollo/client";
import React, {useEffect} from "react";
import {GetTaskIndexDocument} from "@/graph/generated/generated";
import {TradingAccountSelector} from "@/app/tasks/selector";
import {useAppSelector} from "@/store";
import {useDispatch} from "react-redux";
import {refreshTask} from "@/store/taskSlice";
import {AddTask} from "@/app/tasks/add";
import {Loading} from "@/components/loading";
import {useSearchParams} from "next/navigation";
import {Error} from "@/components/error";
import {useTranslation} from "react-i18next";
import {Button} from "@/components/ui/button";
import {ManualKRTask} from "@/lib/ref-url";
import {TaskCards} from "@/app/tasks/cards";
import {RefreshTarget} from "@/store/refresherSlice";

export default function Tasks() {
    const {t} = useTranslation()
    const params = useSearchParams()
    const dispatch = useDispatch()
    const tradingAccountId = params.get('trading_account_id')
    const refresher = useAppSelector((state) => state.refresher)
    const {data, loading, refetch, error} = useQuery(GetTaskIndexDocument, {
        variables: {
            tradingAccountID: tradingAccountId ?? null
        },
        fetchPolicy: "no-cache"
    })

    useEffect(() => {
        if (refresher?.refresh && refresher?.targetName === RefreshTarget.Task) {
            refetch({
                tradingAccountID: refresher.params?.tradingAccountID ?? null
            })
                .then(() => data)
                .then(() => {
                    dispatch(refreshTask({refresh: false}))
                })
        }

    }, [refresher]);

    if (loading) {
        return <Loading/>
    }

    if (error) {
        return <Error message={error.message}/>
    }

    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">
            <div className="flex">
                <h1 className="text-3xl font-bold">
                    {t('tasks.title')}
                </h1>
                <Button className="ml-2" variant="link"
                        onClick={() => {
                            window.open(ManualKRTask, '_blank')
                        }}
                >
                    {t("manual.btn")}
                </Button>
            </div>

            <div className={"mt-6 w-full flex space-x-2"}>
                <div>
                    {/*@ts-ignore*/}
                    <TradingAccountSelector taskIndex={data?.taskIndex}/>
                </div>

                <div className={"flex-grow"}></div>

                <div>
                    {
                        data?.taskIndex?.selectedTradingAccount
                        // @ts-ignore
                        && <AddTask tradingAccount={data?.taskIndex?.selectedTradingAccount}/>
                    }
                </div>
            </div>

            <div className="mt-6">
                {/*@ts-ignore*/}
                <TaskTable tradingAccount={data?.taskIndex?.selectedTradingAccount}/>

                {/*@ts-ignore*/}
                <TaskCards tradingAccount={data?.taskIndex?.selectedTradingAccount}/>
            </div>
        </main>
    )
}

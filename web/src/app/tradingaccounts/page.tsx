"use client";

import {TradingAccountIndexDocument} from "@/graph/generated/generated";
import {useQuery} from "@apollo/client";
import React, {useEffect} from "react";
import {AddTradingAccount} from "@/components/tradingaccounts/add";
import {useAppDispatch, useAppSelector} from "@/store";
import {refreshTradingAccount,} from "@/store/tradingAccountSlice";
import {TradingAccountTable} from "@/components/tradingaccounts/table";
import {Loading} from "@/components/loading";
import {Error} from "@/components/error";
import {Button} from "@/components/ui/button";
import {ManualKRTradingAccount} from "@/lib/ref-url";
import {TradingAccountCards} from "@/components/tradingaccounts/cards";
import {useTranslation} from "@/lib/i18n";


export default function TradingAccounts() {
    const {t} = useTranslation()
    const dispatch = useAppDispatch()
    const tradingAccount = useAppSelector((state) => state.tradingAccount);
    const {data, loading, refetch, error} = useQuery(TradingAccountIndexDocument, {
        fetchPolicy: "no-cache"
    });

    useEffect(() => {
        if (tradingAccount?.refresh) {
            refetch()
                .then(() => data)
                .then(() => {
                    dispatch(refreshTradingAccount(false))
                })
        }
    }, [tradingAccount])

    if (loading) {
        return <Loading/>
    }

    if (error) {
        return <Error message={error.message}/>
    }


    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">
            <div className="flex">
                <h1 className="text-3xl font-bold">{t("trading_account.title")}</h1>
                <Button className="ml-2" variant="link"
                        onClick={() => {
                            window.open(ManualKRTradingAccount, '_blank')
                        }}
                >
                    {t("manual.btn")}
                </Button>
            </div>

            <div className={"w-full flex space-x-2"}>
                <div className={"flex-grow"}></div>
                {/*<Button onClick={*/}
                {/*    () => {*/}
                {/*        dispatch(setTradingAccountTutorial(false))*/}
                {/*    }*/}
                {/*}>*/}
                {/*    Reset Tutorial*/}
                {/*</Button>*/}
                <AddTradingAccount/>
            </div>

            <div className="mt-6">

                {/*@ts-ignore*/}
                <TradingAccountTable tradingAccounts={data?.tradingAccountIndex?.tradingAccounts}/>

                {/*@ts-ignore*/}
                <TradingAccountCards tradingAccounts={data?.tradingAccountIndex?.tradingAccounts}/>

                {/*<TradingAccountTutorial/>*/}
            </div>
        </main>
    )
}



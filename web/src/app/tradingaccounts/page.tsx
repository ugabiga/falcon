"use client";

import {TradingAccountIndexDocument} from "@/graph/generated/generated";
import {useQuery} from "@apollo/client";
import {useEffect} from "react";
import {AddTradingAccount} from "@/app/tradingaccounts/add";
import {useAppDispatch, useAppSelector} from "@/store";
import {refreshTradingAccount,} from "@/store/tradingAccountSlice";
import {TradingAccountTable} from "@/app/tradingaccounts/table";
import {Loading} from "@/components/loading";
import {Error} from "@/components/error";
import {useTranslation} from "react-i18next";


export default function TradingAccounts() {
    const {t} = useTranslation();
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
            <h1 className="text-3xl font-bold">{t("trading_account.title")}</h1>

            <div className={"w-full flex space-x-2"}>
                <div className={"flex-grow"}></div>
                <AddTradingAccount/>
            </div>

            <div className="mt-6">
                {/*@ts-ignore*/}
                <TradingAccountTable tradingAccounts={data?.tradingAccountIndex?.tradingAccounts}/>
            </div>
        </main>
    )
}

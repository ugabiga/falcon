"use client";

import {TaskTable} from "@/app/tasks/table";
import {useQuery} from "@apollo/client";
import {useEffect} from "react";
import {GetTaskIndexDocument} from "@/graph/generated/generated";
import {TradingAccountSelector} from "@/app/tasks/selector";
import {useAppSelector} from "@/store";
import {useDispatch} from "react-redux";
import {refreshTask} from "@/store/taskSlice";
import {AddTask} from "@/app/tasks/add";
import {Loading} from "@/components/loading";
import {useSearchParams} from "next/navigation";
import {Error} from "@/components/error";

export default function Tasks(){
    const params = useSearchParams()
    const tradingAccountId = params.get('trading_account_id')

    const {data, loading, refetch, error} = useQuery(GetTaskIndexDocument,{
        variables: {
            tradingAccountID: tradingAccountId ?? null
        }
    })

    const task = useAppSelector((state) => state.task)
    const dispatch = useDispatch()

    useEffect(() => {
        if (task?.refresh) {
            refetch({
                tradingAccountID: task.tradingAccountID
            })
                .then(() => data)
                .then(() => {
                    dispatch(refreshTask({
                        refresh: false
                    }))
                })
        }

    }, [task]);

    if (loading) {
        return <Loading/>
    }

    if (error) {
        return <Error message={error.message}/>
    }

    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">
            <h1 className="text-3xl font-bold">Tasks</h1>

            <div className={"mt-6 w-full flex space-x-2"}>
                <div>
                    {/*@ts-ignore*/}
                    <TradingAccountSelector taskIndex={data?.taskIndex}/>
                </div>

                <div className={"flex-grow"}></div>

                <div>
                    <AddTask tradingAccountID={data?.taskIndex?.selectedTradingAccount?.id}/>
                </div>
            </div>

            <div className="mt-6">
                {/*@ts-ignore*/}
                <TaskTable tasks={data?.taskIndex?.selectedTradingAccount?.tasks}/>
            </div>
        </main>
    )
}

"use client";

import {TaskTable} from "@/app/tasks/table";
import {useQuery} from "@apollo/client";
import {useEffect, useState} from "react";
import {GetTaskIndexDocument} from "@/graph/generated/generated";
import {TradingAccountSelector} from "@/app/tasks/selector";
import {useAppSelector} from "@/store";
import {useDispatch} from "react-redux";
import {refreshTask} from "@/store/taskSlice";
import {AddTask} from "@/app/tasks/add";

export default function Tasks() {
    const {data, loading, refetch} = useQuery(GetTaskIndexDocument)
    const [selectedTradingAccountId, setSelectedTradingAccountId] = useState<string | null>(null)
    const task = useAppSelector((state) => state.task)
    const dispatch = useDispatch()

    useEffect(() => {
        if (task?.refresh) {
            refetch({
                tradingAccountID: task.tradingAccountID
            })
                .then(r => data)
                .then(r => {
                    dispatch(refreshTask({
                        refresh: false
                    }))
                })
        }

    }, [task]);

    if (loading) {
        return <div>Loading...</div>
    }

    if (!data) {
        return <div>No Data</div>
    }

    if (!data.taskIndex?.selectedTradingAccount) {
        return <div>No Trading Account Selected</div>
    }

    return (
        <main className="min-h-screen p-12">
            <h1 className="text-3xl font-bold">Tasks</h1>

            <div className={"mt-6 w-full flex space-x-2"}>
                <div>
                    {/*@ts-ignore*/}
                    <TradingAccountSelector taskIndex={data.taskIndex}/>
                </div>

                <div className={"flex-grow"}></div>

                <div>
                    <AddTask tradingAccountID={data.taskIndex.selectedTradingAccount.id}/>
                </div>
            </div>

            <div className="mt-6">
                {/*@ts-ignore*/}
                <TaskTable tasks={data.taskIndex.selectedTradingAccount?.tasks}/>
            </div>
        </main>
    )
}

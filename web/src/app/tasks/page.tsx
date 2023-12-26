"use client";

import {TaskTable} from "@/app/tasks/table";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useQuery} from "@apollo/client";
import {useState} from "react";
import {GetTaskIndexDocument} from "@/graph/generated/generated";

export default function Tasks() {
    const {data, loading, refetch} = useQuery(GetTaskIndexDocument)
    const [selectedTradingAccountId, setSelectedTradingAccountId] = useState<string | null>(null)

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

            <div className={"w-full flex space-x-2"}>
                <div className={"flex-grow"}></div>
                <div>
                    <Select defaultValue={data.taskIndex.selectedTradingAccount?.id}
                            onValueChange={(value) => {
                                setSelectedTradingAccountId(value)
                                refetch({
                                    tradingAccountID: value
                                }).then(() => data)
                            }}
                    >
                        <SelectTrigger>
                            <SelectValue placeholder="Select a Exchange"/>
                        </SelectTrigger>
                        <SelectContent>
                            {
                                data.taskIndex.tradingAccounts?.map((tradingAccount) => {
                                    return (
                                        <SelectItem key={tradingAccount.id} value={tradingAccount.id}>
                                            {tradingAccount.identifier}
                                        </SelectItem>
                                    )
                                })
                            }
                        </SelectContent>
                    </Select>
                </div>
            </div>

            <div className="mt-6">
                {/*@ts-ignore*/}
                <TaskTable tasks={data.taskIndex.selectedTradingAccount?.tasks}/>
            </div>
        </main>
    )
}

"use client";

import {TaskTable} from "@/app/tasks/table";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useQuery} from "@apollo/client";
import {GetTradingAccountWithTasksDocument} from "@/graph/generated/generated";
import {useState} from "react";

export default function Tasks() {
    const {data, loading, refetch} = useQuery(GetTradingAccountWithTasksDocument)
    const [selectedTradingAccountId, setSelectedTradingAccountId] = useState<string | null>(null)

    if (loading) {
        return <div>Loading...</div>
    }

    if (!data) {
        return <div>No Data</div>
    }

    if (!data.tradingAccountsWithTasks.selectedTradingAccount) {
        return <div>No Trading Account Selected</div>
    }

    return (
        <main className="min-h-screen p-12">
            <h1 className="text-3xl font-bold">Tasks</h1>

            <div className={"w-full flex space-x-2"}>
                <div className={"flex-grow"}></div>
                <div>
                    <Select defaultValue={data.tradingAccountsWithTasks.selectedTradingAccount?.id}
                            onValueChange={(value) => {
                                setSelectedTradingAccountId(value)
                                refetch({
                                    id: value
                                }).then(() => data)
                            }}
                    >
                        <SelectTrigger>
                            <SelectValue placeholder="Select a Exchange"/>
                        </SelectTrigger>
                        <SelectContent>
                            {
                                data.tradingAccountsWithTasks.tradingAccounts?.map((tradingAccount) => {
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
                <TaskTable tasks={data.tradingAccountsWithTasks.selectedTradingAccount?.tasks}/>
            </div>
        </main>
    )
}

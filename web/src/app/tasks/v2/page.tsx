"use client";

import {useTranslation} from "@/lib/i18n";
import {Button} from "@/components/ui/button";
import {ManualKRTask} from "@/lib/ref-url";
import React, {useEffect} from "react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {useGetApiV1Tasks} from "@/api/endpoints/transformer";
import {V1TaskIndexResponse} from "@/api/model";
import {convertNumberToCryptoSize} from "@/lib/number";
import {convertDayOfWeek, convertHours, convertToNextExecutionTime} from "@/lib/cron-parser";
import {TaskMoreBtn} from "@/components/tasks/more-btn";
import TaskTable from "@/components/tasks/v2/task-table";
import TradingAccountSelector from "@/components/tasks/v2/task-tradingaccount-selector";
import TaskCreate from "@/components/tasks/v2/task-create";
import {RefreshTarget} from "@/store/refresherSlice";
import {refreshTask} from "@/store/taskSlice";
import {useAppSelector} from "@/store";
import {useReceiveRefreshSignal} from "@/lib/use-refresh";
import TaskCards from "@/components/tasks/v2/task-cards";

export default function Tasks() {
    const {t} = useTranslation()
    const {data, refetch} = useGetApiV1Tasks()

    useReceiveRefreshSignal({
        target: RefreshTarget.Task,
        afterReceiveSignal: refetch
    })

    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">
            {/* Header */}
            <div className="flex">
                <h1 className="text-3xl font-bold">
                    {t('tasks.title')}
                </h1>
                <Button
                    className="ml-2"
                    variant="link"
                    onClick={() => {
                        window.open(ManualKRTask, '_blank')
                    }}
                >
                    {t("manual.btn")}
                </Button>
            </div>

            {/* Selector */}
            <div className={"mt-6 w-full flex space-x-2"}>

                <div>
                    <TradingAccountSelector data={data}/>
                </div>

                <div className={"flex-grow"}></div>

                <div>
                    <TaskCreate tradingAccount={data?.selected_trading_account}/>
                </div>
            </div>

            {/* Table */}
            <div className="hidden md:block">
                <TaskTable data={data}/>
            </div>
            <div className="block md:hidden mt-6">
                <TaskCards data={data}/>
            </div>

        </main>
    )
}


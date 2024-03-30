import {V1TaskIndexResponse} from "@/api/model";
import {useTranslation} from "@/lib/i18n";
import {Card, CardContent, CardHeader} from "@/components/ui/card";
import {Label} from "@/components/ui/label";
import {convertNumberToCryptoSize} from "@/lib/number";
import {convertDayOfWeek, convertHours, convertToNextExecutionTime} from "@/lib/cron-parser";
import Spacer from "@/components/spacer";
import {TableCell} from "@/components/ui/table";
import TaskDetail from "@/components/tasks/v2/task-detail";
import {Button} from "@/components/ui/button";
import Link from "next/link";
import React from "react";

export default function TaskCards(
    {
        data
    }: {
        data?: V1TaskIndexResponse
    }
) {
    const {t} = useTranslation();

    function convertSchedule(cronExpression: string): string {
        const hours = convertHours(cronExpression)
        const dayOfWeek = convertDayOfWeek(cronExpression)
        let translatedDayOfWeek = ""

        switch (dayOfWeek) {
            case "everyday":
                translatedDayOfWeek = t("tasks.table.schedule.everyday")
                break
            case "every_weekday":
                translatedDayOfWeek = t("tasks.table.schedule.every_weekday")
                break
            default:
                translatedDayOfWeek =
                    t("tasks.table.schedule.every_week")
                    + " "
                    + dayOfWeek.split(',')
                        .map(day => {
                            return t("common.days." + day)
                        })
                        .join(', ')
        }

        return t("tasks.table.schedule.encoded", {
            hours: hours,
            days: translatedDayOfWeek
        })
    }


    return (
        <div className="space-y-2">
            {
                data?.selected_tasks?.map((task) => (
                    <Card key={task.id}>
                        <CardContent className="grid gap-5 pt-4">
                            <div className="col-span-2">
                                <div className="flex items-center">
                                    <h3 className="text-xl font-semibold tracking-tight">
                                        {task.symbol}
                                    </h3>
                                    <Spacer/>
                                    <Label className={task.is_active ? "text-green-500" : "text-red-500"}>
                                        {t("task.table.is_active.boolean." + task.is_active)}
                                    </Label>
                                </div>
                            </div>

                            <Label className="font-semibold">
                                {t("tasks.table.type")} : {t("tasks.type." + task.type)}
                            </Label>

                            <Label className="text-right font-light">
                                {t("tasks.table.size")} : {task.size + " " + task.symbol}
                            </Label>

                            <Label className="col-span-2 font-light">
                                {t("tasks.table.schedule")} : {convertSchedule(task.cron)}
                            </Label>
                            <Label className="col-span-2 font-light">
                                {t("tasks.table.next_execution_time")} : {convertToNextExecutionTime(task.cron, t("tasks.table.next_execution_time.fail"))}
                            </Label>

                            <div className="col-span-2 flex">
                                <Button variant="outline" asChild>
                                    <Link
                                        href={`/tasks/${task.id}`
                                            + `/history?trading_account_id=${task.trading_account_id}`}
                                    >
                                        {t("tasks.table.history")}
                                    </Link>
                                </Button>

                                <Spacer/>

                                <TaskDetail
                                    variant="secondary"
                                    task={task}
                                    tradingAccount={data.selected_trading_account}
                                />
                            </div>

                        </CardContent>
                    </Card>
                ))
            }
        </div>
    )
}
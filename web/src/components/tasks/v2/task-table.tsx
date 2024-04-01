import {V1TaskIndexResponse} from "@/api/model";
import {useTranslation} from "@/lib/i18n";
import {convertDayOfWeek, convertHours, convertToNextExecutionTime} from "@/lib/cron-parser";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import React from "react";
import TaskDetail from "@/components/tasks/v2/task-detail";
import {TaskMoreBtn} from "@/components/tasks/old/more-btn";
import Link from "next/link";
import {Button} from "@/components/ui/button";

export default function TaskTable(
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
        <div className="mt-6 rounded-md border">
            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead>{t("tasks.table.symbol")}</TableHead>
                        <TableHead>{t("tasks.table.type")}</TableHead>
                        <TableHead>{t("tasks.table.size")}</TableHead>
                        <TableHead>{t("tasks.table.schedule")}</TableHead>
                        <TableHead>{t("tasks.table.next_execution_time")}</TableHead>
                        <TableHead>{t("tasks.table.is_active")}</TableHead>
                        <TableHead
                            className="text-center"
                        >
                            {t("tasks.table.action")}
                        </TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {
                        data?.selected_tasks?.map((task) => (
                            <TableRow key={task.id}>
                                <TableCell>{task.symbol}</TableCell>
                                <TableCell>
                                    {t("tasks.type." + task.type)}
                                </TableCell>
                                <TableCell>
                                    {task.size + " " + task.symbol}
                                </TableCell>
                                <TableCell>
                                    {convertSchedule(task.cron)}
                                </TableCell>
                                <TableCell>
                                    {convertToNextExecutionTime(
                                        task.cron,
                                        t("tasks.table.next_execution_time.fail")
                                    )}
                                </TableCell>
                                <TableCell className={task.is_active ? "text-green-500" : "text-red-500"}>
                                    {t("task.table.is_active.boolean." + task.is_active)}
                                </TableCell>
                                <TableCell className="text-center">
                                    <TaskDetail
                                        variant="link"
                                        task={task}
                                        tradingAccount={data.selected_trading_account}
                                    />
                                    <Button variant="link" asChild>
                                        <Link
                                            href={`/tasks/${task.id}`
                                                + `/history?trading_account_id=${task.trading_account_id}`}
                                        >
                                            {t("tasks.table.history")}
                                        </Link>
                                    </Button>
                                </TableCell>
                            </TableRow>
                        ))
                        || (data?.trading_accounts === null) && (
                            <TableRow>
                                <TableCell colSpan={7} className="font-medium text-center">
                                    {t("tasks.trading_accounts.empty")}
                                </TableCell>
                            </TableRow>
                        )
                        || (
                            <TableRow>
                                <TableCell colSpan={7} className="font-medium text-center">
                                    {t("tasks.table.empty")}
                                </TableCell>
                            </TableRow>
                        )
                    }
                </TableBody>
            </Table>
        </div>
    )
}


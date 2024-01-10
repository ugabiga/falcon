import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {Task} from "@/graph/generated/generated";
import {convertDayOfWeek, convertHours, convertToNextExecutionTime} from "@/lib/cron-parser";
import {useTranslation} from "react-i18next";
import {convertNumberToCryptoSize} from "@/lib/number";
import {TaskMoreBtn} from "@/app/tasks/more-btn";

export function TaskTable({tasks}: { tasks?: Task[] }) {
    const {t} = useTranslation();

    function convertSchedule(cronExpression: string): string {
        const hours = convertHours(cronExpression)
        const dayOfWeek = convertDayOfWeek(cronExpression)
        let translatedDayOfWeek = ""
        console.log("dayOfWeek", dayOfWeek)

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
        <div className="hidden md:block">
            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead>{t("tasks.table.id")}</TableHead>
                        <TableHead>{t("tasks.table.type")}</TableHead>
                        <TableHead>{t("tasks.table.schedule")}</TableHead>
                        <TableHead>{t("tasks.table.symbol")}</TableHead>
                        <TableHead>{t("tasks.table.size")}</TableHead>
                        <TableHead>{t("tasks.table.next_execution_time")}</TableHead>
                        <TableHead>{t("tasks.table.is_active")}</TableHead>
                        <TableHead>{t("tasks.table.action")}</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {
                        !tasks || tasks?.length === 0
                            ? (
                                <TableRow>
                                    <TableCell colSpan={9} className="font-medium text-center">
                                        {t("tasks.table.empty")}
                                    </TableCell>
                                </TableRow>
                            )
                            : tasks?.map((task) => (
                                    <TableRow key={task.id}>
                                        <TableCell>{task.id}</TableCell>
                                        <TableCell>{task.type}</TableCell>
                                        <TableCell>
                                            {convertSchedule(task.cron)}
                                        </TableCell>
                                        <TableCell>{task.symbol}</TableCell>
                                        <TableCell>
                                            {convertNumberToCryptoSize(task.size, task.symbol)}
                                        </TableCell>
                                        <TableCell>
                                            {convertToNextExecutionTime(task.cron, t("tasks.table.next_execution_time.fail"))}
                                        </TableCell>
                                        <TableCell>
                                            {t("task.table.is_active.boolean." + task.isActive)}
                                        </TableCell>
                                        <TableCell>
                                            <TaskMoreBtn task={task}/>
                                        </TableCell>
                                    </TableRow>
                                )
                            )
                    }
                </TableBody>
            </Table>
        </div>
    )
}


function convertNumberToCurrency(value: number, currency: string): string {
    let decimalPlaces = 0
    switch (currency) {
        case 'KRW':
            decimalPlaces = 0
            break
        case 'BTC':
            decimalPlaces = 8
            break
        case 'ETH':
            decimalPlaces = 8
            break
        case 'USDT':
            decimalPlaces = 2
            break
        default:
            decimalPlaces = 2
    }

    try {
        return new Intl.NumberFormat('en-US', {
            style: 'currency',
            currency: currency,
            minimumFractionDigits: decimalPlaces,
        }).format(value)
    } catch (e) {
        return new Intl.NumberFormat('en-US', {
            minimumFractionDigits: decimalPlaces,
        }).format(value) + ' ' + currency
    }
}

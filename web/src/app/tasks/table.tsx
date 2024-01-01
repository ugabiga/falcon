import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {Task} from "@/graph/generated/generated";
import {nextCronDate, parseCronExpression} from "@/lib/cron-parser";
import {EditTask} from "@/app/tasks/edit";
import {Button} from "@/components/ui/button";
import Link from "next/link";
import {convertBooleanToYesNo} from "@/lib/converter";
import {useTranslation} from "react-i18next";


function convertDayOfWeek(value: string): string {
    const result = parseCronExpression(value)
    const daysRaw = result.fields.dayOfWeek.toString()
    const daysOfWeek = ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday'];
    const days = daysRaw.split(',').map(Number);
    if (days.includes(0) && days.includes(7)) {
        days.splice(days.indexOf(0), 1);
        days.splice(days.indexOf(7), 1);
        days.push(0);
    }
    if (days.length === 7) {
        return 'everyday';
    } else if (days.length === 5 && !days.some(day => day === 0 || day === 6)) {
        return 'every_weekday';
    } else {
        return 'every_' + days.map(day => daysOfWeek[day]).join(', ');
    }
}

function convertHours(value: string): string {
    const result = parseCronExpression(value)
    return result.fields.hour.toString()
}

function convertCronToHumanReadable(value: string): { days: string, hours: string } {
    const result = parseCronExpression(value)
    const days = result.fields.dayOfWeek.toString()
    const hours = result.fields.hour.toString()
    return {
        days: convertDayOfWeek(days),
        hours: hours
    }
}

function formatDate(dateTime: Date): string {
    const options: Intl.DateTimeFormatOptions = {
        year: 'numeric',
        month: 'numeric',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric',
        hour12: false, // Use 24-hour format
    };

    const formatter = new Intl.DateTimeFormat('ko-KR', options);
    const formattedString = formatter.format(dateTime);

    return formattedString.replace(/(\d+)\/(\d+)\/(\d+), (\d+):(\d+)/, '$3 $4 $5');
}

function convertToNextExecutionTime(value: string, failMessage?: string) {
    const result = nextCronDate(value)
    if (result === null) {
        return failMessage || 'No next execution time.';
    }
    return formatDate(result)
}

function convertNumberToCryptoSize(value: number, symbol: string): string {
    const decimalPlaces = 5
    return new Intl.NumberFormat('en-US', {
        minimumFractionDigits: decimalPlaces,
    }).format(value) + ' ' + symbol
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

export function TaskTable({tasks}: { tasks?: Task[] }) {
    const {t} = useTranslation();
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">{t("tasks.table.id")}</TableHead>
                    <TableHead>{t("tasks.table.type")}</TableHead>
                    <TableHead>{t("tasks.table.schedule")}</TableHead>
                    <TableHead>{t("tasks.table.symbol")}</TableHead>
                    <TableHead>{t("tasks.table.size")}</TableHead>
                    <TableHead>{t("tasks.table.next_execution_time")}</TableHead>
                    <TableHead>{t("tasks.table.is_active")}</TableHead>
                    <TableHead>{t("tasks.table.action")}</TableHead>
                    <TableHead>{t("tasks.table.more")}</TableHead>
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
                                    {t("tasks.table.next_execution_time.encoded", {
                                        days: t("tasks.table.next_execution_time." + convertDayOfWeek(task.cron)),
                                        hours: convertHours(task.cron)
                                    })}
                                </TableCell>
                                <TableCell>{task.symbol}</TableCell>
                                <TableCell>{convertNumberToCryptoSize(task.size, task.symbol)}</TableCell>
                                <TableCell>{convertToNextExecutionTime(task.cron, t("tasks.table.next_execution_time.fail"))}</TableCell>
                                <TableCell>{t("task.table.is_active.boolean." + task.isActive)}</TableCell>
                                <TableCell>
                                    <EditTask task={task}/>
                                </TableCell>
                                <TableCell>
                                    <Button variant="link" asChild>
                                        <Link href={`/tasks/${task.id}/history`} legacyBehavior>
                                            <a>{t("tasks.table.history")}</a>
                                        </Link>
                                    </Button>
                                </TableCell>
                            </TableRow>
                        ))

                }
            </TableBody>
        </Table>
    )
}

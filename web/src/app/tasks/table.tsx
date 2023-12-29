import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {Task} from "@/graph/generated/generated";
import {nextCronDate, parseCronExpression} from "@/lib/cron-parser";
import {EditTask} from "@/app/tasks/edit";
import {Button} from "@/components/ui/button";
import Link from "next/link";
import {convertBooleanToYesNo} from "@/lib/converter";


function convertDayOfWeek(value: string): string {
    const daysOfWeek = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
    const days = value.split(',').map(Number);
    if (days.includes(0) && days.includes(7)) {
        days.splice(days.indexOf(0), 1);
        days.splice(days.indexOf(7), 1);
        days.push(0);
    }
    if (days.length === 7) {
        return 'Everyday';
    } else if (days.length === 5 && !days.some(day => day === 0 || day === 6)) {
        return 'Every weekday';
    } else {
        return 'Every ' + days.map(day => daysOfWeek[day]).join(', ');
    }
}

function convertCronToHumanReadable(value: string) {
    const result = parseCronExpression(value)
    const days = result.fields.dayOfWeek.toString()
    const hours = result.fields.hour.toString()
    return `${convertDayOfWeek(days)} at ${hours} o'clock.`
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

function convertToNextCronDate(value: string) {
    const result = nextCronDate(value)
    if (result === null) {
        return 'No next execution time.';
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
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">ID</TableHead>
                    <TableHead>Type</TableHead>
                    <TableHead>Schedule</TableHead>
                    <TableHead>Crypto Symbol</TableHead>
                    <TableHead>Investing Size</TableHead>
                    <TableHead>Next Execution Time(24h)</TableHead>
                    <TableHead>Is Active</TableHead>
                    <TableHead>Action</TableHead>
                    <TableHead>More</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    !tasks || tasks?.length === 0
                        ? (
                            <TableRow>
                                <TableCell colSpan={9} className="font-medium text-center">No tasks found.</TableCell>
                            </TableRow>
                        )
                        : tasks?.map((task) => (
                            <TableRow key={task.id}>
                                <TableCell>{task.id}</TableCell>
                                <TableCell>{task.type}</TableCell>
                                <TableCell>{convertCronToHumanReadable(task.cron)}</TableCell>
                                <TableCell>{task.symbol}</TableCell>
                                <TableCell>{convertNumberToCryptoSize(task.size, task.symbol)}</TableCell>
                                <TableCell>{convertToNextCronDate(task.cron)}</TableCell>
                                <TableCell>{convertBooleanToYesNo(task.isActive)}</TableCell>
                                <TableCell>
                                    <EditTask task={task}/>
                                </TableCell>
                                <TableCell>
                                    <Button variant="link" asChild>
                                        <Link href={`/tasks/${task.id}/history`} legacyBehavior>
                                            <a>History</a>
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

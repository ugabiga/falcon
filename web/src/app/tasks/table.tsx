import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {Task} from "@/graph/generated/generated";
import {nextCronDate, parseCronExpression} from "@/lib/cron-parser";
import {EditTask} from "@/app/tasks/edit";

function convertBooleanToYesNo(value: boolean) {
    return value ? "Yes" : "No"
}

function convertToPastHourString(input: string): string {
    const hours = input.split(',').map(Number);

    if (hours.length === 0) {
        return 'No execution times provided.';
    }

    const hoursString = hours.join(', ');
    const lastHour = hours.pop();

    if (hours.length === 0) {
        return `At ${lastHour} o'clock.`;
    }

    return `At ${hoursString}, and ${lastHour} o'clock.`;
}

function convertCronToHumanReadable(value: string) {
    const result = parseCronExpression(value)
    const hours = result.fields.hour.toString()
    return convertToPastHourString(hours)
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

export function TaskTable({tasks}: { tasks?: Task[] }) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">ID</TableHead>
                    <TableHead>Schedule(24h)</TableHead>
                    <TableHead>Type</TableHead>
                    <TableHead>Next Execution Time(24h)</TableHead>
                    <TableHead>Is Active</TableHead>
                    <TableHead>Action</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    tasks?.length === 0 ? (
                            <TableRow>
                                <TableCell colSpan={6} className="font-medium text-center">No tasks found.</TableCell>
                            </TableRow>
                        )
                        : tasks?.map((task) => (
                            <TableRow key={task.id}>
                                <TableCell>{task.id}</TableCell>
                                <TableCell>{convertCronToHumanReadable(task.cron)}</TableCell>
                                <TableCell>{task.type}</TableCell>
                                <TableCell>{convertToNextCronDate(task.cron)}</TableCell>
                                <TableCell>{convertBooleanToYesNo(task.isActive)}</TableCell>
                                <TableCell>
                                    <EditTask task={task}/>
                                </TableCell>
                            </TableRow>
                        ))

                }
            </TableBody>
        </Table>
    )
}

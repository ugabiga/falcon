import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {Task} from "@/graph/generated/generated";

export function TaskTable({tasks}: { tasks?: Task[] }) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">ID</TableHead>
                    <TableHead>Schedule</TableHead>
                    <TableHead>Type</TableHead>
                    <TableHead>Next Execution Time(24h)</TableHead>
                    <TableHead>Is Active</TableHead>
                    <TableHead>Action</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    tasks?.map((task) => (
                        <TableRow key={task.id}>
                            <TableCell>{task.id}</TableCell>
                            <TableCell>{task.cron}</TableCell>
                            <TableCell>{task.type}</TableCell>
                            <TableCell>{task.nextExecutionTime}</TableCell>
                            <TableCell>{task.isActive}</TableCell>
                            <TableCell>Edit</TableCell>
                        </TableRow>
                    ))
                }
                {/*<TableRow>*/}
                {/*    <TableCell>1</TableCell>*/}
                {/*    <TableCell>Every 1 minute</TableCell>*/}
                {/*    <TableCell>DCA</TableCell>*/}
                {/*    <TableCell>2021-10-01 00:00</TableCell>*/}
                {/*    <TableCell>Yes</TableCell>*/}
                {/*    <TableCell>Edit</TableCell>*/}
                {/*</TableRow>*/}
            </TableBody>
        </Table>
    )
}

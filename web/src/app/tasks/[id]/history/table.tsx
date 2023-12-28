import {TaskHistory} from "@/graph/generated/generated";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {convertBooleanToYesNo} from "@/lib/converter";

export function TaskHistoryTable({taskHistories}: { taskHistories?: TaskHistory[] }) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">ID</TableHead>
                    <TableHead>IsSuccess</TableHead>
                    <TableHead>UpdatedAt</TableHead>
                    <TableHead>CreatedAt</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    !taskHistories || taskHistories?.length === 0
                        ? (
                            <TableRow>
                                <TableCell colSpan={6} className="font-medium text-center">
                                    No task histories found.
                                </TableCell>
                            </TableRow>
                        )
                        : taskHistories?.map((taskHistory) => (
                            <TableRow key={taskHistory.id}>
                                <TableCell>{taskHistory.id}</TableCell>
                                <TableCell>{convertBooleanToYesNo(taskHistory.isSuccess)}</TableCell>
                                <TableCell>{taskHistory.updatedAt}</TableCell>
                                <TableCell>{taskHistory.createdAt}</TableCell>
                            </TableRow>
                        ))

                }
            </TableBody>

        </Table>
    )

}
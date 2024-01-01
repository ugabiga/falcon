import {TaskHistory} from "@/graph/generated/generated";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {convertBooleanToYesNo} from "@/lib/converter";
import {useTranslation} from "react-i18next";

export function TaskHistoryTable({taskHistories}: { taskHistories?: TaskHistory[] }) {
    const {t} = useTranslation()
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">{t("task_history.table.id")}</TableHead>
                    <TableHead>{t("task_history.table.is_success")}</TableHead>
                    <TableHead>{t("task_history.table.updated_at")}</TableHead>
                    <TableHead>{t("task_history.table.created_at")}</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    !taskHistories || taskHistories?.length === 0
                        ? (
                            <TableRow>
                                <TableCell colSpan={6} className="font-medium text-center">
                                    {t("task_history.table.empty")}
                                </TableCell>
                            </TableRow>
                        )
                        : taskHistories?.map((taskHistory) => (
                            <TableRow key={taskHistory.id}>
                                <TableCell>{taskHistory.id}</TableCell>
                                <TableCell>{t("task_history.table.is_success.boolean." + taskHistory.isSuccess)}</TableCell>
                                <TableCell>{taskHistory.updatedAt}</TableCell>
                                <TableCell>{taskHistory.createdAt}</TableCell>
                            </TableRow>
                        ))

                }
            </TableBody>

        </Table>
    )

}
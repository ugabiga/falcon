import {TaskHistory} from "@/graph/generated/generated";
import {useTranslation} from "react-i18next";
import {Card, CardContent} from "@/components/ui/card";
import {Label} from "@/components/ui/label";

export function TaskHistoryCards({taskHistories}: { taskHistories?: TaskHistory[] }) {

    const {t} = useTranslation();
    return (
        <div className="block md:hidden">
            {
                !taskHistories || taskHistories?.length === 0
                    ? <div className="font-medium text-center">
                        {t("task_history.table.empty")}
                    </div>
                    : taskHistories?.map((taskHistory) => (
                            <Card key={taskHistory.id}>
                                <div className="grid grid-cols-4 gap-6">
                                    <div className="mt-4 mb-6 ml-6 col-span-3">
                                        <h3 className="scroll-m-20 text-xl font-semibold tracking-tight">
                                            {taskHistory.id}
                                        </h3>
                                    </div>
                                </div>
                                <CardContent className="grid grid-cols-2 gap-6">
                                    <Label>
                                        {t("task_history.table.is_success")} : {t("task_history.table.is_success.boolean." + taskHistory.isSuccess)}
                                    </Label>
                                    <Label className="col-span-2">
                                        {t("task_history.table.log")} : {taskHistory.log}
                                    </Label>
                                    <Label className="col-span-2">
                                        {t("task_history.table.created_at")} : {taskHistory.createdAt}
                                    </Label>
                                </CardContent>
                            </Card>
                        )
                    )
            }
        </div>
    )
}

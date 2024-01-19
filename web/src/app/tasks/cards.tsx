import {Task, TradingAccount} from "@/graph/generated/generated";
import {useTranslation} from "react-i18next";
import {Card, CardContent} from "@/components/ui/card";
import {Label} from "@/components/ui/label";
import {convertNumberToCryptoSize} from "@/lib/number";
import {convertDayOfWeek, convertHours, convertToNextExecutionTime} from "@/lib/cron-parser";
import {TaskMoreBtn} from "@/app/tasks/more-btn";


export function TaskCards(
    {
        tradingAccount
    }: {
        tradingAccount?: TradingAccount
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
        <div className="block md:hidden space-y-2">
            {
                !tradingAccount?.tasks || tradingAccount?.tasks?.length === 0
                    ? <div className="font-medium text-center">
                        {t("tasks.table.empty")}
                    </div>
                    : tradingAccount?.tasks?.map((task) => (
                            <Card key={task.id}>
                                <div className="grid grid-cols-5 gap-6">
                                    <div className="mt-4 mb-6 ml-6 col-span-4">
                                        <h3 className="scroll-m-20 text-xl font-semibold tracking-tight">
                                            {task.id}
                                        </h3>
                                    </div>
                                    <div className="flex space-x-2 mt-2 mr-2">
                                        <div className="flex-grow"></div>
                                        <TaskMoreBtn task={task} tradingAccount={tradingAccount}/>
                                    </div>
                                </div>
                                <CardContent className="grid grid-cols-2 gap-6">
                                    <Label className="col-span-2">
                                        {t("task.table.is_active.boolean." + task.isActive)}
                                    </Label>
                                    <Label>
                                        {t("tasks.table.type")} : {t("tasks.type." + task.type)}
                                    </Label>
                                    <Label>
                                        {t("tasks.table.symbol")} : {task.symbol}
                                    </Label>
                                    <Label>
                                        {t("tasks.table.size")} : {convertNumberToCryptoSize(task.size, task.symbol)}
                                    </Label>
                                    <Label className="col-span-2">
                                        {t("tasks.table.schedule")} : {convertSchedule(task.cron)}
                                    </Label>
                                    <div className="flex flex-col space-y-2 col-span-2">
                                        <Label>
                                            {t("tasks.table.next_execution_time")} : {convertToNextExecutionTime(task.cron, t("tasks.table.next_execution_time.fail"))}
                                        </Label>
                                    </div>
                                </CardContent>
                            </Card>
                        )
                    )
            }
        </div>
    )
}

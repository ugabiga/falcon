import {Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {useTranslation} from "react-i18next";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import React, {useState} from "react";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {convertStringToTaskType} from "@/lib/model";
import {TaskFormFields, TaskFromSchema} from "@/components/tasks/v2/task-form";
import {useMutation} from "@apollo/client";
import {UpdateTaskDocument} from "@/graph/generated/generated";
import {useSendRefreshSignal} from "@/lib/use-refresh";
import {ModelTask, ModelTradingAccount} from "@/api/model";
import {parseCronExpression} from "@/lib/cron-parser";
import {RefreshTarget} from "@/store/refresherSlice";
import {errorToast} from "@/components/toast";
import {translatedError} from "@/lib/error";
import TaskDelete from "@/components/tasks/v2/task-delete";
import Spacer from "@/components/spacer";
import {parseParamsFromData, parseParamsFromTask} from "@/lib/task-params";
import {Checkbox} from "@/components/ui/checkbox";

export default function TaskDetail(
    {
        variant,
        task,
        tradingAccount
    }: {
        variant: "secondary" | "link"
        task: ModelTask
        tradingAccount: ModelTradingAccount
    }
) {
    const {t} = useTranslation()
    const [openDialog, setOpenDialog] = useState(false)
    const form = useForm<z.infer<typeof TaskFromSchema>>({
        resolver: zodResolver(TaskFromSchema),
        defaultValues: {
            type: convertStringToTaskType(task.type!),
            currency: task.currency,
            size: task.size?.toString(),
            symbol: task.symbol,
            days: convertCronToDays(task.cron),
            hours: convertCronToHours(task.cron),
            isActive: task.is_active,
            grid: parseParamsFromTask(task)
        },
    })

    const {sendRefresh} = useSendRefreshSignal()
    const [updateTask] = useMutation(UpdateTaskDocument)

    function onSubmit(data: z.infer<typeof TaskFromSchema>) {

        updateTask({
            variables: {
                tradingAccountID: task.trading_account_id!,
                taskID: task.id!,
                currency: data.currency,
                size: Number(data.size),
                symbol: data.symbol,
                days: data.days.join(','),
                hours: data.hours.join(','),
                type: data.type,
                isActive: data.isActive,
                params: parseParamsFromData(data)
            }
        }).then(() => {
            onCompleteAction()
        }).catch((e) => {
            errorToast(translatedError(t, e.message))
        })
    }

    function onCompleteAction() {
        setOpenDialog(false)
        sendRefresh(RefreshTarget.Task)
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant={variant}>
                    {t("tasks.detail.btn")}
                </Button>
            </DialogTrigger>
            <DialogContent className={"overflow-y-scroll h-[calc(100dvh)] sm:max-w-[500px] md:h-auto md:max-h-screen"}>

                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>
                                {t("tasks.form.edit.title")} ({task.id})
                            </DialogTitle>
                        </DialogHeader>

                        <FormField
                            control={form.control}
                            name="isActive"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>
                                        {t("tasks.form.is_active")}
                                    </FormLabel>

                                    <FormControl>
                                        <div
                                            className="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                                            <div className="items-center flex space-x-2">
                                                <Checkbox
                                                    id="isActive"
                                                    checked={field.value}
                                                    onCheckedChange={(value) => {
                                                        field.onChange(value)
                                                    }}
                                                />
                                                <label
                                                    htmlFor="isActive"
                                                    className="text-sm flex-grow font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                                >
                                                    {t("tasks.form.is_active.desc")}
                                                </label>
                                            </div>
                                        </div>
                                    </FormControl>

                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        <TaskFormFields form={form} tradingAccount={tradingAccount}/>

                        <DialogFooter className="w-full flex gap-2">
                            {/*<Button*/}
                            {/*    variant="outline"*/}
                            {/*    onClick={() => setOpenDialog(false)}>*/}
                            {/*    {t("common.cancel")}*/}
                            {/*</Button>*/}
                            <TaskDelete task={task} onDelete={onCompleteAction}/>
                            <Spacer/>
                            <Button type="submit">
                                {t("tasks.form.submit")}
                            </Button>
                        </DialogFooter>
                    </form>
                </Form>

            </DialogContent>
        </Dialog>
    )
}

function convertCronToHours(cron?: string): string[] {
    if (!cron) return []
    const parsedCron = parseCronExpression(cron)
    return parsedCron.fields.hour.toString().split(',')
}

function convertCronToDays(cron?: string): string[] {
    if (!cron) return []
    const parsedCron = parseCronExpression(cron)
    return parsedCron.fields.dayOfWeek.toString().split(',')
}
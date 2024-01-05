import {useMutation} from "@apollo/client";
import {Task, UpdateTaskDocument} from "@/graph/generated/generated";
import React, {useState} from "react";
import {useAppDispatch} from "@/store";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {Form, FormControl, FormField, FormItem, FormLabel} from "@/components/ui/form";
import {TaskForm, TaskFromSchema} from "@/app/tasks/form";
import {parseCronExpression} from "@/lib/cron-parser";
import {refreshTask} from "@/store/taskSlice";
import {Checkbox} from "@/components/ui/checkbox";
import {Label} from "@/components/ui/label";
import {errorToast} from "@/components/toast";
import {useTranslation} from "react-i18next";

export function EditTask({task}: { task: Task }) {
    const {t} = useTranslation();
    const [updateTask] = useMutation(UpdateTaskDocument)
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof TaskFromSchema>>({
        resolver: zodResolver(TaskFromSchema),
        defaultValues: {
            type: convertStringToTaskType(task.type),
            currency: task.currency,
            size: task.size,
            symbol: task.symbol,
            days: convertCronToDays(task.cron),
            hours: convertCronToHours(task.cron),
            isActive: task.isActive,
            grid: parseGridParams(task)
        },
    })

    function onSubmit(data: z.infer<typeof TaskFromSchema>) {
        updateTask({
            variables: {
                tradingAccountID: task.tradingAccountID,
                taskID: task.id,
                currency: data.currency,
                size: data.size,
                symbol: data.symbol,
                days: data.days,
                hours: data.hours,
                type: data.type,
                isActive: data.isActive,
                params: parseParamsFromData(data)
            }
        }).then(() => {
            setOpenDialog(false)
            dispatch(refreshTask({
                tradingAccountID: task.tradingAccountID,
                refresh: true
            }))
        }).catch((e) => {
            errorToast(e.message)
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="ghost">
                    {t("tasks.form.edit.btn")}
                </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>
                                {t("tasks.form.edit.title")}
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
                                    <div
                                        className="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                                        <FormControl>
                                            <Checkbox
                                                checked={field.value}
                                                onCheckedChange={(value) => {
                                                    field.onChange(value)
                                                }}
                                            />
                                        </FormControl>
                                        <div className="space-y-1 leading-none">
                                            <Label className={"text-sm font-light"}>
                                                {t("tasks.form.is_active.desc")}
                                            </Label>
                                        </div>
                                    </div>
                                </FormItem>
                            )}
                        />

                        <TaskForm form={form}/>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    )
}


function convertCronToHours(cron: string): string {
    const parsedCron = parseCronExpression(cron)
    return parsedCron.fields.hour.toString()
}

function convertCronToDays(cron: string): string {
    const parsedCron = parseCronExpression(cron)
    const result = parsedCron.fields.dayOfWeek.toString()

    if (result === "0,1,2,3,4,5,6,7") {
        return "*"
    }

    return result
}

function convertStringToTaskType(value: string): "DCA" | "Grid" {
    return value === "DCA" ? "DCA" : "Grid"
}

function parseGridParams(task: Task): { gap: number, quantity: number } {
    if (task.type === "Grid") {
        return {
            gap: task.params?.gap,
            quantity: task.params?.quantity
        }
    }

    return {
        gap: 0,
        quantity: 0,
    }
}

function parseParamsFromData(data: z.infer<typeof TaskFromSchema>): { gap: number, quantity: number } | null {
    if (data.type === "Grid") {
        return {
            gap: data.grid?.gap ?? 0,
            quantity: data.grid?.quantity ?? 0,
        }
    }

    return null
}



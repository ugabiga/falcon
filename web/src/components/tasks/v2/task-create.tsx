import {Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {useTranslation} from "react-i18next";
import {Form} from "@/components/ui/form";
import React, {useState} from "react";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {TaskType} from "@/lib/model";
import {TaskFormFields, TaskFromSchema} from "@/components/tasks/v2/task-form";
import {useMutation} from "@apollo/client";
import {CreateTaskDocument} from "@/graph/generated/generated";
import {capture} from "@/lib/posthog";
import {RefreshTarget} from "@/store/refresherSlice";
import {errorToast} from "@/components/toast";
import {translatedError} from "@/lib/error";
import {useSendRefreshSignal} from "@/lib/use-refresh";
import {ModelTradingAccount} from "@/api/model";

export default function TaskCreate(
    {
        tradingAccount
    }: {
        tradingAccount: ModelTradingAccount
    }) {

    const {t} = useTranslation()
    const [openDialog, setOpenDialog] = useState(false)
    const form = useForm<z.infer<typeof TaskFromSchema>>({
        resolver: zodResolver(TaskFromSchema),
        defaultValues: {
            type: TaskType.DCA,
            isActive: true,
        },
    })

    const {sendRefresh} = useSendRefreshSignal()
    const [createTask] = useMutation(CreateTaskDocument);

    function onSubmit(data: z.infer<typeof TaskFromSchema>) {
        createTask({
            variables: {
                tradingAccountID: tradingAccount.id!,
                currency: data.currency,
                size: Number(data.size),
                symbol: data.symbol,
                days: data.days.join(","),
                hours: data.hours.join(","),
                type: data.type,
            }
        }).then(() => {
            capture("Task Added", {
                tradingAccountID: tradingAccount.id,
                taskType: data.type
            })
            setOpenDialog(false)
            form.reset()
            sendRefresh(RefreshTarget.Task)
        }).catch((e) => {
            errorToast(translatedError(t, e.message))
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="outline">
                    {t("tasks.form.add.btn")}
                </Button>
            </DialogTrigger>
            <DialogContent className={"sm:max-w-[500px] overflow-y-scroll h-[calc(100dvh)] sm:h-auto"}>

                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>
                                {t("tasks.form.add.title")}
                            </DialogTitle>
                        </DialogHeader>

                        <TaskFormFields form={form} tradingAccount={tradingAccount}/>

                        <DialogFooter>
                            <Button type="submit" className={"mt-6"}>
                                {t("tasks.form.submit")}
                            </Button>
                        </DialogFooter>
                    </form>
                </Form>

            </DialogContent>
        </Dialog>
    )
}


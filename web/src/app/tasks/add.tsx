import {useMutation} from "@apollo/client";
import {CreateTaskDocument, TradingAccount} from "@/graph/generated/generated";
import React, {useState} from "react";
import {useAppDispatch} from "@/store";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {Form} from "@/components/ui/form";
import {refreshTask} from "@/store/taskSlice";
import {errorToast} from "@/components/toast";
import {TaskForm, TaskFromSchema} from "@/app/tasks/form";
import {useTranslation} from "react-i18next";

export function AddTask({tradingAccount}: { tradingAccount: TradingAccount }) {
    const {t} = useTranslation();
    const dispatch = useAppDispatch()
    const [createTask] = useMutation(CreateTaskDocument);
    const [openDialog, setOpenDialog] = useState(false)
    const form = useForm<z.infer<typeof TaskFromSchema>>({
        resolver: zodResolver(TaskFromSchema),
        defaultValues: {
            hours: "",
            type: "DCA",
            isActive: true,
        },
    })

    function onSubmit(data: z.infer<typeof TaskFromSchema>) {
        console.log("data", data)
        createTask({
            variables: {
                tradingAccountID: tradingAccount.id,
                currency: data.currency,
                size: data.size,
                symbol: data.symbol,
                days: data.days,
                hours: data.hours,
                type: data.type,
            }
        }).then(() => {
            setOpenDialog(false)
            form.reset()
            dispatch(refreshTask({
                tradingAccountID: tradingAccount.id,
                refresh: true
            }))
        }).catch((e) => {
            errorToast(t("error." + e.message))
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="outline">
                    {t("tasks.form.add.btn")}
                </Button>
            </DialogTrigger>
            <DialogContent className={"sm:max-w-[425px] overflow-y-scroll h-[calc(100dvh)] sm:h-auto"}>
                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>
                                {t("tasks.form.add.title")}
                            </DialogTitle>
                        </DialogHeader>

                        <TaskForm form={form} tradingAccount={tradingAccount}/>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    )

}

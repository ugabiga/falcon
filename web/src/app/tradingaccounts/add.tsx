import {useMutation} from "@apollo/client";
import {CreateTradingAccountDocument} from "@/graph/generated/generated";
import React, {useState} from "react";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {TradingAccountForm, TradingAccountFormSchema} from "@/app/tradingaccounts/form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {Form} from "@/components/ui/form";
import {useAppDispatch} from "@/store";
import {refreshTradingAccount} from "@/store/tradingAccountSlice";
import {errorToast} from "@/components/toast";
import {useTranslation} from "react-i18next";

export function AddTradingAccount() {
    const {t} = useTranslation();
    const [createTradingAccount] = useMutation(CreateTradingAccountDocument);
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof TradingAccountFormSchema>>({
        resolver: zodResolver(TradingAccountFormSchema),
        defaultValues: {
            name: "",
            exchange: "upbit",
            key: "",
            secret: "",
        },
    })

    function onSubmit(data: z.infer<typeof TradingAccountFormSchema>) {
        if (data.secret === undefined || data.secret === null) {
            errorToast(t("trading_account.add.error.secret"))
            return
        }

        createTradingAccount({
            variables: {
                name: data.name,
                exchange: data.exchange,
                key: data.key,
                secret: data.secret ?? "",
            }
        }).then(() => {
            setOpenDialog(false)
            dispatch(refreshTradingAccount(true))
        }).catch((e) => {
            errorToast(t("error." + e.message))
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="outline">{t("trading_account.add.btn")}</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>{t("trading_account.add.title")}</DialogTitle>
                        </DialogHeader>

                        <TradingAccountForm form={form}/>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    )
}


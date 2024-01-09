import {useMutation} from "@apollo/client";
import {TradingAccount, UpdateTradingAccountDocument} from "@/graph/generated/generated";
import React, {useState} from "react";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {Form} from "@/components/ui/form";
import {useAppDispatch} from "@/store";
import {refreshTradingAccount} from "@/store/tradingAccountSlice";
import {TradingAccountForm, TradingAccountFormSchema} from "@/app/tradingaccounts/form";
import {errorToast} from "@/components/toast";
import {useTranslation} from "react-i18next";
import {DropdownMenuItem} from "@/components/ui/dropdown-menu";


export function EditTradingAccount(
    {tradingAccount}: { tradingAccount: TradingAccount }
) {
    const {t} = useTranslation();
    const [updateTradingAccount] = useMutation(UpdateTradingAccountDocument);
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof TradingAccountFormSchema>>({
        resolver: zodResolver(TradingAccountFormSchema),
        defaultValues: {
            name: tradingAccount.name,
            exchange: tradingAccount.exchange,
            key: tradingAccount.key,
            secret: "",
        },
    })

    function onSubmit(data: z.infer<typeof TradingAccountFormSchema>) {
        updateTradingAccount({
            variables: {
                id: tradingAccount.id,
                ...data
            }
        }).then(() => {
            setOpenDialog(false)
            dispatch(refreshTradingAccount(true))
        }).catch((e) => {
            errorToast(t("error."+ e.message))
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                    {t("trading_account.edit.btn")}
                </DropdownMenuItem>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>{t("trading_account.edit.title")}</DialogTitle>
                        </DialogHeader>

                        <TradingAccountForm form={form}/>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    )
}


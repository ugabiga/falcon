import {DeleteTradingAccountDocument, TradingAccount} from "@/graph/generated/generated";
import {DropdownMenuItem} from "@/components/ui/dropdown-menu";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger
} from "@/components/ui/alert-dialog";
import {useMutation} from "@apollo/client";
import {errorToast} from "@/components/toast";

export function DeleteTradingAccount(
    {tradingAccount}: { tradingAccount: TradingAccount }
) {
    const {t} = useTranslation();
    const [openDialog, setOpenDialog] = useState(false)
    const [deleteTradingAccount] = useMutation(DeleteTradingAccountDocument);

    const handleDelete = () => {
        deleteTradingAccount({
            variables: {
                id: tradingAccount.id
            }
        }).then(() => {
            console.log("success");
        }).catch(error => {
            errorToast(error.message)
        })
    }

    return (
        <AlertDialog open={openDialog} onOpenChange={setOpenDialog}>
            <AlertDialogTrigger asChild>
                <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                    {t("trading_account.delete.btn")}
                </DropdownMenuItem>
            </AlertDialogTrigger>
            <AlertDialogContent>
                <AlertDialogHeader>
                    <AlertDialogTitle>
                        {t("trading_account.delete.title")}
                    </AlertDialogTitle>
                    <AlertDialogDescription>
                        {t("trading_account.delete.description")}
                    </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter className="flex">
                    <AlertDialogCancel>
                        {t("trading_account.delete.cancel")}
                    </AlertDialogCancel>
                    <div className="flex-grow"/>
                    <AlertDialogAction
                        className="btn btn-danger"
                        onClick={() => handleDelete()}
                    >
                        {t("trading_account.delete.yes")}
                    </AlertDialogAction>
                </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
    )
}
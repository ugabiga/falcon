"use client";

import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow,} from "@/components/ui/table"
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog"

import {GetTradingAccountsDocument, TradingAccount} from "@/graph/generated/generated";
import {useQuery} from "@apollo/client";
import {Button} from "@/components/ui/button";
import {useEffect} from "react";
import {AddTradingAccount} from "@/app/tradingaccounts/add";
import {EditTradingAccount} from "@/app/tradingaccounts/edit";
import {useAppDispatch, useAppSelector} from "@/store";
import {refreshTradingAccount,} from "@/store/tradingAccountSlice";

function camelize(str: string) {
    return str.replace(/^\w|[A-Z]|\b\w/g, function (word, index) {
        return index === 0 ? word.toUpperCase() : word.toLowerCase();
    }).replace(/\s+/g, '');
}

export default function TradingAccounts() {
    const {data, loading, refetch} = useQuery(GetTradingAccountsDocument);
    const tradingAccount = useAppSelector((state) => state.tradingAccount);
    const dispatch = useAppDispatch()

    useEffect(() => {
        if (tradingAccount?.refresh) {
            console.log(tradingAccount)
            refetch()
                .then(r => data)
                .then(r => {
                    dispatch(refreshTradingAccount(false))
                })
        }
    }, [tradingAccount])

    if (loading || !data) {
        return <div>Loading...</div>;
    }

    return (
        <main className="min-h-screen p-12">
            <h1 className="text-3xl font-bold">Trading Accounts</h1>

            <div className={"w-full flex space-x-2"}>
                <div className={"flex-grow"}></div>
                <AddTradingAccount/>
            </div>

            <div className="mt-6">
                {/*@ts-ignore*/}
                <TradingAccountTable tradingAccounts={data.tradingAccounts}/>
            </div>
        </main>
    )
}

function TradingAccountTable({tradingAccounts}: { tradingAccounts?: TradingAccount[] }) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">Exchange</TableHead>
                    <TableHead>Currency</TableHead>
                    <TableHead>Identifier</TableHead>
                    <TableHead>IP</TableHead>
                    <TableHead>Action</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    tradingAccounts?.map((tradingAccount) => (
                        <TableRow key={tradingAccount.id}>
                            <TableCell className="font-medium">{camelize(tradingAccount.exchange)}</TableCell>
                            <TableCell>{tradingAccount.currency}</TableCell>
                            <TableCell>{tradingAccount.identifier}</TableCell>
                            <TableCell>{tradingAccount.ip}</TableCell>
                            <TableCell>
                                <EditTradingAccount tradingAccount={tradingAccount}/>
                            </TableCell>
                        </TableRow>
                    ))
                    ??
                    <TableRow>
                        <TableCell className="font-medium">No trading accounts found</TableCell>
                    </TableRow>
                }
            </TableBody>
        </Table>
    )

}

function ClearTradingAccountDialog() {
    return (
        <AlertDialog>
            <AlertDialogTrigger asChild>
                <Button variant="outline">Clear All</Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
                <AlertDialogHeader>
                    <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                    <AlertDialogDescription>
                        This action cannot be undone. This will delete your all accounts.
                    </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                    <AlertDialogCancel>Cancel</AlertDialogCancel>
                    <AlertDialogAction
                        // onClick={onClearAll}
                    >
                        Continue
                    </AlertDialogAction>
                </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
    )

}

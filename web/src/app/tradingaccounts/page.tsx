"use client";

import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
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
import {zodResolver} from "@hookform/resolvers/zod";
import * as z from "zod";
import {
    Dialog,
    DialogContent,
    DialogDescription, DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"

import {CreateTradingAccountDocument, GetTradingAccountsDocument, TradingAccount} from "@/graph/generated/generated";
import {useMutation, useQuery} from "@apollo/client";
import {Button} from "@/components/ui/button";
import {useState} from "react";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useForm} from "react-hook-form";
import {Input} from "@/components/ui/input";

function camelize(str: string) {
    return str.replace(/^\w|[A-Z]|\b\w/g, function (word, index) {
        return index === 0 ? word.toUpperCase() : word.toLowerCase();
    }).replace(/\s+/g, '');
}

export default function TradingAccounts() {
    const {data, loading} = useQuery(GetTradingAccountsDocument);
    const [openAdd, setOpenAdd] = useState(false)
    const [openImport, setOpenImport] = useState(false)

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

function AddTradingAccount() {
    const [createTradingAccount] = useMutation(CreateTradingAccountDocument);
    const [openDialog, setOpenDialog] = useState(false)

    const formSchema = z.object({
        exchange: z
            .string({
                required_error: "Please enter a exchange",
            })
            .min(1, {
                message: "Please enter a exchange",
            }),
        currency: z
            .string({
                required_error: "Please enter a currency",
            })
            .min(1, {
                message: "Please enter a currency",
            }),
        identifier: z
            .string({
                required_error: "Please enter a identifier",
            })
            .min(1, {
                message: "Please enter a identifier",
            }),
        credential: z
            .string({
                required_error: "Please enter a credential",
            })
            .min(1, {
                message: "Please enter a credential",
            }),
    })
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            exchange: "",
            currency: "",
            identifier: "",
            credential: "",
        },
    })

    function onSubmit(data: z.infer<typeof formSchema>) {
        createTradingAccount({
            variables: {
                exchange: data.exchange,
                currency: data.currency,
                identifier: data.identifier,
                credential: data.credential,
            }
        }).then(() => {
            setOpenDialog(false)
            form.reset()
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="outline">Add</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <Form {...form}>
                    <form className={"grid gap-2 py-4"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader>
                            <DialogTitle>Add Trading Account</DialogTitle>
                        </DialogHeader>

                        <FormField
                            control={form.control}
                            name="exchange"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Exchange</FormLabel>
                                    <FormControl>
                                        <Input placeholder="exchange" {...field} />
                                    </FormControl>
                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        <FormField
                            control={form.control}
                            name="currency"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Currency</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Currency" {...field} />
                                    </FormControl>
                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        <FormField
                            control={form.control}
                            name="identifier"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Identifier</FormLabel>
                                    <FormControl>
                                        <Input placeholder="identifier" {...field} />
                                    </FormControl>
                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        <FormField
                            control={form.control}
                            name="credential"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Credential</FormLabel>
                                    <FormControl>
                                        <Input placeholder="credential" {...field} />
                                    </FormControl>
                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        {/* Submit */}
                        <DialogFooter className={"mt-4"}>
                            <Button type="submit">Save changes</Button>
                        </DialogFooter>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
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

function TradingAccountTable({
                                 tradingAccounts
                             }: {
    tradingAccounts?: TradingAccount[]
}) {

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
                            <TableCell>Action</TableCell>
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

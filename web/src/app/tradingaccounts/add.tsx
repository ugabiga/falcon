import {useMutation} from "@apollo/client";
import {CreateTradingAccountDocument} from "@/graph/generated/generated";
import {useState} from "react";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {AddTradingAccountFormSchema} from "@/app/tradingaccounts/form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {useAppDispatch} from "@/store";
import {refreshTradingAccount} from "@/store/tradingAccountSlice";
import {errorToast} from "@/components/toast";


export function AddTradingAccount() {
    const [createTradingAccount] = useMutation(CreateTradingAccountDocument);
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof AddTradingAccountFormSchema>>({
        resolver: zodResolver(AddTradingAccountFormSchema),
        defaultValues: {
            name: "",
            exchange: "upbit",
            currency: "KRW",
            identifier: "",
            credential: "",
        },
    })

    function onSubmit(data: z.infer<typeof AddTradingAccountFormSchema>) {
        createTradingAccount({
            variables: {
                name: data.name,
                exchange: data.exchange,
                identifier: data.identifier,
                credential: data.credential,
            }
        }).then(() => {
            setOpenDialog(false)
            form.reset()
            dispatch(refreshTradingAccount(true))
        }).catch((e) => {
            errorToast(e.message)
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="outline">Add</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>Add Trading Account</DialogTitle>
                        </DialogHeader>

                        <FormField
                            control={form.control}
                            name="name"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Name</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Name" {...field} />
                                    </FormControl>
                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        <FormField
                            control={form.control}
                            name="exchange"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Exchange</FormLabel>
                                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                                        <FormControl>
                                            <SelectTrigger>
                                                <SelectValue placeholder="Select a Exchange"/>
                                            </SelectTrigger>
                                        </FormControl>
                                        <SelectContent>
                                            <SelectItem value="upbit">Upbit</SelectItem>
                                            <SelectItem value="binance">Binance</SelectItem>
                                        </SelectContent>
                                    </Select>
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
                                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                                        <FormControl>
                                            <SelectTrigger>
                                                <SelectValue placeholder="Select a currency"/>
                                            </SelectTrigger>
                                        </FormControl>
                                        <SelectContent>
                                            <SelectItem value="KRW">KRW</SelectItem>
                                            <SelectItem value="USD">USD</SelectItem>
                                        </SelectContent>
                                    </Select>
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
                                        <Input placeholder="Identifier" {...field} />
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
                                        <Input type="password" placeholder="Credential" {...field} />
                                    </FormControl>
                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        {/* Submit */}
                        <DialogFooter>
                            <Button type="submit" className={"mt-6"}>Save changes</Button>
                        </DialogFooter>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    )
}

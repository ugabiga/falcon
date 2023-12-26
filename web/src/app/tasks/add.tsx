import {useMutation} from "@apollo/client";
import {CreateTradingAccountDocument} from "@/graph/generated/generated";
import {useState} from "react";
import {useAppDispatch} from "@/store";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {AddTradingAccountFormSchema} from "@/app/tradingaccounts/form";
import {zodResolver} from "@hookform/resolvers/zod";
import {refreshTradingAccount} from "@/store/tradingAccountSlice";
import {Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {Input} from "@/components/ui/input";
import {AddTaskForm} from "@/app/tasks/form";

export function AddTask() {
    // const [createTradingAccount] = useMutation(CreateTradingAccountDocument);
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof AddTaskForm>>({
        resolver: zodResolver(AddTaskForm),
        defaultValues: {
            name: "upbit",
        },
    })

    function onSubmit(data: z.infer<typeof AddTaskForm>) {
        // createTradingAccount({
        //     variables: {
        //         exchange: data.exchange,
        //         currency: data.currency,
        //         identifier: data.identifier,
        //         credential: data.credential,
        //     }
        // }).then(() => {
        //     setOpenDialog(false)
        //     form.reset()
        //     dispatch(refreshTradingAccount(true))
        // })
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

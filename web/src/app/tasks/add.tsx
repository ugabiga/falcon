import {useMutation} from "@apollo/client";
import {CreateTaskDocument} from "@/graph/generated/generated";
import React, {useState} from "react";
import {useAppDispatch} from "@/store";
import {useForm} from "react-hook-form";
import * as z from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {Input} from "@/components/ui/input";
import {AddTaskForm} from "@/app/tasks/form";
import {refreshTask} from "@/store/taskSlice";
import {errorToast} from "@/components/toast";
import {DaysOfWeekSelector} from "@/components/days-of-week-selector";

export function AddTask({tradingAccountID}: { tradingAccountID?: string }) {
    if (!tradingAccountID) {
        return null
    }

    const [createTask] = useMutation(CreateTaskDocument);
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof AddTaskForm>>({
        resolver: zodResolver(AddTaskForm),
        defaultValues: {
            hours: "",
            type: "DCA"
        },
    })

    function onSubmit(data: z.infer<typeof AddTaskForm>) {
        createTask({
            variables: {
                tradingAccountID: tradingAccountID!,
                currency: data.currency,
                days: data.days,
                hours: data.hours,
                type: data.type,
            }
        }).then(() => {
            setOpenDialog(false)
            form.reset()
            dispatch(refreshTask({
                tradingAccountID: tradingAccountID,
                refresh: true
            }))
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
                            <DialogTitle>Add Task</DialogTitle>
                        </DialogHeader>

                        <FormField
                            control={form.control}
                            name="type"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Type</FormLabel>
                                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                                        <FormControl>
                                            <SelectTrigger>
                                                <SelectValue placeholder="Select a Type"/>
                                            </SelectTrigger>
                                        </FormControl>
                                        <SelectContent>
                                            <SelectItem value="DCA">DCA</SelectItem>
                                            <SelectItem value="Grid">Grid</SelectItem>
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
                            name="days"
                            render={({field}) => (
                                <FormItem className="min-h-12">
                                    <FormLabel>Days</FormLabel>
                                    <FormMessage/>
                                    <DaysOfWeekSelector selectedDaysInString={field.value} onChange={
                                        (days) => {
                                            console.log(days)
                                            field.onChange(days)
                                        }
                                    }/>
                                </FormItem>
                            )}
                        />

                        <FormField
                            control={form.control}
                            name="hours"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Execution Hours(Format 1,5,13,23)</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Hours" {...field} />
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

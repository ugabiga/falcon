import {useMutation} from "@apollo/client";
import {Task, UpdateTaskDocument} from "@/graph/generated/generated";
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
import {UpdateTaskForm} from "@/app/tasks/form";
import {parseCronExpression} from "@/lib/cron-parser";
import {refreshTask} from "@/store/taskSlice";
import {Checkbox} from "@/components/ui/checkbox";
import {Label} from "@/components/ui/label";
import {errorToast} from "@/components/toast";


function convertCronToHours(cron: string): string {
    const parsedCron = parseCronExpression(cron)
    return parsedCron.fields.hour.toString()
}

function convertStringToTaskType(value: string): "DCA" | "Grid" {
    return value === "DCA" ? "DCA" : "Grid"
}

export function EditTask({task}: { task: Task }) {
    const [updateTask] = useMutation(UpdateTaskDocument)
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof UpdateTaskForm>>({
        resolver: zodResolver(UpdateTaskForm),
        defaultValues: {
            hours: convertCronToHours(task.cron),
            type: convertStringToTaskType(task.type),
            isActive: task.isActive
        },
    })

    function onSubmit(data: z.infer<typeof UpdateTaskForm>) {
        updateTask({
            variables: {
                id: task.id,
                hours: data.hours,
                type: data.type,
                isActive: task.isActive
            }
        }).then(() => {
            setOpenDialog(false)
            form.reset()
            dispatch(refreshTask({
                tradingAccountID: task.tradingAccountID,
                refresh: true
            }))
        }).catch((e) => {
            errorToast(e.message)
        })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="ghost">Edit</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <Form {...form}>
                    <form className={"grid gap-2 py-4 space-y-2"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader className="mb-2">
                            <DialogTitle>Edit task</DialogTitle>
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

                        <FormField
                            control={form.control}
                            name="isActive"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>
                                        Is Active
                                    </FormLabel>
                                    <div
                                        className="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                                        <FormControl>
                                            <Checkbox
                                                checked={field.value}
                                                onChange={field.onChange}
                                            />
                                        </FormControl>
                                        <div className="space-y-1 leading-none">
                                            <Label className={"text-sm font-light"}>
                                                If checked, this task will be executed
                                            </Label>
                                        </div>
                                    </div>
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

import {useMutation} from "@apollo/client";
import {CreateTaskDocument, Task} from "@/graph/generated/generated";
import {useState} from "react";
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
import {parseCronExpression} from "@/lib/cron-parser";


function convertCronToHours(cron: string): string {
    const parsedCron = parseCronExpression(cron)
    return parsedCron.fields.hour.toString()
}

function convertStringToTaskType(value: string): "DCA" | "Grid" {
    return value === "DCA" ? "DCA" : "Grid"
}

export function EditTask({task}: { task: Task }) {
    const [createTask] = useMutation(CreateTaskDocument);
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof AddTaskForm>>({
        resolver: zodResolver(AddTaskForm),
        defaultValues: {
            hours: convertCronToHours(task.cron),
            type: convertStringToTaskType(task.type)
        },
    })

    function onSubmit(data: z.infer<typeof AddTaskForm>) {
        // createTask({
        //     variables: {
        //         tradingAccountID: tradingAccountID,
        //         hours: data.hours,
        //         type: data.type,
        //     }
        // }).then(() => {
        //     setOpenDialog(false)
        //     form.reset()
        //     dispatch(refreshTask({
        //         tradingAccountID: tradingAccountID,
        //         refresh: true
        //     }))
        // })
    }

    return (
        <Dialog open={openDialog} onOpenChange={setOpenDialog}>
            <DialogTrigger asChild>
                <Button variant="outline">Edit</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <Form {...form}>
                    <form className={"grid gap-2 py-4"}
                          onSubmit={form.handleSubmit(onSubmit)}
                    >
                        <DialogHeader>
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

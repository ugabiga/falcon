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
import {DaysOfWeekSelector} from "@/components/days-of-week-selector";
import {NumericFormatInput} from "@/components/numeric-format-input";


function convertCronToHours(cron: string): string {
    const parsedCron = parseCronExpression(cron)
    return parsedCron.fields.hour.toString()
}

function convertCronToDays(cron: string): string {
    const parsedCron = parseCronExpression(cron)
    const result = parsedCron.fields.dayOfWeek.toString()

    if (result === "0,1,2,3,4,5,6,7") {
        return "*"
    }

    return result
}

function convertStringToTaskType(value: string): "DCA" | "Grid" {
    return value === "DCA" ? "DCA" : "Grid"
}

function parseGridParams(task: Task): { gap: number, quantity: number } {
    if (task.type === "Grid") {
        return {
            gap: task.params?.gap,
            quantity: task.params?.quantity
        }
    }

    return {
        gap: 0,
        quantity: 0,
    }
}

function parseParamsFromData(data: z.infer<typeof UpdateTaskForm>): { gap: number, quantity: number } | null {
    if (data.type === "Grid") {
        return {
            gap: data.grid.gap,
            quantity: data.grid.quantity
        }
    }

    return null
}


export function EditTask({task}: { task: Task }) {
    const [updateTask] = useMutation(UpdateTaskDocument)
    const [openDialog, setOpenDialog] = useState(false)
    const dispatch = useAppDispatch()

    const form = useForm<z.infer<typeof UpdateTaskForm>>({
        resolver: zodResolver(UpdateTaskForm),
        defaultValues: {
            type: convertStringToTaskType(task.type),
            currency: task.currency,
            size: task.size,
            symbol: task.symbol,
            days: convertCronToDays(task.cron),
            hours: convertCronToHours(task.cron),
            isActive: task.isActive,
            grid: parseGridParams(task)
        },
    })

    function onSubmit(data: z.infer<typeof UpdateTaskForm>) {
        updateTask({
            variables: {
                id: task.id,
                currency: data.currency,
                size: data.size,
                symbol: data.symbol,
                days: data.days,
                hours: data.hours,
                type: data.type,
                isActive: data.isActive,
                params: parseParamsFromData(data)
            }
        }).then(() => {
            setOpenDialog(false)
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
                            name="symbol"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>Crypto Symbol</FormLabel>
                                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                                        <FormControl>
                                            <SelectTrigger>
                                                <SelectValue placeholder="Select a crypto currency"/>
                                            </SelectTrigger>
                                        </FormControl>
                                        <SelectContent>
                                            <SelectItem value="BTC">BTC</SelectItem>
                                            <SelectItem value="ETH">ETH</SelectItem>
                                        </SelectContent>
                                    </Select>
                                    <FormMessage/>
                                </FormItem>
                            )}
                        />

                        <FormField
                            control={form.control}
                            name="size"
                            render={({field}) => (
                                <FormItem>
                                    <FormLabel>
                                        Investing Size {form.watch("symbol")}
                                    </FormLabel>
                                    <FormControl>
                                        <NumericFormatInput
                                            value={field.value}
                                            thousandSeparator={true}
                                            allowNegative={false}
                                            allowLeadingZeros={false}
                                            fixedDecimalScale={false}
                                            suffix={" " + form.watch("symbol")}
                                            onValueChange={(values) => {
                                                field.onChange(values.floatValue)
                                            }}
                                        />

                                    </FormControl>
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

                        {form.watch("type") === "Grid" && <GridFormFields form={form}/>}

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
                                                onCheckedChange={(value) => {
                                                    field.onChange(value)
                                                }}
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

function GridFormFields({form}: { form: ReturnType<typeof useForm<z.infer<typeof UpdateTaskForm>>> }) {
    return (
        <>
            <FormField
                control={form.control}
                name="grid.gap"
                render={({field}) => (
                    <FormItem>
                        <FormLabel>Grid Gap (%)</FormLabel>
                        <FormControl>
                            <Input type="number"
                                   min={0}
                                   placeholder="Grid gap"
                                   value={field.value}
                                   onChange={(e) => {
                                       field.onChange(Number(e.target.value))
                                   }}
                            />
                        </FormControl>
                        <FormMessage/>
                    </FormItem>
                )}
            />

            <FormField
                control={form.control}
                name="grid.quantity"
                render={({field}) => (
                    <FormItem>
                        <FormLabel>Grid quantity</FormLabel>
                        <FormControl>
                            <Input type="number" placeholder="Grid quantity"
                                   value={field.value}
                                   onChange={(e) => {
                                       field.onChange(Number(e.target.value))
                                   }}
                            />
                        </FormControl>
                        <FormMessage/>
                    </FormItem>
                )}
            />
        </>
    )
}

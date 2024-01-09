import * as z from "zod";
import {useForm} from "react-hook-form";
import {useTranslation} from "react-i18next";
import {TradingAccountFormSchema} from "@/app/tradingaccounts/form";
import {FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {NumericFormatInput} from "@/components/numeric-format-input";
import {DaysOfWeekSelector} from "@/components/days-of-week-selector";
import {Input} from "@/components/ui/input";
import React from "react";
import {DialogFooter} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";

export const TaskFromSchema = z.object({
    currency: z
        .string({
            required_error: "Please enter a currency",
        }),
    size: z
        .number({
            required_error: "Please enter a currency size",
        }),
    symbol: z
        .string({
            required_error: "Please enter a exchange",
        })
        .min(1, {
            message: "Please enter a exchange",
        }),
    days: z
        .string({
            required_error: "Please enter a exchange",
        })
        .min(1, {
            message: "Please enter a exchange",
        }),
    hours: z
        .string({
            required_error: "Please enter a exchange",
        })
        .min(1, {
            message: "Please enter a exchange",
        }),
    type: z
        .enum(["DCA", "Grid"]),
    isActive: z
        .boolean({
            required_error: "Please enter a exchange",
        }),
    grid: z
        .object({
            gap: z.number({}),
            quantity: z.number({}),
        })
        .optional()
})


export function TaskForm({form}: {
    form: ReturnType<typeof useForm<z.infer<typeof TaskFromSchema>>>
}) {
    const {t} = useTranslation();

    return <>
        <FormField
            control={form.control}
            name="type"
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("tasks.form.type")}
                    </FormLabel>
                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                        <FormControl>
                            <SelectTrigger>
                                <SelectValue placeholder={t("tasks.form.type")}/>
                            </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                            <SelectItem value="DCA">DCA</SelectItem>
                            <SelectItem value="Grid" disabled>Grid</SelectItem>
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
                    <FormLabel>
                        {t("tasks.form.currency")}
                    </FormLabel>
                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                        <FormControl>
                            <SelectTrigger>
                                <SelectValue placeholder={t("tasks.form.currency")}/>
                            </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                            <SelectItem value="KRW">KRW</SelectItem>
                            <SelectItem value="USD" disabled>USD</SelectItem>
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
                    <FormLabel>
                        {t("tasks.form.symbol")}
                    </FormLabel>
                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                        <FormControl>
                            <SelectTrigger>
                                <SelectValue placeholder={t("tasks.form.symbol")}/>
                            </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                            <SelectItem value="BTC">BTC</SelectItem>
                            <SelectItem value="ETH" disabled>ETH</SelectItem>
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
                        {t("tasks.form.size")} {form.watch("symbol")}
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
                    <FormLabel>
                        {t("tasks.form.days")}
                    </FormLabel>
                    <FormMessage/>
                    <DaysOfWeekSelector
                        selectedDaysInString={field.value}
                        placeholder={t("tasks.form.days.placeholder")}
                        onChange={
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
                    <FormLabel>
                        {t("tasks.form.hours")}
                    </FormLabel>
                    <FormControl>
                        <Input placeholder={t("tasks.form.hours")} {...field} />
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        {form.watch("type") === "Grid" && <GridFormFields form={form}/>}

        {/* Submit */}
        <DialogFooter>
            <Button type="submit" className={"mt-6"}>
                {t("tasks.form.submit")}
            </Button>
        </DialogFooter>
    </>
}

function GridFormFields({form}: { form: ReturnType<typeof useForm<z.infer<typeof TaskFromSchema>>> }) {
    const {t} = useTranslation();
    return (
        <>
            <FormField
                control={form.control}
                name="grid.gap"
                render={({field}) => (
                    <FormItem>
                        <FormLabel>
                            {t("tasks.form.grid.gap")}
                        </FormLabel>
                        <FormControl>
                            <Input type="number"
                                   min={0}
                                   placeholder={t("tasks.form.grid.gap")}
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
                        <FormLabel>
                            {t("tasks.form.grid.quantity")}
                        </FormLabel>
                        <FormControl>
                            <Input type="number" placeholder={t("tasks.form.grid.quantity")}
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

import {useForm} from "react-hook-form";
import {z} from "zod";
import {useTranslation} from "react-i18next";
import {TaskFromSchema} from "@/components/tasks/v2/task-form";
import SeparationHeader from "@/components/separation-header";
import React from "react";
import {FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Checkbox} from "@/components/ui/checkbox";
import {Label} from "@/components/ui/label";

export function TaskFormBuyingGridFields(
    {
        form,
        // tradingAccount
    }: {
        form: ReturnType<typeof useForm<z.infer<typeof TaskFromSchema>>>,
        // tradingAccount: any
    }) {
    const {t} = useTranslation();

    return <>
        <SeparationHeader name={t("tasks.form.buying-grid.title")}/>

        <FormField
            control={form.control}
            name="grid.gap_percent"
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("tasks.form.grid.gap")}
                    </FormLabel>
                    <FormControl>
                        <Input
                            type={"number"}
                            inputMode="decimal"
                            pattern="\d*.?\d*"
                            value={field.value || ""}
                            step={1}
                            placeholder={t("tasks.form.grid.gap.description")}
                            onChange={(e) => {
                                field.onChange(e.target.value)
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
                        <Input
                            placeholder={t("tasks.form.grid.quantity.description")}
                            type={"number"}
                            inputMode="decimal"
                            pattern="\d*"
                            value={field.value || ""}
                            step={1}
                            onChange={(e) => {
                                field.onChange(e.target.value)
                            }}
                        />
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        <FormField
            control={form.control}
            name="grid.use_incremental_size"
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("tasks.form.grid.use_incremental_size")}
                    </FormLabel>

                    <FormControl>
                        <div className="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                            <div className="items-center flex space-x-2">
                                <Checkbox
                                    id="use_incremental_size"
                                    checked={field.value}
                                    onCheckedChange={(value) => {
                                        field.onChange(value)
                                    }}
                                />
                                <label
                                    htmlFor="use_incremental_size"
                                    className="text-sm flex-grow font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                >
                                    {t("tasks.form.grid.use_incremental_size.description")}
                                </label>
                            </div>
                        </div>
                    </FormControl>

                    <FormMessage/>
                </FormItem>
            )}
        />

        <FormField
            control={form.control}
            name="grid.incremental_size"
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("tasks.form.grid.incremental_size")}
                    </FormLabel>
                    <FormControl>
                        <Input
                            type={"number"}
                            inputMode="decimal"
                            pattern="\d*.?\d*"
                            value={field.value || ""}
                            placeholder={t("tasks.form.grid.incremental_size.description")}
                            onChange={(e) => {
                                field.onChange(e.target.value)
                            }}
                        />
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        <FormField
            control={form.control}
            name="grid.delete_previous_orders"
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("tasks.form.grid.delete_previous_orders")}
                    </FormLabel>

                    <FormControl>
                        <div className="flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                            <div className="items-center flex space-x-2">
                                <Checkbox
                                    id="delete_previous_orders"
                                    checked={field.value}
                                    onCheckedChange={(value) => {
                                        field.onChange(value)
                                    }}
                                />
                                <label
                                    htmlFor="delete_previous_orders"
                                    className="text-sm flex-grow font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                >
                                    {t("tasks.form.grid.delete_previous_orders.description")}
                                </label>
                            </div>
                        </div>
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />
    </>
}
import * as z from "zod";
import {useForm} from "react-hook-form";
import {useTranslation} from "react-i18next";
import {FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {DaysOfWeekSelector} from "@/components/days-of-week-selector";
import {Input} from "@/components/ui/input";
import React from "react";
import {DialogFooter} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {useConvertSizeToCurrency} from "@/hooks/convert-size-to-currency";
import {TradingAccount} from "@/graph/generated/generated";
import {ExchangeList} from "@/lib/exchanges";
import {TaskType} from "@/lib/model";
import {Checkbox} from "@/components/ui/checkbox";
import {Label} from "@/components/ui/label";

export const TaskFromSchema = z.object({
    currency: z
        .string(),
    size: z
        .number(),
    symbol: z
        .string(),
    days: z
        .string(),
    hours: z
        .string(),
    type: z
        .nativeEnum(TaskType),
    isActive: z
        .boolean(),
    grid: z
        .object({
            gap_percent: z.number({}),
            quantity: z.number({}),
            use_incremental_size: z.boolean({}),
            incremental_size: z.number({}).optional(),
            delete_previous_orders: z.boolean({}),
        })
        .optional()
})

export function TaskForm(
    {
        form,
        tradingAccount,
    }: {
        form: ReturnType<typeof useForm<z.infer<typeof TaskFromSchema>>>
        tradingAccount: TradingAccount
    }) {
    const {t} = useTranslation();
    const {fetchConvertedTotal, convertedTotal} = useConvertSizeToCurrency()

    const handleSizeOnChange = () => {
        const symbol = form.watch("symbol")
        const currency = form.watch("currency")
        const size = form.watch("size")

        const canGetTicker = symbol && currency && size;
        if (!canGetTicker) {
            return;
        }

        fetchConvertedTotal(symbol, currency, size)
    }
    const handleOnTypeChange = () => {
        const type = form.watch("type")
        const gap_percent = form.watch("grid.gap_percent")
        if (type == TaskType.BuyingGrid && !gap_percent) {
            form.setValue("grid", {
                use_incremental_size: false,
                delete_previous_orders: true
            } as TaskGridParams)
        }
    }

    const addComma = (price: string) => {
        return price?.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    };

    return <>
        <FormField
            control={form.control}
            name="type"
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("tasks.form.type")}
                    </FormLabel>
                    <Select
                        onValueChange={(value) => {
                            field.onChange(value)
                            handleOnTypeChange()
                        }}
                        defaultValue={field.value}
                    >
                        <FormControl>
                            <SelectTrigger>
                                <SelectValue placeholder={t("tasks.form.type")}/>
                            </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                            <SelectItem value={TaskType.DCA}>{t("tasks.type.dca")}</SelectItem>
                            <SelectItem value={TaskType.BuyingGrid}>{t("tasks.type.buying_grid")}</SelectItem>
                        </SelectContent>
                    </Select>
                    <p className="text-sm text-gray-500 mb-2">
                        {
                            form.watch("type") === TaskType.DCA
                                ? t("tasks.form.type.dca.description")
                                : t("tasks.form.type.buying_grid.description")
                        }
                    </p>
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
                            {
                                ExchangeList
                                    .find((exchange) => exchange.value === tradingAccount.exchange)
                                    ?.supportCurrencies
                                    .map((currency, index) => {
                                        return <SelectItem key={index} value={currency}>{currency}</SelectItem>
                                    })
                            }
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
                    <Select
                        // onValueChange={field.onChange}
                        onValueChange={(value) => {
                            field.onChange(value)
                            handleSizeOnChange()
                        }}
                        defaultValue={field.value}>
                        <FormControl>
                            <SelectTrigger>
                                <SelectValue placeholder={t("tasks.form.symbol")}/>
                            </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                            <SelectItem value="BTC">Bitcoin BTC</SelectItem>
                            <SelectItem value="ETH">Ethereum ETH</SelectItem>
                            <SelectItem value="SOL">Solana SOL</SelectItem>
                            <SelectItem value="XRP">XRP</SelectItem>
                            <SelectItem value="ADA">Cardano ADA</SelectItem>
                            <SelectItem value="AVAX">Avalanche AVAX</SelectItem>
                            <SelectItem value="TRX">TRON TRX</SelectItem>
                            <SelectItem value="MATIC">Polygon MATIC</SelectItem>
                            <SelectItem value="ARB">Arbitrum ARB</SelectItem>
                            <SelectItem value="INJ">Injective Protocol INJ</SelectItem>
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
                        {t("tasks.form.size")} {form.watch("symbol")} {convertedTotal && `(${convertedTotal} ${form.watch("currency")})`}
                    </FormLabel>
                    <FormControl>
                        <Input
                            type={"number"}
                            value={field.value || 0}
                            onChange={(e) => {
                                field.onChange(Number(e.target.value))
                                handleSizeOnChange()
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
                        <Input
                            placeholder={t("tasks.form.hours.placeholder")}
                            {...field}
                        />
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        {form.watch("type") === TaskType.BuyingGrid && <GridFormFields form={form}/>}

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
                name="grid.quantity"
                render={({field}) => (
                    <FormItem>
                        <FormLabel>
                            {t("tasks.form.grid.quantity")}
                        </FormLabel>
                        <FormControl>
                            <Input type="number"
                                   step={1}
                                   placeholder={t("tasks.form.grid.quantity.description")}
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
                name="grid.gap_percent"
                render={({field}) => (
                    <FormItem>
                        <FormLabel>
                            {t("tasks.form.grid.gap")}
                        </FormLabel>
                        <FormControl>
                            <Input type="number"
                                   step={1}
                                   placeholder={t("tasks.form.grid.gap.description")}
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
                name="grid.use_incremental_size"
                render={({field}) => (
                    <FormItem>
                        <FormLabel>
                            {t("tasks.form.grid.use_incremental_size")}
                        </FormLabel>
                        <FormControl>
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
                                        {t("tasks.form.grid.use_incremental_size.description")}
                                    </Label>
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
                            <Input type="number"
                                   placeholder={t("tasks.form.grid.incremental_size.description")}
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
                name="grid.delete_previous_orders"
                render={({field}) => (
                    <FormItem>
                        <FormLabel>
                            {t("tasks.form.grid.delete_previous_orders")}
                        </FormLabel>
                        <FormControl>
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
                                        {t("tasks.form.grid.delete_previous_orders.description")}
                                    </Label>
                                </div>
                            </div>
                        </FormControl>
                        <FormMessage/>
                    </FormItem>
                )}
            />

        </>
    )
}

export interface TaskGridParams {
    gap_percent: number,
    quantity: number,
    use_incremental_size: boolean,
    incremental_size: number,
    delete_previous_orders: boolean

}

export function parseParamsFromData(data: z.infer<typeof TaskFromSchema>): TaskGridParams | null {
    if (data.type === TaskType.BuyingGrid) {
        return {
            gap_percent: data.grid?.gap_percent ?? 0,
            quantity: data.grid?.quantity ?? 0,
            use_incremental_size: data.grid?.use_incremental_size ?? false,
            incremental_size: data.grid?.incremental_size ?? 0,
            delete_previous_orders: data.grid?.delete_previous_orders ?? true
        }
    }

    return null
}



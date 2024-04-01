import {useForm} from "react-hook-form";
import {z} from "zod";
import {TaskType} from "@/lib/model";
import {FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import React from "react";
import {useTranslation} from "react-i18next";
import {ExchangeList} from "@/lib/exchanges";
import {Input} from "@/components/ui/input";
import {useConvertSizeToCurrency} from "@/hooks/convert-size-to-currency";
import {ToggleGroup, ToggleGroupItem} from "@/components/ui/toggle-group";
import SeparationHeader from "@/components/separation-header";
import {TaskFormBuyingGridFields} from "@/components/tasks/v2/task-form-grid";

const HOURS = Array.from({length: 24}, (_, i) => i)

const DAYS = [
    {
        value: "1",
        label: "monday"
    },
    {
        value: "2",
        label: "tuesday"
    },
    {
        value: "3",
        label: "wednesday"
    },
    {
        value: "4",
        label: "thursday"
    },
    {
        value: "5",
        label: "friday"
    },
    {
        value: "6",
        label: "saturday"
    },
    {
        value: "0",
        label: "sunday",
    }
]

export const TaskFromSchema = z.object({
    type: z.nativeEnum(TaskType),
    currency: z.string(),
    size: z.string()
        .regex(/^\d+\.?\d*$/, "Size must be a number"),
    symbol: z.string(),
    days: z.array(
        z.string()
    ),
    hours: z.array(
        z.string().regex(/^\d+$/, "Priority must be a number"),
    ),
    isActive: z.boolean(),
    grid: z.object({
        gap_percent: z.string({})
            .regex(/^\d+\.?\d*$/, "Gap percent must be a number"),
        quantity: z.string({})
            .regex(/^\d+$/, "Quantity must be a number"),
        use_incremental_size: z.boolean({}),
        incremental_size: z.string({})
            .regex(/^\d+\.?\d*$/, "Incremental size must be a number")
            .optional(),
        delete_previous_orders: z.boolean({}),
    }).optional()
})

export function TaskFormFields(
    {
        form,
        tradingAccount
    }: {
        form: ReturnType<typeof useForm<z.infer<typeof TaskFromSchema>>>,
        tradingAccount: any
    }) {
    const {t} = useTranslation();

    const {fetchConvertedTotal, convertedTotal} = useConvertSizeToCurrency()

    const handleSizeOnChange = () => {
        const symbol = form.watch("symbol")
        const currency = form.watch("currency")
        const size = form.watch("size")
        const numSize = Number(size)

        const canGetTicker = symbol && currency && size;
        if (!canGetTicker) {
            return;
        }

        fetchConvertedTotal(symbol, currency, numSize)
    }

    return <>
        <SeparationHeader name={t("tasks.form.basic.title")}/>

        <FormField
            control={form.control}
            name={"type"}
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("tasks.form.type")}
                    </FormLabel>
                    <Select
                        defaultValue={field.value}
                        onValueChange={(value) => {
                            field.onChange(value)
                        }}
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
                    <FormDescription>
                        {
                            form.watch("type") === TaskType.DCA
                                ? t("tasks.form.type.dca.description")
                                : t("tasks.form.type.buying_grid.description")
                        }
                    </FormDescription>
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
                        {t("tasks.form.size")}
                        {
                            form.watch("symbol")} {convertedTotal && `(${convertedTotal} ${form.watch("currency")})`
                    }
                    </FormLabel>
                    <FormControl>
                        <Input
                            placeholder={t("tasks.form.size")}
                            type={"number"}
                            inputMode="decimal"
                            pattern="\d*.?\d*"
                            value={field.value || ""}
                            onChange={(e) => {
                                field.onChange(e.target.value)
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
                <FormItem>
                    <FormLabel asChild>
                        <legend>
                            {t("tasks.form.days")}
                        </legend>
                    </FormLabel>
                    <FormControl>
                        <div className="py-1">
                            <ToggleGroup
                                className="grid grid-cols-4 gap-2"
                                type="multiple"
                                variant="outline"
                                value={field.value}
                                onValueChange={(value) => {
                                    field.onChange(value)
                                }}
                            >
                                {
                                    DAYS.map((inputValues) => {
                                        return (
                                            <ToggleGroupItem
                                                className="hover:bg-inherit"
                                                value={inputValues.value}
                                                key={inputValues.value}
                                            >
                                                {t("common.days." + inputValues.label)}
                                            </ToggleGroupItem>
                                        )
                                    })
                                }
                            </ToggleGroup>
                        </div>
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        <FormField
            control={form.control}
            name="hours"
            render={({field}) => (
                <FormItem>
                    <FormLabel asChild>
                        <legend>
                            {t("tasks.form.hours")}
                        </legend>
                    </FormLabel>
                    <FormControl>
                        <div className="py-1">
                            <ToggleGroup
                                className="grid grid-cols-6 gap-2"
                                variant="outline"
                                type="multiple"
                                value={field.value}
                                onValueChange={(value) => {
                                    field.onChange(value)
                                }}
                            >
                                {HOURS.map((v) => {
                                    return <ToggleGroupItem
                                        className="hover:bg-inherit"
                                        key={v.toString()}
                                        value={v.toString()}
                                    >
                                        {v.toString().padStart(2, "0") + ":00"}
                                    </ToggleGroupItem>
                                })}
                            </ToggleGroup>
                        </div>
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        {form.watch("type") === TaskType.BuyingGrid && <TaskFormBuyingGridFields form={form}/>}

    </>
}
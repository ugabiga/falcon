import * as z from "zod";
import {useForm} from "react-hook-form";
import {FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {DialogFooter} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";
import {useTranslation} from "react-i18next";
import React from "react";
import {ExchangeList} from "@/lib/exchanges";


export const TradingAccountFormSchema = z.object({
    name: z
        .string({
            required_error: "Please enter a name",
        })
        .min(1, {
            message: "Please enter a name",
        }),
    exchange: z
        .string({
            required_error: "Please enter a exchange",
        })
        .min(1, {
            message: "Please enter a exchange",
        }),
    key: z
        .string({
            required_error: "Please enter a identifier",
        })
        .min(1, {
            message: "Please enter a identifier",
        }),
    secret: z
        .union([
            z.string()
                .length(0),
            z.string()
                .min(1)
        ])
        .optional()
        .transform((val) => val === "" ? undefined : val),
})

export function TradingAccountForm({form}: {
    form: ReturnType<typeof useForm<z.infer<typeof TradingAccountFormSchema>>>
}) {
    const {t} = useTranslation();

    return <>
        <FormField
            control={form.control}
            name="name"
            render={({field}) => (
                <FormItem>
                    <FormLabel>{t("trading_account.form.name")}</FormLabel>
                    <FormControl>
                        <Input
                            placeholder={t("trading_account.form.name_placeholder")}
                            {...field}
                        />
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
                    <FormLabel>{t("trading_account.form.exchange")}</FormLabel>
                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                        <FormControl>
                            <SelectTrigger>
                                <SelectValue placeholder={t("trading_account.form.exchange")}/>
                            </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                            {ExchangeList.map((item, index) => {
                                return <SelectItem key={index} value={item.value}>
                                    {t(`common.exchange.${item.value}`)}
                                </SelectItem>
                            })}
                        </SelectContent>
                    </Select>
                    <FormMessage/>
                </FormItem>
            )}
        />

        <FormField
            control={form.control}
            name="key"
            render={({field}) => (
                <FormItem>
                    <FormLabel>
                        {t("trading_account.form.key")}
                    </FormLabel>
                    <FormControl>
                        <Input
                            type="password"
                            placeholder={t("trading_account.form.key_placeholder")}
                            {...field}
                        />
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        <FormField
            control={form.control}
            name="secret"
            render={({field}) => (
                <FormItem>
                    <FormLabel>{t("trading_account.form.secret")}</FormLabel>
                    <FormControl>
                        <Input
                            type="password"
                            placeholder={t("trading_account.form.secret_placeholder")}
                            {...field}
                        />
                    </FormControl>
                    <FormMessage/>
                </FormItem>
            )}
        />

        {/* Submit */}
        <DialogFooter>
            <Button type="submit" className={"mt-6"}>{t("trading_account.form.submit")}</Button>
        </DialogFooter>
    </>
}

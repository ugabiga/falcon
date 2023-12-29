import * as z from "zod";

export const AddTradingAccountFormSchema = z.object({
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
        .string({
            required_error: "Please enter a credential",
        })
        .min(1, {
            message: "Please enter a credential",
        }),
})

export const EditTradingAccountFormSchema = z.object({
    name: z
        .string({
            required_error: "Please enter a name",
        }),
    exchange: z
        .string({
            required_error: "Please enter a exchange",
        }),
    key: z
        .string({
            required_error: "Please enter a identifier",
        }),
    secret: z
        .string({
            required_error: "Please enter a credential",
        }),
})

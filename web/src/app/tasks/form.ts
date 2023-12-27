import * as z from "zod";

export const AddTaskForm= z.object({
    hours: z
        .string({
            required_error: "Please enter a exchange",
        })
        .min(1, {
            message: "Please enter a exchange",
        }),
    type: z
        .enum(["DCA", "Grid"])

})

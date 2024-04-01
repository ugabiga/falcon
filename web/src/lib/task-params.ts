import {z} from "zod";
import {TaskFromSchema} from "@/components/tasks/v2/task-form";
import {TaskType} from "@/lib/model";
import {Task} from "@/graph/generated/generated";
import {TaskGridInnerParams} from "@/components/tasks/old/form";
import {ModelTask} from "@/api/model";

export interface TaskGridInputParams {
    gap_percent: string
    quantity: string,
    use_incremental_size: boolean,
    incremental_size: string,
    delete_previous_orders: boolean
}

export interface TaskGridParams {
    gap_percent: number
    quantity: number,
    use_incremental_size: boolean,
    incremental_size: number,
    delete_previous_orders: boolean
}

export function parseParamsFromData(data: z.infer<typeof TaskFromSchema>): TaskGridParams | null {
    switch (data.type) {
        case TaskType.BuyingGrid:
            return {
                gap_percent: Number(data.grid?.gap_percent) ?? 0,
                quantity: Number(data.grid?.quantity) ?? 0,
                use_incremental_size: data.grid?.use_incremental_size ?? false,
                incremental_size: Number(data.grid?.incremental_size) ?? 0,
                delete_previous_orders: data.grid?.delete_previous_orders ?? true
            }
        default:
            return null
    }
}

export function parseParamsFromTask(task: ModelTask): TaskGridInputParams {
    if (task.type === TaskType.BuyingGrid) {
        return {
            gap_percent: String(task.params?.gap_percent) ?? 0,
            quantity: String(task.params?.quantity),
            use_incremental_size: task.params?.use_incremental_size ?? false,
            incremental_size: String(task.params?.incremental_size) ?? 0,
            delete_previous_orders: task.params?.delete_previous_orders ?? true,
        }
    }

    return {
        gap_percent: "0",
        quantity: "0",
        use_incremental_size: false,
        incremental_size: "0",
        delete_previous_orders: true,
    }
}


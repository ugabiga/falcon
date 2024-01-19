export enum TaskType {
    DCA = "dca",
    BuyingGrid = "buying_grid",
}
export function convertStringToTaskType(value: string): TaskType {
    return value as TaskType
}

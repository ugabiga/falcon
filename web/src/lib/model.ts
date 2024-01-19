export enum TaskType {
    DCA = "dca",
    LongGrid = "buying_grid",
}
export function convertStringToTaskType(value: string): TaskType {
    return value as TaskType
}

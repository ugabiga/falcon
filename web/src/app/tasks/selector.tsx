import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {TaskIndex} from "@/graph/generated/generated";
import {useDispatch} from "react-redux";
import {refreshTask} from "@/store/taskSlice";

export function TradingAccountSelector({taskIndex}: { taskIndex: TaskIndex }) {
    const dispatch = useDispatch()

    return (
        <Select defaultValue={taskIndex.selectedTradingAccount?.id}
                onValueChange={(value) => {
                    dispatch(refreshTask({
                        tradingAccountID: value,
                        refresh: true
                    }))
                }}
        >
            <SelectTrigger>
                <SelectValue placeholder="Select a Exchange"/>
            </SelectTrigger>
            <SelectContent>
                {
                    taskIndex.tradingAccounts?.map((tradingAccount) => {
                        return (
                            <SelectItem key={tradingAccount.id} value={tradingAccount.id}>
                                {tradingAccount.identifier}
                            </SelectItem>
                        )
                    })
                }
            </SelectContent>
        </Select>
    )
}
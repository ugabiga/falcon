import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {TaskIndex} from "@/graph/generated/generated";
import {useDispatch} from "react-redux";
import {refreshTask} from "@/store/taskSlice";
import {useTranslation} from "react-i18next";

export function TradingAccountSelector({taskIndex}: { taskIndex?: TaskIndex }) {
    const {t} = useTranslation()
    const dispatch = useDispatch()

    if (!taskIndex) {
        return null
    }

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
                <SelectValue placeholder={t('tasks.select_trading_account.placeholder')}/>
            </SelectTrigger>
            <SelectContent>
                {
                    taskIndex.tradingAccounts?.map((tradingAccount) => {
                        return (
                            <SelectItem key={tradingAccount.id} value={tradingAccount.id}>
                                {tradingAccount.name}
                            </SelectItem>
                        )
                    })
                }
            </SelectContent>
        </Select>
    )
}
import {V1TaskIndexResponse} from "@/api/model";
import {useTranslation} from "@/lib/i18n";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import React from "react";
import {useSendRefreshSignal} from "@/lib/use-refresh";
import {RefreshTarget} from "@/store/refresherSlice";
import {refreshTask} from "@/store/taskSlice";
import {useDispatch} from "react-redux";

export default function TradingAccountSelector({data}: { data?: V1TaskIndexResponse }) {
    const {t} = useTranslation()
    const dispatch = useDispatch()

    if (!data) {
        return null
    }

    return (
        <Select
            defaultValue={data?.selected_trading_account?.id}
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
                    data.trading_accounts?.map((account) => (
                        <SelectItem key={account.id} value={String(account.id)}>
                            {account.name}
                        </SelectItem>
                    ))
                }
            </SelectContent>
        </Select>
    )
}

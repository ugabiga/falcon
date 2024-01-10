import {TradingAccount} from "@/graph/generated/generated";
import {Label} from "@/components/ui/label";
import {Card, CardContent, CardTitle} from "@/components/ui/card";
import {camelize, trim} from "@/lib/str";
import {useTranslation} from "react-i18next";
import {TradingAccountMoreBtn} from "@/app/tradingaccounts/more-btn";

export function TradingAccountCards({tradingAccounts}: { tradingAccounts?: TradingAccount[] }) {
    const {t} = useTranslation();

    return (
        <div className="block md:hidden space-y-2">
            {
                !tradingAccounts || tradingAccounts?.length === 0
                    ? <div className="font-medium text-center">
                        {t("trading_account.table.empty")}
                    </div>
                    : tradingAccounts?.map((tradingAccount) => (
                            <Card key={tradingAccount.id}>
                                <div className="grid grid-cols-2 gap-6">
                                    <div className="mt-4 mb-6 ml-6">
                                        <CardTitle>
                                            {tradingAccount.name}
                                        </CardTitle>
                                    </div>
                                    <div className="flex space-x-2 mt-2 mr-2">
                                        <div className="flex-grow"></div>
                                        <TradingAccountMoreBtn tradingAccount={tradingAccount}/>
                                    </div>
                                </div>
                                <CardContent className="grid grid-cols-2 gap-6">
                                    <Label>
                                        {t("trading_account.table.exchange")} : {camelize(tradingAccount.exchange)}
                                    </Label>
                                    <Label>
                                        {t("trading_account.table.key")} : {trim(tradingAccount.key, 4)}
                                    </Label>
                                    <Label>
                                        {t("trading_account.table.ip")} : {tradingAccount.ip}
                                    </Label>
                                </CardContent>
                            </Card>
                        )
                    )
            }
        </div>
    )

}

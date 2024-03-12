import {TradingAccount} from "@/graph/generated/generated";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {trim} from "@/lib/str";
import {useTranslation} from "react-i18next";
import {TradingAccountMoreBtn} from "@/components/tradingaccounts/more-btn";


export function TradingAccountTable({tradingAccounts}: { tradingAccounts?: TradingAccount[] }) {
    const {t} = useTranslation();

    return (
        <div className="hidden md:block">
            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead>{t("trading_account.table.name")}</TableHead>
                        <TableHead>{t("trading_account.table.exchange")}</TableHead>
                        <TableHead>{t("trading_account.table.key")}</TableHead>
                        <TableHead>{t("trading_account.table.action")}</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {
                        !tradingAccounts || tradingAccounts?.length === 0
                            ? <TableRow>
                                <TableCell colSpan={6} className="font-medium text-center">
                                    {t("trading_account.table.empty")}
                                </TableCell>
                            </TableRow>
                            : tradingAccounts?.map((tradingAccount) => (
                                <TableRow key={tradingAccount.id}>
                                    <TableCell className="font-medium">
                                        {tradingAccount.name}
                                    </TableCell>
                                    <TableCell>
                                        {t(`common.exchange.${tradingAccount.exchange}`)}
                                    </TableCell>
                                    <TableCell>{trim(tradingAccount.key, 4)}</TableCell>
                                    <TableCell>
                                        <TradingAccountMoreBtn tradingAccount={tradingAccount}/>
                                    </TableCell>
                                </TableRow>
                            ))
                    }
                </TableBody>
            </Table>
        </div>
    )
}

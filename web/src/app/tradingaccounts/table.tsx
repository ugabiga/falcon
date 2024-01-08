import {TradingAccount} from "@/graph/generated/generated";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {EditTradingAccount} from "@/app/tradingaccounts/edit";
import {camelize} from "@/lib/str";
import {useTranslation} from "react-i18next";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {MoreHorizontal} from "lucide-react";
import {Button} from "@/components/ui/button";
import {DeleteTradingAccount} from "@/app/tradingaccounts/delete";

function trim(str: string, length: number) {
    if (str.length <= length) {
        return str;
    }
    return str.trim().slice(0, length) + '...';
}

export function TradingAccountTable({tradingAccounts}: { tradingAccounts?: TradingAccount[] }) {
    const {t} = useTranslation();

    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">{t("trading_account.table.name")}</TableHead>
                    <TableHead>{t("trading_account.table.exchange")}</TableHead>
                    <TableHead>{t("trading_account.table.key")}</TableHead>
                    <TableHead>{t("trading_account.table.ip")}</TableHead>
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
                                <TableCell className="font-medium">{tradingAccount.name}</TableCell>
                                <TableCell>{camelize(tradingAccount.exchange)}</TableCell>
                                <TableCell>{trim(tradingAccount.key, 4)}</TableCell>
                                <TableCell>{tradingAccount.ip}</TableCell>
                                <TableCell>
                                    <DropdownMenu>
                                        <DropdownMenuTrigger asChild>
                                            <Button variant="ghost" className="h-8 w-8 p-0">
                                                <MoreHorizontal className={"h-4 w-4"}/>
                                            </Button>
                                        </DropdownMenuTrigger>
                                        <DropdownMenuContent align="end">
                                            <EditTradingAccount tradingAccount={tradingAccount}/>
                                            <DeleteTradingAccount tradingAccount={tradingAccount}/>
                                        </DropdownMenuContent>
                                    </DropdownMenu>
                                </TableCell>
                            </TableRow>
                        ))
                }
            </TableBody>
        </Table>
    )
}

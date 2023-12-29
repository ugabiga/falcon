import {TradingAccount} from "@/graph/generated/generated";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {EditTradingAccount} from "@/app/tradingaccounts/edit";
import {camelize} from "@/lib/str";

function trim(str: string, length: number) {
    if (str.length <= length) {
        return str;
    }
    return str.trim().slice(0, length) + '...';
}

export function TradingAccountTable({tradingAccounts}: { tradingAccounts?: TradingAccount[] }) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">Name</TableHead>
                    <TableHead>Exchange</TableHead>
                    <TableHead>Key</TableHead>
                    <TableHead>IP</TableHead>
                    <TableHead>Action</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    !tradingAccounts || tradingAccounts?.length === 0
                        ? <NoData/>
                        : tradingAccounts?.map((tradingAccount) => (
                            <TableRow key={tradingAccount.id}>
                                <TableCell className="font-medium">{tradingAccount.name}</TableCell>
                                <TableCell>{camelize(tradingAccount.exchange)}</TableCell>
                                <TableCell>{trim(tradingAccount.key, 4)}</TableCell>
                                <TableCell>{tradingAccount.ip}</TableCell>
                                <TableCell>
                                    <EditTradingAccount tradingAccount={tradingAccount}/>
                                </TableCell>
                            </TableRow>
                        ))
                }
            </TableBody>
        </Table>
    )
}

function NoData() {
    return (
        <TableRow>
            <TableCell colSpan={6} className="font-medium text-center">
                No Trading Account found.
            </TableCell>
        </TableRow>
    )
}

import {TradingAccount} from "@/graph/generated/generated";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table";
import {EditTradingAccount} from "@/app/tradingaccounts/edit";
import {camelize} from "@/lib/str";

export function TradingAccountTable({tradingAccounts}: { tradingAccounts?: TradingAccount[] }) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">Name</TableHead>
                    <TableHead>Exchange</TableHead>
                    <TableHead>Currency</TableHead>
                    <TableHead>Identifier</TableHead>
                    <TableHead>IP</TableHead>
                    <TableHead>Action</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {
                    tradingAccounts?.length === 0 ? (
                            <TableRow>
                                <TableCell colSpan={6} className="font-medium text-center">
                                    No Trading Account found.
                                </TableCell>
                            </TableRow>
                        )
                        : tradingAccounts?.map((tradingAccount) => (
                            <TableRow key={tradingAccount.id}>
                                <TableCell className="font-medium">{tradingAccount.name}</TableCell>
                                <TableCell>{camelize(tradingAccount.exchange)}</TableCell>
                                <TableCell>{tradingAccount.currency}</TableCell>
                                <TableCell>{tradingAccount.identifier}</TableCell>
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

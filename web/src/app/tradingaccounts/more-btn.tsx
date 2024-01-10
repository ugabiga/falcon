import {TradingAccount} from "@/graph/generated/generated";
import {DropdownMenu, DropdownMenuContent, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {MoreHorizontal} from "lucide-react";
import {EditTradingAccount} from "@/app/tradingaccounts/edit";
import {DeleteTradingAccount} from "@/app/tradingaccounts/delete";


export function TradingAccountMoreBtn({tradingAccount}: { tradingAccount: TradingAccount }) {
    return (
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
    )
}

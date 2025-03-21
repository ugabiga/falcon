import {Task, TradingAccount} from "@/graph/generated/generated";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {MoreHorizontal} from "lucide-react";
import {EditTask} from "@/components/tasks/old/edit";
import {DeleteTask} from "@/components/tasks/old/delete";
import Link from "next/link";
import {useTranslation} from "react-i18next";


export function TaskMoreBtn(
    {
        task,
        tradingAccount
    }: {
        task: Task
        tradingAccount: TradingAccount
    }
) {
    const {t} = useTranslation();

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="ghost" className="h-8 w-8 p-0">
                    <MoreHorizontal className={"h-4 w-4"}/>
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
                <EditTask task={task} tradingAccount={tradingAccount}/>
                <DropdownMenuItem>
                    <Link
                        href={`/tasks/${task.id}/history?trading_account_id=${task.tradingAccountID}`}
                        legacyBehavior>
                        {t("tasks.table.history")}
                    </Link>
                </DropdownMenuItem>
                <DeleteTask task={task}/>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

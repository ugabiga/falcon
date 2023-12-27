import {usePathname} from "next/navigation";
import Link from "next/link";
import {cn} from "@/lib/utils";
import {NavigationRightMenu} from "@/components/navigation-right-menu";
import {Avatar, AvatarImage} from "@/components/ui/avatar";

export function NavigationMainBar() {
    const pathname = usePathname()

    return (
        <div className="hidden md:flex md:max-w-[1200px] w-full mx-auto flex justify-between">
            <nav className="pt-4 pb-4 flex items-center gap-6 text-sm w-full">
                <Link
                    href="/"
                    className={cn(
                        "transition-colors hover:text-foreground/80",
                        pathname === "/" ? "text-foreground" : "text-foreground/60"
                    )}
                >
                    <div className="flex items-center">
                        <Avatar>
                            <AvatarImage src="/logo1.png" alt="avatar"/>
                        </Avatar>
                        <span className="ml-2">Falcon</span>
                    </div>
                </Link>
                <Link
                    href="/tradingaccounts"
                    className={cn(
                        "transition-colors hover:text-foreground/80",
                        pathname === "/tradingaccounts" ? "text-foreground" : "text-foreground/60"
                    )}
                >
                    Trading Accounts
                </Link>
                <Link
                    href="/tasks"
                    className={cn(
                        "transition-colors hover:text-foreground/80",
                        pathname === "/tasks" ? "text-foreground" : "text-foreground/60"
                    )}
                >
                    Tasks
                </Link>
                <div className={"flex-grow"}></div>
                <NavigationRightMenu/>
            </nav>
        </div>
    )
}

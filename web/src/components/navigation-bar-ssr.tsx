import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {Avatar, AvatarFallback} from "@/components/ui/avatar";
import Link from "next/link";
import {getServerSession} from "next-auth";
import {LanguagesIcon} from "lucide-react";
import {icon} from "@/components/styles";
import {DarkModeButton} from "@/components/navigation-bar-darkmode";
import {useTranslation} from "@/lib/i18n-server";

export async function NavigationBarSsr() {

    return (
        <div className={"flex"}>
            <div>
                <DarkModeButton/>
                <LanguageMenu/>
            </div>
            <SessionMenu/>
        </div>
    )
}

export function LanguageMenu() {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="ghost">
                    <LanguagesIcon className={icon()}/>
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
                <DropdownMenuItem>
                    <Link href={"/language/en"}>
                        English
                    </Link>
                </DropdownMenuItem>
                <DropdownMenuItem>
                    <Link href={"/language/ko"}>
                        한국어 (Korean)
                    </Link>
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

async function SessionMenu() {
    const session = await getServerSession()
    const {t} = await useTranslation()

    if (session) {
        return (
            <div className={"ml-2"}>
                <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                        <Button className="rounded-full" size="icon" variant="ghost">
                            <Avatar className="h-9 w-9">
                                <AvatarFallback>{convertToInitials(session.user?.name ?? "")}</AvatarFallback>
                            </Avatar>
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                        <DropdownMenuItem>
                            <Link href={"/users"}>
                                {t("nav.profile")}
                            </Link>
                        </DropdownMenuItem>
                        <DropdownMenuItem>
                            <Link href={"/api/auth/signout"}>
                                {t("nav.sign_out")}
                            </Link>
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
        )
    } else {
        return (
            <Link href={"/auth/signin"}>
                Login
            </Link>
        )
    }
}

function convertToInitials(name: string) {
    const [first, last] = name.split(" ")

    let initial = "User"
    try {
        if (!first) {
            return initial
        }

        if (!last) {
            return first[0]
        }

        if (first.length > 1 && last.length > 1) {
            return first[0] + last[0]
        }

        if (first.length > 1) {
            return first[0]
        }
    } catch (e) {
        console.log("Error while converting name to initials", e)
        return initial
    }

    return initial
}

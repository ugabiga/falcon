import {useTheme} from "next-themes";
import {LanguagesIcon, Moon, Sun} from "lucide-react";
import {useEffect, useState} from "react";
import {icon} from "@/components/styles";
import {signIn, signOut, useSession} from "next-auth/react";
import {Avatar, AvatarFallback} from "@/components/ui/avatar";
import {Button} from "@/components/ui/button";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import {useRouter} from "next/navigation";
import {useTranslation} from "react-i18next";
import {useChangeI18nLanguage} from "@/lib/i18n-client";
import Link from "next/link";


export function NavigationRightMenu() {
    const [isClient, setIsClient] = useState(false)

    // Fix : Content does not match server-rendered HTML
    useEffect(() => {
        setIsClient(true)
    }, [])

    return (
        <div className={"flex"}>
            {
                isClient
                && <div>
                    <LanguageMenu/>
                    <DarkModeButton/>
                </div>
            }
            <SessionMenu/>
        </div>
    )
}

function DarkModeButton() {
    const {theme, setTheme} = useTheme()

    return (
        <Button variant="ghost" onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
            {theme === "dark" ? <Sun className={icon()}/> : <Moon className={icon()}/>}
        </Button>
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

// function LanguageMenu() {
//     const {changeLanguage} = useChangeI18nLanguage()
//
//     return (
//         <DropdownMenu>
//             <DropdownMenuTrigger asChild>
//                 <Button variant="ghost">
//                     <LanguagesIcon className={icon()}/>
//                 </Button>
//             </DropdownMenuTrigger>
//             <DropdownMenuContent>
//                 <DropdownMenuItem onClick={() => changeLanguage("en")}>
//                     English
//                 </DropdownMenuItem>
//                 <DropdownMenuItem onClick={() => changeLanguage("ko")}>
//                     한국어 (Korean)
//                 </DropdownMenuItem>
//             </DropdownMenuContent>
//         </DropdownMenu>
//
//     )
// }

function SessionMenu() {
    const {t} = useTranslation()
    const {data: session} = useSession();
    const router = useRouter()

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
                        <DropdownMenuItem onClick={
                            () => {
                                router.push("/users")
                            }
                        }>
                            {t("nav.profile")}
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => signOut({redirect: false}).then(() => {
                            router.push("/")
                        })}>
                            {t("nav.sign_out")}
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
        )
    } else {
        return (
            <Button variant="ghost" onClick={() => signIn()}>
                {t("nav.sign_in")}
            </Button>
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

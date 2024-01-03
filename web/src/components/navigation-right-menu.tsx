import {useTheme} from "next-themes";
import {Moon, Sun} from "lucide-react";
import {useEffect, useState} from "react";
import {icon} from "@/components/styles";
import {signIn, signOut, useSession} from "next-auth/react";
import {Avatar, AvatarFallback} from "@/components/ui/avatar";
import {Button} from "@/components/ui/button";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import {useRouter} from "next/navigation";
import {useTranslation} from "react-i18next";


export function NavigationRightMenu() {
    const {theme, setTheme} = useTheme()
    const [isClient, setIsClient] = useState(false)

    // Fix : Content does not match server-rendered HTML
    useEffect(() => {
        setIsClient(true)
    }, [])

    return (
        <div className={"flex"}>
            {isClient
                && <Button variant="ghost" onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
                    {theme === "dark" ? <Sun className={icon()}/> : <Moon className={icon()}/>}
                </Button>
            }
            <SessionMenu/>
        </div>
    )
}

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
        initial = first[0] + last[0]
    } catch (e) {
        try {
            initial = first[0]
        } catch (e) {
            return initial
        }
    }

    return initial
}

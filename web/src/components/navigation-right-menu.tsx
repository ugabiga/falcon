import {useTheme} from "next-themes";
import {Moon, Sun} from "lucide-react";
import {useEffect, useState} from "react";
import {icon} from "@/components/styles";
import {signIn, signOut, useSession} from "next-auth/react";
import {Avatar, AvatarFallback} from "@/components/ui/avatar";
import {Button} from "@/components/ui/button";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";


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
    const {data: session} = useSession();
    // const session = true

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
                                // @ts-ignore
                                window.location.href = "/users"
                            }
                        }>
                            Profile
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => signOut()}>
                            Logout
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
        )
    } else {
        return (
            <Button variant="ghost" onClick={() => signIn()}>Login</Button>
        )
    }

}

function convertToInitials(name: string) {
    const [first, last] = name.split(" ")
    return `${first[0]}${last[0]}`
}

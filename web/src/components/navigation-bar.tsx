"use client"


import {useTheme} from "next-themes";
import {
    NavigationMenu,
    NavigationMenuLink,
    NavigationMenuList,
    navigationMenuTriggerStyle
} from "@/components/ui/navigation-menu";
import Link from "next/link";
import {Toggle} from "@/components/ui/toggle";
import {Moon, Sun} from "lucide-react";
import {User} from "@/graph/generated/generated";
import React, {useState} from "react";
import {Label} from "@/components/ui/label";
import {Button} from "@/components/ui/button";
import {useAppSelector} from "@/store";
import {Spacer} from "@/components/ui/Spacer";

export function NavigationBar() {
    const {theme, setTheme} = useTheme()
    const [isClient, setIsClient] = useState(false)
    const user = useAppSelector(state => state.user)


    return (
        <div className="md:max-w-[1200px] overflow-auto w-full mx-auto flex justify-between">
            <NavigationMenu>
                <NavigationMenuList>
                    <Link href={"/"} legacyBehavior passHref>
                        <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                            Home
                        </NavigationMenuLink>
                    </Link>
                </NavigationMenuList>
            </NavigationMenu>

            <Spacer/>

            <div className={"flex items-center"}>
                {
                    user.isLogged
                        ? <Button>Logout</Button>
                        : <Button>Login</Button>
                }
                {/*<Label>Hi, {user.name}</Label>*/}
                {/*{*/}
                {/*    user.isLogged*/}
                {/*        ? <Label>{user.name}</Label>*/}
                {/*        : <Button>Login</Button>*/}
                {/*}*/}
            </div>

            <Toggle onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
                {
                    isClient && theme === "dark"
                        ? <Sun className={"h-[1.2rem] w-[1.2rem transition-all"}/>
                        : <Moon className={"h-[1.2rem] w-[1.2rem transition-all"}/>
                }
            </Toggle>
        </div>
    )
}
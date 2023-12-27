"use client";

import {useTheme} from "next-themes";
import {cn} from "@/lib/utils"
import {
    NavigationMenu,
    NavigationMenuItem,
    NavigationMenuLink,
    NavigationMenuList,
    navigationMenuTriggerStyle
} from "@/components/ui/navigation-menu";
import Link from "next/link";
import {Toggle} from "@/components/ui/toggle";
import {Moon, Sun} from "lucide-react";
import React, {useEffect, useState} from "react";
import {icon} from "@/components/styles";
import {signIn, signOut, useSession} from "next-auth/react";
import {Avatar, AvatarFallback} from "@/components/ui/avatar";
import {Button} from "@/components/ui/button";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";


export function NavigationBar() {
    const {theme, setTheme} = useTheme()
    const [isClient, setIsClient] = useState(false)

    // Fix : Content does not match server-rendered HTML
    useEffect(() => {
        setIsClient(true)
    }, [])

    return (
        <div className="mt-2 mb-2 md:max-w-[1200px] overflow-auto w-full mx-auto flex justify-between">
            <NavigationMenu>
                <NavigationMenuList>
                    <NavigationMenuItem>
                        <Link href="/" legacyBehavior passHref>
                            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                                Home
                            </NavigationMenuLink>
                        </Link>
                    </NavigationMenuItem>
                    <NavigationMenuItem>
                        <Link href="/tradingaccounts" legacyBehavior passHref>
                            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                                Trading Accounts
                            </NavigationMenuLink>
                        </Link>
                    </NavigationMenuItem>
                    <NavigationMenuItem>
                        <Link href="/tasks" legacyBehavior passHref>
                            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                                Tasks
                            </NavigationMenuLink>
                        </Link>
                    </NavigationMenuItem>
                </NavigationMenuList>
            </NavigationMenu>
            <div className="flex items-center">
                {
                    isClient
                    && <Toggle onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
                        {theme === "dark" ? <Sun className={icon()}/> : <Moon className={icon()}/>}
                    </Toggle>
                }
                <SessionMenu/>
            </div>
        </div>
    )
}

function convertToInitials(name: string) {
    const [first, last] = name.split(" ")
    return `${first[0]}${last[0]}`
}

function SessionMenu() {
    const {data: session} = useSession();
    // const session = true

    if (session) {
        return (
            <div className={"mr-2"}>
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

const ListItem = React.forwardRef<
    React.ElementRef<"a">,
    React.ComponentPropsWithoutRef<"a">
>(({className, title, children, ...props}, ref) => {
    return (
        <li>
            <NavigationMenuLink asChild>
                <a
                    ref={ref}
                    className={cn(
                        "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
                        className
                    )}
                    {...props}
                >
                    <div className="text-sm font-medium leading-none">{title}</div>
                    <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
                        {children}
                    </p>
                </a>
            </NavigationMenuLink>
        </li>
    )
})
ListItem.displayName = "ListItem"

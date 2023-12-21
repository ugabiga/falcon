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

export function NavigationBar() {
    const {theme, setTheme} = useTheme()

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
            <Toggle onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
                {
                    theme === "dark"
                        ? <Sun
                            className="h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"/>
                        : <Moon
                            className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"/>
                }
            </Toggle>
        </div>
    )
}
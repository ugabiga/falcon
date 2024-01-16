"use client";

import {Button} from "@/components/ui/button";
import {Sun, Moon, LanguagesIcon} from "lucide-react";
import {icon} from "@/components/styles";
import React, {useEffect, useState} from "react";
import {useTheme} from "next-themes";
import {useChangeI18nLanguage} from "@/lib/i18n-client";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";

export function DarkModeButton() {
    const {theme, setTheme} = useTheme()
    const [isClient, setIsClient] = useState(false)

    //Fix : Content does not match server-rendered HTML
    useEffect(() => {
        setIsClient(true)
    }, [])

    return (
        <>
            {
                isClient
                    ?
                        <Button variant="ghost" onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
                            {theme === "dark" ? <Sun className={icon()}/> : <Moon className={icon()}/>}
                        </Button>
                    : <div></div>
            }
        </>
    )
}

export function LanguageMenu() {
    const {changeLanguage} = useChangeI18nLanguage()

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="ghost">
                    <LanguagesIcon className={icon()}/>
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
                <DropdownMenuItem onClick={() => changeLanguage("en")}>
                    English
                </DropdownMenuItem>
                <DropdownMenuItem onClick={() => changeLanguage("ko")}>
                    한국어 (Korean)
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>

    )
}

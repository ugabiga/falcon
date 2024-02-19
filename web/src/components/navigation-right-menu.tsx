"use client";

import {LanguagesIcon, Moon, Sun} from "lucide-react";
import React, {useEffect, useState} from "react";
import {icon} from "@/components/styles";
import {signIn, signOut, useSession} from "next-auth/react";
import {Avatar, AvatarFallback} from "@/components/ui/avatar";
import {Button} from "@/components/ui/button";
import {DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from "@/components/ui/dropdown-menu";
import {useRouter} from "next/navigation";
import {useTranslation} from "react-i18next";
import Link from "next/link";
import {convertToInitials} from "@/lib/converter";
import {useTheme} from "next-themes";
import {resetPostHog, setPostHogUser} from "@/lib/posthog";


export function NavigationRightMenu() {
    return (
        <div className={"flex"}>
            <DarkModeButton/>
            <LanguageMenu/>
            <SessionMenu/>
        </div>
    )
}

export function DarkModeButton() {
    const {theme, setTheme} = useTheme()
    const [isClient, setIsClient] = useState(false)

    //Fix : Content does not match server-rendered HTML
    useEffect(() => {
        setIsClient(true)
    }, [])

    if (!isClient) {
        return (
            <div></div>
        )
    }

    return (
        <Button variant="ghost" onClick={() => setTheme(theme === "dark" ? "light" : "dark")}>
            {theme === "dark" ? <Sun className={icon()}/> : <Moon className={icon()}/>}
        </Button>
    )
}

export function LanguageMenu() {
    const [isClient, setIsClient] = useState(false)

    useEffect(() => {
        setIsClient(true)
    }, [])

    if (!isClient) {
        return (
            <div></div>
        )
    }

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

function SessionMenu() {
    const {t} = useTranslation()
    const {data: session} = useSession();
    const router = useRouter()

    if (!session) {
        return (
            <Button variant="ghost" onClick={() => signIn()}>
                {t("nav.sign_in")}
            </Button>
        )
    }

    setPostHogUser(session.user?.name)

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
                        resetPostHog()
                        router.push("/")
                    })}>
                        {t("nav.sign_out")}
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )

}


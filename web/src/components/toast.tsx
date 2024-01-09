"use client";


import {toast} from "sonner";
import {transformErrorMessage} from "@/lib/error";
import {useTranslation} from "react-i18next";
import {signOut} from "next-auth/react";

export function normalToast({message}: { message: string }) {
    toast(message, {
        position: "top-right",
        action: {
            label: "Close",
            onClick: () => {
            }
        },
        duration: 2000,
    })
}

export function errorToast(message: string) {
    toast.error(message, {
        position: "top-right",
        action: {
            label: "Close",
            onClick: () => {
            }
        },
        duration: 2000,
        onAutoClose: () => {
            switch (message) {
                case "Response not successful: Received status code 401":
                    return signOut({redirect: true}).then()
                default:
                    return
            }
        }
    })
}


"use client";


import {toast} from "sonner";
import {transformErrorMessage} from "@/lib/error";
import {useTranslation} from "react-i18next";

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
    })
}


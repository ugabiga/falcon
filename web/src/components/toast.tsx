"use client";


import {toast} from "sonner";
import {transformErrorMessage} from "@/lib/error";

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
    const errorMessage = transformErrorMessage(message)
    toast.error(errorMessage, {
        position: "top-right",
        action: {
            label: "Close",
            onClick: () => {
            }
        },
        duration: 2000,
    })
}


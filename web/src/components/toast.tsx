"use client";


import {toast} from "sonner";

export function customToast({message}: { message: string }) {
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

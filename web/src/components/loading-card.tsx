"use client"

import React, {useEffect, useState} from "react";
import {Icons} from "@/components/icons";

export function LoadingCard() {
    const [showLoading, setShowLoading] = useState(false);

    useEffect(() => {
        const timer = setTimeout(() => {
            setShowLoading(true);
        }, 300);

        return () => {
            clearTimeout(timer);
        };
    }, []);

    if (!showLoading) {
        return null
    }

    return (
        <div className="w-full flex flex-col justify-center items-center p-12">
            <Icons.spinner className="h-8 w-8 animate-spin"/>
        </div>
    )
}

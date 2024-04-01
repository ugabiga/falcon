"use client"

import {TableCell, TableRow} from "@/components/ui/table";
import React, {useEffect, useState} from "react";
import {Icons} from "@/components/icons";

export function LoadingTableCell({colSpan}: { colSpan: number }) {
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
        return (
            <TableRow>
            </TableRow>
        )
    }

    return (
        <TableRow>
            <TableCell colSpan={colSpan} className="text-center">
                <div className="w-full flex flex-col justify-center items-center">
                    <Icons.spinner className="h-8 w-8 animate-spin"/>
                </div>
            </TableCell>
        </TableRow>
    )
}

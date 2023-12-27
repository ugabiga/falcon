import {useEffect, useState} from 'react';
import {Icons} from "@/components/icons";

export function Loading() {
    const [showLoading, setShowLoading] = useState(false);

    useEffect(() => {
        const timer = setTimeout(() => {
            setShowLoading(true);
        }, 200);

        return () => {
            clearTimeout(timer);
        };
    }, []);

    if (!showLoading) {
        return null;
    }

    return (
        <div
            className="fixed top-0 h-screen w-full bg-opacity-70 backdrop-blur-sm flex flex-col justify-center items-center z-50">
            <Icons.spinner className="mr-2 h-4 w-4 animate-spin"/>
            Loading...
        </div>
    );
}
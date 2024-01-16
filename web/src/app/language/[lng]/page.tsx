"use client";

import {useCookies} from "react-cookie";
import {cookieName} from "@/lib/i18n";
import {useEffect} from "react";
import {Loading} from "@/components/loading";

export default function Language({params}: { params: { lng: string } }) {
    const [cookies, setCookie] = useCookies([cookieName])

    useEffect(() => {
        setCookie(cookieName, params.lng, {path: '/'})
        //Redirect to home page
        window.location.href = "/"
    }, [params])

    return (
        <div>
            <Loading/>
        </div>
    )
}
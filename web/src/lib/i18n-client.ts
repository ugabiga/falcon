import {useCookies} from "react-cookie";
import {useEffect, useState} from "react";
import i18n, {cookieName} from "@/lib/i18n";


export function useSetupI18n() {
    const [cookies, setCookie] = useCookies([cookieName])
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        if (!cookies.i18next) {
            // setCookie(cookieName, i18n.language, {path: '/'})
            setLoading(false)
            return
        }
        i18n.changeLanguage(cookies.i18next).then()
        setLoading(false)
    }, [cookies.i18next]);

    return {loading}
}

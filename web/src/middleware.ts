import {NextRequest, NextResponse} from 'next/server'
import acceptLanguage from 'accept-language'
import {cookieName, fallbackLng, languages} from "@/lib/i18n";

acceptLanguage.languages(languages)

export function middleware(req: NextRequest) {
    if (req.nextUrl.pathname.indexOf('icon') > -1 || req.nextUrl.pathname.indexOf('chrome') > -1) return NextResponse.next()

    let lng: string | undefined | null
    if (req.cookies.has(cookieName)) lng = acceptLanguage.get(req.cookies.get(cookieName)?.value)
    if (!lng) lng = acceptLanguage.get(req.headers.get('Accept-Language'))
    if (!lng) lng = fallbackLng

    const res = NextResponse.next();
    res.cookies.set(cookieName, lng)

    return res;
}

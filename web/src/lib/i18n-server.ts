import {createInstance} from 'i18next'
import {cookies, headers} from 'next/headers'
import acceptLanguage from 'accept-language'
import usTranslations from "@/translations/en_US.json";
import krTranslations from "@/translations/ko_KR.json";

export const fallbackLng = 'ko'
export const languages = [fallbackLng, 'en']
export const cookieName = 'i18next'

export const resources = {
    en: {
        translation: usTranslations,
    },
    ko: {
        translation: krTranslations,
    },
};

const initI18next = async (lng: string) => {
    // on server side we create a new instance for each render, because during compilation everything seems to be executed in parallel
    const i18nInstance = createInstance()
    await i18nInstance
        .init({
            resources,
            lng: lng,
            fallbackLng: fallbackLng,
            debug: false,
            interpolation: {escapeValue: true},
            returnObjects: true,
            returnEmptyString: true,
            returnNull: true,
        })
    return i18nInstance
}

acceptLanguage.languages(languages)

export function detectLanguage() {
    const reqCookies = cookies()
    const readOnlyHeaders = headers()
    let lng
    const nextUrlHeader = readOnlyHeaders.has('next-url') && readOnlyHeaders.get('next-url')
    if (nextUrlHeader && nextUrlHeader.indexOf(`"lng":"`) > -1) {
        const qsObj = JSON.parse(nextUrlHeader.substring(nextUrlHeader.indexOf('{'), nextUrlHeader.indexOf(`}`) + 1))
        lng = qsObj.lng
    }
    if (!lng && reqCookies.has(cookieName)) lng = acceptLanguage.get(reqCookies.get(cookieName)!.value)
    if (!lng) lng = acceptLanguage.get(readOnlyHeaders.get('Accept-Language'))
    if (!lng) lng = fallbackLng
    return lng
}

export async function useTranslation() {
    const lng = detectLanguage()
    const i18nextInstance = await initI18next(lng)
    return {
        t: i18nextInstance.t,
        i18n: i18nextInstance
    }
}

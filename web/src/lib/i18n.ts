import i18n, {FlatNamespace, KeyPrefix} from "i18next";
import {FallbackNs, initReactI18next, useTranslation as useTranslationOrg} from "react-i18next";
import usTranslations from '@/translations/en_US.json';
import krTranslations from '@/translations/ko_KR.json';

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


i18n.use(initReactI18next)
    .init({
        resources,
        lng: fallbackLng,
        fallbackLng: fallbackLng,
        debug: false,
        interpolation: {escapeValue: true},
        returnObjects: true,
        returnEmptyString: true,
        returnNull: true,
    }).then()


export function useTranslation<
    Ns extends FlatNamespace,
    KPrefix extends KeyPrefix<FallbackNs<Ns>> = undefined
>(
    ns?: Ns,
    options: { keyPrefix?: KPrefix } = {}
) {
    return useTranslationOrg(ns, options)
}

export default i18n;
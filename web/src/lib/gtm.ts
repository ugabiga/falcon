type WindowWithDataLayer = Window & {
    dataLayer: Record<string, any>[]
}

declare const window: WindowWithDataLayer

export const GTM_ID = process.env.NEXT_PUBLIC_GTM

export const pageView = (url: string) => {
    if (typeof window.dataLayer !== "undefined") {
        window.dataLayer.push({
            event: "pageview",
            page: url,
        })
    } else {
        console.log({
            event: "pageview",
            page: url,
        })
    }
}
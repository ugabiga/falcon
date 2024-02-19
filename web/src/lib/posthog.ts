import posthog from "posthog-js";

export function capture(eventName: string, properties: any = undefined) {
    console.log('PostHog event', eventName, properties);
    posthog.capture(eventName, properties);
}

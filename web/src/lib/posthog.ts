import posthog from "posthog-js";

export function capture(eventName: string, properties: any = undefined) {
    console.log('PostHog event', eventName, properties);
    posthog.capture(eventName, properties);
}

export function setPostHogUser(id: string, name: string | null | undefined) {
    posthog.identify(`user_${id}`, {
        name: name,
    });
}

export function resetPostHog() {
    posthog.capture('logout');
    posthog.reset();
}
import {TFunction} from "i18next";

export function transformErrorMessage(message?: string) {
    if (!message) {
        return "An unknown error occurred."
    }

    if (message === "Failed to fetch") {
        return "Failed to connect to the server. Please try again later."
    }

    if (message.includes("ent:")) {
        return message.split("ent:")[1]
    }

    if (message === "exceed_limit") {
        return "You have reached the maximum number of items you can create."
    }

    return message;
}

export function translatedError(t: TFunction<"translation", undefined>, message: string) {
    if (message.includes("size_not_satisfied_minimum_size")) {
        const size = message.split("#")[1]

        return t("error.size_not_satisfied_minimum_size", {size})
    }

    return t("error." + message)
}

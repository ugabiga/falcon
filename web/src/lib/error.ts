export function transformErrorMessage(message: string) {
    if (message === "Failed to fetch") {
        return "Failed to connect to the server. Please try again later."
    }

    if (message.includes("ent:")) {
        return message.split("ent:")[1]
    }

    return message;
}

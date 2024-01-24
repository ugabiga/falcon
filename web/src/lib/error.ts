export function transformErrorMessage(message: string) {

    if (message === "Failed to fetch") {
        return "Failed to connect to the server. Please try again later."
    }

    if (message.includes("ent:")) {
        return message.split("ent:")[1]
    }

    if (message === "exceed_limit") {
        return "You have reached the maximum number of items you can create."
    }

    if (message.includes("error.size_not_satisfied_minimum_size")) {
        const size = message.split("#")[1]
            .replace("-", " ")

        return `The minimum size is ${size}`
    }

    return message;
}

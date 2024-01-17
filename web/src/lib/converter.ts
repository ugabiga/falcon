function convertBooleanToYesNo(value: boolean) {
    return value ? "Yes" : "No"
}


function convertToInitials(name: string) {
    const [first, last] = name.split(" ")

    let initial = "User"
    try {
        if (!first) {
            return initial
        }

        if (!last) {
            return first[0]
        }

        if (first.length > 1 && last.length > 1) {
            return first[0] + last[0]
        }

        if (first.length > 1) {
            return first[0]
        }
    } catch (e) {
        console.log("Error while converting name to initials", e)
        return initial
    }

    return initial
}

export {
    convertBooleanToYesNo,
    convertToInitials
}

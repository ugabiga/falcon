export function Error({message}: { message: string }) {
    const errorMessage = transformErrorMessage(message)
    return (
        <div className="h-screen w-full flex flex-col justify-center items-center space-y-4">
            <h1 className="text-3xl font-bold text-red-500">Error</h1>
            <h2 className="text-2xl font-bold text-red-500">{errorMessage}</h2>
        </div>
    )
}

function transformErrorMessage(message: string) {
    if (message === "Failed to fetch") {
        return "Failed to connect to the server. Please try again later."
    }
    return message;
}
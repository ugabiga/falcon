import {transformErrorMessage} from "@/lib/error";
import {Button} from "@/components/ui/button";
import {signIn} from "next-auth/react";

export function Error({message}: { message: string }) {

    switch (message) {
        case "Response not successful: Received status code 401":
            return UnAuthorizedError()
    }

    const errorMessage = transformErrorMessage(message)
    return (
        <div className="h-screen w-full flex flex-col justify-center items-center space-y-4">
            <h1 className="text-3xl font-bold text-red-500">Error</h1>
            <h2 className="text-2xl font-bold text-red-500">{errorMessage}</h2>
        </div>
    )
}


function UnAuthorizedError() {
    return (
        <div className="h-screen w-full flex flex-col justify-center items-center space-y-4">
            <h1 className="text-3xl font-bold"> You are not authorized to access this page</h1>
            <h2 className="text-2xl font-bold"> Please sign in to continue</h2>
            <Button onClick={() => signIn()}>
                Sign in
            </Button>
        </div>
    )
}

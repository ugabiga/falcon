import {transformErrorMessage} from "@/lib/error";
import {Button} from "@/components/ui/button";
import {signIn, signOut} from "next-auth/react";
import {useTranslation} from "react-i18next";
import {ErrorType} from "@/api/mutator/custom-instance";
import {HandlerAPIError} from "@/api/model";

export function Error({error}: { error: ErrorType<HandlerAPIError> }) {
    if (!error.response) {
        return DefaultError()
    }

    switch (error.response.data.message) {
        case "Unauthorized":
            return UnAuthorizedError()
        default:
            return DefaultError(error.response.data.message)
    }

}

function DefaultError(message?: string) {
    const errorMessage = transformErrorMessage(message)
    return (
        <div className="h-screen w-full flex flex-col justify-center items-center space-y-4">
            <h1 className="text-3xl font-bold text-red-500">Error</h1>
            <h2 className="text-2xl font-bold text-red-500">{errorMessage}</h2>
        </div>
    )
}

function UnAuthorizedError() {
    const {t} = useTranslation()

    signOut({redirect: false}).then()

    return (
        <div className="h-screen w-full flex flex-col justify-center items-center space-y-4">
            <h2 className="text-2xl font-bold"> {t("error.unauthorized.action")}</h2>
            <Button onClick={() => signIn()}>
                {t("error.unauthorized.action.sign_in")}
            </Button>
        </div>
    )
}

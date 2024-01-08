import {transformErrorMessage} from "@/lib/error";
import {Button} from "@/components/ui/button";
import {signIn, signOut} from "next-auth/react";
import {useTranslation} from "react-i18next";

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
    signOut({redirect: false}).then()

    const {t} = useTranslation()
    return (
        <div className="h-screen w-full flex flex-col justify-center items-center space-y-4">
            <h1 className="text-3xl font-bold"> {t("error.unauthorized")}</h1>
            <h2 className="text-2xl font-bold"> {t("error.unauthorized.action")}</h2>
            <Button onClick={() => signIn()}>
                {t("error.unauthorized.action.sign_in")}
            </Button>
        </div>
    )
}

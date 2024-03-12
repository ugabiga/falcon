"use client"

import {
    useGetApiV1UsersMe,
    useGetAuthProtected,
    useGetAuthProtectedSuspense,
    usePutApiV1UsersMe
} from "@/api/endpoints/transformer";
import {Button} from "@/components/ui/button";
import {Loading} from "@/components/loading";
import {Input} from "@/components/ui/input";
import {useState} from "react";
import {Error} from "@/components/error";

export default function Dev() {
    const [textInput, setTextInput] = useState<string>("")
    const {data, isLoading, error, refetch} = useGetApiV1UsersMe({
        query: {
            enabled: false,
        },
    })
    const {mutate: updateUserProfile} = usePutApiV1UsersMe({
        mutation: {
            onSuccess: (data) => {
                refetch().then()
                setTextInput("")
            }
        }
    })

    if (process.env.NODE_ENV === "production") {
        return null
    }

    if (isLoading) {
        return <Loading/>
    }

    if (error) {
        return <Error message={error.message}/>
    }

    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">

            <h1 className="text-3xl font-semibold mb-6">
                {process.env.NODE_ENV.charAt(0).toUpperCase() + process.env.NODE_ENV.slice(1)} Environment
            </h1>

            <div>
                <h2 className="text-xl font-semibold mb-2">Actions</h2>

                <div className="flex space-x-4 mb-6">

                    <Button
                        onClick={() => {
                            refetch().then()
                        }}
                    >
                        Fetch
                    </Button>

                    <Input
                        placeholder="Text Input"
                        onChange={(e) => {
                            setTextInput(e.target.value)
                        }}
                    />

                    <Button
                        onClick={() => {
                            updateUserProfile({
                                data: {
                                    name: textInput,
                                },
                            })
                        }}
                    >
                        Mutate
                    </Button>
                </div>
            </div>

            <div>
                <h2 className="text-xl font-semibold mb-2">Data</h2>

                <div className="rounded-lg bg-zinc-900 p-4">
                <pre>{
                    JSON.stringify(data, null, 2) || "null"
                }</pre>
                </div>
            </div>
        </main>
    )
}

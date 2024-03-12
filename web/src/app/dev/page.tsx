"use client"

import {useGetAuthProtected, useGetAuthProtectedSuspense} from "@/api/endpoints/transformer";
import {Button} from "@/components/ui/button";

export default function Dev() {
    const {data, isLoading, error, refetch} = useGetAuthProtected({
        query: {
            enabled: false,
        },
    })

    // If in production, return null
    if (process.env.NODE_ENV === "production") {
        return null
    }

    if (isLoading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Error: {error.message}</div>
    }

    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">

            <h1 className="text-3xl font-semibold mb-4">
                {/* Capitalize the first letter */}
                {process.env.NODE_ENV.charAt(0).toUpperCase() + process.env.NODE_ENV.slice(1)} Environment
            </h1>

            <div className="mb-4">
                <pre>{JSON.stringify(data, null, 2)}</pre>
            </div>

            {/*   Click button and show response in json format */}
            <Button
                onClick={() => {
                    refetch().then()
                }}
            >
                Click me
            </Button>

        </main>
    )
}

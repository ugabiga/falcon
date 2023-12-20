"use client";

import Image from 'next/image'
import {useQuery} from "@apollo/client";
import {GetUserDocument} from "@/graph/generated/generated";

export default function Home() {
    const {data, loading} = useQuery(GetUserDocument, {
        variables: {
            id: "1"
        }
    })

    if (loading) {
        return <div>Loading...</div>
    }

    // if (!data) {
    //     return <div>No data</div>
    // }

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <h1 className="text-6xl font-bold">Welcome to {data?.user.name}</h1>
        </main>
    )
}

"use client";

import {signIn} from "next-auth/react";
import {Button} from "@/components/ui/button";

export default function Home() {
    // const {data, loading} = useQuery(GetUserDocument, {
    //     variables: {
    //         id: "1"
    //     }
    // })
    //
    // if (loading) {
    //     return <div>Loading...</div>
    // }

    // if (!data) {
    //     return <div>No data</div>
    // }

    const handleProtected = () => {
        fetch("http://localhost:8080/auth/protected", {
            credentials: "include"
        })
            .then(res => res.json())
            .then(data => console.log(data))
            .catch(err => console.log(err))
    }

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <h1 className="text-6xl font-bold">Welcome </h1>
            <Button variant={"ghost"} onClick={() => signIn()}>Sign in</Button>
            <Button variant={"ghost"} onClick={() => handleProtected()}>Protected</Button>
        </main>
    )
}

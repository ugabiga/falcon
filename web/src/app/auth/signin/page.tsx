"use client";

import {Button} from "@/components/ui/button";
import {signIn} from "next-auth/react";

export default function SignIn() {
    return (
        <main className="min-h-screen p-12 ">
            <Button variant="ghost" onClick={() => signIn()}>
                sign in
            </Button>
        </main>
    )
}
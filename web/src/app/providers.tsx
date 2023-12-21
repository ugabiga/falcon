'use client';

import React from "react";
import {ApolloProvider} from "@apollo/client";
import {client} from "@/graph/client";
import {ThemeProvider} from "next-themes";

export default function Providers({children}: { children: React.ReactNode }) {
    return (
        <ThemeProvider attribute="class" defaultTheme={"system"} enableSystem>
            <ApolloProvider client={client}>
                {children}
            </ApolloProvider>
        </ThemeProvider>
    )
}
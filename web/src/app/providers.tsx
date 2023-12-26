'use client';

import React from "react";
import {ApolloProvider} from "@apollo/client";
import {client} from "@/graph/client";
import {ThemeProvider} from "next-themes";
import {Provider} from "react-redux";
import {store} from "@/store";

export default function Providers({children}: { children: React.ReactNode }) {
    return (
        <ThemeProvider attribute="class" defaultTheme={"system"} enableSystem>
            <ApolloProvider client={client}>
                <Provider store={store}>
                    {children}
                </Provider>
            </ApolloProvider>
        </ThemeProvider>
    )
}
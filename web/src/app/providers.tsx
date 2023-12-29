'use client';

import React from "react";
import {ApolloProvider} from "@apollo/client";
import {client} from "@/graph/client";
import {ThemeProvider} from "next-themes";
import {Provider} from "react-redux";
import {store} from "@/store";
import {getServerSession} from "next-auth";
import {I18nextProvider} from "react-i18next";
import i18n from "@/lib/i18n";

export default function Providers({children}: { children: React.ReactNode }) {

    return (
        <ThemeProvider attribute="class" defaultTheme={"system"} enableSystem>
            <I18nextProvider i18n={i18n}>
                <ApolloProvider client={client}>
                    <Provider store={store}>
                        {children}
                    </Provider>
                </ApolloProvider>
            </I18nextProvider>
        </ThemeProvider>
    )
}
'use client';

import React from "react";
import {ApolloProvider} from "@apollo/client";
import {client} from "@/graph/client";
import {ThemeProvider} from "next-themes";
import {Provider} from "react-redux";
import {persistor, store} from "@/store";
import {useSetupI18n} from "@/lib/i18n-client";
import {PersistGate} from "redux-persist/integration/react";
import {CSPostHogProvider} from "@/lib/posthog-provider";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

const queryClient = new QueryClient();
export default function Providers({children}: { children: React.ReactNode }) {
    const {loading} = useSetupI18n();

    return (
        <ThemeProvider attribute="class" defaultTheme={"system"} enableSystem>
            <ApolloProvider client={client}>
                <QueryClientProvider client={queryClient}>
                    <Provider store={store}>
                        <PersistGate loading={null} persistor={persistor}>
                            <CSPostHogProvider>
                                {children}
                            </CSPostHogProvider>
                        </PersistGate>
                    </Provider>
                </QueryClientProvider>
            </ApolloProvider>
        </ThemeProvider>
    )
}
import NextAuth from "next-auth";
import GoogleProvider from "next-auth/providers/google";
import {cookies} from "next/headers";

const googleClientId = process.env.GOOGLE_CLIENT_ID || "";
const googleClientSecret = process.env.GOOGLE_CLIENT_SECRET || "";
const jwtCookieName = process.env.JWT_COOKIE_NAME || "falcon.access_token";
const apiUrl = process.env.NEXT_PUBLIC_API_URL

const handler = NextAuth({
    logger: {
        debug: console.log,
        error: console.error,
        warn: console.warn,
        info: console.info
    },
    providers: [
        GoogleProvider({
            clientId: googleClientId,
            clientSecret: googleClientSecret,
            httpOptions: {
                timeout: 40000
            },
            authorization: {
                params: {
                    prompt: "consent",
                    access_type: "offline",
                    response_type: "code"
                }
            },
        })
    ],
    events: {
        signOut: async (message) => {
            cookies().delete(jwtCookieName)
        }
    },
    callbacks: {
        async session({session, token}) {
            const jwtToken = cookies().get(jwtCookieName)

            if (jwtToken === undefined) {
                return {
                    ...session,
                    user: {
                        name: null,
                        email: null,
                        image: null,
                    },
                    expires: new Date().toISOString(),
                }
            }

            const resp = await fetch(apiUrl + "/auth/protected", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer " + jwtToken?.value,
                },
            })

            if (resp.status !== 200) {
                return {
                    ...session,
                    user: {
                        name: null,
                        email: null,
                        image: null,
                    },
                    expires: new Date().toISOString(),
                }
            }

            return session;
        },
        async signIn({user, account, profile, email, credentials}) {
            if (!user || !account) {
                return false;
            }

            if (!account?.access_token) {
                return false;
            }

            const resp = await fetch(apiUrl + "/auth/signin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    type: "google",
                    account_id: user.id,
                    access_token: account.access_token,
                })
            })

            if (resp.status !== 200) {
                console.error("Error", resp.status, resp.statusText)
                return false;
            }

            const data = await resp.json()
            const {token} = data

            cookies().set(
                jwtCookieName,
                token,
                {
                    domain: process.env.NEXT_PUBLIC_COOKIE_DOMAIN ?? "",
                    httpOnly: true,
                    maxAge: 60 * 60 * 24, //24 hours
                }
            );

            return true;
        },
    }
})

export {handler as GET, handler as POST}
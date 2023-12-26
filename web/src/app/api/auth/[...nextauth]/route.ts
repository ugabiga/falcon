import NextAuth from "next-auth";
import GoogleProvider from "next-auth/providers/google";
import {cookies} from "next/headers";

const googleClientId = process.env.GOOGLE_CLIENT_ID || "";
const googleClientSecret = process.env.GOOGLE_CLIENT_SECRET || "";

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
            authorization: {
                params: {
                    prompt: "consent",
                    access_type: "offline",
                    response_type: "code"
                }
            }
        })
    ],
    events: {
        signOut: async (message) => {
            cookies().delete("access_token",);
        }
    },
    callbacks: {
        async signIn({user, account, profile, email, credentials}) {
            if (!user || !account) {
                return false;
            }

            if (!account?.access_token) {
                return false;
            }

            console.log(account)


            const resp = await fetch("http://localhost:8080/auth/signin", {
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
            console.log(resp)

            const data = await resp.json()
            const {token} = data


            cookies().set(
                "falcon.access_token",
                token,
                {
                    httpOnly: true,
                    maxAge: 36000,
                }
            );
            //
            // const respCSRF = await api.getCSRF()
            //
            // cookies().set(
            //     "back.csrf_token",
            //     respCSRF.csrfToken!,
            //     {
            //         httpOnly: true,
            //         maxAge: 36000,
            //     }
            // );
            //
            // user.id = resp.user.id.toString()
            // user.name = resp.user.firstName + " " + resp.user.lastName

            return true;
        },
    }
})

export {handler as GET, handler as POST}
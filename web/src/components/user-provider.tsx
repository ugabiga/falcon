"use client";

import React from "react";
import {useQuery} from "@apollo/client";
import {GetUserDocument} from "@/graph/generated/generated";
import {NavigationBar} from "@/components/navigation-bar";
import {useAppDispatch} from "@/store";
import {set} from "@/store/userSlice";


export function UserProvider({children}: { children: React.ReactNode }) {
    const dispatch = useAppDispatch()
    const {data, loading} = useQuery(GetUserDocument, {
        variables: {
            id: "1"
        }
    })

    if (loading) {
        return null
    }

    console.log(data)

    dispatch(set({
        name: "test",
        isLogged: true
    }))
    // dispatch(set(
    //     data
    //         ? {
    //             name: data.user.name,
    //             isLogged: true
    //         }
    //         : {
    //             name: "",
    //             isLogged: false
    //         }
    // ))

    return (
        <>
            <NavigationBar/>
            {children}
        </>
    )
}
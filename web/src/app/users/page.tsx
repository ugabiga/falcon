"use client";

import {UpdateUserDocument, UserIndexDocument} from "@/graph/generated/generated";
import {useMutation, useQuery} from "@apollo/client";
import {Label} from "@/components/ui/label";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {Button} from "@/components/ui/button";
import {useEffect, useState} from "react";
import {Loading} from "@/components/loading";
import {errorToast, normalToast} from "@/components/toast";
import {Error} from "@/components/error";
import {useTranslation} from "react-i18next";


export default function Users() {
    const {t} = useTranslation()
    const {data, loading, error} = useQuery(UserIndexDocument)
    const [updateUser] = useMutation(UpdateUserDocument)
    const [user, setUser] = useState({
        name: "",
        timezone: ""
    })

    useEffect(() => {
        setUser({
            name: data?.userIndex.user.name || "",
            timezone: data?.userIndex.user.timezone || ""
        })
    }, [data]);

    if (loading) {
        return <Loading/>
    }

    if (error) {
        return <Error message={error.message}/>
    }

    if (!data) {
        return <div>No data</div>
    }


    const handleOnSave = () => {
        updateUser({
            variables: {
                name: user.name,
                timezone: user.timezone
            }
        }).then(() => {
            normalToast({message: t("users.profile.saved")})
        }).catch((e) => {
            errorToast(t("error." + e.message))
        })
    }

    return (
        <main className="min-h-screen p-12 ">
            <div className="w-full text-center">
                <h1 className="text-4xl font-bold">
                    {t("users.profile.title")}
                </h1>
            </div>

            <div className={" mt-6 space-y-6 w-full grid place-items-center"}>
                <div className="grid w-full max-w-sm items-center gap-1.5">
                    <Label htmlFor="name">{t("users.profile.name")}</Label>
                    <Input type="name" id="name" defaultValue={data.userIndex.user.name} onChange={(e) => {
                        return setUser({
                            ...user,
                            name: e.target.value
                        })
                    }}/>
                </div>

                <div className="grid w-full max-w-sm items-center gap-1.5">
                    <Label htmlFor="name">{t("users.profile.timezone.title")}</Label>
                    <Select defaultValue={data.userIndex.user.timezone}
                            onValueChange={(value) => {
                                return setUser({
                                    ...user,
                                    timezone: value
                                })
                            }}>
                        <SelectTrigger>
                            <SelectValue placeholder="Timezone"/>
                        </SelectTrigger>
                        <SelectContent>
                            <SelectItem value="default" disabled>{t("users.profile.timezone.placeholder")}</SelectItem>
                            <SelectItem value="Asia/Seoul">{t("users.profile.timezone.asia-seoul")}</SelectItem>
                            <SelectItem value="UTC">{t("users.profile.timezone.utc")}</SelectItem>
                        </SelectContent>
                    </Select>
                </div>

                <div className="grid w-full max-w-sm">
                    <Button onClick={handleOnSave}>
                        {t("users.profile.save")}
                    </Button>
                </div>
            </div>
        </main>
    )
}
"use client";

import {GetUserDocument, UpdateUserDocument} from "@/graph/generated/generated";
import {useMutation, useQuery} from "@apollo/client";
import {Label} from "@/components/ui/label";
import {Input} from "@/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {Button} from "@/components/ui/button";
import {useEffect, useState} from "react";
import {useToast} from "@/components/ui/use-toast";

export default function Users() {
    const { toast } = useToast()
    const {data, loading} = useQuery(GetUserDocument)
    const [updateUser] = useMutation(UpdateUserDocument)
    const [user, setUser] = useState({
        name: "",
        timezone: ""
    })

    useEffect(() => {
        setUser({
            name: data?.user.name || "",
            timezone: data?.user.timezone || ""
        })
    }, [data]);

    if (loading) {
        return <div></div>
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
            toast({
                title: "Success",
                description: "Your profile has been updated",
            })
        })
    }

    return (
        <main className="min-h-screen p-12">
            <h1 className="text-4xl font-bold">Profile</h1>
            <div className={"mt-6 space-y-6 w-full"}>

                <div className="grid w-full max-w-sm items-center gap-1.5">
                    <Label htmlFor="name">Name</Label>
                    <Input type="name" id="name" defaultValue={data.user.name} onChange={(e) => {
                        return setUser({
                            ...user,
                            name: e.target.value
                        })
                    }}/>
                </div>

                <div className="grid w-full max-w-sm items-center gap-1.5">
                    <Label htmlFor="name">Timezone</Label>
                    <Select defaultValue={data.user.timezone}
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
                            <SelectItem value="default" disabled>Choose your timezone</SelectItem>
                            <SelectItem value="Asia/Seoul">Seoul</SelectItem>
                            <SelectItem value="UTC">UTC</SelectItem>
                        </SelectContent>
                    </Select>
                </div>

                <div className="grid w-full max-w-sm">
                    <Button onClick={handleOnSave}>
                        Save
                    </Button>
                </div>

            </div>
        </main>
    )
}
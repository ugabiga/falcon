"use client";

import {TaskHistoryTable} from "@/app/tasks/[id]/history/table";
import {useQuery} from "@apollo/client";
import {GetTaskHistoryIndexDocument} from "@/graph/generated/generated";
import {Loading} from "@/components/loading";
import {Button} from "@/components/ui/button";
import {useRouter} from "next/navigation";
import {Error} from "@/components/error";
import {useTranslation} from "react-i18next";

export default function TaskHistory({params}: { params: { id: string } }) {
    const {t} = useTranslation()
    const router = useRouter()
    const {data, loading, error} = useQuery(GetTaskHistoryIndexDocument, {
        variables: {
            taskID: params.id
        }
    })

    if (loading) {
        return <Loading/>
    }

    if (error) {
        return <Error message={error.message}/>
    }

    return (
        <>
            <div className={"mt-2 w-full flex"}>
                <Button variant="link" onClick={() => {
                    if (data?.taskHistoryIndex?.task?.tradingAccountID == null) {
                        router.push("/tasks")
                        return
                    }

                    router.push("/tasks?trading_account_id=" + data.taskHistoryIndex?.task?.tradingAccountID)
                }}>
                    {t("task_history.back.btn")}
                </Button>
            </div>
            <main className="min-h-screen mt-6 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">
                <h1 className="text-3xl font-bold">
                    {t("task_history.title")}
                </h1>
                <div className="mt-6">
                    {/*@ts-ignore*/}
                    <TaskHistoryTable taskHistories={data?.taskHistoryIndex?.taskHistories}/>
                </div>
            </main>
        </>
    )
}
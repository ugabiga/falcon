"use client";

import {TaskHistoryTable} from "@/app/tasks/[id]/history/table";
import {useQuery} from "@apollo/client";
import {GetTaskHistoryIndexDocument} from "@/graph/generated/generated";
import {Loading} from "@/components/loading";

export default function TaskHistory({params}: { params: { id: string } }) {
    const {data, loading} = useQuery(GetTaskHistoryIndexDocument, {
        variables: {
            taskID: params.id
        }
    })

    if (loading) {
        return <Loading/>
    }

    return (
        <main className="min-h-screen mt-12 pr-4 pl-4 md:max-w-[1200px] overflow-auto w-full mx-auto">
            <h1 className="text-3xl font-bold">
                Task History (Task ID : {params.id})
            </h1>
            <div className="mt-6">
                {/*@ts-ignore*/}
                <TaskHistoryTable taskHistories={data?.taskHistoryIndex?.taskHistories}/>
            </div>
        </main>
    )
}
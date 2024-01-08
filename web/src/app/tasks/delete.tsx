import {useTranslation} from "react-i18next";
import {useAppDispatch} from "@/store";
import {useMutation} from "@apollo/client";
import {DeleteTaskDocument, Task} from "@/graph/generated/generated";
import React, {useState} from "react";
import {refreshTask} from "@/store/taskSlice";
import {errorToast} from "@/components/toast";
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger
} from "@/components/ui/alert-dialog";
import {DropdownMenuItem} from "@/components/ui/dropdown-menu";

export function DeleteTask(
    {task}: { task: Task }
) {
    const {t} = useTranslation();
    const dispatch = useAppDispatch()
    const [deleteTask] = useMutation(DeleteTaskDocument);
    const [openDialog, setOpenDialog] = useState(false)

    const handleDelete = () => {
        deleteTask({
            variables: {
                taskID: task.id,
                tradingAccountID: task.tradingAccountID
            }
        }).then(() => {
            setOpenDialog(false)
            dispatch(refreshTask({
                tradingAccountID: task.tradingAccountID,
                refresh: true
            }))
        }).catch(error => {
            errorToast(error.message)
        })
    }

    return (
        <AlertDialog open={openDialog} onOpenChange={setOpenDialog}>
            <AlertDialogTrigger asChild>
                <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                    {t("task.delete.btn")}
                </DropdownMenuItem>
            </AlertDialogTrigger>
            <AlertDialogContent>
                <AlertDialogHeader>
                    <AlertDialogTitle>
                        {t("task.delete.title")}
                    </AlertDialogTitle>
                </AlertDialogHeader>
                <AlertDialogFooter className="flex">
                    <AlertDialogCancel>
                        {t("task.delete.cancel")}
                    </AlertDialogCancel>
                    <div className="flex-grow"/>
                    <AlertDialogAction
                        className="btn btn-danger"
                        onClick={() => handleDelete()}
                    >
                        {t("task.delete.yes")}
                    </AlertDialogAction>
                </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
    )
}
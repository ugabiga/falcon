import {ModelTask} from "@/api/model";
import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import {Button} from "@/components/ui/button";
import {useMutation} from "@apollo/client";
import {DeleteTaskDocument} from "@/graph/generated/generated";
import {errorToast} from "@/components/toast";
import {useTranslation} from "react-i18next";

export default function TaskDelete(
    {
        task,
        onDelete
    }: {
        task: ModelTask,
        onDelete: () => void
    }) {
    const {t} = useTranslation();
    const [deleteTask] = useMutation(DeleteTaskDocument);

    function onSubmit() {
        deleteTask({
            variables: {
                taskID: task.id!,
                tradingAccountID: task.trading_account_id!
            }
        }).then(() => {
            onDelete()
        }).catch(e => {
            errorToast(t("error." + e.message))
        })
    }

    return (
        <AlertDialog>
            <AlertDialogTrigger asChild>
                <Button variant="destructive">
                    {t("task.delete.btn")}
                </Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
                <AlertDialogHeader>
                    <AlertDialogTitle>
                        {t("task.delete.title")}
                    </AlertDialogTitle>
                </AlertDialogHeader>
                <AlertDialogFooter>
                    <AlertDialogCancel>
                        {t("task.delete.cancel")}
                    </AlertDialogCancel>
                    <AlertDialogAction
                        onClick={onSubmit}
                    >
                        {t("task.delete.yes")}
                    </AlertDialogAction>
                </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
    )
}

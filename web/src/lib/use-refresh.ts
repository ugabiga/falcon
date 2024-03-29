import {useEffect} from "react";
import {RefreshTarget, setRefresher, setRefresherNone} from "@/store/refresherSlice";
import {useAppDispatch, useAppSelector} from "@/store";

export function useReceiveRefreshSignal({target, afterReceiveSignal}: {
    target: RefreshTarget,
    afterReceiveSignal: () => void
}) {
    const dispatch = useAppDispatch()
    const refresher = useAppSelector((state) => state.refresher)

    function receiverRefreshSignal(inputTarget: RefreshTarget, inputAfterReceiveSignal: () => void) {
        target = inputTarget
        afterReceiveSignal = inputAfterReceiveSignal
    }

    useEffect(() => {
        if (refresher.targetName === target) {
            afterReceiveSignal()
            dispatch(setRefresherNone())
        }
    }, [refresher])

    return {
        receiverRefreshSignal
    }
}

export function useSendRefreshSignal() {
    const dispatch = useAppDispatch()

    function sendRefresh(target: RefreshTarget) {
        dispatch(setRefresher({targetName: target, refresh: true}))
    }

    return {
        sendRefresh
    }
}

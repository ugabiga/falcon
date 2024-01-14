import {configureStore} from "@reduxjs/toolkit";
import {TypedUseSelectorHook, useDispatch, useSelector} from "react-redux";
import userSlice from "@/store/userSlice";
import tradingAccountSlice from "@/store/tradingAccountSlice";
import taskSlice from "@/store/taskSlice";
import refresherSlice from "@/store/refresherSlice";


export const store = configureStore({
    reducer: {
        user: userSlice.reducer,
        tradingAccount: tradingAccountSlice.reducer,
        task: taskSlice.reducer,
        refresher: refresherSlice.reducer,
    }
})


export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector
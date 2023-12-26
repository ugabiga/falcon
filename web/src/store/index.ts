import {configureStore} from "@reduxjs/toolkit";
import {TypedUseSelectorHook, useDispatch, useSelector} from "react-redux";
import userSlice from "@/store/userSlice";
import tradingAccountSlice from "@/store/tradingAccountSlice";


export const store = configureStore({
    reducer: {
        user: userSlice.reducer,
        tradingAccount: tradingAccountSlice.reducer,
    }
})


export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector
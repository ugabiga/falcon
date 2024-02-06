import {combineReducers, configureStore} from "@reduxjs/toolkit";
import {TypedUseSelectorHook, useDispatch, useSelector} from "react-redux";
import userSlice from "@/store/userSlice";
import tradingAccountSlice from "@/store/tradingAccountSlice";
import taskSlice from "@/store/taskSlice";
import refresherSlice from "@/store/refresherSlice";
import storage from "redux-persist/lib/storage";
import {FLUSH, PAUSE, PERSIST, persistReducer, persistStore, PURGE, REGISTER, REHYDRATE} from "redux-persist";
import {thunk} from "redux-thunk";


const rootReducer = combineReducers({
    user: userSlice.reducer,
    tradingAccount: tradingAccountSlice.reducer,
    task: taskSlice.reducer,
    refresher: refresherSlice.reducer,
})

const persistConfig = {
    key: 'root',
    storage: storage,
    whitelist: []
}

const persistedReducer = persistReducer(persistConfig, rootReducer)

export const store = configureStore({
    reducer: persistedReducer,
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({
            serializableCheck: {
                ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER]
            }
        }).concat(thunk)
})

export const persistor = persistStore(store)
export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector
import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export enum RefreshTarget {
    None = '',
    Task = 'task',
}

export interface Refresh {
    targetName?: RefreshTarget,
    refresh: boolean,
    params?: any
}

const initialState: Refresh = {
    targetName: RefreshTarget.None,
    refresh: false,
}

const refresherSlice = createSlice({
    name: 'task',
    initialState,
    reducers: {
        refreshTask: (state, action: PayloadAction<{
            tradingAccountID?: string,
            refresh: boolean
        }>) => {
            state.targetName = RefreshTarget.Task;
            state.refresh = action.payload.refresh;
            state.params = {
                tradingAccountID: action.payload.tradingAccountID
            }
        },
        setRefresher: (state, action: PayloadAction<Refresh>) => {
            state.targetName = action.payload.targetName
        },
        setRefresherNone: (state) => {
            state.targetName = RefreshTarget.None
            state.refresh = false
        }
    }
});

export const {
    refreshTask,
    setRefresher,
    setRefresherNone
} = refresherSlice.actions;
export default refresherSlice;
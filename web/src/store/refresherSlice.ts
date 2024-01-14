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
        }
    }
});

export const {refreshTask} = refresherSlice.actions;
export default refresherSlice;
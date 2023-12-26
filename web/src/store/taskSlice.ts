import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export interface Task {
    refresh: boolean,
    tradingAccountID?: string,
}

const initialState: Task = {
    refresh: false,
}

const taskSlice = createSlice({
    name: 'task',
    initialState,
    reducers: {
        refreshTask: (state, action: PayloadAction<{
            tradingAccountID?: string,
            refresh: boolean
        }>) => {
            state.refresh = action.payload.refresh;
            state.tradingAccountID = action.payload.tradingAccountID;
        }
    }
});

export const {refreshTask} = taskSlice.actions;
export default taskSlice;
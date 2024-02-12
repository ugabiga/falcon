import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export interface TutorialState {
    tradingAccountTutorial: boolean,
    taskTutorial: boolean,
}

const initialState: TutorialState = {
    tradingAccountTutorial: false,
    taskTutorial: false,
}

const tutorialSlice = createSlice({
    name: 'tutorial',
    initialState,
    reducers: {
        setTradingAccountTutorial: (state, action: PayloadAction<boolean>) => {
            state.tradingAccountTutorial = action.payload;
        },
        setTaskTutorial: (state, action: PayloadAction<boolean>) => {
            state.taskTutorial = action.payload;
        }
    }
});

export const {
    setTradingAccountTutorial,
    setTaskTutorial
} = tutorialSlice.actions;
export default tutorialSlice;
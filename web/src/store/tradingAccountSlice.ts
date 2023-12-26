import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export interface TradingAccount {
    refresh: boolean,
}

const initialState: TradingAccount = {
    refresh: false,
}

const tradingAccountSlice = createSlice({
    name: 'tradingAccount',
    initialState,
    reducers: {
        refreshTradingAccount: (state, action: PayloadAction<boolean>) => {
            state.refresh = action.payload;
        }
    }
});

export const {refreshTradingAccount} = tradingAccountSlice.actions;
export default tradingAccountSlice;
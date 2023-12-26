import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export interface User {
    isLogged: boolean,
    name: string
}

const initialState: User = {
    isLogged: false,
    name: ""
}

const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        set: (state, action: PayloadAction<User>) => {
            state = action.payload
        }
    }
});

export const {set} = userSlice.actions;
export default userSlice
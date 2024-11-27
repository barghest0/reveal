import { createSlice } from "@reduxjs/toolkit";
import { UserAuth } from "shared/api/user";
import { loginUserAuth } from "./actions/login-user";
import { StatusFlag } from "shared/index";
import { getProfile } from "./actions/profile-user";
import { registrationUser } from "./actions/registration-user";



const initialState: UserAuth = {
    data: undefined,
    status: null,
    loading: false
}


export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
        .addCase(loginUserAuth.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(loginUserAuth.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.loading = false;
            state.data = action.payload;
        })
        .addCase(loginUserAuth.rejected, (state) => {
            state.status = StatusFlag.Rejected;
            state.loading = false;
        })

        builder
        .addCase(getProfile.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(getProfile.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.loading = false;
            state.data = action.payload;
        })
        .addCase(getProfile.rejected, (state) => {
            state.status = StatusFlag.Rejected;
            state.loading = false;
        })

        builder
        .addCase(registrationUser.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(registrationUser.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.loading = false;
            state.data = action.payload;
        })
        .addCase(registrationUser.rejected, (state) => {
            state.status = StatusFlag.Rejected;
            state.loading = false;
        })
    },
})
import { createAsyncThunk } from "@reduxjs/toolkit";
import { UserApi } from "shared/api/user";




export const getProfile = createAsyncThunk(
    "user/loginUserAuth",
    async () => {
        const response = await UserApi.getUserProfile();
        return response
    }
);
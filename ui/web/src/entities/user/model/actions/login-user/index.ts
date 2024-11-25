import { createAsyncThunk } from "@reduxjs/toolkit";
import { UserApi } from "shared/api/user";



export const loginUserAuth = createAsyncThunk(
    "user/loginUserAuth",
    async ({email, password}: {email: string, password: string}) => {
        const response = await UserApi.login(email, password);
        console.log(response)
        return response
    }
);
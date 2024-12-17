import { createAsyncThunk } from "@reduxjs/toolkit";
import { UserApi } from "shared/api/user";




export const registrationUser = createAsyncThunk(
    "user/registrationUser",
    async ({email, password}: {email: string, password: string}) => {
        const response = await UserApi.registrationUser(email, password);
        return response
    }
)
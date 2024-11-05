import { createAsyncThunk } from "@reduxjs/toolkit";
import { CreateCartApi } from "entities/cart/api";


export const CreateCart = createAsyncThunk(
    'cart/createCart',
    async () => {
        const response = await CreateCartApi();
        return response
    }
) 
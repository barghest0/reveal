import { createAsyncThunk } from "@reduxjs/toolkit";
import { createCartApi } from "features/button-add-cart/api";

export const createCart = createAsyncThunk(
    'cart/createCart',
    async () => {
        const response = await createCartApi();
        return response
    }
) 
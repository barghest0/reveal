import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "shared/api/cart";





export const getProductsCart = createAsyncThunk(
    'cart/getProductsCart',
    async (userId: number) => {
        const response = await CartApi.getCartApi(userId)
        return response
    }
)
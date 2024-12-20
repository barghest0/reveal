import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "shared/api/cart";





export const getProductsCart = createAsyncThunk(
    'cart/getProductsCart',
    async () => {
        const response = await CartApi.getCartApi()
        return response
    }
)
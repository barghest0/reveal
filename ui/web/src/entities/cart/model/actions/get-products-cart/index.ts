import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "shared/api/cart";





export const getProductsCart = createAsyncThunk(
    'cart/getProductsCart',
    async (cart: number) => {
        const response = await CartApi.getCartApi(cart)
        return response
    }
)
import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "../../request-api/cart-api";




export const getProductsCart = createAsyncThunk(
    'cart/getProductsCart',
    async (cart: number) => {
        const response = await CartApi.getCartApi(cart)
        return response
    }
)
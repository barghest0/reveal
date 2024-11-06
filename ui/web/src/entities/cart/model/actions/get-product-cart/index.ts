import { createAsyncThunk } from "@reduxjs/toolkit";
import { Cart, CartApi } from "entities/cart";





export const getProductsCart = createAsyncThunk(
    'cart/getProductsCart',
    async (cart: number) => {
        const response = await CartApi.getCartApi(cart)
        return response
    }
)
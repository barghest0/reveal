import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "shared/api/cart";



export const deleteProductFromCart = createAsyncThunk(
    "cart/deleteProductFromCart",
    async(productId: number) => {
        const response = await CartApi.removeProductToCartApi(productId)
        return response
    }
)
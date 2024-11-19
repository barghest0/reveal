import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "shared/api/cart";



export const deleteProductFromCart = createAsyncThunk(
    "cart/deleteProductFromCart",
    async({userId, productId}: {userId: number, productId: number}) => {
        const response = await CartApi.removeProductToCartApi(userId, productId)
        return response
    }
)
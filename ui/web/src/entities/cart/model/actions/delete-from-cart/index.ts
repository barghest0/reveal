import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "shared/api/cart";



export const deleteProductFromCart = createAsyncThunk(
    "cart/deleteProductFromCart",
    async({cartId, productId}: {cartId: number, productId: number}) => {
        const response = await CartApi.removeProductToCartApi(cartId, productId)
        return response
    }
)
import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartItem } from "../../types";
import { CartApi } from "entities/cart/api";



export const addProductToCart = createAsyncThunk(
    'cart/addProductToCart',
    async (cartId: number) => {
        // const response = await CartApi.addProductToCartApi(cartId)
        // return response
    }
)


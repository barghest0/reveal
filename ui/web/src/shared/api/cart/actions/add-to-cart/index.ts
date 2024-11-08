import { createAsyncThunk } from "@reduxjs/toolkit";
import { ProductItem } from "../../types";
import { CartApi } from "../../request-api/cart-api";


export const addProductToCart = createAsyncThunk(
    'cart/addProductToCart',
    async ({cartId, product}: {cartId:number; product: ProductItem}) => { 
        const response = await CartApi.addProductToCartApi(cartId, product)
        return response
    }
)


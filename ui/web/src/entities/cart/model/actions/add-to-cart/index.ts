import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartApi } from "shared/api/cart";
import { ProductItem } from "shared/api/products";



export const addProductToCart = createAsyncThunk(
    "cart/addProductToCart",
    async ({cartId, product}: {cartId:number; product: ProductItem}) => { 
        const response = await CartApi.addProductToCartApi(cartId, product)
        return response
    }
)


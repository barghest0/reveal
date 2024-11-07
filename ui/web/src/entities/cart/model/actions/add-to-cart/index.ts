import { createAsyncThunk } from "@reduxjs/toolkit";
import { ProductItem } from "entities/card-product";
import { CartApi } from "entities/cart/api";



export const addProductToCart = createAsyncThunk(
    'cart/addProductToCart',
    async ({cartId, product}: {cartId:number; product: ProductItem}) => {
        console.log("THUNK", product)
        const response = await CartApi.addProductToCartApi(cartId, product)
        return response
    }
)


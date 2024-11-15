import { createAsyncThunk } from "@reduxjs/toolkit";
import { Cart, CartApi } from "shared/api/cart";
import { ProductItem } from "shared/api/products";

export const addProductToCart = createAsyncThunk(
    "cart/addProductToCart",
    async ({userId, product}: {userId:number, product: ProductItem}) => {
        const response = await CartApi.addProductToCartApi(userId, product);
        return response
    }    

)



import { createAsyncThunk } from "@reduxjs/toolkit";
import { Cart, CartApi } from "shared/api/cart";
import { ProductItem } from "shared/api/products";

export const addProductToCart = createAsyncThunk(
    "cart/addProductToCart",
    async (product: ProductItem) => {
        const response = await CartApi.addProductToCartApi(product);
        return response
    }    

)



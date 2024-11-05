import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartItem, CartProduct } from "../../types";
import { CartApi } from "entities/cart/api";



export const addProductToCart = createAsyncThunk(
    'cart/addProductToCart',
    async (product: CartItem, {getState, dispatch}) => {
        const state = getState() as { cart: CartProduct};
        let cart = state.cart.data;
        if (cart) {
            const response = await CartApi.addProductToCartApi(cart.id, product);
            return response;
        }
        
    }
)


import { createAsyncThunk } from "@reduxjs/toolkit";
import { CartItem } from "features/button-add-cart/api";
import { cartProduct } from "features/button-add-cart/model";
import { createCart } from "../create-cart";
import { AddProductToCart } from "features/button-add-cart";

export const addProductToCart = createAsyncThunk(
    'cart/addProductToCart',
    async (product: CartItem, {getState, dispatch}) => {
        const state = getState() as { cart: cartProduct};
        let cart = state.cart.data;

        // Проверяем, существует ли корзина
        if (!cart) {
            cart = await dispatch(createCart()).unwrap();
        }

        const response = await AddProductToCart(cart.id, product);
        return response;
    }
)


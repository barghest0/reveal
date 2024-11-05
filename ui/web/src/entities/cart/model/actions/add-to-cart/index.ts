import { createAsyncThunk } from "@reduxjs/toolkit";
import { CreateCart } from "../create-cart";
import { CartItem, CartProduct } from "../../types";
import { AddProductToCartApi } from "entities/cart/api";


export const AddProductToCart = createAsyncThunk(
    'cart/addProductToCart',
    async (product: CartItem, {getState, dispatch}) => {
        const state = getState() as { cart: CartProduct};
        let cart = state.cart.data;

        // Проверяем, существует ли корзина
        if (!cart) {
            cart = await dispatch(CreateCart()).unwrap();
        }

        const response = await AddProductToCartApi(cart.id, product);
        return response;
    }
)


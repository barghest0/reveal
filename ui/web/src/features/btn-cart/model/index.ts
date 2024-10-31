import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { addProductToCartApi, Cart, CartItem, createCartApi } from "../api";
import { StatusFlag } from "shared/types/status-flag";


export interface cartProduct {
    data: Cart | undefined,
    status: null | StatusFlag,
    loading: boolean
}

export const createCart = createAsyncThunk(
    'cart/createCart',
    async () => {
        const response = await createCartApi();
        return response
    }
)

export const addProductToCart = createAsyncThunk(
    'cart/addProductToCart',
    async (product: CartItem, {getState, dispatch}) => {
        const state = getState() as { cart: cartProduct};
        let cart = state.cart.data;

        // Проверяем, существует ли корзина
        if (!cart) {
            cart = await dispatch(createCart()).unwrap();
        }

        const response = await addProductToCartApi(cart.id, product);
        return response;
    }
)

const initialState: cartProduct = {
    status: null,
    data: undefined,
    loading: false,
}

export const cartSlice = createSlice({
    name: 'cart',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
        .addCase(createCart.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(createCart.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.data = action.payload;
            state.loading = false;
        })
        .addCase(createCart.rejected, (state) => {
            state.status = StatusFlag.Rejected;
            state.loading = false;
        })
        .addCase(addProductToCart.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(addProductToCart.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.loading = false;
            if (state.data) {
                state.data.products.push(action.payload);
            }
        })
        .addCase(addProductToCart.rejected, (state) => {
            state.status = StatusFlag.Rejected;
            state.loading = false
        })
    }
})

const reducer = cartSlice.reducer;
export default reducer;
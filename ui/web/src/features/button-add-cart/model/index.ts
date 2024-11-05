
import { Cart } from "../api";
import { StatusFlag } from "shared/types/status/status-flag";
import { createSlice } from "@reduxjs/toolkit";
import { createCart } from "entities/cart";


export interface cartProduct {
    data: Cart | undefined,
    status: null | StatusFlag,
    loading: boolean
}

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
        // .addCase(addProductToCart.pending, (state) => {
        //     state.status = StatusFlag.Pending;
        //     state.loading = true;
        // })
        // .addCase(addProductToCart.fulfilled, (state, action) => {
        //     state.status = StatusFlag.Fulfilled;
        //     state.loading = false;
        //     if (state.data) {
        //         state.data.products.push(action.payload);
        //     }
        // })
        // .addCase(addProductToCart.rejected, (state) => {
        //     state.status = StatusFlag.Rejected;
        //     state.loading = false
        // })
    }
})

export const cartReducer = cartSlice.reducer;

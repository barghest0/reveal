

import { StatusFlag } from "shared/types/status/status-flag";
import { createSlice } from "@reduxjs/toolkit";
import { CartProduct, CreateCart } from "entities/cart";



const initialState: CartProduct = {
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
        .addCase(CreateCart.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(CreateCart.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.data = action.payload;
            state.loading = false;
        })
        .addCase(CreateCart.rejected, (state) => {
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

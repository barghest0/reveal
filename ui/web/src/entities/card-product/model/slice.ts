import { createSlice } from "@reduxjs/toolkit";
import { ProductsCard } from "./types";
import { getProductsCard } from "./actions/get-products";
import { StatusFlag } from "shared/index";



const initialState: ProductsCard = {
    data: undefined,
    status: null,
    loading: false

}

export const productsSlice = createSlice({
    name: "products",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
        .addCase(getProductsCard.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(getProductsCard.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.loading = false;
            state.data = action.payload;
        })
        .addCase(getProductsCard.rejected, (state) => {
            state.status = StatusFlag.Rejected;
            state.loading = false;
        })
    }
})
export const productReducer = productsSlice.reducer
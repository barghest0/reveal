
import { createSlice } from "@reduxjs/toolkit";
import { CartProduct,} from "entities/cart";
import { StatusFlag } from "shared/index";
import { addProductToCart, GetCartType, getProductsCart } from "shared/api/cart";
import { CartTypes } from "shared/api/cart/types";


const initialState: GetCartType<CartTypes.CART_PRODUCT> = {
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
        .addCase(getProductsCart.pending, (state) => {
            state.status = StatusFlag.Pending;
            state.loading = true;
        })
        .addCase(getProductsCart.fulfilled, (state, action) => {
            state.status = StatusFlag.Fulfilled;
            state.loading = false;
            state.data = action.payload; 
        })
        .addCase(getProductsCart.rejected, (state) => {
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
                state.data.Products.push(action.payload);
            }
        })
        .addCase(addProductToCart.rejected, (state) => {
            state.status = StatusFlag.Rejected;
            state.loading = false
        })
    }
})

export const cartReducer = cartSlice.reducer;

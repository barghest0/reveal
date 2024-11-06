
import { createSlice } from "@reduxjs/toolkit";
import { CartProduct,} from "entities/cart";
import { getProductsCart } from "./actions/get-product-cart";
import { StatusFlag } from "shared/index";
// import { addProductToCart } from "entities/cart/model/actions/add-to-cart";



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

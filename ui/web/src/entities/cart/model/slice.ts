
import { createSlice } from "@reduxjs/toolkit";
import { StatusFlag } from "shared/index";
import { CartProduct } from "shared/api/cart/types";
import { getProductsCart } from "./actions/get-products-cart";
import { addProductToCart } from "./actions/add-to-cart";
import { ProductsCard } from "shared/api/products";
import { getProductsCard } from "./actions/get-products-card";
import { deleteProductFromCart } from "./actions/delete-from-cart";

type ProductsCartState = {
    cart: CartProduct;
    products: ProductsCard
}

const initialState: ProductsCartState = {
    cart: {
        status: null,
        data: undefined,
        loading: false,
    },
    products: {
        status: null,
        data: undefined,
        loading: false,
    }
    
}

export const productsCartSlice = createSlice({
    name: 'cart',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
        .addCase(getProductsCart.pending, (state) => {
            state.cart.status = StatusFlag.Pending;
            state.cart.loading = true;
        })
        .addCase(getProductsCart.fulfilled, (state, action) => {
            state.cart.status = StatusFlag.Fulfilled;
            state.cart.loading = false;
            state.cart.data = action.payload; 
        })
        .addCase(getProductsCart.rejected, (state) => {
            state.cart.status = StatusFlag.Rejected;
            state.cart.loading = false;
        })
        .addCase(addProductToCart.pending, (state) => {
            state.cart.status = StatusFlag.Pending;
            state.cart.loading = true;
        })
        .addCase(addProductToCart.fulfilled, (state, action) => {
            state.cart.status = StatusFlag.Fulfilled;
            state.cart.loading = false;
            if (state.cart.data) {
                state.cart.data.products.push(action.payload);
            }
        })
        .addCase(addProductToCart.rejected, (state) => {
            state.cart.status = StatusFlag.Rejected;
            state.cart.loading = false
        })

        builder
        .addCase(getProductsCard.pending, (state) => {
            state.products.status = StatusFlag.Pending;
            state.products.loading = true;
        })
        .addCase(getProductsCard.fulfilled, (state, action) => {
            state.products.status = StatusFlag.Fulfilled;
            state.products.loading = false;
            state.products.data = action.payload;
        })
        .addCase(getProductsCard.rejected, (state) => {
            state.products.status = StatusFlag.Rejected;
            state.products.loading = false;
        })

        builder
        .addCase(deleteProductFromCart.pending, (state) => {
            state.cart.status = StatusFlag.Pending;
        })
        .addCase(deleteProductFromCart.fulfilled, (state, action) => {
            state.cart.status = StatusFlag.Fulfilled;
            const existingProduct = state.cart.data?.products.findIndex(product => product.product_id === action.payload.product_id);
            console.log("REDUCE", existingProduct)
            if (existingProduct)
            state.cart.data?.products.splice(existingProduct, 1);
        })
        .addCase(deleteProductFromCart.rejected, (state) => {
            state.cart.status = StatusFlag.Rejected;

        })
        
    }
})

export const productsCartReducer = productsCartSlice.reducer;

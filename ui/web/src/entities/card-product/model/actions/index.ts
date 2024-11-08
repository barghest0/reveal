import { createAsyncThunk } from "@reduxjs/toolkit";
import { CardProductsApi } from "shared/api/products";


export const getProductsCard = createAsyncThunk(
    'products/getProductsCard',
    async () => {
        const response = await CardProductsApi.getProducts()
        return response
    }
)
import { createAsyncThunk } from "@reduxjs/toolkit";
import { CardProductsApi } from "entities/card-product/api";



export const getProductsCard = createAsyncThunk(
    'products/getProductsCard',
    async () => {
        const response = await CardProductsApi.getProducts()
        return response
    }
)
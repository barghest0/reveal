import axios from "axios"
import { Cart, CartItem } from "../types";
import { ProductItem } from "shared/api/products";

const baseUrl = "http://localhost:8083/cart"
const config = {
    headers: {
        'Authorization': `Bearer ${process.env.TOKEN}`,
    }
}

export namespace CartApi {
    export const getCartApi = async () => {
        try {
            const response = await axios.get<Cart>(`${baseUrl}`, config);
            return response.data;
        }
        catch (error) {
            console.log("Ошибка получение корзины: ", error);
            throw error;
        }
    }

    export const addProductToCartApi = async (product: ProductItem) => {
        const productItem = {
            product_id: product.id,
            name: product.name,
            price: product.price,
            description: product.description
        }
        try {
            const response = await axios.post(`${baseUrl}/products`, productItem, config);
            return response.data;
        }
        catch (error) {
            console.log("Ошибка добавление товара в корзину: ", error);
            throw error;
        }
    }

    export const removeProductToCartApi = async (productId: number) => {
        try {
            const response = await axios.delete<CartItem>(`${baseUrl}/products/${productId}`, config);
            return response.data;
        }
        catch (error) {
            console.log('Ошибка при удаление товара из корзины: ', error);
            throw error;
        }
    }     
}

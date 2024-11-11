import axios from "axios"
import { Cart, CartItem } from "../types";
import { ProductItem } from "shared/api/products";

const baseUrl = "http://localhost:8083/cart"


export namespace CartApi {
    export const getCartApi = async (userId: number) => {
        try {
            const response = await axios.get<Cart>(`${baseUrl}/${userId}`);
            return response.data;
        }
        catch (error) {
            console.log("Ошибка получение корзины: ", error);
            throw error;
        }
    }

    export const addProductToCartApi = async (cartId: number, product: ProductItem) => {
        const productData = {
            product_id: product.id,
            price: product.price,
            name: product.name
        }

        try {
            const response = await axios.post(`${baseUrl}/${cartId}/products`, productData);
            return response.data;
        }
        catch (error) {
            console.log("Ошибка добавление товара в корзину: ", error);
            throw error;
        }
    }

    export const removeProductToCartApi = async (cartId: number, productId: number) => {
        try {
            console.log("API DELETE", productId)
            const response = await axios.delete<CartItem>(`${baseUrl}/${cartId}/products/${productId}`);
            console.log(response)
            return response.data;
        }
        catch (error) {
            console.log('Ошибка при удаление товара из корзины: ', error);
            throw error;
        }
    }     
}

import axios from "axios"
import { Cart, CartItem } from "../types";
import { ProductItem } from "shared/api/products";

const baseUrl = "http://localhost:8083/cart"


export namespace CartApi {
    export const getCartApi = async (userId: number) => {
        console.log(userId, "GetCartApi")
        try {
            const response = await axios.get<Cart>(`${baseUrl}/${userId}`);
            return response.data;
        }
        catch (error) {
            console.log("Ошибка получение корзины: ", error);
            throw error;
        }
    }

    export const addProductToCartApi = async (userId: number, product: ProductItem) => {
        try {
            const response = await axios.post(`${baseUrl}/${userId}/products`, product);
            return response.data;
        }
        catch (error) {
            console.log("Ошибка добавление товара в корзину: ", error);
            throw error;
        }
    }

    export const removeProductToCartApi = async (cartId: number, productId: number) => {
        console.log(cartId, productId, "REMOVEPRODUCTAPi")
        try {
            const response = await axios.delete<CartItem>(`${baseUrl}/${cartId}/products/${productId}`);
            return response.data;
        }
        catch (error) {
            console.log('Ошибка при удаление товара из корзины: ', error);
            throw error;
        }
    }     
}

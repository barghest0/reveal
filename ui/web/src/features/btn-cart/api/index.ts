import axios from "axios"


const baseUrl = "http://localhost:8083/cart"

export interface Cart {
    id: number,
    userId: number,
    products: CartItem[]
}

export interface CartItem {
    id: number,
    cartId: number,
    productId: number,
    quantity: number,
    price: number
}

export const createCartApi = async () => {
    try {
        const response = await axios.post<Cart>(baseUrl);
        return response.data;
    }
    catch (error) {
        console.log("Ошибка при создании корзины: ", error);
        throw error;
    }
}

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

export const addProductToCartApi = async (cartId: number, product: CartItem) => {
    try {
        const response = await axios.put(`${baseUrl}/products`, product);
        return response.data
    }
    catch (error) {
        console.log("Ошибка добавление товара в корзину: ", error);
        throw error;
    }
}

 export const removeProductToCartApi = async (cartId: number, productId: number) => {
    try {
        const response = await axios.delete<CartItem>(`${baseUrl}/products/${productId}`);
        return response.data;
    }
    catch (error) {
        console.log('Ошибка при удаление товара из корзины: ', error);
        throw error;
    }
}
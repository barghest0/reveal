import axios from "axios"


const baseUrl = "http://localhost:8083/cart"

interface Cart {
    id: number,
    userId: number,
    items: CartItem[]
}

interface CartItem {
    id: number,
    cartId: number,
    productId: number,
    quantity: number,
    price: number
}

const createCart = async () => {
    try {
        const response = await axios.post<Cart>(baseUrl);
        return response.data;
    }
    catch (error) {
        console.log("Ошибка при создании корзины: ", error);
        throw error;
    }
}

const getCart = async (userId: number) => {
    try {
        const response = await axios.get<Cart>(`${baseUrl}/`);
        return response.data
    }
    catch (error) {
        console.log("Ошибка получение корзины: ", error);
        throw error;
    }
}

const addProductToCart = async (cartId: number, productId: number) => {
    try {

    }
}
import axios from "axios";


const baseUrl = "http://localhost:8082/products"

export namespace CardProductsApi {
    export const getProducts = async () => {
        try {
            const response = await axios.get(baseUrl);
            return response.data;
        }
        catch (error) {
            console.log("Ошибка получение товаров: ", error);
            throw error;
        }
    }

}
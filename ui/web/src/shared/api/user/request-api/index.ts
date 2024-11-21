import axios from "axios"
import { User } from "../types"





const baseUrl = "http://localhost:8081/users"
const config = {
    headers: {
        'Authorization': `Bearer ${process.env.TOKEN}`,
    }
}

export namespace UserApi {
    export const login = async (email: string, password: string) => {
        try {
            const response = await axios.post<User>(`${baseUrl}/login`, {email, password});
            return response.data;
        }
        catch (error) {
            console.log("Ошибка авторизации: ", error);
            throw error;
        }
    };
    export const getUserProfile = async () => {
        try {
            const response = await axios.get<User>(`${baseUrl}/profile`, config);
            return response.data
        }
        catch (error) {
            console.log("Ошибка получения пользователя: ", error);
            throw error;
        }
    }
}

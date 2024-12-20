import axios from "axios"
import { User } from "../types"
import { getToken, saveToken } from "shared/lib/session";





const baseUrl = "http://localhost:8081/users"
const config = {
    headers: {
        'Authorization': `Bearer ${getToken()}`,
    }
}

export namespace UserApi {
    export const login = async (email: string, password: string) => {
        try {
            const response = await axios.post<User>(`${baseUrl}/login`, {email, password});
            
            let token = response.data.token;
            if (typeof token === 'string') { 
                token = token.replace('Bearer ', '');
                saveToken(token)
            }
            
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
    };
    export const registrationUser = async (email: string, password: string) => {
        try {
            const response = await axios.post<User>(`${baseUrl}/register`, {email, password});
            return response.data;
        }   
        catch (error) {
            console.log("Ошибка регистрации: ", error);
            throw error;
        }
    }
}

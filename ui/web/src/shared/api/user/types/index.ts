import { StatusFlag } from "shared/types/status/status-flag";



export type User = {
    id: number;
    name: string;
    email: string;
    password: string;
    token: string;
}

export type UserAuth = {
    data: User | undefined;
    status: StatusFlag | null;
    loading: boolean;
}
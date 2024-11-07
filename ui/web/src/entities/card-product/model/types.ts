import { StatusFlag } from "shared/index";

export interface Products {
    products: ProductItem[]
}

export interface ProductItem {
    id: number;
    name: string;
    price: number;
}

export interface ProductsCard {
    data: ProductItem[] | undefined;
    status: null | StatusFlag;
    loading: boolean;
}
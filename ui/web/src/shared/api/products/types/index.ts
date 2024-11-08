import { StatusFlag } from "shared/types/status/status-flag"

export type Products = {
    products: ProductItem[]
}

export type ProductItem = {
    id: number,
    name: string,
    price: number,
}

export type ProductsCard = {
    data: ProductItem[] | undefined,
    status: null | StatusFlag,
    loading: boolean,
}


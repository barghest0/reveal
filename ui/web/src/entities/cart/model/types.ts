import { ProductItem } from "entities/card-product"
import { StatusFlag } from "shared/types/status/status-flag"

export interface Cart {
    id: number,
    userId: number,
    Products: CartItem[]
}

export interface CartItem extends ProductItem{
    id: number,
    cart_id: number,
    product_id: number,
    quantity: number,
}



export interface CartProduct {
    data: Cart | undefined,
    status: null | StatusFlag,
    loading: boolean
}


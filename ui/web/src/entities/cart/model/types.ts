import { StatusFlag } from "shared/types/status/status-flag"

export interface Cart {
    id: number,
    userId: number,
    Products: CartItem[]
}

export interface CartItem {
    id: number,
    cartId: number,
    product_id: number,
    quantity: number,
    price: number
}



export interface CartProduct {
    data: Cart | undefined,
    status: null | StatusFlag,
    loading: boolean
}


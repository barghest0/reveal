import { ProductItem } from "shared/api/products"
import { StatusFlag } from "shared/types/status/status-flag"

export type Cart = {
    id: number,
    userId: number,
    Products: CartItem[]
}

export type CartItem = ProductItem & {
    id: number,
    cart_id: number,
    product_id: number,
    quantity: number
}


export type CartProduct = {
    data: Cart | undefined,
    status: null | StatusFlag,
    loading: boolean,
}


import { StatusFlag } from "shared/types/status/status-flag"

export enum CartTypes {
    PRODUCTS = "PRODUCTS",
    PRODUCT_ITEM = "PRODUCT_ITEM",
    PRODUCTS_CARD = "PRODUCTS_CARD",
    CART = "CART",
    CART_ITEM = "CART_ITEM",
    CART_PRODUCT = "CART_PRODUCT"
}

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

type CartTypeMap = {
    [CartTypes.PRODUCTS]: Products;
    [CartTypes.PRODUCT_ITEM]: ProductItem;
    [CartTypes.PRODUCTS_CARD]: ProductsCard;
    [CartTypes.CART]: Cart;
    [CartTypes.CART_ITEM]: CartItem;
    [CartTypes.CART_PRODUCT]: CartProduct;
}

export type GetCartType<T extends CartTypes> = CartTypeMap[T];
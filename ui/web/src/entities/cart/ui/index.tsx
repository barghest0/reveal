import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";
import { getProductsCart } from "../model/actions/get-products-cart";

interface Props {

}

export const ProductsCart: React.FC <Props> = () => {
    const dispatch = useAppDispatch();
    const {data, status, loading} = useAppSelector((state) => state.productsCart.cart);

    useEffect(() => {
        dispatch(getProductsCart(1))
    }, [])
    

    return (
        <div>  
            <h1>Cart:</h1>
            {data?.Products.map((product) => (
                <p>Product ID: {product.product_id}</p>
            ))}
        </div>
    )
}
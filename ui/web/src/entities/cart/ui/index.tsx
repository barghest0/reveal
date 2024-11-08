import React, { useEffect } from "react";
import { getProductsCart } from "shared/api/cart";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";

interface Props {

}

export const ProductsCart: React.FC <Props> = () => {
    const dispatch = useAppDispatch();
    const {data, status, loading} = useAppSelector((state) => state.cart);

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
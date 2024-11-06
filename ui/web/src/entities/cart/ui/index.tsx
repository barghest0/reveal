import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";
import { getProductsCart } from "../model/actions/get-product-cart";
import { StatusFlag } from "shared/index";


interface Props {

}

export const ProductsCart: React.FC <Props> = () => {
    const dispatch = useAppDispatch();
    const cart = useAppSelector((state) => state.cart);

    useEffect(() => {
        if (cart.data?.userId) {
            dispatch(getProductsCart(cart.data.userId))
        }
    }, [dispatch])

    function getPr() {
        if (cart.data?.userId) {
            dispatch(getProductsCart(cart.data.userId))
        }
    }
    console.log(cart)
    if (cart.loading) return <div>Loading...</div>

    if (cart.status === StatusFlag.Rejected) return <div>Error loading products</div>

    return (
        <>  
            
            <button onClick={() => getPr}>GEt</button>
            <h1>Cart products</h1>
            {cart.data && cart.data.products.length > 0 ? (
                <ul>
                    {cart.data.products.map(product => (
                        <li key={product.id}>{product.price}</li>
                    ))}
                </ul>
            ) : (
                <div>No products in cart</div>
            )}
        </>
    )
}
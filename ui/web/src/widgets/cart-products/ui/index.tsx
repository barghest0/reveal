import { styled } from "@mui/material";
import { ProductCart } from "entities/cart";
import { getProductsCart } from "entities/cart/model/actions/get-products-cart";
import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";



export const ProductsCartList= () => {
    const dispatch = useAppDispatch();
    const {data, status, loading} = useAppSelector((state) => state.productsCart.cart);

    useEffect(() => {
        dispatch(getProductsCart(1))
    }, [])
    return (
        <CartContainer>
            {data?.products.map(product => (
                
                <ProductCart
                    id={product.product_id}
                    name={product.name}
                    price={product.price}
                    cartId={product.cart_id}
                />
            ))}
        </CartContainer>
    )
}

const CartContainer = styled("div")({
    display: 'grid',
    gridTemplateColumns: "1fr",
    gap: 20    
})
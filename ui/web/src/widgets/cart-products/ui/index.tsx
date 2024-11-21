import { styled } from "@mui/material";
import { CartProduct, getProductsCart } from "entities/product";
import { ButtonDeleteFromCart } from "features/cart";

import { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";



export const ProductsCartList= () => {
    const dispatch = useAppDispatch();
    const {data, status, loading} = useAppSelector((state) => state.productsCart.cart);

    
    
    useEffect(() => {
        dispatch(getProductsCart())
    }, [data?.products])
    
    return (
        <CartContainer>
            {data?.products.map(product => (
                <CartProduct key={product.product_id}
                    id={product.product_id}
                    name={product.name}
                    price={product.price}
                    userId={product.cart_id}
                    deleteButton={<ButtonDeleteFromCart id={product.product_id}/>}
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
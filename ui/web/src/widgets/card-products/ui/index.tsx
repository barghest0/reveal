import { Card, styled } from "@mui/material";
import { CardProduct } from "entities/card-product";
import { getProductsCard } from "entities/cart";
import { AddProductToCart } from "features/button-add-cart";
import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";


export const CardProductsList = () => {
    const dispatch = useAppDispatch();
    const {data, loading, status} = useAppSelector(state => state.productsCart.products);

    useEffect(() => {
        dispatch(getProductsCard())        
    }, [])     
   
    return (
        <ProductsContainer>
            {data?.map(product => (
                <CardProduct key={product.id}
                    id={product.id}
                    name={product.name}
                    price={product.price}
                    ButtonAdd={<AddProductToCart product={product}/>}/>
            ))}
        </ProductsContainer>
    )
}

const ProductsContainer = styled("div")({
        display: 'grid',
        gridTemplateColumns: "repeat(auto-fit, minmax(250px, 1fr))",
        gap: 20
})


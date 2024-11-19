import { Button } from "@mui/material";
import { addProductToCart } from "entities/product/model/actions/add-to-cart";
import { getProductsCart } from "entities/product/model/actions/get-products-cart";
import React, { useEffect } from "react";
import { ProductItem } from "shared/api/products";

import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";

interface Props {
    product?: ProductItem | undefined;
}

export const AddProductToCart: React.FC<Props> = ({product}) => {
    const dispatch = useAppDispatch();
    const userId = useAppSelector(state => state.productsCart.cart.data?.user_id)
    useEffect(() => {
        dispatch(getProductsCart(1))
    }, [])
   
    const handleAddProduct = (userId: number, product: ProductItem) => {
        if (userId)
        dispatch(addProductToCart({userId, product}))
    }

    return (
        <>
            {userId && product && <Button variant="contained" onClick={() => handleAddProduct(userId, product)}>Add to cart</Button>}
        </>
    )
}
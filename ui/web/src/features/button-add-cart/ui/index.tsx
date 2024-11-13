import { Button } from "@mui/material";
import { addProductToCart } from "entities/cart/model/actions/add-to-cart";
import { getProductsCart } from "entities/cart/model/actions/get-products-cart";
import React, { useEffect } from "react";
import { ProductItem } from "shared/api/products";

import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";

interface Props {
    product?: ProductItem | undefined;
}

export const AddProductToCart: React.FC<Props> = ({product}) => {
    const dispatch = useAppDispatch();
    const user_id = useAppSelector(state => state.productsCart)
    useEffect(() => {
        dispatch(getProductsCart(1))
    }, [])
    console.log(user_id, 'button')
    const handleAddProduct = (user_id: number, product: ProductItem) => {
        console.log(user_id, product)
        if (user_id)
        dispatch(addProductToCart({user_id, product}))
    }

    return (
        <>
            {/* {user_id && product && <Button variant="contained" onClick={() => handleAddProduct(user_id, product)}>Add to cart</Button>} */}
        </>
    )
}
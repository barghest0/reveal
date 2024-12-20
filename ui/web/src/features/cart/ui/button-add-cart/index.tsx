import { Button } from "@mui/material";
import { addProductToCart } from "entities/product/model/actions/add-to-cart";
import { getProductsCart } from "entities/product/model/actions/get-products-cart";
import React, { useEffect } from "react";
import { ProductItem } from "shared/api/products";

import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";

type Props = {
    product?: ProductItem | undefined;
}

export const AddProductToCart: React.FC<Props> = ({product}) => {
    const dispatch = useAppDispatch();

    const handleAddProduct = (product: ProductItem) => {
        dispatch(addProductToCart(product))
    }

    return (
        <>
            {product && <Button variant="contained" onClick={() => handleAddProduct(product)}>Add to cart</Button>}
        </>
    )
}
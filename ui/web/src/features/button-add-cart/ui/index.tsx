import { Button } from "@mui/material";
import { addProductToCart } from "entities/cart/model/actions/add-to-cart";
import React from "react";
import { ProductItem } from "shared/api/products";

import { useAppDispatch } from "shared/types/hooks/hook";

interface Props {
    cartId: number;
    product: ProductItem | undefined;
}

export const AddProductToCart: React.FC<Props> = ({cartId, product}) => {
    const dispatch = useAppDispatch();

    const handleAddProduct = (cartId:number, product:ProductItem) => {
        dispatch(addProductToCart({cartId, product}))
    }

    return (
        <>
            {product && <Button variant="contained" onClick={() => handleAddProduct(cartId, product)}>Add to cart</Button>}
        </>
    )
}
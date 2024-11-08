import { Button } from "@mui/material";
import React from "react";
import { addProductToCart } from "shared/api/cart";
import { ProductItem } from "shared/api/cart/types";
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
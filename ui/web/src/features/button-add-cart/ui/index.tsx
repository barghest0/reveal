import { Button } from "@mui/material";
import { ProductItem } from "entities/card-product/model/types";
import React from "react";
import { useAppDispatch } from "shared/types/hooks/hook";

interface Props {
    card: ProductItem | undefined
}

export const AddProductToCart: React.FC<Props> = ({card}) => {
    const dispatch = useAppDispatch();

    const handleAddProduct = (card:ProductItem) => {
        
    }

    return (
        <>
            <Button variant="contained">Add to cart</Button>
        </>
    )
}
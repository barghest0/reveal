import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";
import { getProductsCard } from "../model/actions";
import { Card } from "@mui/material";
import { createUseStyles } from "react-jss";
import { AddProductToCart } from "features/button-add-cart";

interface Props {

}

export const ProductsCard: React.FC<Props> = () => {
    const dispatch = useAppDispatch();
    const {data, loading, status} = useAppSelector(state => state.products);

    useEffect(() => {
        dispatch(getProductsCard())        
    }, [])

    const styles = useStyles();

    
    return (
        <div className={styles.cardWrapper}>
            {data?.map((product) => (
                <Card variant="outlined">
                    <h1>Product id: {product.id}</h1>
                    {/* <AddProductToCart card={product} cartId={1}/> */}
                </Card>
            ))}
            
        </div>
    )
}

const useStyles = createUseStyles({
    cardWrapper: {
        display: 'grid',
        width: '100%',
        gridTemplateColumns: "1fr 1fr 1fr",
        gap: 20
    },
})

import React, { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";
import { Card } from "@mui/material";
import { createUseStyles } from "react-jss";
import { AddProductToCart } from "features/button-add-cart";
import { getProductsCard } from "../model/actions";


interface Props {

}

export const ProductsCard: React.FC<Props> = () => {
    const dispatch = useAppDispatch();
    const {data, loading, status} = useAppSelector(state => state.productsCart.products);

    useEffect(() => {
        dispatch(getProductsCard())        
    }, [])

    const styles = useStyles();

    // console.log(data?.map((item) => item))
    
    return (
        <div className={styles.cardWrapper}>
            {data?.map((product) => (
                <Card variant="outlined">
                    <h1>Product id: {product.id}</h1>
                    <AddProductToCart cartId={1} product={product}/>
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

import React from "react";
import { styled } from '@mui/material/styles';

type CardProps = {
    id: number;
    name: string;
    price: number;
    ButtonAdd: React.ReactNode;
}



export const CardProduct: React.FC<CardProps> = ({id, name, price, ButtonAdd}) => { 
    return (
        <ProductContainer>
            <ImageWrapper src="https://avatars.mds.yandex.net/i?id=ecad8571eba37fec382dd2490e8f24b5_l-5524081-images-thumbs&n=13"/>
            <InfoWrapper>
                <h1>{id}</h1>
                <p>{name}</p>
                <p>{price}</p>
            </InfoWrapper>
            
            <>{ButtonAdd}</>
        </ProductContainer>
    )
}

const ProductContainer = styled("div")({
    display: "grid",
    borderRadius: 8,
    boxShadow: "0px 0px 21px -3px #966a57",
    overflow: "hidden",
    cursor: "pointer"
})

const ImageWrapper = styled("img")({
    width: "100%",
    height: "100%",
    borderRadius: 8
})

const InfoWrapper = styled("div")({
    padding: 8,
    border: 1
})


import React from "react";
import { styled } from "@mui/material";

type ProductCartProps = {
    id: number;
    userId: number;
    name: string;
    price: number;
    deleteButton: React.ReactNode;
}

export const CartProduct: React.FC <ProductCartProps> = ({id, userId, name, price, deleteButton}) => {

    return (
        <ProductContainer>  
            <ImageWrapper src="https://avatars.mds.yandex.net/i?id=ecad8571eba37fec382dd2490e8f24b5_l-5524081-images-thumbs&n=13"/>
            <InfoWrapper>
                <NameTitle>{name} Заглушка</NameTitle>
                <h1>{id}</h1>
                <p>{name}</p>
                <p>{price}</p>
                <DeleteBlock>{deleteButton}</DeleteBlock>
            </InfoWrapper>
        </ProductContainer>
    )
}

const ProductContainer = styled("div")({
    display: "grid",
    gridTemplateColumns: "1fr 1fr",
    borderRadius: 22,
    boxShadow: "0px 0px 21px -3px #966a57",
    overflow: 'hidden',
    cursor: 'pointer'

})

const ImageWrapper = styled("img")({
    width:'100%',
    height: '100%',
    borderRadius: 8
})

const InfoWrapper = styled("div")({
    padding: 8
})

const NameTitle = styled("span")({
    color: "#966a57" 
})

const DeleteBlock = styled("div")({

})
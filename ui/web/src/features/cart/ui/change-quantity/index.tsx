import { Button, styled } from "@mui/material"
import RemoveIcon from '@mui/icons-material/Remove';
import AddIcon from '@mui/icons-material/Add';

type Props = {

}

export const ButtonQuantityToCart: React.FC<Props> = ({}) => {

    return (
        <Container>
            <ButtonWrapper>
                <Button style={buttonStyle} variant="text">
                    <RemoveIcon style={IconStyle}/>
                </Button>
            </ButtonWrapper>

            <TextCount>1</TextCount>

            <ButtonWrapper>
                <Button style={buttonStyle} variant="text">
                    <AddIcon style={IconStyle}/>
                </Button>
            </ButtonWrapper> 
        </Container>
    )
}

const Container = styled("div")({
    display: 'grid',
    gridTemplateColumns: "repeat(auto-fill, minmax(56px, 1fr))",

})

const ButtonWrapper = styled("div")({
    height: 56,
    minWidth: 56,
})

const TextCount = styled("span")({
    alignSelf: 'center',
    textAlign: 'center',
    whiteSpace: 'nowrap',
    fontSize: 16,
    fontFamily: 'GtRegular',
    fontWeight: 300
})



const buttonStyle = {
    minWidth: 56,
    height: 56,
    borderRadius: 12,
    background: "rgba(0, 150, 255, 0.08)",
    color: 'black',
}

const IconStyle = {
    color: "rgba(0, 91, 255, 1)"
}
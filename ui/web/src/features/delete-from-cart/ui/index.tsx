import { IconButton } from "@mui/material"
import DeleteIcon from '@mui/icons-material/Delete';
import { useAppDispatch } from "shared/types/hooks/hook";
import { deleteProductFromCart } from "entities/cart/model/actions/delete-from-cart";

type DeleteFromCartProps = {
    cartId: number;
    id: number;
}

export const ButtonDeleteFromCart:React.FC<DeleteFromCartProps> = ({cartId, id}) => {
    const dispatch = useAppDispatch();
    
    const handleDeleteProduct = (cartId: number, id: number) => {
        dispatch(deleteProductFromCart({cartId, productId:id}))
    }

    return (
        <IconButton 
            aria-label="delete" 
            sx={{width: 50, height: 50}} 
            onClick={() => handleDeleteProduct(cartId, id)}>

            <DeleteIcon fontSize="inherit" sx={{width: "100%", height: "100%"}}/>
        </IconButton>
   ) 
}


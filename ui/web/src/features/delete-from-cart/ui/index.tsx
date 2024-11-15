import { IconButton } from "@mui/material"
import DeleteIcon from '@mui/icons-material/Delete';
import { useAppDispatch } from "shared/types/hooks/hook";
import { deleteProductFromCart } from "entities/cart/model/actions/delete-from-cart";

type DeleteFromCartProps = {
    userId: number;
    id: number;
}

export const ButtonDeleteFromCart:React.FC<DeleteFromCartProps> = ({userId, id}) => {
    const dispatch = useAppDispatch();
    
    const handleDeleteProduct = (userId: number, id: number) => {
        dispatch(deleteProductFromCart({userId, productId:id}))
        
    }

    return (
        <IconButton 
            aria-label="delete" 
            sx={{width: 50, height: 50}} 
            onClick={() => handleDeleteProduct(userId, id)}>

            <DeleteIcon fontSize="inherit" sx={{width: "100%", height: "100%"}}/>
        </IconButton>
   ) 
}


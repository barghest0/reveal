import { IconButton } from "@mui/material"
import DeleteIcon from '@mui/icons-material/Delete';
import { useAppDispatch } from "shared/types/hooks/hook";
import { deleteProductFromCart } from "entities/product/model/actions/delete-from-cart";

type DeleteFromCartProps = {
    id: number;
}

export const ButtonDeleteFromCart:React.FC<DeleteFromCartProps> = ({id}) => {
    const dispatch = useAppDispatch();
    
    const handleDeleteProduct = (id: number) => {
        dispatch(deleteProductFromCart(id))
    }

    return (
        <IconButton 
            aria-label="delete" 
            sx={{width: 50, height: 50}} 
            onClick={() => handleDeleteProduct(id)}>

            <DeleteIcon fontSize="inherit" sx={{width: "100%", height: "100%"}}/>
        </IconButton>
   ) 
}


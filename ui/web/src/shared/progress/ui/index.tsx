import { Box, CircularProgress } from "@mui/material"



type Props = {
    loading: boolean,
    content: string
}


export const ProgressCircular: React.FC<Props> = ({loading, content}) => {
    
    return (
        <Box sx={{ display: 'flex' }}>
            {loading ? <CircularProgress sx={{color: 'red'}}/> : content}
        </Box>
    )
}
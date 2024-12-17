
import { FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput, TextField } from "@mui/material"
import React from "react";


type Props = {
    email?: string;
    setEmail: (email: string) => void;
}

export const LoginField: React.FC<Props> = ({email, setEmail}) => {

    return (
        <FormControl sx={{ width: '100%'}} variant="outlined">
            <InputLabel htmlFor="outlined-adornment-password" sx={{fontFamily: "GtRegular"}}>Email</InputLabel>
            <OutlinedInput
                id="outlined-adornment-password"
                onChange={(e) => setEmail(e.target.value)}
                value={email}
                type='email'
                sx={{fontFamily: "GtRegular"}}
                label="Email"
                required
            />
        </FormControl>
    )
}
import { Visibility, VisibilityOff } from "@mui/icons-material";
import { FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput, TextField } from "@mui/material"
import React from "react";


type Props = {
    password?: string;
    setPassword: (password: string) => void;
}

export const PasswordField: React.FC<Props> = ({password, setPassword}) => {
    const [showPassword, setShowPassword] = React.useState(false);


    const handleClickShowPassword = () => setShowPassword((show) => !show);

    return (
        <>

            <FormControl sx={{ width: '100%'}} variant="outlined">
                <InputLabel htmlFor="outlined-adornment-password" sx={{fontFamily: "GtRegular"}}>Password</InputLabel>
                <OutlinedInput
                    id="outlined-adornment-password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    type={showPassword ? 'text' : 'password'}
                    sx={{fontFamily: "GtRegular"}}
                    endAdornment={
                        <InputAdornment position="end">
                            <IconButton
                                aria-label={
                                    showPassword ? 'hide the password' : 'display the password'
                                }
                                onClick={handleClickShowPassword}
                                edge="end"
                            >
                            {showPassword ? <VisibilityOff /> : <Visibility />}
                            </IconButton>
                        </InputAdornment>
                        }
                    label="Password"
                    required
                />
            </FormControl>
        </>
    )
}
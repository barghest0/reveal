import { Button, styled } from "@mui/material"
import { loginUserAuth } from "entities/user/model/actions/login-user";
import { LoginField, PasswordField } from "features/user"
import React, { useState } from "react"
import { getToken } from "shared/lib/session";
import { useAppDispatch } from "shared/types/hooks/hook";



export const AuthForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const dispatch = useAppDispatch();
    
    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()
       
        // let email = "kakafonia@inbox.ru"
        // let password = "kakafon"
        dispatch(loginUserAuth({email, password}))
        console.log(getToken())
    }

    
    return (
        <Container onSubmit={handleSubmit}>
            <h1>Authorization</h1>
            <LoginField setEmail={setEmail}/>
            <PasswordField setPassword={setPassword}/>
            <Button type="submit">Log in</Button>
            {error && <p>{error}</p>}
        </Container>
    )
}


const Container = styled("form")({
    maxWidth: 320,
    padding: 10,
    border: "1px solid black",
    borderRadius: 22,
    display: 'grid',
    gridTemplateRows: '1fr 1fr 1fr 1fr',
    alignItems: 'center',
    gridRow: 20
})


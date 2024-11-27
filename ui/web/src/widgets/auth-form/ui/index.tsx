import { Button, styled } from "@mui/material"
import { loginUserAuth } from "entities/user/model/actions/login-user";
import { registrationUser } from "entities/user/model/actions/registration-user";
import { LoginField, PasswordField } from "features/user"
import React, { useState } from "react"
import { useAppDispatch } from "shared/types/hooks/hook";



export const AuthForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [reg, setReg] = useState(false);
    const dispatch = useAppDispatch();

   

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        reg 
            ? dispatch(registrationUser({email, password}))
            : dispatch(loginUserAuth({email, password}))
    }

    
    return (
        <>
            {reg 
                ? (
                    <Container onSubmit={handleSubmit}>
                        <h1>Registation</h1>
                        <LoginField setEmail={setEmail}/>
                        <PasswordField setPassword={setPassword}/>

                        <ButtonWrapper>
                            <Button type="submit">Registation</Button>
                            <Button onClick={() => setReg(false)}>Log in</Button>
                        </ButtonWrapper>
                        
                        {error && <p>{error}</p>}
                    </Container>
                ) 
                :
                (
                    <Container onSubmit={handleSubmit}>
                        <h1>Authorization</h1>
                        <LoginField setEmail={setEmail}/>
                        <PasswordField setPassword={setPassword}/>

                        <ButtonWrapper>
                            <Button type="submit">Log in</Button>
                            <Button onClick={() => setReg(true)}>Register</Button>
                        </ButtonWrapper>
                        
                        {error && <p>{error}</p>}
                    </Container>
                )
            }
        </>            
        
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

const ButtonWrapper = styled("div")({
    display: 'flex',
    alignItems: 'center',
    justifyContent: "space-between"
})


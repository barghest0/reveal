
import { Button, styled } from "@mui/material"
import { loginUserAuth } from "entities/user/model/actions/login-user";
import { registrationUser } from "entities/user/model/actions/registration-user";
import { LoginField, PasswordField } from "features/user"
import React, { useState } from "react"
import { ProgressCircular } from "shared/progress/ui";
import { useAppDispatch, useAppSelector } from "shared/types/hooks/hook";
import CloseRoundedIcon from '@mui/icons-material/CloseRounded';


export const AuthForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [reg, setReg] = useState(false);
    const {data, loading, status} = useAppSelector(state => state.user)
    const dispatch = useAppDispatch();

   

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        console.log(reg)
        reg ? dispatch(registrationUser({email, password})) : dispatch(loginUserAuth({email, password}))
    }

    
    
    return (
        <Container>
            {reg ? <h1>Registation</h1> : <h1>Login</h1>}
            <CloseButton><CloseRoundedIcon sx={{width: 20, height: 20}}/></CloseButton>
            {reg 
                ? (
                    <Wrapper onSubmit={handleSubmit}>
                        <FieldsWrapper>
                            <LoginField setEmail={setEmail}/>
                            <PasswordField setPassword={setPassword}/>
                        </FieldsWrapper>
                        

                        <ButtonWrapper>
                            <Button type="submit" variant="contained" sx={{fontFamily: "GtRegular"}}>
                                <ProgressCircular loading={loading} content="Registration"/>    
                            </Button>
                            <SpliteLine>or</SpliteLine>
                            <Button onClick={() => setReg(false)} sx={{fontFamily: "GtRegular"}}>Log in</Button>
                        </ButtonWrapper>
                        
                        {error && <p>{error}</p>}
                    </Wrapper>
                ) 
                :
                (
                    <Wrapper onSubmit={handleSubmit}>
                        <FieldsWrapper>
                            <LoginField setEmail={setEmail}/>
                            <PasswordField setPassword={setPassword}/>
                        </FieldsWrapper>
                        

                        <ButtonWrapper>
                            <Button type="submit" variant="contained" sx={{fontFamily: "GtRegular"}}>
                                <ProgressCircular loading={loading} content="Log in"/>
                            </Button>
                            <SpliteLine>or</SpliteLine>
                            <Button onClick={() => setReg(true)} sx={{fontFamily: "GtRegular"}}>Registration</Button>
                        </ButtonWrapper>
                        
                        {error && <p>{error}</p>}
                    </Wrapper>
                )
            }
        </Container>            
        
    )
}


const Container = styled("div")({
    maxWidth: 320,
    padding: 10,
    border: "1px solid black",
    borderRadius: 22,
    position: 'relative'
})

const Wrapper = styled("form")({
    maxWidth: 320,
    padding: 10,
    display: 'grid',
    gridTemplateRows: '1fr 1fr',
    alignItems: 'center',
    gridRow: 20    
})

const FieldsWrapper = styled("div")({
    display: 'grid',
    gap: 20
})

const ButtonWrapper = styled("div")({
    display: 'grid',
    gap: 20
})

const CloseButton = styled("div")({
    width: 32,
    height: 32,
    borderRadius: 22,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    background: 'rgba(204, 214, 228, 0.4)',
    position: 'absolute',
    top: 10,
    right: 10,
    cursor: 'pointer',
})

const SpliteLine = styled("div")({
    width: "max-content",
    display: 'flex',
    alignItems: 'center',
    position: "relative",
    textAlign: 'center',
    padding: '0, 16px',
    color: 'rgba(0,0,0,0.6)',
    margin: '16px auto',
    '&::before': {
        content: '""',
        height: 1,
        backgroundColor: "rgba(0,26,52,.2)",
        position: "absolute",
        top: 10,
        width: 64,
        left: "-80px"
    },
    '&::after': {
        content: '""',
        height: 1,
        backgroundColor: "rgba(0,26,52,.2)",
        position: "absolute",
        top: 10,
        width: 64,
        right: "-80px"
    }
})
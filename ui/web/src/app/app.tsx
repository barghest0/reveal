import { LoginField, PasswordField } from "features/user";
import { useEffect } from "react";
import { createUseStyles } from "react-jss";
import { Link } from "react-router-dom";
import { clearToken, getToken } from "shared/lib/session";
import { AuthForm } from "widgets/auth-form";

export const App = () => {

    const styles = useStyles();
    return (
        <div className={styles.container}>
            <Link to="/cartPage">
                <p>Go to cart</p>
            </Link>
            <Link to="/mainPage">
                <p>Go to main</p>
            </Link>
            <AuthForm/>
        </div>
    )
}

const useStyles = createUseStyles({
    container: {
        padding: 20,
        rowGap: 20,
        minHeight: "100vh",
        // background: '#756148'
    },
    cards: {
        marginBottom: 20
    }
})
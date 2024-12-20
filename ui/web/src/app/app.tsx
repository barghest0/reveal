
import { GlobalStyles } from "@mui/material";
import { createUseStyles } from "react-jss";
import { Link } from "react-router-dom";
import { AuthForm } from "widgets/auth-form";





export const App = () => {

    const styles = useStyles();
    return (
        <div className={styles.container}>
            {globalStyles}
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
    },
    cards: {
        marginBottom: 20
    }
})
const globalStyles = (
  <GlobalStyles
    styles={(theme) => ({
      html: {
        height: "100%",
      },
      body: {
        margin: 0,
        padding: theme.spacing(2),
        backgroundColor: theme.palette.background.default,
        color: theme.palette.text.primary,
        fontFamily: "GtRegular"
      }
    })}
  />
)
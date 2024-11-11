
import { createUseStyles } from "react-jss";
import { Link } from "react-router-dom";
import { CardProductsList } from "widgets/card-products";
import { ProductsCartList } from "widgets/cart-products";

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
            <ProductsCartList/>
        </div>
    )
}

const useStyles = createUseStyles({
    container: {
        padding: 20,
        rowGap: 20,
    },
    cards: {
        marginBottom: 20
    }
})
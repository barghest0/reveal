import { ProductsCard } from "entities/card-product";
import { ProductsCart } from "entities/cart";
import { createUseStyles } from "react-jss";

export const App = () => {

    const styles = useStyles();
    return (
        <div className={styles.container}>
            <div className={styles.cards}>
                <ProductsCard/>
            </div>
            <ProductsCart/>
        </div>
    )
}

const useStyles = createUseStyles({
    container: {
        padding: 20,
        rowGap: 20,
        backgroundColor: 'green'
    },
    cards: {
        marginBottom: 20
    }
})
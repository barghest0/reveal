package widgets.ProductsList

import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import entities.product.Product
import entities.product.ProductCard
import features.AddToCart.AddToCartButton

@Composable
fun ProductsList(navController: NavController, products: List<Product>) {

  LazyColumn {
    items(products) { product ->
      ProductCard(
              product = product,
              onClick = { navController.navigate("product/${product.id}") }
      ) { AddToCartButton(product) }
    }
  }
}

package widgets.CartProductsList

import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import entities.cart.CartItem
import entities.product.CartProductCard
import features.cart.RemoveFromCartButton

@Composable
fun CartProductsList(navController: NavController, items: List<CartItem>) {

  LazyColumn {
    items(items) { item ->
      CartProductCard(
              product = item.product,
              onClick = { navController.navigate("product/${item.product.id}") }
      ) { RemoveFromCartButton(item.product.id) }
    }
  }
}

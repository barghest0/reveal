package widgets.CartProductsList

import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import entities.product.CartProductCard
import features.cart.CartViewModel
import features.cart.RemoveFromCartButton

@Composable
fun CartProductsList(navController: NavController, viewModel: CartViewModel) {

  LazyColumn {
    items(viewModel.products) { item ->
      CartProductCard(
              product = item.product,
              onClick = { navController.navigate("product/${item.product.id}") }
      ) { RemoveFromCartButton(item.product.id, viewModel) }
    }
  }
}

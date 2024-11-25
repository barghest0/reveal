package features.AddToCart

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.platform.LocalContext
import entities.product.Product
import features.cart.CartRepository
import features.cart.CartViewModel
import shared.session.PreferencesManager

@Composable
fun AddToCartButton(product: Product) {
  val context = LocalContext.current
  val token = remember { PreferencesManager(context).getToken() }
  val viewModel = remember { CartViewModel(CartRepository(token)) }
  val isAddedToCart = viewModel.isProductExist(product.id)

  Button(
          onClick = {
            if (!isAddedToCart) {
              viewModel.addToCart(product)
            }
          },
  ) { Text(if (isAddedToCart) "В корзине" else "Добавить в корзину") }
}

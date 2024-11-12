package features.AddToCart

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import entities.cart.CartRepository
import entities.cart.CartViewModel
import entities.product.Product

@Composable
fun AddToCartButton(product: Product) {

  val viewModel = remember { CartViewModel(CartRepository()) }
  val isAddedToCart = viewModel.isProductExist(product.id)

  Button(
          onClick = {
            if (!isAddedToCart) {
              viewModel.addToCart(product)
            }
          },
  ) { Text(if (isAddedToCart) "В корзине" else "Добавить в корзину") }
}

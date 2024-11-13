package features.RemoveFromCart

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import entities.cart.CartRepository
import entities.cart.CartViewModel

@Composable
fun RemoveFromCartButton(product_id: Int) {

  val viewModel = remember { CartViewModel(CartRepository()) }

  Button(
          onClick = { viewModel.removeFromCart(product_id) },
  ) { Text("Удалить") }
}

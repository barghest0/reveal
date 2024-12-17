package features.cart

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable

@Composable
fun RemoveFromCartButton(product_id: Int, viewModel: CartViewModel) {

  Button(
          onClick = { viewModel.removeFromCart(product_id) },
  ) { Text("Удалить") }
}

package features.RemoveFromCart

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.platform.LocalContext
import entities.cart.CartRepository
import entities.cart.CartViewModel
import shared.session.PreferencesManager

@Composable
fun RemoveFromCartButton(product_id: Int) {

  val context = LocalContext.current
  val token = remember { PreferencesManager(context).getToken() }
  val viewModel = remember { CartViewModel(CartRepository(token)) }

  Button(
          onClick = { viewModel.removeFromCart(product_id) },
  ) { Text("Удалить") }
}

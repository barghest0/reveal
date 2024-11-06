package pages.cart

import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.navigation.NavController
import entities.cart.CartRepository
import entities.cart.CartViewModel
import kotlinx.coroutines.*
import shared.ui.layout.ScreenLayout

@Composable
fun CartScreen(
        navController: NavController,
) {
  val viewModel = remember { CartViewModel(CartRepository()) }

  ScreenLayout {
    LazyColumn {
      items(viewModel.products.toList()) { cart_item -> Text("${cart_item.product_id}") }
    }
  }
}

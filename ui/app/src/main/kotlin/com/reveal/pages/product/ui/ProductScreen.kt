package pages.product

import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.navigation.NavController
import kotlinx.coroutines.*

@Composable
fun ProductScreen(navController: NavController, productId: String?) {
  Text("$productId}")
}

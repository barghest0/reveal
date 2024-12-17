package pages.product

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.navigation.NavController
import entities.product.ProductRepository
import entities.product.ProductViewModel
import kotlinx.coroutines.*
import shared.ui.layout.ScreenLayout

@Composable
fun ProductScreen(navController: NavController, productId: String) {
  val viewModel = remember { ProductViewModel(ProductRepository()) }

  LaunchedEffect(Unit) { viewModel.fetchProduct(productId) }

  val product = viewModel.product.value // Используем collectAsState для получения значения

  ScreenLayout {
    if (product == null) {
      Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
        CircularProgressIndicator()
      }
    } else {
      ProductDetails(product)
    }
  }
}

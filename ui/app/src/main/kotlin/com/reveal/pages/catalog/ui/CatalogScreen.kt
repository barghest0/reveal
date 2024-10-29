package pages.catalog

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.navigation.NavController
import entities.product.ProductRepository
import entities.product.ProductViewModel
import kotlinx.coroutines.*
import shared.ui.layout.ScreenLayout
import widgets.ProductsList.ProductsList

@Composable
fun CatalogScreen(
        navController: NavController,
) {
  val viewModel = ProductViewModel(ProductRepository())

  LaunchedEffect(Unit) {
    viewModel.fetchProducts() // Получаем данные при первом запуске
  }

  ScreenLayout { ProductsList(navController, viewModel.products) }
}

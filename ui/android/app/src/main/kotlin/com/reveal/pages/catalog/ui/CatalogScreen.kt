package pages.catalog

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.remember
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
  val viewModel = remember { ProductViewModel(ProductRepository()) }

  LaunchedEffect(Unit) { viewModel.fetchProducts() }

  ScreenLayout { ProductsList(navController, viewModel.products) }
}

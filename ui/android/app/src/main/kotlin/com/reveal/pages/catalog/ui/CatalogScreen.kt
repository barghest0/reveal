package pages.catalog

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.navigation.NavController
import entities.product.ProductRepository
import entities.product.ProductViewModel
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.*
import kotlinx.serialization.json.Json
import shared.ui.layout.ScreenLayout
import widgets.ProductsList.ProductsList

var client =
        HttpClient(CIO) { install(ContentNegotiation) { json(Json { ignoreUnknownKeys = true }) } }

@Composable
fun CatalogScreen(
        navController: NavController,
) {
  val viewModel = ProductViewModel(ProductRepository(client))

  LaunchedEffect(Unit) {
    viewModel.fetchProducts() // Получаем данные при первом запуске
  }

  ScreenLayout { ProductsList(viewModel.products) }
}

package pages.catalog

import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.navigation.NavController
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.*
import shared.ui.layout.ScreenLayout

var client = HttpClient(CIO) { install(ContentNegotiation) { json() } }

@Composable
fun CatalogScreen(
        navController: NavController,
) {

  ScreenLayout { Text(text = "Каталог") }
}

package app

import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import pages.catalog.CatalogScreen
import pages.login.LoginScreen
import pages.product.ProductScreen
import pages.profile.ProfileScreen
import pages.registration.RegistrationScreen
import widgets.NavigationBar.NavigationBar

@Composable
fun AppNavigation(navController: NavHostController) {
  Surface(color = MaterialTheme.colorScheme.background) {
    Scaffold(
            modifier = Modifier.fillMaxSize().padding(top = 16.dp),
            bottomBar = { NavigationBar(navController) }
    ) { _innerPadding ->
      NavHost(navController = navController, startDestination = "catalog") {
        composable("catalog") { CatalogScreen(navController) }
        composable("registration") { RegistrationScreen(navController) }
        composable("profile") { ProfileScreen(navController) }
        composable("login") { LoginScreen(navController) }
        composable("product/{productId}") { backStackEntry ->
          val productId = backStackEntry.arguments?.getString("productId")
          ProductScreen(navController, productId)
        }
      }
    }
  }
}

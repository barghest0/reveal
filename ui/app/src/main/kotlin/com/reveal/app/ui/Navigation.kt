package app

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.runtime.Composable
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import pages.login.LoginScreen
import pages.profile.ProfileScreen
import pages.registration.RegistrationScreen
import widgets.NavigationBar

@Composable
fun AppNavigation(navController: NavHostController) {
  Surface(color = MaterialTheme.colorScheme.background) {
    Scaffold(bottomBar = { NavigationBar(navController) }) { innerPadding ->
      NavHost(navController = navController, startDestination = "registration") {
        composable("registration") { RegistrationScreen(navController) }
        composable("login") { LoginScreen(navController) }
        composable("profile") { ProfileScreen(navController) }
      }
    }
  }
}

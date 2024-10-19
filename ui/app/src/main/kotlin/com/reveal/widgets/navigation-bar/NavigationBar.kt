package widgets

import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.navigation.NavHostController

@Composable
fun NavigationBar(navController: NavHostController) {
  NavigationBar {
    NavigationBarItem(
            icon = { /* Ваша иконка */},
            label = { Text("registration") },
            selected =
                    false, // Здесь можно использовать состояние для выделения выбранного элемента
            onClick = { navController.navigate("registration") }
    )
    NavigationBarItem(
            icon = { /* Ваша иконка */},
            label = { Text("login") },
            selected = false,
            onClick = { navController.navigate("login") }
    )
    NavigationBarItem(
            icon = { /* Ваша иконка */},
            label = { Text("Profile") },
            selected = false,
            onClick = { navController.navigate("profile") }
    )
  }
}

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
            label = { Text("Каталог") },
            selected =
                    false, // Здесь можно использовать состояние для выделения выбранного элемента
            onClick = { navController.navigate("catalog") }
    )

    NavigationBarItem(
            icon = { /* Ваша иконка */},
            label = { Text("Profile") },
            selected = false,
            onClick = { navController.navigate("profile") }
    )
  }
}

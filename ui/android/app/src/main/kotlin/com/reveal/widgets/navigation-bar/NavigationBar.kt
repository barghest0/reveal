package widgets.NavigationBar

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.navigation.NavHostController

@Composable
fun NavigationBar(navController: NavHostController) {
  NavigationBar(modifier = Modifier.height(80.dp)) {
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

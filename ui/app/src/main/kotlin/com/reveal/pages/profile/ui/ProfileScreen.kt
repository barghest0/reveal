package pages.profile

import android.content.Context
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import entities.user.UserRepository
import features.profile.ProfileViewModel
import features.profile.ProfileViewModelFactory
import kotlinx.coroutines.*
import shared.session.PreferencesManager
import shared.ui.layout.ScreenLayout

@Composable
fun ProfileScreen(
        navController: NavController,
        context: Context = LocalContext.current,
        tokenManager: PreferencesManager = PreferencesManager(context),
        viewModel: ProfileViewModel =
                viewModel(factory = ProfileViewModelFactory(UserRepository(), tokenManager))
) {

  val profile by viewModel.profileState

  ScreenLayout {
    if (profile != null) {
      Text(text = "Имя: ${profile?.name}")
    }
    Button(
            onClick = {
              tokenManager.clearToken()
              navController.navigate("login")
            }
    ) { Text("Logout") }
  }
}

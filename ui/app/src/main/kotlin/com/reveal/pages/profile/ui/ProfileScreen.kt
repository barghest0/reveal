package pages.profile

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import features.profile.ProfileViewModel
import features.profile.ProfileViewModelFactory
import features.profile.UserProfileRepository
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.*
import shared.session.PreferencesManager

var client = HttpClient(CIO) { install(ContentNegotiation) { json() } }

@Composable
fun ProfileScreen(
        navController: NavController,
        viewModel: ProfileViewModel =
                viewModel(
                        factory =
                                ProfileViewModelFactory(
                                        UserProfileRepository(
                                                client,
                                                PreferencesManager(LocalContext.current)
                                        ),
                                )
                )
) {

  val profile by viewModel.profileState

  Box(modifier = Modifier.fillMaxSize()) {
    if (profile != null) {
      Text(text = "Имя: ${profile?.name}")
      // Отобразите другую информацию профиля
    } else {
      Text(text = "Загрузка профиля...")
    }
  }
}

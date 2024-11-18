package pages.profile

import android.content.Context
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
        viewModel: ProfileViewModel =
                viewModel(
                        factory =
                                ProfileViewModelFactory(
                                        UserRepository(),
                                        PreferencesManager(context)
                                )
                )
) {

  val profile by viewModel.profileState

  ScreenLayout {
    if (profile != null) {
      Text(text = "Имя: ${profile?.name}")
    }
  }
}

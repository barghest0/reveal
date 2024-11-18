package features.profile

import androidx.compose.runtime.State
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import entities.user.UserProfile
import entities.user.UserRepository
import kotlinx.coroutines.launch
import shared.session.PreferencesManager

class ProfileViewModel(
        private val userRepository: UserRepository,
        private val tokenManager: PreferencesManager
) : ViewModel() {

  private val _profileState = mutableStateOf<UserProfile?>(null)
  val profileState: State<UserProfile?> = _profileState

  init {
    fetchUserProfile()
  }

  private fun fetchUserProfile() {
    viewModelScope.launch {
      var token = tokenManager.getToken()
      if (token != null) {
        val profile = userRepository.getProfile(token)
        _profileState.value = profile
      }
    }
  }
}

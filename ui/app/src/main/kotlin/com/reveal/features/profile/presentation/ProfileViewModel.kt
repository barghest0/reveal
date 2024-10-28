package features.profile

import androidx.compose.runtime.State
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import entities.user.UserProfile
import entities.user.UserRepository
import kotlinx.coroutines.launch

class ProfileViewModel(private val userRepository: UserRepository) : ViewModel() {

  private val _profileState = mutableStateOf<UserProfile?>(null)
  val profileState: State<UserProfile?> = _profileState

  init {
    fetchUserProfile()
  }

  private fun fetchUserProfile() {
    viewModelScope.launch {
//      val profile = userRepository.getProfile()
//      _profileState.value = profile
    }
  }
}

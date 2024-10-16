import androidx.compose.runtime.State
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.reveal.entities.user.model.UserProfile
import kotlinx.coroutines.launch

class ProfileViewModel(private val userProfileRepository: UserProfileRepository) : ViewModel() {

  private val _profileState = mutableStateOf<UserProfile?>(null)
  val profileState: State<UserProfile?> = _profileState

  init {
    fetchUserProfile()
  }

  private fun fetchUserProfile() {
    viewModelScope.launch {
      val profile = userProfileRepository.getProfile()
      _profileState.value = profile
    }
  }
}

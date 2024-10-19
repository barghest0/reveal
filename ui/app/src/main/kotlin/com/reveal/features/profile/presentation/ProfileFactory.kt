package features.profile 

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import features.profile.ProfileViewModel

class ProfileViewModelFactory(private val userProfileRepository: UserProfileRepository) :
        ViewModelProvider.Factory {
  override fun <T : ViewModel> create(modelClass: Class<T>): T {
    if (modelClass.isAssignableFrom(ProfileViewModel::class.java)) {
      @Suppress("UNCHECKED_CAST") return ProfileViewModel(userProfileRepository) as T
    }
    throw IllegalArgumentException("Unknown ViewModel class")
  }
}

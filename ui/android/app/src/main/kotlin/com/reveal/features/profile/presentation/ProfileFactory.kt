package features.profile

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import entities.user.UserRepository

class ProfileViewModelFactory(private val userRepository: UserRepository) :
        ViewModelProvider.Factory {
  override fun <T : ViewModel> create(modelClass: Class<T>): T {
    if (modelClass.isAssignableFrom(ProfileViewModel::class.java)) {
      @Suppress("UNCHECKED_CAST") return ProfileViewModel(userRepository) as T
    }
    throw IllegalArgumentException("Unknown ViewModel class")
  }
}

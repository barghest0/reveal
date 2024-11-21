<<<<<<< HEAD
package features.profile 

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import features.profile.ProfileViewModel

class ProfileViewModelFactory(private val userProfileRepository: UserProfileRepository) :
        ViewModelProvider.Factory {
  override fun <T : ViewModel> create(modelClass: Class<T>): T {
    if (modelClass.isAssignableFrom(ProfileViewModel::class.java)) {
      @Suppress("UNCHECKED_CAST") return ProfileViewModel(userProfileRepository) as T
=======
package features.profile

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import entities.user.UserRepository
import shared.session.PreferencesManager

class ProfileViewModelFactory(
        private val userRepository: UserRepository,
        private val tokenManager: PreferencesManager
) : ViewModelProvider.Factory {
  override fun <T : ViewModel> create(modelClass: Class<T>): T {
    if (modelClass.isAssignableFrom(ProfileViewModel::class.java)) {
      @Suppress("UNCHECKED_CAST") return ProfileViewModel(userRepository, tokenManager) as T
>>>>>>> cart-ui
    }
    throw IllegalArgumentException("Unknown ViewModel class")
  }
}

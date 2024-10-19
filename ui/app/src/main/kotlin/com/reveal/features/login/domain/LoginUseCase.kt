package features.login

import entities.user.UserRepository

class LoginUseCase(private val userRepository: UserRepository) {
  suspend fun login(name: String, password: String): String? {
    if (name.isEmpty() || password.isEmpty()) {
      return null
    }
    return userRepository.login(name, password)
  }
}

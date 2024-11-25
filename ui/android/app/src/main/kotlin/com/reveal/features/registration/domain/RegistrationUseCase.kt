package features.registration

import entities.user.User
import entities.user.UserRepository

class RegistrationUseCase(private val userRepository: UserRepository) {
  suspend fun execute(name: String, email: String, password: String): Boolean {
    if (name.isEmpty() || email.isEmpty() || password.isEmpty()) {
      return false
    }
    var user = User(name, email, password)
    return userRepository.register(user)
  }
}

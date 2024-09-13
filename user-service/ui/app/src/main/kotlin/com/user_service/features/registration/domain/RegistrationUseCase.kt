package com.user_service.features.registration.domain

import com.user_service.entities.user.data.model.User
import com.user_service.entities.user.data.model.repository.UserRepository

class RegistrationUseCase(private val userRepository: UserRepository) {
  fun execute(name: String, email: String, password: String): Boolean {
    if (name.isEmpty() || email.isEmpty() || password.isEmpty()) {
      return false
    }
    var user = User(name, email, password)
    return userRepository.register(user)
  }
}

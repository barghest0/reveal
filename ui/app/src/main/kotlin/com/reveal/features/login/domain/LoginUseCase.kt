package com.reveal.features.login.domain

import com.reveal.entities.user.model.repository.UserRepository

class LoginUseCase(private val userRepository: UserRepository) {
  suspend fun login(name: String, password: String): String? {
    if (name.isEmpty() || password.isEmpty()) {
      return null
    }
    return userRepository.login(name, password)
  }
}

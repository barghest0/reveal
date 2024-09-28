package com.user_service.features.login.domain

import com.user_service.entities.user.model.User
import com.user_service.entities.user.model.repository.UserRepository

class LoginUseCase(private val userRepository: UserRepository) {
  suspend fun login(name: String, password: String): Boolean {
    if (name.isEmpty() || password.isEmpty()) {
      return false
    }
    return userRepository.login(name,password)
  }

}

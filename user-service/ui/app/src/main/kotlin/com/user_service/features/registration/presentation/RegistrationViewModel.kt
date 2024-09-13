package com.user_service.features.registration.presentation

import androidx.compose.runtime.mutableStateOf
import com.user_service.features.registration.domain.RegistrationUseCase

class RegistrationViewModel(private val registrationUseCase: RegistrationUseCase) {
  val uiState = mutableStateOf(RegistrationUiState())

  fun register(name: String, email: String, password: String) {
    val isSuccess = registrationUseCase.execute(name, email, password)
    if (isSuccess) {
      uiState.value = uiState.value.copy(name, email, password, error = null)
    } else {
      uiState.value = uiState.value.copy(error = "Registration failed")
    }
  }
}

data class RegistrationUiState(
        val name: String = "",
        val email: String = "",
        val password: String = "",
        val error: String? = null
)

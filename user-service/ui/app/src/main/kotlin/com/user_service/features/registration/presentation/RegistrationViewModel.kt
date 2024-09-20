package com.user_service.features.registration.presentation

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.user_service.features.registration.domain.RegistrationUseCase
import kotlinx.coroutines.isActive
import kotlinx.coroutines.launch

class RegistrationViewModel(private val registrationUseCase: RegistrationUseCase) : ViewModel() {
  val uiState = mutableStateOf(RegistrationUiState())

  fun register(name: String, email: String, password: String) {
    viewModelScope.launch {
      if (isActive) {

        val isSuccess = registrationUseCase.execute(name, email, password)
        uiState.value = uiState.value.copy(error = isSuccess)
        // if (isSuccess) {
        //   uiState.value = uiState.value.copy(error = null)
        // } else {
        //   uiState.value = uiState.value.copy(error = "Registration failed")
        // }
      }
    }
  }

  fun updateField(field: String, value: String) {
    uiState.value =
        when (field) {
          "name" -> uiState.value.copy(name = value)
          "email" -> uiState.value.copy(email = value)
          "password" -> uiState.value.copy(password = value)
          else -> uiState.value
        }
  }
}

data class RegistrationUiState(
    val name: String = "",
    val email: String = "",
    val password: String = "",
    val error: String? = null
)

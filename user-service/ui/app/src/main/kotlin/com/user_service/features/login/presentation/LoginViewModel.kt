package com.user_service.features.login.presentation

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.user_service.features.login.domain.LoginUseCase
import kotlinx.coroutines.isActive
import kotlinx.coroutines.launch

class LoginViewModel(private val useCase: LoginUseCase) : ViewModel() {
  val uiState = mutableStateOf(LoginUiState())

  fun login(name: String, password: String) {
    viewModelScope.launch {
      if (isActive) {
        val isSuccess = useCase.login(name, password)
        uiState.value =
            uiState.value.copy(error = if (isSuccess) null else "Login failed", success = isSuccess)
      }
    }
  }

  fun updateField(field: String, value: String) {
    uiState.value =
        when (field) {
          "name" -> uiState.value.copy(name = value)
          "password" -> uiState.value.copy(password = value)
          else -> uiState.value
        }
  }
}

data class LoginUiState(
    val name: String = "",
    val password: String = "",
    val error: String? = null,
    val success: Boolean = false
)

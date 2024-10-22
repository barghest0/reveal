package features.login

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.isActive
import kotlinx.coroutines.launch
import shared.session.PreferencesManager

class LoginViewModel(
        private val useCase: LoginUseCase,
        val preferencesManager: PreferencesManager
) : ViewModel() {
  val uiState = mutableStateOf(LoginUiState())

  fun login(name: String, password: String) {
    viewModelScope.launch {
      if (isActive) {
        val token = useCase.login(name, password)
        if (token != null) {
          uiState.value = uiState.value.copy(success = true)

          preferencesManager.saveToken(token)
        } else {
          uiState.value = uiState.value.copy(error = "Login failed")
        }
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

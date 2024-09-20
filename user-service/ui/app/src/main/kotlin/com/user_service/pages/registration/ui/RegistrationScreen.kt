package com.user_service.pages.registration.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.tooling.preview.Preview
import com.user_service.entities.user.data.model.repository.UserRepository
import com.user_service.features.registration.domain.RegistrationUseCase
import com.user_service.features.registration.presentation.RegistrationViewModel
import com.user_service.features.registration.ui.RegistrationForm

@Composable
fun RegistrationScreen(
        viewModel: RegistrationViewModel = remember {
          RegistrationViewModel(RegistrationUseCase(UserRepository()))
        }
) {
  val state by viewModel.uiState

  Column {
    RegistrationForm(
            name = state.name,
            email = state.email,
            password = state.password,
            updateField = {field, value -> viewModel.updateField(field,value)},
            onRegisterClicked = { name, email, password ->
              viewModel.register(name, email, password)
            }
    )

    // Отображение сообщений об ошибках
    state.error?.let { Text(text = it, color = MaterialTheme.colorScheme.error) }
  }
}


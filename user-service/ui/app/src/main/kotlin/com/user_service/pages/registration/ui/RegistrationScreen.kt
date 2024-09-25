package com.user_service.pages.registration.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.navigation.NavHostController
import com.user_service.entities.user.model.repository.UserRepository
import com.user_service.features.registration.domain.RegistrationUseCase
import com.user_service.features.registration.presentation.RegistrationViewModel
import com.user_service.features.registration.ui.RegistrationForm
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.serialization.kotlinx.json.json

@Composable
fun RegistrationScreen(
    navController: NavHostController,
    viewModel: RegistrationViewModel = remember {
      RegistrationViewModel(
          RegistrationUseCase(
              UserRepository(HttpClient(CIO) { install(ContentNegotiation) { json() } })))
    }
) {
  val state by viewModel.uiState

  Box(modifier = Modifier.fillMaxSize()) {
    Column(
        modifier = Modifier.align(Alignment.Center),
        horizontalAlignment = Alignment.CenterHorizontally) {
          RegistrationForm(
              name = state.name,
              email = state.email,
              password = state.password,
              updateField = { field, value -> viewModel.updateField(field, value) },
              onRegisterClicked = { name, email, password ->
                viewModel.register(name, email, password)
              })

          if (state.success) {
            LaunchedEffect(Unit) {
              navController.navigate("login") { popUpTo("registration") { inclusive = true } }
            }
          }

          state.error?.let { Text(text = it, color = MaterialTheme.colorScheme.error) }
        }
  }
}

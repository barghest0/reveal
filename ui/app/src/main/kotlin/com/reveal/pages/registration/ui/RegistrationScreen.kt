package pages.registration

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.text.ClickableText
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.TextDecoration
import androidx.compose.ui.text.withStyle
import androidx.navigation.NavHostController
import entities.user.UserRepository
import features.registration.RegistrationForm
import features.registration.RegistrationUseCase
import features.registration.RegistrationViewModel
import shared.ui.layout.ScreenLayout

@Composable
fun RegistrationScreen(
        navController: NavHostController,
        viewModel: RegistrationViewModel = remember {
          RegistrationViewModel(RegistrationUseCase(UserRepository()))
        }
) {
  val state by viewModel.uiState

  ScreenLayout {
    Column() {
      RegistrationForm(
              name = state.name,
              email = state.email,
              password = state.password,
              updateField = { field, value -> viewModel.updateField(field, value) },
              onClickRegister = { name, email, password ->
                viewModel.register(name, email, password)
              }
      )

      val annotatedText = buildAnnotatedString {
        withStyle(style = SpanStyle(color = Color.White)) { append("Уже есть аккаунт? ") }
        pushStringAnnotation(tag = "login", annotation = "login")
        withStyle(
                style = SpanStyle(color = Color.White, textDecoration = TextDecoration.Underline)
        ) { append("Войти") }
        pop()
      }

      ClickableText(
              text = annotatedText,
              onClick = { offset ->
                annotatedText
                        .getStringAnnotations(tag = "login", start = offset, end = offset)
                        .firstOrNull()
                        ?.let {
                          navController.navigate("login") // Переход на экран логина
                        }
              }
      )

      if (state.success) {
        LaunchedEffect(Unit) {
          navController.navigate("login") { popUpTo("registration") { inclusive = true } }
        }
      }

      state.error?.let { Text(text = it, color = MaterialTheme.colorScheme.error) }
    }
  }
}

package pages.login

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.text.ClickableText
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.TextDecoration
import androidx.compose.ui.text.withStyle
import androidx.navigation.NavHostController
import entities.user.UserRepository
import features.login.LoginForm
import features.login.LoginUseCase
import features.login.LoginViewModel
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.serialization.kotlinx.json.json
import shared.session.PreferencesManager
import shared.ui.layout.ScreenLayout

var client = HttpClient(CIO) { install(ContentNegotiation) { json() } }

@Composable
fun LoginScreen(navController: NavHostController) {
  val context = LocalContext.current
  val viewModel: LoginViewModel = remember {
    LoginViewModel(LoginUseCase(UserRepository(client)), PreferencesManager(context))
  }

  val state by viewModel.uiState

  ScreenLayout {
    Column() {
      LoginForm(
              name = state.name,
              password = state.password,
              updateField = { field, value -> viewModel.updateField(field, value) },
              onClickLogin = { name, password -> viewModel.login(name, password) }
      )

      val annotatedText = buildAnnotatedString {
        withStyle(style = SpanStyle(color = Color.White)) { append("Нет аккаунта? ") }
        pushStringAnnotation(tag = "registration", annotation = "registration")
        withStyle(
                style = SpanStyle(color = Color.White, textDecoration = TextDecoration.Underline)
        ) { append("Регистрация") }
        pop()
      }

      ClickableText(
              text = annotatedText,
              onClick = { offset ->
                annotatedText
                        .getStringAnnotations(tag = "registration", start = offset, end = offset)
                        .firstOrNull()
                        ?.let {
                          navController.navigate("registration") // Переход на экран регистрации
                        }
              }
      )

      if (state.success) {
        LaunchedEffect(Unit) {
          navController.navigate("profile") { popUpTo("login") { inclusive = true } }
        }
      }

      state.error?.let { Text(text = it, color = MaterialTheme.colorScheme.error) }
    }
  }
}

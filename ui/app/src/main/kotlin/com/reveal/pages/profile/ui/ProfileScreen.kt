package pages.profile

import androidx.compose.foundation.text.ClickableText
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.TextDecoration
import androidx.compose.ui.text.withStyle
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import entities.user.UserRepository
import features.profile.ProfileViewModel
import features.profile.ProfileViewModelFactory
import kotlinx.coroutines.*
import shared.ui.layout.ScreenLayout

@Composable
fun ProfileScreen(
        navController: NavController,
        viewModel: ProfileViewModel = viewModel(factory = ProfileViewModelFactory(UserRepository()))
) {

  val profile by viewModel.profileState

  ScreenLayout {
    if (profile != null) {
      Text(text = "Имя: ${profile?.name}")
      // Отобразите другую информацию профиля
    } else {
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
    }
  }
}

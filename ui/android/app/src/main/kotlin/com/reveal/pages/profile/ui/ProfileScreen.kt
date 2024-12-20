package pages.profile

<<<<<<< HEAD
=======
<<<<<<< HEAD
import androidx.compose.foundation.text.ClickableText
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.TextDecoration
import androidx.compose.ui.text.withStyle
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import features.profile.ProfileViewModel
import features.profile.ProfileViewModelFactory
import features.profile.UserProfileRepository
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.serialization.kotlinx.json.json
=======
>>>>>>> web-auth
import android.content.Context
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import entities.user.UserRepository
import features.profile.ProfileViewModel
import features.profile.ProfileViewModelFactory
<<<<<<< HEAD
=======
>>>>>>> cart-ui
>>>>>>> web-auth
import kotlinx.coroutines.*
import shared.session.PreferencesManager
import shared.ui.layout.ScreenLayout

<<<<<<< HEAD
=======
<<<<<<< HEAD
var client = HttpClient(CIO) { install(ContentNegotiation) { json() } }

@Composable
fun ProfileScreen(
        navController: NavController,
        viewModel: ProfileViewModel =
                viewModel(
                        factory =
                                ProfileViewModelFactory(
                                        UserProfileRepository(
                                                client,
                                                PreferencesManager(LocalContext.current)
                                        ),
                                )
                )
=======
>>>>>>> web-auth
@Composable
fun ProfileScreen(
        navController: NavController,
        context: Context = LocalContext.current,
        tokenManager: PreferencesManager = PreferencesManager(context),
        viewModel: ProfileViewModel =
                viewModel(factory = ProfileViewModelFactory(UserRepository(), tokenManager))
<<<<<<< HEAD
=======
>>>>>>> cart-ui
>>>>>>> web-auth
) {

  val profile by viewModel.profileState

  ScreenLayout {
    if (profile != null) {
      Text(text = "Имя: ${profile?.name}")
<<<<<<< HEAD
=======
<<<<<<< HEAD
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
=======
>>>>>>> web-auth
    }
    Button(
            onClick = {
              tokenManager.clearToken()
              navController.navigate("login")
            }
    ) { Text("Logout") }
<<<<<<< HEAD
=======
>>>>>>> cart-ui
>>>>>>> web-auth
  }
}

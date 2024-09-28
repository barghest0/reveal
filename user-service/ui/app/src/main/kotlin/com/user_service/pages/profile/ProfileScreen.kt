import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.mutableStateOf
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.navigation.NavHostController
import com.user_service.entities.user.model.User
import com.user_service.entities.user.model.repository.UserRepository
import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.plugins.cookies.HttpCookies
import io.ktor.client.plugins.cookies.cookies
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.*

@Composable
fun ProfileScreen(
    navController: NavHostController,
) {
  val user = mutableStateOf<User?>(null)

  LaunchedEffect(Unit) {
    val client =
        HttpClient(CIO) {
          install(ContentNegotiation) { json() }
          install(HttpCookies) {}
        }
    val repository = UserRepository(client)
    val profile = repository.getProfile()
    
    println("user.value ${profile}")
    if(profile !=null){

      user.value = user.value?.copy(name = profile.name)
      }
  }
  if (user.value != null) {
    Box(modifier = Modifier.fillMaxSize()) {
      Column(
          modifier = Modifier.align(Alignment.Center),
          horizontalAlignment = Alignment.CenterHorizontally) {
            Text("${user.value?.name}")
          }
    }
  }
}

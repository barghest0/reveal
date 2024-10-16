import com.reveal.entities.user.model.UserProfile
import com.reveal.shared.session.PreferencesManager
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.request.get
import io.ktor.client.request.header
import io.ktor.http.HttpHeaders
import io.ktor.util.InternalAPI

class UserProfileRepository(
        private val client: HttpClient,
        private val tokenManager: PreferencesManager
) {

  @OptIn(InternalAPI::class)
  suspend fun getProfile(): UserProfile? {
    return try {
      val token = tokenManager.getToken()
      val response =
              client.get("http://192.168.3.2/users/profile") {
                header(HttpHeaders.Authorization, "Bearer $token")
              }
      val string = response.body<String>()
      println("string $string")
      val body = response.body<UserProfile>()
      println("response body $body")
      body
    } catch (exception: Exception) {
      println("exception $exception")
      null
    }
  }
}

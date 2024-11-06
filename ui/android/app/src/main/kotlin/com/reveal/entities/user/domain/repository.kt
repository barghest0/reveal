package entities.user

import io.ktor.client.call.body
import io.ktor.client.request.get
import io.ktor.client.request.post
import io.ktor.http.ContentType
import io.ktor.http.contentType
import io.ktor.util.InternalAPI
import kotlinx.serialization.Serializable
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import shared.api.HTTPClient

class UserRepository() {
  // @OptIn(InternalAPI::class)
  // suspend fun getProfile(): UserProfile? {
  //   return try {
  //     val tokenManager = PreferencesManager(LocalContext.current)
  //     val token = tokenManager.getToken()
  //     val response =
  //             HTTPClient.client.get("http://192.168.3.2/users/profile") {
  //               header(HttpHeaders.Authorization, "Bearer $token")
  //             }
  //     val string = response.body<String>()
  //     println("string $string")
  //     val body = response.body<UserProfile>()
  //     println("response body $body")
  //     body
  //   } catch (exception: Exception) {
  //     println("exception $exception")
  //     null
  //   }
  // }
  @OptIn(InternalAPI::class)
  suspend fun register(user: User): Boolean {
    return try {
      val response =
              HTTPClient.client.post("http://192.168.3.2/users/register") {
                contentType(ContentType.Application.Json)
                body = Json.encodeToString(user)
              }

      val text = response.body<String>()

      println("register ${response} ${text}")
      true
    } catch (exception: Exception) {

      println(exception)
      false
    }
  }

  @OptIn(InternalAPI::class)
  suspend fun login(name: String, password: String): String? {
    return try {

      val response =
              HTTPClient.client.post("http://192.168.3.2/users/login") {
                contentType(ContentType.Application.Json)
                body = Json.encodeToString(Credentials(name, password))
              }

      response.headers["Authorization"]?.removePrefix("Bearer ")
    } catch (exception: Exception) {
      println(exception)
      return null
    }
  }
}

@Serializable data class Credentials(val name: String, val password: String)

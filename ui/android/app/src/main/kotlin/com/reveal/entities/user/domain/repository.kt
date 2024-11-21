package entities.user

import io.ktor.client.call.body
import io.ktor.client.request.get
import io.ktor.client.request.header
import io.ktor.client.request.post
import io.ktor.client.statement.bodyAsText
import io.ktor.http.ContentType
import io.ktor.http.HttpHeaders
import io.ktor.http.contentType
import io.ktor.util.InternalAPI
import kotlinx.serialization.Serializable
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import shared.api.HTTPClient

class UserRepository() {
  @OptIn(InternalAPI::class)
  suspend fun getProfile(token: String): UserProfile? {
    return try {
      val response =
              HTTPClient.client.get("http://192.168.3.2/users/profile") {
                header(HttpHeaders.Authorization, "Bearer $token")
              }
      val body = Json.decodeFromString<UserProfile>(response.bodyAsText())
      println("response body $body")
      body
    } catch (exception: Exception) {
      println("exception $exception")
      null
    }
  }
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
      println("LOGIN RESPONSE ${response}")

      response.headers["Authorization"]?.removePrefix("Bearer ")
    } catch (exception: Exception) {
      println("LOGIN EXCEPTION ${exception}")
      return null
    }
  }
}

@Serializable data class Credentials(val name: String, val password: String)

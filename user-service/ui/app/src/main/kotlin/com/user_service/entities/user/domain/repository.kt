package com.user_service.entities.user.model.repository

import com.user_service.entities.user.model.User
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.request.get
import io.ktor.client.request.post
import io.ktor.http.ContentType
import io.ktor.http.contentType
import io.ktor.util.InternalAPI
import kotlinx.serialization.Serializable
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

class UserRepository(private val client: HttpClient) {
  @OptIn(InternalAPI::class)
  suspend fun register(user: User): Boolean {
    return try {
      val response =
          client.post("http://192.168.3.2:8080/register") {
            contentType(ContentType.Application.Json)
            body = Json.encodeToString(user)
          }

      val text = response.body<String>()

      println(text)
      true
    } catch (exception: Exception) {

      println(exception)
      false
    }
  }

  @OptIn(InternalAPI::class)
  suspend fun login(name: String, password: String): Boolean {
    return try {

      val response =
          client.post("http://192.168.3.2:8080/login") {
            contentType(ContentType.Application.Json)
            body = Json.encodeToString(LoginRequest(name, password))
          }

      val text = response.body<String>()

      println(text)
      true
    } catch (exception: Exception) {

      println(exception)
      false
    }
  }

  @OptIn(InternalAPI::class)
  suspend fun getProfile(): User? {
    return try {
      val response = client.get("http://192.168.3.2:8080/profile") {}
      println("response $response")
      response.body<User>()
    } catch (exception: Exception) {
      println("exception $exception")
      null
    }
  }
}

@Serializable data class LoginRequest(val name: String, val password: String)

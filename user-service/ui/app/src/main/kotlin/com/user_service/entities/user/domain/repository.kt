package com.user_service.entities.user.model.repository

import com.user_service.entities.user.model.User
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.request.post
import io.ktor.http.ContentType
import io.ktor.http.contentType
import io.ktor.util.InternalAPI
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
}

package com.user_service.entities.user.data.model.repository

import com.user_service.entities.user.data.model.User
import io.ktor.client.HttpClient
import io.ktor.client.request.get
import io.ktor.client.statement.HttpResponse
import io.ktor.client.statement.bodyAsText

class UserRepository(private val client: HttpClient) {
  suspend fun register(user: User): String {
    try {
      client.use {
        val response: HttpResponse =
            client.get("http://10.0.2.2:8080/users/2")

        val text = response.bodyAsText()

        println(text)

        return text
      }
    } catch (exception: Exception) {
      println(exception)

      return "${exception}"
    } finally {

      client.close()
    }
  }
}

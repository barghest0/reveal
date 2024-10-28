package shared.api

import io.ktor.client.HttpClient
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.get
import kotlinx.serialization.*
import kotlinx.serialization.json.Json

object HTTPClient {
  val client: HttpClient by lazy {
    HttpClient(CIO) { install(ContentNegotiation) { (Json { ignoreUnknownKeys = true }) } }
  }
}

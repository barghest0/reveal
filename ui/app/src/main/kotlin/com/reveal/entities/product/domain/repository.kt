package entities.product

import io.ktor.client.HttpClient
import io.ktor.client.request.get
import io.ktor.client.statement.bodyAsText
import io.ktor.util.InternalAPI
import kotlinx.serialization.json.Json

class ProductRepository(private val client: HttpClient) {
  @OptIn(InternalAPI::class)
  suspend fun getAllProducts(): List<Product>? {
    return try {
      val response = client.get("http://192.168.3.2/products")

      val products = Json.decodeFromString<List<Product>>(response.bodyAsText())

      products
    } catch (exception: Exception) {

      println(exception)
      null
    }
  }
}

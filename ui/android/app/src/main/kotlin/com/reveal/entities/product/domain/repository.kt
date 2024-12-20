package entities.product

<<<<<<< HEAD
import io.ktor.client.HttpClient
=======
>>>>>>> cart-ui
import io.ktor.client.request.get
import io.ktor.client.statement.bodyAsText
import io.ktor.util.InternalAPI
import kotlinx.serialization.json.Json
<<<<<<< HEAD

class ProductRepository(private val client: HttpClient) {
  @OptIn(InternalAPI::class)
  suspend fun getAllProducts(): List<Product>? {
    return try {
      val response = client.get("http://192.168.3.2/products")

      val products = Json.decodeFromString<List<Product>>(response.bodyAsText())

=======
import shared.api.HTTPClient

class ProductRepository() {
  @OptIn(InternalAPI::class)
  suspend fun getAllProducts(): List<Product>? {
    return try {
      val response = HTTPClient.client.get("http://192.168.3.2/products")

      val products = Json.decodeFromString<List<Product>>(response.bodyAsText())

      println(products)

>>>>>>> cart-ui
      products
    } catch (exception: Exception) {

      println(exception)
      null
    }
  }
<<<<<<< HEAD
=======

  @OptIn(InternalAPI::class)
  suspend fun getProduct(id: String): Product? {
    return try {
      val response = HTTPClient.client.get("http://192.168.3.2/products/${id}")
      val product = Json.decodeFromString<Product>(response.bodyAsText())

      println(product)

      product
    } catch (exception: Exception) {

      println(exception)
      null
    }
  }
>>>>>>> cart-ui
}

package entities.cart

import io.ktor.client.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.HttpHeaders
import io.ktor.http.contentType
import io.ktor.http.isSuccess
import io.ktor.util.InternalAPI
import kotlinx.serialization.json.Json
import kotlinx.serialization.serializer
import shared.api.HTTPClient

class CartRepository(val token: String?) {
  private val json = Json { ignoreUnknownKeys = true }

  @OptIn(InternalAPI::class)
  suspend fun addToCart(dto: CartItemDTO): CartItem? {
    return try {

      val response: HttpResponse =
              HTTPClient.client.post("http://192.168.3.2/cart/products") {
                contentType(io.ktor.http.ContentType.Application.Json)
                body = json.encodeToString(CartItemDTO.serializer(), dto)
                header(HttpHeaders.Authorization, "Bearer $token")
              }

      val cart_item = Json.decodeFromString<CartItem>(response.bodyAsText())

      cart_item
    } catch (exception: Exception) {
      println(exception)
      null
    }
  }

  @OptIn(InternalAPI::class)
  suspend fun getCart(): Cart? {
    return try {
      val response: HttpResponse =
              HTTPClient.client.get("http://192.168.3.2/cart") {
                contentType(io.ktor.http.ContentType.Application.Json)
                header(HttpHeaders.Authorization, "Bearer $token")
              }
      println("CART response $response")
      val cart = Json { ignoreUnknownKeys = true }.decodeFromString<Cart>(response.bodyAsText())
      println("CART $cart")

      cart
    } catch (exception: Exception) {
      println("CART EXCEPTION $exception")
      null
    }
  }

  suspend fun removeFromCart(product_id: Int): Boolean {
    val response: HttpResponse =
            HTTPClient.client.delete("http://192.168.3.2/cart/products/${product_id}") {
              header(HttpHeaders.Authorization, "Bearer $token")
            }

    return response.status.isSuccess()
  }

  // suspend fun getCartItems(): List<CartItem> {
  //     val response: HttpResponse = client.get("https://yourapi.com/cart")
  //     return json.decodeFromString(response.bodyAsText())
  // }

  // suspend fun updateQuantity(cartItemId: Int, quantity: Int): Boolean {
  //     val response: HttpResponse = client.put("https://yourapi.com/cart/$cartItemId") {
  //         contentType(io.ktor.http.ContentType.Application.Json)
  //         body = json.encodeToString(mapOf("quantity" to quantity))
  //     }
  //     return response.status.isSuccess()
  // }

  // suspend fun clearCart(): Boolean {
  //     val response: HttpResponse = client.delete("https://yourapi.com/cart")
  //     return response.status.isSuccess()
  // }
}

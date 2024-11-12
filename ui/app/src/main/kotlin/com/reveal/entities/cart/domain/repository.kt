package entities.cart

import io.ktor.client.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.contentType
import io.ktor.http.isSuccess
import io.ktor.util.InternalAPI
import kotlinx.serialization.json.Json
import shared.api.HTTPClient

class CartRepository() {
  private val json = Json { ignoreUnknownKeys = true }

  @OptIn(InternalAPI::class)
  suspend fun addToCart(dto: CartItemDTO): Boolean {
    val response: HttpResponse =
            HTTPClient.client.post("http://192.168.3.2/cart/1/products") {
              contentType(io.ktor.http.ContentType.Application.Json)
              body = json.encodeToString(CartItemDTO.serializer(), dto)
            }

    return response.status.isSuccess()
  }

  @OptIn(InternalAPI::class)
  suspend fun getCart(): Cart? {
    return try {
      val response: HttpResponse =
              HTTPClient.client.get("http://192.168.3.2/cart/1") {
                contentType(io.ktor.http.ContentType.Application.Json)
              }
      val cart = Json { ignoreUnknownKeys = true }.decodeFromString<Cart>(response.bodyAsText())
      println(cart)

      cart
    } catch (exception: Exception) {
      println(exception)
      null
    }
  }

  // suspend fun getCartItems(): List<CartItem> {
  //     val response: HttpResponse = client.get("https://yourapi.com/cart")
  //     return json.decodeFromString(response.bodyAsText())
  // }

  // suspend fun removeFromCart(cartItemId: Int): Boolean {
  //     val response: HttpResponse = client.delete("https://yourapi.com/cart/$cartItemId")
  //     return response.status.isSuccess()
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

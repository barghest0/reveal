package entities.cart

import io.ktor.client.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.contentType
import io.ktor.http.isSuccess
import io.ktor.util.InternalAPI
import kotlinx.serialization.json.Json
import kotlinx.serialization.serializer
import shared.api.HTTPClient

class CartRepository() {
  private val json = Json { ignoreUnknownKeys = true }

  @OptIn(InternalAPI::class)
  suspend fun addToCart(cartItem: CartItem): Boolean {
    val response: HttpResponse =
            HTTPClient.client.post("http://192.168.3.2/cart") {
              contentType(io.ktor.http.ContentType.Application.Json)
              body = json.encodeToString(CartItem.serializer(), cartItem)
            }
    return response.status.isSuccess()
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

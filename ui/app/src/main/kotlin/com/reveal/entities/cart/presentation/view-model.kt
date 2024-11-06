package entities.cart

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import entities.product.Product
import kotlinx.coroutines.launch

class CartViewModel(private val cartRepository: CartRepository) : ViewModel() {
  val products = mutableStateListOf<CartItem>()

  fun addToCart(product: Product) {
    val cartItem = CartItem(product_id = product.id, quantity = 1, price = product.price)
    viewModelScope.launch { cartRepository.addToCart(cartItem) }
  }

  fun getCart() {
    viewModelScope.launch {
      var cart = cartRepository.getCart()
      if (cart != null) {
        products.addAll(cart.Products)
      }
    }
  }
}

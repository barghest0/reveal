package entities.cart

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import entities.product.Product
import kotlinx.coroutines.launch

class CartViewModel(private val cartRepository: CartRepository) : ViewModel() {
  val products = mutableStateListOf<CartItem>()

  init {
    // Загружаем корзину при инициализации
    getCart()
  }

  fun addToCart(product: Product) {
    val cartItem = CartItemDTO(product_id = product.id, quantity = 1)
    viewModelScope.launch {
      val is_success = cartRepository.addToCart(cartItem)
      // if (is_success) {
      //   products.add(cartItem)
      // }
    }
  }

  fun getCart() {
    viewModelScope.launch {
      var cart = cartRepository.getCart()
      if (cart != null) {
        products.clear()
        products.addAll(cart.products)
      }
    }
  }

  fun isProductExist(product_id: Int): Boolean {
    return products.any { it.product_id == product_id }
  }
}

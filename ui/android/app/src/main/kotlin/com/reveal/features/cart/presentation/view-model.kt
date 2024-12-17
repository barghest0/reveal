package features.cart

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import entities.product.Product
import kotlinx.coroutines.launch

class CartViewModel(private val cartRepository: CartRepository) : ViewModel() {
  private val _products = mutableStateListOf<CartItem>()
  val products: List<CartItem>
    get() = _products

  init {
    getCart()
  }

  fun addToCart(product: Product) {
    val quantity = 1
    val cartItem =
            CartItemDTO(
                    product_id = product.id,
                    quantity = quantity,
                    price = product.price * quantity
            )
    viewModelScope.launch {
      val cart_item = cartRepository.addToCart(cartItem)
      if (cart_item != null) {
        _products.add(cart_item)
      }
    }
  }

  fun removeFromCart(product_id: Int) {
    viewModelScope.launch {
      val is_success = cartRepository.removeFromCart(product_id)
      if (is_success) {
        _products.removeIf { it.product_id == product_id }
      }
    }
  }

  fun getCart() {
    viewModelScope.launch {
      val cart = cartRepository.getCart()
      if (cart != null) {
        _products.clear()
        _products.addAll(cart.products)
      }
    }
  }

  fun isProductExist(product_id: Int): Boolean {
    return _products.any { it.product_id == product_id }
  }
}

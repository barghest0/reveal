package entities.cart

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import entities.product.Product
import kotlinx.coroutines.launch

class CartViewModel(private val cartRepository: CartRepository) : ViewModel() {
  val products = mutableStateListOf<CartItem>()
  val existed_products = mutableStateListOf<Int>()

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
      products.add(cart_item)
    }
  }

  fun getCart() {
    viewModelScope.launch {
      var cart = cartRepository.getCart()
      if (cart != null) {
        products.clear()
        products.addAll(cart.products)
        existed_products.clear()
        existed_products.addAll(cart.products.map { it.product_id })
      }
    }
  }

  fun isProductExist(product_id: Int): Boolean {
    return products.any { it.product_id == product_id }
  }
}

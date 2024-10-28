package entities.cart

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import entities.product.Product
import kotlinx.coroutines.launch

class CartViewModel(private val cartRepository: CartRepository) : ViewModel() {
  // Список товаров в корзине
  val cartItems = mutableStateListOf<CartItem>()

  // Функция для добавления товара в корзину
  fun addToCart(product: Product) {
    val cartItem =
            CartItem(
                    id = 0,
                    cart_id = 0,
                    product_id = product.id,
                    quantity = 1,
                    price = product.price
            )
    viewModelScope.launch {
      if (cartRepository.addToCart(cartItem)) {
        cartItems.add(cartItem) // Обновляем состояние корзины
      }
    }
  }
}

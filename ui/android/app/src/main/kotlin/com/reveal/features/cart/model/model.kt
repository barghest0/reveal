package features.cart

import entities.product.CartProduct
import kotlinx.serialization.Serializable

@Serializable data class Cart(val id: Int, val user_id: Int, val products: List<CartItem>)

@Serializable
data class CartItem(
        val id: Int,
        val cart_id: Int,
        val product_id: Int,
        val quantity: Int,
        val price: Double,
        val product: CartProduct
)

@Serializable
data class CartItemDTO(
        val product_id: Int,
        val quantity: Int,
        val price: Int,
)

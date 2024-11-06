package entities.cart

import kotlinx.serialization.Serializable

@Serializable
data class CartItem(
        val id: Int,
        val cart_id: Int,
        val product_id: Int,
        val quantity: Int,
        val price: Double
)

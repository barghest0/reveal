package entities.cart

import kotlinx.serialization.Serializable

@Serializable data class Cart(val id: Int, val user_id: Int, val products: List<CartItem>)

@Serializable
data class CartProduct(val id: Int, val name: String, val description: String, val price: Int)

@Serializable
data class CartItem(
        val product_id: Int,
        val quantity: Int,
        val price: Double,
        val product: CartProduct
)

@Serializable
data class CartItemDTO(
        val product_id: Int,
        val quantity: Int,
)

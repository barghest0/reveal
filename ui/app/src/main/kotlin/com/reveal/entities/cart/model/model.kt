package entities.cart

import kotlinx.serialization.Serializable

@Serializable data class Cart(val id: Int, val user_id: Int, val Products: List<CartItem>)

@Serializable data class CartItem(val product_id: Int, val quantity: Int, val price: Double)

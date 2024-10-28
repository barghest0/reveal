package entities.product

import kotlinx.serialization.Serializable

@Serializable
data class Product(
        val id: Int,
        val name: String,
        val description: String,
        val price: Double,
        val seller_id: Int,
        val created_at: String,
        val updated_at: String
)

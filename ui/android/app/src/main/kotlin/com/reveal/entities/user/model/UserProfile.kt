package entities.user

import kotlinx.serialization.Serializable

@Serializable
data class UserProfile(val id: Int, val name: String, val email: String, val password: String)

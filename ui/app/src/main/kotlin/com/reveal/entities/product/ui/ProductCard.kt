package entities.product

import androidx.compose.foundation.layout.Box
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable

@Composable
fun ProductCard(product: Product) {
  Box() { Text("${product.name}") }
}

package widgets.ProductsList

import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.runtime.Composable
import entities.product.Product
import entities.product.ProductCard

@Composable
fun ProductsList(products: List<Product>) {
  LazyColumn { items(products) { product -> ProductCard(product = product) } }
}

package entities.product

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

class ProductViewModel(private val repository: ProductRepository) : ViewModel() {
  private val _products = mutableStateListOf<Product>()
  val products: List<Product>
    get() = _products

  fun fetchProducts() {
    viewModelScope.launch {
      withContext(Dispatchers.IO) {
        try {
          val fetchedProducts = repository.getAllProducts()
          if (fetchedProducts != null) {
            _products.clear()
            _products.addAll(fetchedProducts)
          }
        } catch (e: Exception) {
          println(e)
        }
      }
    }
  }
}

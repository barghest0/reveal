package entities.product

import androidx.compose.runtime.mutableStateListOf
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

class ProductViewModel(private val repository: ProductRepository) : ViewModel() {

  private val _products = mutableStateListOf<Product>()
  val products: List<Product>
    get() = _products

  var product = mutableStateOf<Product?>(null)
    private set

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

  fun fetchProduct(id: String) {
    viewModelScope.launch {
      withContext(Dispatchers.IO) {
        try {
          val fetchedProduct = repository.getProduct(id)
          product.value = fetchedProduct
        } catch (e: Exception) {
          println(e)
        }
      }
    }
  }
}

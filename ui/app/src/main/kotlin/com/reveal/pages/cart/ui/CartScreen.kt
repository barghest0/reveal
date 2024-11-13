package pages.cart

import androidx.compose.runtime.Composable
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import entities.cart.CartRepository
import entities.cart.CartViewModel
import kotlinx.coroutines.*
import shared.ui.layout.ScreenLayout
import widgets.CartProductsList.CartProductsList

@Composable
fun CartScreen(
        navController: NavController,
        viewModel: CartViewModel = CartViewModel(CartRepository())
) {

  ScreenLayout { CartProductsList(navController, viewModel.products) }
}

// какашечкии пуки каки каки пуки

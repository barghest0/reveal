package pages.cart

import android.content.Context
import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import entities.cart.CartRepository
import entities.cart.CartViewModel
import kotlinx.coroutines.*
import shared.session.PreferencesManager
import shared.ui.layout.ScreenLayout
import widgets.CartProductsList.CartProductsList

@Composable
fun CartScreen(
        navController: NavController,
        context: Context = LocalContext.current,
        viewModel: CartViewModel =
                CartViewModel(CartRepository(PreferencesManager(context).getToken()))
) {

  ScreenLayout { CartProductsList(navController, viewModel.products) }
}

// какашечкии пуки каки каки пуки

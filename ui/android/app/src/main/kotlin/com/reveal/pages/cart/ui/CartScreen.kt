package pages.cart

import android.content.Context
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.platform.LocalContext
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import features.cart.CartRepository
import features.cart.CartViewModel
import kotlinx.coroutines.*
import shared.session.PreferencesManager
import shared.ui.layout.ScreenLayout
import widgets.CartProductsList.CartProductsList

@Composable
fun CartScreen(
        navController: NavController,
        context: Context = LocalContext.current,
        viewModel: CartViewModel = remember {
          CartViewModel(CartRepository(PreferencesManager(context).getToken()))
        }
) {

  ScreenLayout { CartProductsList(navController, viewModel) }
}

// какашечкии пуки каки каки пуки

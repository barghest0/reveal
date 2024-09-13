package com.user_service

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import com.user_service.features.registration.ui.RegistrationScreen
import com.user_service.shared.ui.theme.UserserviceTheme

class MainActivity : ComponentActivity() {
  override fun onCreate(savedInstanceState: Bundle?) {
    super.onCreate(savedInstanceState)
    enableEdgeToEdge()
    setContent {
      UserserviceTheme {
        Surface(color = MaterialTheme.colorScheme.background) {
          RegistrationScreen()
          // Create the NavController
          // val navController = rememberNavController()
          //
          // // Set up the NavHost
          // AppNavigation(navController = navController)
        }
      }
    }
  }
}

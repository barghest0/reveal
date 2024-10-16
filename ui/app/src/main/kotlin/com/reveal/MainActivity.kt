package com.reveal

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.navigation.compose.rememberNavController
import com.reveal.app.navigation.ui.AppNavigation
import com.reveal.shared.ui.theme.RevealTheme

class MainActivity : ComponentActivity() {
  override fun onCreate(savedInstanceState: Bundle?) {
    super.onCreate(savedInstanceState)
    enableEdgeToEdge()
    setContent {
      RevealTheme {
        Surface(color = MaterialTheme.colorScheme.background) {
          // Create the NavController
          val navController = rememberNavController()

          // Set up the NavHost
          AppNavigation(navController = navController)
        }
      }
    }
  }
}

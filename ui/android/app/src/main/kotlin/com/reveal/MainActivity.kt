package com.reveal

import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.navigation.compose.rememberNavController
import app.AppNavigation
import com.google.firebase.messaging.FirebaseMessaging
import shared.navigation.NavigationController
import shared.ui.theme.RevealTheme

class MainActivity : ComponentActivity() {
  private val navigtaion = NavigationController()
  override fun onCreate(savedInstanceState: Bundle?) {
    super.onCreate(savedInstanceState)
    enableEdgeToEdge()
    FirebaseMessaging.getInstance().token.addOnCompleteListener { task ->
      if (task.isSuccessful) {
        // Получаем токен устройства
        val token = task.result
        Log.d("FCM", "Токен устройства: $token")

        // Отправить токен на сервер
        sendTokenToServer(token)
      } else {
        // Ошибка получения токена
        Log.w("FCM", "Не удалось получить токен", task.exception)
      }
    }
    setContent {
      RevealTheme {
        Surface(color = MaterialTheme.colorScheme.background) {
          val navController = rememberNavController()
          navigtaion.controller = navController

          AppNavigation(navController)
        }
      }
    }
  }

  private fun sendTokenToServer(token: String) {
    // Здесь вы можете отправить токен на ваш сервер с помощью Retrofit, OkHttp или другого
    // механизма.
    Log.d("FCM", "Отправка токена на сервер: $token")
  }
}

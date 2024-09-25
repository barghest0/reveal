package com.user_service.app.navigation.ui

import androidx.compose.runtime.Composable
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import com.user_service.pages.login.ui.LoginScreen
import com.user_service.pages.registration.ui.RegistrationScreen

@Composable
fun AppNavigation(navController: NavHostController) {
  NavHost(navController = navController, startDestination = "registration") {
    composable("registration") { RegistrationScreen(navController) }
    composable("login") { LoginScreen() }
  }
}

package com.reveal.app.navigation.ui

import ProfileScreen
import androidx.compose.runtime.Composable
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import com.reveal.pages.login.ui.LoginScreen
import com.reveal.pages.registration.ui.RegistrationScreen

@Composable
fun AppNavigation(navController: NavHostController) {
  NavHost(navController = navController, startDestination = "registration") {
    composable("registration") { RegistrationScreen(navController) }
    composable("login") { LoginScreen(navController) }
    composable("profile") { ProfileScreen(navController) }
  }
}

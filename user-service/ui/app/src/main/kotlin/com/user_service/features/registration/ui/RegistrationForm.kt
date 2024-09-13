package com.user_service.features.registration.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.tooling.preview.Preview

@Composable
fun RegistrationForm(
        username: String,
        email: String,
        password: String,
        onRegisterClicked: (String, String, String) -> Unit
) {
  Column {
    OutlinedTextField(
            value = username,
            onValueChange = { /* Handle username change */},
            label = { /* Label for username field */}
    )
    OutlinedTextField(
            value = email,
            onValueChange = { /* Handle email change */},
            label = { /* Label for email field */}
    )
    OutlinedTextField(
            value = password,
            onValueChange = { /* Handle password change */},
            label = { /* Label for password field */},
    )
    Button(onClick = { onRegisterClicked(username, email, password) }) { Text("Register") }
  }
}

@Preview
@Composable
fun PreviewRegistrationForm() {
  RegistrationForm(username = "", email = "", password = "", onRegisterClicked = { _, _, _ -> })
}

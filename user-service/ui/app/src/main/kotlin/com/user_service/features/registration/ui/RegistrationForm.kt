package com.user_service.features.registration.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.tooling.preview.Preview

@Composable
fun RegistrationForm(
        name: String,
        email: String,
        password: String,
        updateField: (String, String)->Unit,
        onRegisterClicked: (String, String, String) -> Unit
) {
  Column {
    OutlinedTextField(
            value = name,
            onValueChange = {name -> updateField("name", name) },
            label = { /* Label for username field */}
    )
    OutlinedTextField(
            value = email,
            onValueChange = { email -> updateField("email", email) },
            label = { /* Label for email field */}
    )
    OutlinedTextField(
            value = password,
            onValueChange = { password -> updateField("password", password)},
            label = { /* Label for password field */},
    )
    Button(onClick = { onRegisterClicked(name, email, password) }) { Text("Register") }
  }
}

@Preview
@Composable
fun PreviewRegistrationForm() {
  RegistrationForm(name = "", email = "", password = "", updateField = {_,_ ->}, onRegisterClicked = { _, _, _ -> })
}

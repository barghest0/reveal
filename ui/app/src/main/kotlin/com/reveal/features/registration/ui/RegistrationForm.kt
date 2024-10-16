package com.reveal.features.registration.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable

@Composable
fun RegistrationForm(
    name: String,
    email: String,
    password: String,
    updateField: (String, String) -> Unit,
    onClickRegister: (String, String, String) -> Unit
) {
  Column() {
    OutlinedTextField(
        value = name,
        onValueChange = { name -> updateField("name", name) },
        label = { Text("Name") })
    OutlinedTextField(
        value = email,
        onValueChange = { email -> updateField("email", email) },
        label = { Text("Email") })
    OutlinedTextField(
        value = password,
        onValueChange = { password -> updateField("password", password) },
        label = { Text("Password") },
    )
    Button(onClick = { onClickRegister(name, email, password) }) { Text("Register") }
  }
}

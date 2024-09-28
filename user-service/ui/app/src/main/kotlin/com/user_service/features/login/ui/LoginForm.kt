package com.user_service.features.login.ui

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable

@Composable
fun LoginForm(
    name: String,
    password: String,
    updateField: (String, String) -> Unit,
    onClickLogin: (String, String) -> Unit
) {
  Column() {
    OutlinedTextField(
        value = name,
        onValueChange = { name -> updateField("name", name) },
        label = { Text("Name") })
    OutlinedTextField(
        value = password,
        onValueChange = { password -> updateField("password", password) },
        label = { Text("Password") },
    )
    Button(onClick = { onClickLogin(name, password) }) { Text("Login") }
  }
}

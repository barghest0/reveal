package shared.ui.fields

import androidx.compose.material3.OutlinedTextField
import androidx.compose.runtime.Composable

@Composable
fun TextInput(value: String, onValueChange: (String) -> Unit, label: @Composable (() -> Unit)?) {
  OutlinedTextField(value = value, onValueChange = onValueChange, label = label)
}

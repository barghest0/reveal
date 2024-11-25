package shared.ui.buttons

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable

@Composable
fun AppButton(onClick: () -> Unit, label: String) {
  Button(onClick = onClick) { Text(text = label) }
}

package shared.ui.layout

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp

@Composable
fun ScreenLayout(padding: Dp = 16.dp, content: @Composable () -> Unit) {
  Column(modifier = Modifier.fillMaxSize().padding(horizontal = padding, vertical = padding)) {
    content()
  }
}

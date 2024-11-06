package pages.product

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import entities.product.Product

@Composable
fun ProductDetails(product: Product) {
  MaterialTheme {
    Scaffold() { innerPadding ->
      Column(modifier = Modifier.fillMaxSize().padding(innerPadding).padding(16.dp)) {
        Spacer(modifier = Modifier.height(16.dp))

        // Название продукта
        Text(
                text = product.name,
                fontSize = 24.sp,
                fontWeight = FontWeight.Bold,
                color = Color.Black,
                modifier = Modifier.padding(bottom = 4.dp)
        )

        // Цена продукта
        Text(
                text = "${product.price} ₽",
                fontSize = 20.sp,
                color = Color(0xFFD32F2F),
                fontWeight = FontWeight.Bold,
                modifier = Modifier.padding(bottom = 8.dp)
        )

        // Описание продукта
        Text(
                text = product.description,
                fontSize = 16.sp,
                color = Color.Gray,
                modifier = Modifier.padding(bottom = 16.dp)
        )
        // Кнопка добавления в корзину
        Button(
                onClick = {
                  // Логика добавления в корзину
                },
                modifier = Modifier.fillMaxWidth().height(56.dp),
                shape = RoundedCornerShape(12.dp),
                colors = ButtonDefaults.buttonColors(containerColor = Color(0xFF6200EE))
        ) { Text(text = "Добавить в корзину", fontSize = 18.sp, color = Color.White) }
      }
    }
  }
}

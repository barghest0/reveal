package com.reveal

import android.app.NotificationChannel
import android.app.NotificationManager
import android.content.Context
import android.os.Build
import androidx.core.app.NotificationCompat
import com.google.firebase.messaging.FirebaseMessagingService
import com.google.firebase.messaging.RemoteMessage

class CloudMessagingService : FirebaseMessagingService() {

  override fun onMessageReceived(remoteMessage: RemoteMessage) {
    println("RemoteMessage ${remoteMessage.data.isNotEmpty()}")
    // Проверка на данные
    if (remoteMessage.data.isNotEmpty()) {
      // Обработка данных
      val data = remoteMessage.data
      // Логирование или другие действия
      println("Data payload: $data")
    }

    // Проверка на уведомление
    remoteMessage.notification?.let {
      val title = it.title ?: "No Title"
      val body = it.body ?: "No Message"
      println("Notification title: $title, body: $body")
      showNotification(title, body)
    }
  }

  private fun showNotification(title: String, body: String) {
    val notificationManager = getSystemService(Context.NOTIFICATION_SERVICE) as NotificationManager

    // Для Android 8.0 и выше необходимо создать канал уведомлений
    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
      val channel =
              NotificationChannel("orders", "Notifications", NotificationManager.IMPORTANCE_DEFAULT)
                      .apply { description = "Channel for order notifications" }

      // Создаем канал уведомлений, если его еще нет
      notificationManager.createNotificationChannel(channel)
    }

    // Создание уведомления с иконкой
    val notification =
            NotificationCompat.Builder(this, "orders")
                    .setContentTitle(title)
                    .setContentText(body)
                    .setSmallIcon(R.drawable.ic_notification) // Иконка уведомления
                    .setPriority(
                            NotificationCompat.PRIORITY_DEFAULT
                    ) // Задаем приоритет уведомления
                    .setDefaults(NotificationCompat.DEFAULT_ALL) // Включает звук, вибрацию и др.
                    .setAutoCancel(true) // Уведомление исчезает при клике
                    .build()

    println("Notification $notification")

    notificationManager.notify(0, notification)
  }
}

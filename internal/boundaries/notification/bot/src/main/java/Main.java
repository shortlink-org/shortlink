import com.rabbitmq.client.DeliverCallback;
import http.LinkApp;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.telegram.telegrambots.meta.TelegramBotsApi;
import org.telegram.telegrambots.meta.exceptions.TelegramApiException;
import org.telegram.telegrambots.updatesreceivers.DefaultBotSession;

import java.io.IOException;
import java.net.URISyntaxException;
import java.security.KeyManagementException;
import java.security.NoSuchAlgorithmException;
import java.util.concurrent.TimeoutException;

public class Main {

  public static void main(String[] args) throws TelegramApiException, IOException, TimeoutException, URISyntaxException, NoSuchAlgorithmException, KeyManagementException {
    Logger logger = LoggerFactory.getLogger(Main.class);

    // Init API
//    LinkApp api = new LinkApp();

    // Subscribe on AMQP, exchange: 'shortlink'
//    DeliverCallback deliverCallback = (consumerTag, delivery) -> {
//      String message = new String(delivery.getBody(), "UTF-8");
//      System.out.println(" [x] Received '" + message + "'");
//    };

    // RabbitMQ
//    RabbitMQ rabbitmq = new RabbitMQ(deliverCallback);

    // Instantiate Telegram Bots API
//    TelegramBotsApi telegramBotsApi = new TelegramBotsApi(DefaultBotSession.class);

    // Registration our bot
//    try {
//      telegramBotsApi.registerBot(new TelegramBot(rabbitmq, api));
//    } catch (TelegramApiException e) {
//      e.printStackTrace();
//    }

    logger.info("ShortLinkBot successfully started!");
  }

}

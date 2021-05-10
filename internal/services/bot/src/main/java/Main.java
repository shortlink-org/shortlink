import org.apache.log4j.LogManager;
import org.apache.log4j.Logger;
import org.telegram.telegrambots.meta.TelegramBotsApi;
import org.telegram.telegrambots.meta.exceptions.TelegramApiException;
import org.telegram.telegrambots.updatesreceivers.DefaultBotSession;

public class Main {
  final static Logger logger = LogManager.getLogger(Main.class);

  public static void main(String[] args) throws TelegramApiException {
    // Instantiate Telegram Bots API
    TelegramBotsApi telegramBotsApi = new TelegramBotsApi(DefaultBotSession.class);

    // Register our bot
    try {
      telegramBotsApi.registerBot(new TelegramBot());
    } catch (TelegramApiException e) {
      e.printStackTrace();
    }

    logger.info("ShortLinkBot successfully started!");
  }

}

import org.telegram.telegrambots.bots.TelegramLongPollingBot;
import org.telegram.telegrambots.meta.api.methods.send.SendMessage;
import org.telegram.telegrambots.meta.api.objects.Update;
import org.telegram.telegrambots.meta.api.objects.replykeyboard.ReplyKeyboardMarkup;
import org.telegram.telegrambots.meta.api.objects.replykeyboard.buttons.KeyboardRow;
import org.telegram.telegrambots.meta.exceptions.TelegramApiException;

import java.util.ArrayList;
import java.util.List;

public class TelegramBot extends TelegramLongPollingBot {
  private String TELEGRAM_BOT_TOKEN = System.getenv("TELEGRAM_BOT_TOKEN");
  private String TELEGRAM_BOT_USERNAME = System.getenv("TELEGRAM_BOT_USERNAME");

  @Override
  public void onUpdateReceived(Update update) {
    // We check if the update has a message and the message has text
    if (update.hasMessage() && update.getMessage().hasText()) {
      // Set variables
      Long chatId = update.getMessage().getChatId();
      String message_text = update.getMessage().getText();

      // Create a message object
      SendMessage message = new SendMessage();
      message.setChatId(String.valueOf(chatId));

      // Create ReplyKeyboardMarkup object
      ReplyKeyboardMarkup keyboardMarkup = new ReplyKeyboardMarkup();
      // Create the keyboard (list of keyboard rows)
      List<KeyboardRow> keyboard = new ArrayList<>();
      // Create a keyboard row
      KeyboardRow row = new KeyboardRow();

      switch (message_text) {
        case "/get":
          message.setText("get link");

          row.add("/list");
          row.add("/create");
          row.add("/update");
          row.add("/delete");
          row.add("/help");
          break;
        case "/list":
          message.setText("get list");

          row.add("/list");
          row.add("/create");
          row.add("/update");
          row.add("/delete");
          row.add("/help");
          break;
        case "/add":
          message.setText("add a new link");

          row.add("/get");
          row.add("/list");
          row.add("/update");
          row.add("/delete");
          row.add("/help");
          break;
        case "/update":
          message.setText("update link");

          row.add("/get");
          row.add("/list");
          row.add("/create");
          row.add("/delete");
          row.add("/help");
          break;
        case "/delete":
          message.setText("delete link");

          row.add("/get");
          row.add("/list");
          row.add("/create");
          row.add("/update");
          row.add("/help");
          break;
        default:
        case "/help":
          message.setText("helps");

          row.add("/get");
          row.add("/list");
          row.add("/create");
          row.add("/update");
          row.add("/delete");
          row.add("/help");
          break;
      }

      // Add the first row to the keyboard
      keyboard.add(row);

      // Set the keyboard to the markup
      keyboardMarkup.setKeyboard(keyboard);

      // Add it to the message
      message.setReplyMarkup(keyboardMarkup);

      try {
        // Sending our message object to user
        execute(message);
      } catch (TelegramApiException e) {
        e.printStackTrace();
      }
    }
  }

  @Override
  public String getBotUsername() {
    // Return bot username
    return TELEGRAM_BOT_TOKEN;
  }

  @Override
  public String getBotToken() {
    // Return bot token from BotFather
    return TELEGRAM_BOT_TOKEN;
  }
}

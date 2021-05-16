import http.Link;
import http.LinkApp;
import org.jetbrains.annotations.NotNull;
import org.telegram.telegrambots.bots.TelegramLongPollingBot;
import org.telegram.telegrambots.meta.api.methods.send.SendMessage;
import org.telegram.telegrambots.meta.api.objects.Update;
import org.telegram.telegrambots.meta.api.objects.replykeyboard.ReplyKeyboardMarkup;
import org.telegram.telegrambots.meta.api.objects.replykeyboard.buttons.KeyboardRow;
import org.telegram.telegrambots.meta.exceptions.TelegramApiException;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class TelegramBot extends TelegramLongPollingBot {
  private String TELEGRAM_BOT_TOKEN = System.getenv("TELEGRAM_BOT_TOKEN");
  private String TELEGRAM_BOT_USERNAME = System.getenv("TELEGRAM_BOT_USERNAME");

  private RabbitMQ rabbitmq;
  private LinkApp api;

  public TelegramBot(@NotNull RabbitMQ rabbitmq, LinkApp api) {
    this.rabbitmq = rabbitmq;
    this.api = api;
  }

  @Override
  public void onUpdateReceived(@NotNull Update update) {
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

      // Parse command from telegram CLIENT
      ArrayList<String> userCommand = this.parseMessage(message_text);

      switch (userCommand.get(0)) {
        case "/get":
          message.setText("get link");
          break;
        case "/list":
          message.setText("get list");
          break;
        case "/add":
          try {
            // TODO: validate and check URL
            Link resp = this.api.AddLink(userCommand.get(1));
            message.setText("add a new link: " + resp.getHash());
          } catch (IOException e) {
            e.printStackTrace();
            message.setText("error create a new new link: " + e.getMessage());
          }
          break;
        case "/update":
          message.setText("update link");
          break;
        case "/delete":
          message.setText("delete link");
          break;
        default:
        case "/help":
          message.setText("helps");
          break;
      }

      // Add the first row to the keyboard
      this.addDefaultCommand(row);
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
    return TELEGRAM_BOT_USERNAME;
  }

  @Override
  public String getBotToken() {
    // Return bot token from BotFather
    return TELEGRAM_BOT_TOKEN;
  }

  private void addDefaultCommand(@NotNull KeyboardRow row) {
    row.add("/get");
    row.add("/list");
    row.add("/add");
    row.add("/update");
    row.add("/delete");
    row.add("/help");
  }

  @NotNull
  private ArrayList<String> parseMessage(@NotNull String message) {
    // TODO: refactoring

    ArrayList<String> resp = new ArrayList<String>();
    String[] list = message.split(" ");

    // get command
    resp.add(list[0]);

    String[] contentArr = Arrays.copyOfRange(list, 1, list.length);
    // merge content
    String content = String.join(" ", contentArr);
    resp.add(content);

    return resp;
  }
}

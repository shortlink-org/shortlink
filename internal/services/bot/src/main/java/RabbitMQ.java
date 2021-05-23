import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;
import com.rabbitmq.client.DeliverCallback;
import org.jetbrains.annotations.NotNull;

import java.io.IOException;
import java.net.URISyntaxException;
import java.security.KeyManagementException;
import java.security.NoSuchAlgorithmException;
import java.util.concurrent.TimeoutException;

public class RabbitMQ {
  private final static String EXCHANGE_NAME = "shortlink";
  private final static String QUEUE_NAME = "shortlink-bot";
  private String MQ_RABBIT_URI = System.getenv("MQ_RABBIT_URI");
  private Channel channel;

  public RabbitMQ(DeliverCallback deliverCallback) throws IOException, TimeoutException, URISyntaxException, NoSuchAlgorithmException, KeyManagementException {
    ConnectionFactory factory = new ConnectionFactory();

    factory.setUri(MQ_RABBIT_URI);
    Connection connection = factory.newConnection();
    this.channel = connection.createChannel();

    this.channel.exchangeDeclare(EXCHANGE_NAME, "fanout");
    this.channel.queueDeclare(QUEUE_NAME, true, false, false, null);
    this.channel.queueBind(QUEUE_NAME, EXCHANGE_NAME, "");

    this.channel.basicConsume(QUEUE_NAME, true, deliverCallback, consumerTag -> { });
  }

  public void Send(@NotNull String message) throws IOException {
    this.channel.basicPublish(EXCHANGE_NAME, "", null, message.getBytes("UTF-8"));
  }
}

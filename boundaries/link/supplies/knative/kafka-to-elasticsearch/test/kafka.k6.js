import { check } from 'k6';
import { writer, produce, reader, consume, createTopic } from 'k6/x/kafka';

const bootstrapServers = ['shortlink-kafka-bootstrap.kafka:9092'];
const kafkaTopic = 'shortlink.link.event.new';

createTopic(bootstrapServers[0], kafkaTopic);

const producer = writer(bootstrapServers, kafkaTopic);

export default function () {
  const messages = [
    {
      value: JSON.stringify({
        title: 'Load Testing SQL Databases with k6',
        url: 'https://k6.io/blog/load-testing-sql-databases-with-k6/',
        locale: 'en',
      }),
    },
  ];

  const error = produce(producer, messages);
  check(error, {
    'is sent': (err) => err == undefined,
  });
}

export function teardown(data) {
  producer.close();
}

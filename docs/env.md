|Name | Default Value | Description |
|---|---|---|
| "BOT_SLACK_WEBHOOK" | "YOUR_WEBHOOK_URL_HERE" | Your webhook URL |
| "BOT_SMTP_FROM" | "example@site.com" |  |
| "BOT_SMTP_PASS" | "YOUR_PASSWORD" |  |
| "BOT_SMTP_TO" | "EMAIL_USER" |  |
| "BOT_SMTP_HOST" | "smtp.gmail.com" |  |
| "BOT_SMTP_ADDR" | "smtp.gmail.com:587" |  |
| "BOT_TELEGRAM_WEBHOOK" | "YOUR_WEBHOOK_URL_HERE" | Your webhook URL |
| "BOT_TELEGRAM_CHAT_ID" | 123 | Your chat ID |
| "BOT_TELEGRAM_DEBUG_MODE" | false |  |
| "LOG_LEVEL" | logger.INFO_LEVEL |  |
| "LOG_TIME_FORMAT" | time.RFC3339Nano |  |
| "TRACER_SERVICE_NAME" | "ShortLink" |  |
| "TRACER_URI" | "localhost:6831" |  |
| "MQ_ENABLED" | "false" |  |
| "LOG_LEVEL" | logger.INFO_LEVEL |  |
| "LOG_TIME_FORMAT" | time.RFC3339Nano |  |
| "TRACER_SERVICE_NAME" | "ShortLink" |  |
| "TRACER_URI" | "localhost:6831" |  |
| "MQ_ENABLED" | "false" |  |
| "MQ_TYPE" | "kafka" |  |
| "MQ_KAFKA_URI" | "localhost:9092" |  |
| "MQ_KAFKA_CONSUMER_GROUP" | "shortlink" |  |
| "STORE_TYPE" | "ram" |  |
| "STORE_BADGER_PATH" | "/tmp/links.badger" |  |
| "STORE_CASSANDRA_URI" | "localhost:9042" |  |
| "STORE_DGRAPH_URI" | "localhost:9080" |  |
| "STORE_LEVELDB_PATH" | "/tmp/links.db" |  |
| "STORE_MONGODB_URI" | "mongodb://localhost:27017" |  |
| "STORE_MYSQL_URI" | "shortlink:shortlink@(localhost:3306)/shortlink?parseTime=true" |  |
| "STORE_POSTGRES_URI" | dbinfo |  |
| "STORE_REDIS_URI" | "localhost:6379" |  |
| "STORE_SCYLLA_URI" | "localhost:9042" |  |
| "STORE_SQLITE_PATH" | "/tmp/links.sqlite" |  |
| "API_TYPE" | "http-chi" |  |
| "API_PORT" | 7070 |  |
| "API_TIMEOUT" | 60 |  |

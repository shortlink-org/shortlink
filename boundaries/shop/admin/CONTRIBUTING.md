# Contributing

### Getting started

We use Makefile for build and deploy.

```bash
$> make help # show help message with all commands and targets
```

#### Config

**Database (PostgreSQL)**

| Name              | Description       | Default value |
|-------------------|-------------------|---------------|
| POSTGRES_DB       | Database name     | shortlink     |
| POSTGRES_USER     | Database user     | postgres      |
| POSTGRES_PASSWORD | Database password | shortlink     |
| POSTGRES_HOST     | Database host     | localhost     |
| POSTGRES_PORT     | Database port     | 5432          |

**Cache (Redis)**

| Name       | Description | Default value            |
|------------|-------------|--------------------------|
| REDIS_URL  | Redis URL   | redis://localhost:6379/1 |

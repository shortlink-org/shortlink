# Contributing

### Getting started

We use Makefile for build and deploy.

```bash
$> make help # show help message with all commands and targets
```

#### Config

**Authentication**

| Name         | Description  | Default value                         |
|--------------|--------------|---------------------------------------|
| ORY_SDK_URL  | ORY SDK URL  | http://127.0.0.1:4433                 |
| ORY_UI_URL   | ORY UI URL   | http://127.0.0.1:3000/next/auth       |
| LOGIN_URL    | Login URL    | http://127.0.0.1:3000/next/auth/login |

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


### Debug mode

- [Plugin readme](https://django-debug-toolbar.readthedocs.io/en/latest/installation.html)

#### How enable debug mode?

For enable debug mode you need to set cookie `debug_enable=true` in your browser and was authenticated.

```
# Enable debug mode
document.cookie = "debug_enable=true; path=/";

# Disable debug mode
document.cookie = "debug_enable=false; path=/";
```

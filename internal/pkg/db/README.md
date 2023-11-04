# DataBase drivers

This package contains the database drivers for use in ShortLink services.

### URI format

We use the following format for the database URI:

![URI FORMAT](./docs/URI_FORMAT.png)

### Graceful shutdown

Safely terminate database interactions by closing the associated Context:

  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  // Utilize ctx in your database tasks
  ```

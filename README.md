# shortlink

Shortlink service

### HTTP API

+ [Postman link](./docs/shortlink.postman_collection.json)

- GET /:URL #=> 200 or 404
- POST / {"url":"example.com"} #=> 201
- DELETE / {"url":"example.com"} #=> 200
- GET /s/:URL #=> 301
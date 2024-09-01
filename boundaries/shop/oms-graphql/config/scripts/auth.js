function onRequest({ request }) {
  request.headers["user-id"] = "3e173751-8840-4b0d-8065-fbea88357cc4"

  return {
    request
  }
}

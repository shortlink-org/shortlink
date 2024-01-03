import http from 'k6/http'

export default function () {
  http.get('https://shortlink.best/api/links');
}

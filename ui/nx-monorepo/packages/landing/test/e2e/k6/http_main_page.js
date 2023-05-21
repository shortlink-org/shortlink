import { check } from 'k6'
import http from 'k6/http'

export default function () {
  const res = http.get('http://shortlink-landing.shortlink:8080')
  check(res, {
    'is status 200': (r) => r.status === 200,
  })
}

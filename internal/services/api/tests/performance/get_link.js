import http from 'k6/http';
import {check, sleep} from 'k6';

export let options = {
  vus: 10,
  duration: '30s',
};

export default function() {
  let res = http.get('http://localhost:7070/api/link/83873db19');
  check(res, {
    'status is 200': r => r.status === 200,
  });
  sleep(0.1);
}

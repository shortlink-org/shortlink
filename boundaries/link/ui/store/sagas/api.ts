import { Api, ApiConfig } from '../api/Api'

const config: ApiConfig = {
  baseUrl: '/api',
}

export default new Api(config)
